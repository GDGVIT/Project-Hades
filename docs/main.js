require.config({
  paths: {
    bootstrap: './vendor/bootstrap.min',
    diffMatchPatch: './vendor/diff_match_patch.min',
    handlebars: './vendor/handlebars.min',
    handlebarsExtended: './utils/handlebars_helper',
    jquery: './vendor/jquery.min',
    locales: './locales/locale',
    lodash: './vendor/lodash.custom.min',
    pathToRegexp: './vendor/path-to-regexp/index',
    prettify: './vendor/prettify/prettify',
    semver: './vendor/semver.min',
    utilsSampleRequest: './utils/send_sample_request',
    webfontloader: './vendor/webfontloader',
    list: './vendor/list.min',
  },
  shim: {
    bootstrap: {
      deps: ['jquery'],
    },
    diffMatchPatch: {
      exports: 'diff_match_patch',
    },
    handlebars: {
      exports: 'Handlebars',
    },
    handlebarsExtended: {
      deps: ['jquery', 'handlebars'],
      exports: 'Handlebars',
    },
    prettify: {
      exports: 'prettyPrint',
    },
  },
  urlArgs: `v=${(new Date()).getTime()}`,
  waitSeconds: 15,
});

require([
  'jquery',
  'lodash',
  'locales',
  'handlebarsExtended',
  './api_project.js',
  './api_data.js',
  'prettify',
  'utilsSampleRequest',
  'semver',
  'webfontloader',
  'bootstrap',
  'pathToRegexp',
  'list',
], function ($, _, locale, Handlebars, apiProject, apiData, prettyPrint, sampleRequest, semver, WebFont) {
  // load google web fonts
  loadGoogleFontCss();

  let { api } = apiData;

  //
  // Templates
  //
  const templateHeader = Handlebars.compile($('#template-header').html());
  const templateFooter = Handlebars.compile($('#template-footer').html());
  const templateArticle = Handlebars.compile($('#template-article').html());
  const templateCompareArticle = Handlebars.compile($('#template-compare-article').html());
  const templateGenerator = Handlebars.compile($('#template-generator').html());
  const templateProject = Handlebars.compile($('#template-project').html());
  const templateSections = Handlebars.compile($('#template-sections').html());
  const templateSidenav = Handlebars.compile($('#template-sidenav').html());

  //
  // apiProject defaults
  //
  if (!apiProject.template) apiProject.template = {};

  if (apiProject.template.withCompare == null) apiProject.template.withCompare = true;

  if (apiProject.template.withGenerator == null) apiProject.template.withGenerator = true;

  if (apiProject.template.forceLanguage) locale.setLanguage(apiProject.template.forceLanguage);

  // Setup jQuery Ajax
  $.ajaxSetup(apiProject.template.jQueryAjaxSetup);

  //
  // Data transform
  //
  // grouped by group
  const apiByGroup = _.groupBy(api, entry => entry.group);

  // grouped by group and name
  const apiByGroupAndName = {};
  $.each(apiByGroup, (index, entries) => {
    apiByGroupAndName[index] = _.groupBy(entries, entry => entry.name);
  });

  //
  // sort api within a group by title ASC and custom order
  //
  const newList = [];
  const umlauts = {
    ä: 'ae', ü: 'ue', ö: 'oe', ß: 'ss',
  }; // TODO: remove in version 1.0
  $.each(apiByGroupAndName, (index, groupEntries) => {
    // get titles from the first entry of group[].name[] (name has versioning)
    let titles = [];
    $.each(groupEntries, (titleName, entries) => {
      const { title } = entries[0];
      if (title !== undefined) {
        title.toLowerCase().replace(/[äöüß]/g, $0 => umlauts[$0]);
        titles.push(`${title}#~#${titleName}`); // '#~#' keep reference to titleName after sorting
      }
    });
    // sort by name ASC
    titles.sort();

    // custom order
    if (apiProject.order) titles = sortByOrder(titles, apiProject.order, '#~#');

    // add single elements to the new list
    titles.forEach((name) => {
      const values = name.split('#~#');
      const key = values[1];
      groupEntries[key].forEach((entry) => {
        newList.push(entry);
      });
    });
  });
  // api overwrite with ordered list
  api = newList;

  //
  // Group- and Versionlists
  //
  let apiGroups = {};
  const apiGroupTitles = {};
  let apiVersions = {};
  apiVersions[apiProject.version] = 1;

  $.each(api, (index, entry) => {
    apiGroups[entry.group] = 1;
    apiGroupTitles[entry.group] = entry.groupTitle || entry.group;
    apiVersions[entry.version] = 1;
  });

  // sort groups
  apiGroups = Object.keys(apiGroups);
  apiGroups.sort();

  // custom order
  if (apiProject.order) apiGroups = sortByOrder(apiGroups, apiProject.order);

  // sort versions DESC
  apiVersions = Object.keys(apiVersions);
  apiVersions.sort(semver.compare);
  apiVersions.reverse();

  //
  // create Navigationlist
  //
  const nav = [];
  apiGroups.forEach((group) => {
    // Mainmenu entry
    nav.push({
      group,
      isHeader: true,
      title: apiGroupTitles[group],
    });

    // Submenu
    let oldName = '';
    api.forEach((entry) => {
      if (entry.group === group) {
        if (oldName !== entry.name) {
          nav.push({
            title: entry.title,
            group,
            name: entry.name,
            type: entry.type,
            version: entry.version,
          });
        } else {
          nav.push({
            title: entry.title,
            group,
            hidden: true,
            name: entry.name,
            type: entry.type,
            version: entry.version,
          });
        }
        oldName = entry.name;
      }
    });
  });

  /**
     * Add navigation items by analyzing the HTML content and searching for h1 and h2 tags
     * @param nav Object the navigation array
     * @param content string the compiled HTML content
     * @param index where to insert items
     * @return boolean true if any good-looking (i.e. with a group identifier) <h1> tag was found
     */
  function add_nav(nav, content, index) {
    let found_level1 = false;
    if (!content) {
      return found_level1;
    }
    const topics = content.match(/<h(1|2).*?>(.+?)<\/h(1|2)>/gi);
    if (topics) {
      topics.forEach((entry) => {
        const level = entry.substring(2, 3);
        const title = entry.replace(/<.+?>/g, ''); // Remove all HTML tags for the title
        const entry_tags = entry.match(/id="api-([^\-]+)(?:-(.+))?"/); // Find the group and name in the id property
        const group = (entry_tags ? entry_tags[1] : null);
        const name = (entry_tags ? entry_tags[2] : null);
        if (level == 1 && title && group) {
          nav.splice(index, 0, {
            group,
            isHeader: true,
            title,
            isFixed: true,
          });
          index++;
          found_level1 = true;
        }
        if (level == 2 && title && group && name) {
          nav.splice(index, 0, {
            group,
            name,
            isHeader: false,
            title,
            isFixed: false,
            version: '1.0',
          });
          index++;
        }
      });
    }
    return found_level1;
  }

  // Mainmenu Header entry
  if (apiProject.header) {
    var found_level1 = add_nav(nav, apiProject.header.content, 0); // Add level 1 and 2 titles
    if (!found_level1) { // If no Level 1 tags were found, make a title
      nav.unshift({
        group: '_',
        isHeader: true,
        title: (apiProject.header.title == null) ? locale.__('General') : apiProject.header.title,
        isFixed: true,
      });
    }
  }

  // Mainmenu Footer entry
  if (apiProject.footer) {
    const last_nav_index = nav.length;
    var found_level1 = add_nav(nav, apiProject.footer.content, nav.length); // Add level 1 and 2 titles
    if (!found_level1 && apiProject.footer.title != null) { // If no Level 1 tags were found, make a title
      nav.splice(last_nav_index, 0, {
        group: '_footer',
        isHeader: true,
        title: apiProject.footer.title,
        isFixed: true,
      });
    }
  }

  // render pagetitle
  const title = apiProject.title ? apiProject.title : `apiDoc: ${apiProject.name} - ${apiProject.version}`;
  $(document).attr('title', title);

  // remove loader
  $('#loader').remove();

  // render sidenav
  const fields = {
    nav,
  };
  $('#sidenav').append(templateSidenav(fields));

  // render Generator
  $('#generator').append(templateGenerator(apiProject));

  // render Project
  _.extend(apiProject, { versions: apiVersions });
  $('#project').append(templateProject(apiProject));

  // render apiDoc, header/footer documentation
  if (apiProject.header) $('#header').append(templateHeader(apiProject.header));

  if (apiProject.footer) $('#footer').append(templateFooter(apiProject.footer));

  //
  // Render Sections and Articles
  //
  const articleVersions = {};
  let content = '';
  apiGroups.forEach((groupEntry) => {
    const articles = [];
    let oldName = '';
    var fields = {};
    let title = groupEntry;
    let description = '';
    articleVersions[groupEntry] = {};

    // render all articles of a group
    api.forEach((entry) => {
      if (groupEntry === entry.group) {
        if (oldName !== entry.name) {
          // determine versions
          api.forEach((versionEntry) => {
            if (groupEntry === versionEntry.group && entry.name === versionEntry.name) {
              if (!articleVersions[entry.group].hasOwnProperty(entry.name)) {
                articleVersions[entry.group][entry.name] = [];
              }
              articleVersions[entry.group][entry.name].push(versionEntry.version);
            }
          });
          fields = {
            article: entry,
            versions: articleVersions[entry.group][entry.name],
          };
        } else {
          fields = {
            article: entry,
            hidden: true,
            versions: articleVersions[entry.group][entry.name],
          };
        }

        // add prefix URL for endpoint
        if (apiProject.url) fields.article.url = apiProject.url + fields.article.url;

        addArticleSettings(fields, entry);

        if (entry.groupTitle) title = entry.groupTitle;

        // TODO: make groupDescription compareable with older versions (not important for the moment)
        if (entry.groupDescription) description = entry.groupDescription;

        articles.push({
          article: templateArticle(fields),
          group: entry.group,
          name: entry.name,
        });
        oldName = entry.name;
      }
    });

    // render Section with Articles
    var fields = {
      group: groupEntry,
      title,
      description,
      articles,
    };
    content += templateSections(fields);
  });
  $('#sections').append(content);

  // Bootstrap Scrollspy
  $(this).scrollspy({ target: '#scrollingNav', offset: 18 });

  // Content-Scroll on Navigation click.
  $('.sidenav').find('a').on('click', function (e) {
    e.preventDefault();
    const id = $(this).attr('href');
    if ($(id).length > 0) $('html,body').animate({ scrollTop: parseInt($(id).offset().top) }, 400);
    window.location.hash = $(this).attr('href');
  });

  // Quickjump on Pageload to hash position.
  if (window.location.hash) {
    var id = window.location.hash;
    if ($(id).length > 0) $('html,body').animate({ scrollTop: parseInt($(id).offset().top) }, 0);
  }

  /**
     * Check if Parameter (sub) List has a type Field.
     * Example: @apiSuccess          varname1 No type.
     *          @apiSuccess {String} varname2 With type.
     *
     * @param {Object} fields
     */
  function _hasTypeInFields(fields) {
    let result = false;
    $.each(fields, (name) => {
      result = result || _.some(fields[name], item => item.type);
    });
    return result;
  }

  /**
     * On Template changes, recall plugins.
     */
  function initDynamic() {
    // Bootstrap popover
    $('button[data-toggle="popover"]').popover().click((e) => {
      e.preventDefault();
    });

    const version = $('#version strong').html();
    $('#sidenav li').removeClass('is-new');
    if (apiProject.template.withCompare) {
      $(`#sidenav li[data-version='${version}']`).each(function () {
        const group = $(this).data('group');
        const name = $(this).data('name');
        const { length } = $(`#sidenav li[data-group='${group}'][data-name='${name}']`);
        const index = $(`#sidenav li[data-group='${group}'][data-name='${name}']`).index($(this));
        if (length === 1 || index === (length - 1)) $(this).addClass('is-new');
      });
    }

    // tabs
    $('.nav-tabs-examples a').click(function (e) {
      e.preventDefault();
      $(this).tab('show');
    });
    $('.nav-tabs-examples').find('a:first').tab('show');

    // sample request switch
    $('.sample-request-switch').click(function (e) {
      const name = `.${$(this).attr('name')}-fields`;
      $(name).addClass('hide');
      $(this).parent().next(name).removeClass('hide');
    });

    // call scrollspy refresh method
    $(window).scrollspy('refresh');

    // init modules
    sampleRequest.initDynamic();
  }
  initDynamic();

  // Pre- / Code-Format
  prettyPrint();

  //
  // HTML-Template specific jQuery-Functions
  //
  // Change Main Version
  $('#versions li.version a').on('click', function (e) {
    e.preventDefault();

    const selectedVersion = $(this).html();
    $('#version strong').html(selectedVersion);

    // hide all
    $('article').addClass('hide');
    $('#sidenav li:not(.nav-fixed)').addClass('hide');

    // show 1st equal or lower Version of each entry
    $('article[data-version]').each(function (index) {
      const group = $(this).data('group');
      const name = $(this).data('name');
      const version = $(this).data('version');

      if (semver.lte(version, selectedVersion)) {
        if ($(`article[data-group='${group}'][data-name='${name}']:visible`).length === 0) {
          // enable Article
          $(`article[data-group='${group}'][data-name='${name}'][data-version='${version}']`).removeClass('hide');
          // enable Navigation
          $(`#sidenav li[data-group='${group}'][data-name='${name}'][data-version='${version}']`).removeClass('hide');
          $(`#sidenav li.nav-header[data-group='${group}']`).removeClass('hide');
        }
      }
    });

    // show 1st equal or lower Version of each entry
    $('article[data-version]').each(function (index) {
      const group = $(this).data('group');
      $(`section#api-${group}`).removeClass('hide');
      if ($(`section#api-${group} article:visible`).length === 0) {
        $(`section#api-${group}`).addClass('hide');
      } else {
        $(`section#api-${group}`).removeClass('hide');
      }
    });

    initDynamic();
  });

  // compare all article with their predecessor
  $('#compareAllWithPredecessor').on('click', changeAllVersionCompareTo);

  // change version of an article
  $('article .versions li.version a').on('click', changeVersionCompareTo);

  // compare url-parameter
  $.urlParam = function (name) {
    const results = new RegExp(`[\\?&amp;]${name}=([^&amp;#]*)`).exec(window.location.href);
    return (results && results[1]) ? results[1] : null;
  };

  if ($.urlParam('compare')) {
    // URL Paramter ?compare=1 is set
    $('#compareAllWithPredecessor').trigger('click');

    if (window.location.hash) {
      var id = window.location.hash;
      $('html,body').animate({ scrollTop: parseInt($(id).offset().top) - 18 }, 0);
    }
  }

  /**
     * Initialize search
     */
  const options = {
    valueNames: ['nav-list-item'],
  };
  const endpointsList = new List('scrollingNav', options);

  /**
     * Set initial focus to search input
     */
  $('#scrollingNav .sidenav-search input.search').focus();

  /**
     * Detect ESC key to reset search
     */
  $(document).keyup((e) => {
    if (e.keyCode === 27) $('span.search-reset').click();
  });

  /**
     * Search reset
     */
  $('span.search-reset').on('click', () => {
    $('#scrollingNav .sidenav-search input.search')
      .val('')
      .focus();
    endpointsList.search();
  });

  /**
     * Change version of an article to compare it to an other version.
     */
  function changeVersionCompareTo(e) {
    e.preventDefault();

    const $root = $(this).parents('article');
    const selectedVersion = $(this).html();
    const $button = $root.find('.version');
    const currentVersion = $button.find('strong').html();
    $button.find('strong').html(selectedVersion);

    const group = $root.data('group');
    const name = $root.data('name');
    const version = $root.data('version');

    const compareVersion = $root.data('compare-version');

    if (compareVersion === selectedVersion) return;

    if (!compareVersion && version == selectedVersion) return;

    if (compareVersion && articleVersions[group][name][0] === selectedVersion || version === selectedVersion) {
      // the version of the entry is set to the highest version (reset)
      resetArticle(group, name, version);
    } else {
      const $compareToArticle = $(`article[data-group='${group}'][data-name='${name}'][data-version='${selectedVersion}']`);

      let sourceEntry = {};
      let compareEntry = {};
      $.each(apiByGroupAndName[group][name], (index, entry) => {
        if (entry.version === version) sourceEntry = entry;
        if (entry.version === selectedVersion) compareEntry = entry;
      });

      const fields = {
        article: sourceEntry,
        compare: compareEntry,
        versions: articleVersions[group][name],
      };

      // add unique id
      // TODO: replace all group-name-version in template with id.
      fields.article.id = `${fields.article.group}-${fields.article.name}-${fields.article.version}`;
      fields.article.id = fields.article.id.replace(/\./g, '_');

      fields.compare.id = `${fields.compare.group}-${fields.compare.name}-${fields.compare.version}`;
      fields.compare.id = fields.compare.id.replace(/\./g, '_');

      var entry = sourceEntry;
      if (entry.parameter && entry.parameter.fields) fields._hasTypeInParameterFields = _hasTypeInFields(entry.parameter.fields);

      if (entry.error && entry.error.fields) fields._hasTypeInErrorFields = _hasTypeInFields(entry.error.fields);

      if (entry.success && entry.success.fields) fields._hasTypeInSuccessFields = _hasTypeInFields(entry.success.fields);

      if (entry.info && entry.info.fields) fields._hasTypeInInfoFields = _hasTypeInFields(entry.info.fields);

      var entry = compareEntry;
      if (fields._hasTypeInParameterFields !== true && entry.parameter && entry.parameter.fields) fields._hasTypeInParameterFields = _hasTypeInFields(entry.parameter.fields);

      if (fields._hasTypeInErrorFields !== true && entry.error && entry.error.fields) fields._hasTypeInErrorFields = _hasTypeInFields(entry.error.fields);

      if (fields._hasTypeInSuccessFields !== true && entry.success && entry.success.fields) fields._hasTypeInSuccessFields = _hasTypeInFields(entry.success.fields);

      if (fields._hasTypeInInfoFields !== true && entry.info && entry.info.fields) fields._hasTypeInInfoFields = _hasTypeInFields(entry.info.fields);

      const content = templateCompareArticle(fields);
      $root.after(content);
      const $content = $root.next();

      // Event on.click re-assign
      $content.find('.versions li.version a').on('click', changeVersionCompareTo);

      // select navigation
      $(`#sidenav li[data-group='${group}'][data-name='${name}'][data-version='${currentVersion}']`).addClass('has-modifications');

      $root.remove();
      // TODO: on change main version or select the highest version re-render
    }

    initDynamic();
  }

  /**
     * Compare all currently selected Versions with their predecessor.
     */
  function changeAllVersionCompareTo(e) {
    e.preventDefault();
    $('article:visible .versions').each(function () {
      const $root = $(this).parents('article');
      const currentVersion = $root.data('version');
      let $foundElement = null;
      $(this).find('li.version a').each(function () {
        const selectVersion = $(this).html();
        if (selectVersion < currentVersion && !$foundElement) $foundElement = $(this);
      });

      if ($foundElement) $foundElement.trigger('click');
    });
    initDynamic();
  }

  /**
     * Sort the fields.
     */
  function sortFields(fields_object) {
    $.each(fields_object, (key, fields) => {
      const reversed = fields.slice().reverse();

      const max_dot_count = Math.max.apply(null, reversed.map(item => item.field.split('.').length - 1));

      for (var dot_count = 1; dot_count <= max_dot_count; dot_count++) {
        reversed.forEach((item, index) => {
          const parts = item.field.split('.');
          if (parts.length - 1 == dot_count) {
            const fields_names = fields.map(item => item.field);
            if (parts.slice(1).length >= 1) {
              const prefix = parts.slice(0, parts.length - 1).join('.');
              const prefix_index = fields_names.indexOf(prefix);
              if (prefix_index > -1) {
                fields.splice(fields_names.indexOf(item.field), 1);
                fields.splice(prefix_index + 1, 0, item);
              }
            }
          }
        });
      }
    });
  }

  /**
     * Add article settings.
     */
  function addArticleSettings(fields, entry) {
    // add unique id
    // TODO: replace all group-name-version in template with id.
    fields.id = `${fields.article.group}-${fields.article.name}-${fields.article.version}`;
    fields.id = fields.id.replace(/\./g, '_');

    if (entry.header && entry.header.fields) {
      sortFields(entry.header.fields);
      fields._hasTypeInHeaderFields = _hasTypeInFields(entry.header.fields);
    }

    if (entry.parameter && entry.parameter.fields) {
      sortFields(entry.parameter.fields);
      fields._hasTypeInParameterFields = _hasTypeInFields(entry.parameter.fields);
    }

    if (entry.error && entry.error.fields) {
      sortFields(entry.error.fields);
      fields._hasTypeInErrorFields = _hasTypeInFields(entry.error.fields);
    }

    if (entry.success && entry.success.fields) {
      sortFields(entry.success.fields);
      fields._hasTypeInSuccessFields = _hasTypeInFields(entry.success.fields);
    }

    if (entry.info && entry.info.fields) {
      sortFields(entry.info.fields);
      fields._hasTypeInInfoFields = _hasTypeInFields(entry.info.fields);
    }

    // add template settings
    fields.template = apiProject.template;
  }

  /**
     * Render Article.
     */
  function renderArticle(group, name, version) {
    let entry = {};
    $.each(apiByGroupAndName[group][name], (index, currentEntry) => {
      if (currentEntry.version === version) entry = currentEntry;
    });
    const fields = {
      article: entry,
      versions: articleVersions[group][name],
    };

    addArticleSettings(fields, entry);

    return templateArticle(fields);
  }

  /**
     * Render original Article and remove the current visible Article.
     */
  function resetArticle(group, name, version) {
    const $root = $(`article[data-group='${group}'][data-name='${name}']:visible`);
    const content = renderArticle(group, name, version);

    $root.after(content);
    const $content = $root.next();

    // Event on.click muss neu zugewiesen werden (sollte eigentlich mit on automatisch funktionieren... sollte)
    $content.find('.versions li.version a').on('click', changeVersionCompareTo);

    $(`#sidenav li[data-group='${group}'][data-name='${name}'][data-version='${version}']`).removeClass('has-modifications');

    $root.remove();
  }

  /**
     * Load google fonts.
     */
  function loadGoogleFontCss() {
    WebFont.load({
      active() {
        // Update scrollspy
        $(window).scrollspy('refresh');
      },
      google: {
        families: ['Source Code Pro', 'Source Sans Pro:n4,n6,n7'],
      },
    });
  }

  /**
     * Return ordered entries by custom order and append not defined entries to the end.
     * @param  {String[]} elements
     * @param  {String[]} order
     * @param  {String}   splitBy
     * @return {String[]} Custom ordered list.
     */
  function sortByOrder(elements, order, splitBy) {
    const results = [];
    order.forEach((name) => {
      if (splitBy) {
        elements.forEach((element) => {
          const parts = element.split(splitBy);
          const key = parts[1]; // reference keep for sorting
          if (key == name) results.push(element);
        });
      } else {
        elements.forEach((key) => {
          if (key == name) results.push(name);
        });
      }
    });
    // Append all other entries that ar not defined in order
    elements.forEach((element) => {
      if (results.indexOf(element) === -1) results.push(element);
    });
    return results;
  }
});
