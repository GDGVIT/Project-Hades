define([
  'locales',
  'handlebars',
  'diffMatchPatch',
], (locale, Handlebars, DiffMatchPatch) => {
  /**
     * Return a text as markdown.
     * Currently only a little helper to replace apidoc-inline Links (#Group:Name).
     * Should be replaced with a full markdown lib.
     * @param string text
     */
  Handlebars.registerHelper('markdown', (text) => {
    if (!text) {
      return text;
    }
    text = text.replace(/((\[(.*?)\])?\(#)((.+?):(.+?))(\))/mg, (match, p1, p2, p3, p4, p5, p6) => {
      const link = p3 || `${p5}/${p6}`;
      return `<a href="#api-${p5}-${p6}">${link}</a>`;
    });
    return text;
  });

  /**
     * start/stop timer for simple performance check.
     */
  let timer;
  Handlebars.registerHelper('startTimer', (text) => {
    timer = new Date();
    return '';
  });

  Handlebars.registerHelper('stopTimer', (text) => {
    console.log(new Date() - timer);
    return '';
  });

  /**
     * Return localized Text.
     * @param string text
     */
  Handlebars.registerHelper('__', text => locale.__(text));

  /**
     * Console log.
     * @param mixed obj
     */
  Handlebars.registerHelper('cl', (obj) => {
    console.log(obj);
    return '';
  });

  /**
     * Replace underscore with space.
     * @param string text
     */
  Handlebars.registerHelper('underscoreToSpace', text => text.replace(/(_+)/g, ' '));

  /**
     *
     */
  Handlebars.registerHelper('assign', function (name) {
    if (arguments.length > 0) {
      const type = typeof (arguments[1]);
      let arg = null;
      if (type === 'string' || type === 'number' || type === 'boolean') arg = arguments[1];
      Handlebars.registerHelper(name, () => arg);
    }
    return '';
  });

  /**
     *
     */
  Handlebars.registerHelper('nl2br', text => _handlebarsNewlineToBreak(text));

  /**
     *
     */
  Handlebars.registerHelper('if_eq', function (context, options) {
    let compare = context;
    // Get length if context is an object
    if (context instanceof Object && !(options.hash.compare instanceof Object)) compare = Object.keys(context).length;

    if (compare === options.hash.compare) return options.fn(this);

    return options.inverse(this);
  });

  /**
     *
     */
  Handlebars.registerHelper('if_gt', function (context, options) {
    let compare = context;
    // Get length if context is an object
    if (context instanceof Object && !(options.hash.compare instanceof Object)) compare = Object.keys(context).length;

    if (compare > options.hash.compare) return options.fn(this);

    return options.inverse(this);
  });

  /**
     *
     */
  const templateCache = {};
  Handlebars.registerHelper('subTemplate', function (name, sourceContext) {
    if (!templateCache[name]) templateCache[name] = Handlebars.compile($(`#template-${name}`).html());

    const template = templateCache[name];
    const templateContext = $.extend({}, this, sourceContext.hash);
    return new Handlebars.SafeString(template(templateContext));
  });

  /**
     *
     */
  Handlebars.registerHelper('toLowerCase', value => ((value && typeof value === 'string') ? value.toLowerCase() : ''));

  /**
     *
     */
  Handlebars.registerHelper('splitFill', (value, splitChar, fillChar) => {
    const splits = value.split(splitChar);
    return new Array(splits.length).join(fillChar) + splits[splits.length - 1];
  });

  /**
     * Convert Newline to HTML-Break (nl2br).
     *
     * @param {String} text
     * @returns {String}
     */
  function _handlebarsNewlineToBreak(text) {
    return (`${text}`).replace(/([^>\r\n]?)(\r\n|\n\r|\r|\n)/g, '$1' + '<br>' + '$2');
  }

  /**
     *
     */
  Handlebars.registerHelper('each_compare_list_field', (source, compare, options) => {
    const fieldName = options.hash.field;
    const newSource = [];
    if (source) {
      source.forEach((entry) => {
        const values = entry;
        values.key = entry[fieldName];
        newSource.push(values);
      });
    }

    const newCompare = [];
    if (compare) {
      compare.forEach((entry) => {
        const values = entry;
        values.key = entry[fieldName];
        newCompare.push(values);
      });
    }
    return _handlebarsEachCompared('key', newSource, newCompare, options);
  });

  /**
     *
     */
  Handlebars.registerHelper('each_compare_keys', (source, compare, options) => {
    const newSource = [];
    if (source) {
      const sourceFields = Object.keys(source);
      sourceFields.forEach((name) => {
        const values = {};
        values.value = source[name];
        values.key = name;
        newSource.push(values);
      });
    }

    const newCompare = [];
    if (compare) {
      const compareFields = Object.keys(compare);
      compareFields.forEach((name) => {
        const values = {};
        values.value = compare[name];
        values.key = name;
        newCompare.push(values);
      });
    }
    return _handlebarsEachCompared('key', newSource, newCompare, options);
  });

  /**
     *
     */
  Handlebars.registerHelper('each_compare_field', (source, compare, options) => _handlebarsEachCompared('field', source, compare, options));

  /**
     *
     */
  Handlebars.registerHelper('each_compare_title', (source, compare, options) => _handlebarsEachCompared('title', source, compare, options));

  /**
     *
     */
  Handlebars.registerHelper('reformat', (source, type) => {
    if (type == 'json') {
      try {
        return JSON.stringify(JSON.parse(source.trim()), null, '    ');
      } catch (e) {

      }
    }
    return source;
  });

  /**
     *
     */
  Handlebars.registerHelper('showDiff', (source, compare, options) => {
    let ds = '';
    if (source === compare) {
      ds = source;
    } else {
      if (!source) return compare;

      if (!compare) return source;

      const d = diffMatchPatch.diff_main(compare, source);
      diffMatchPatch.diff_cleanupSemantic(d);
      ds = diffMatchPatch.diff_prettyHtml(d);
      ds = ds.replace(/&para;/gm, '');
    }
    if (options === 'nl2br') ds = _handlebarsNewlineToBreak(ds);

    return ds;
  });

  /**
     *
     */
  function _handlebarsEachCompared(fieldname, source, compare, options) {
    const dataList = [];
    var index = 0;
    if (source) {
      source.forEach((sourceEntry) => {
        let found = false;
        if (compare) {
          compare.forEach((compareEntry) => {
            if (sourceEntry[fieldname] === compareEntry[fieldname]) {
              const data = {
                typeSame: true,
                source: sourceEntry,
                compare: compareEntry,
                index,
              };
              dataList.push(data);
              found = true;
              index++;
            }
          });
        }
        if (!found) {
          const data = {
            typeIns: true,
            source: sourceEntry,
            index,
          };
          dataList.push(data);
          index++;
        }
      });
    }

    if (compare) {
      compare.forEach((compareEntry) => {
        let found = false;
        if (source) {
          source.forEach((sourceEntry) => {
            if (sourceEntry[fieldname] === compareEntry[fieldname]) found = true;
          });
        }
        if (!found) {
          const data = {
            typeDel: true,
            compare: compareEntry,
            index,
          };
          dataList.push(data);
          index++;
        }
      });
    }

    let ret = '';
    const { length } = dataList;
    for (var index in dataList) {
      if (index == (length - 1)) dataList[index]._last = true;
      ret += options.fn(dataList[index]);
    }
    return ret;
  }

  var diffMatchPatch = new DiffMatchPatch();

  /**
     * Overwrite Colors
     */
  DiffMatchPatch.prototype.diff_prettyHtml = function (diffs) {
    const html = [];
    const pattern_amp = /&/g;
    const pattern_lt = /</g;
    const pattern_gt = />/g;
    const pattern_para = /\n/g;
    for (let x = 0; x < diffs.length; x++) {
      const op = diffs[x][0]; // Operation (insert, delete, equal)
      const data = diffs[x][1]; // Text of change.
      const text = data.replace(pattern_amp, '&amp;').replace(pattern_lt, '&lt;')
        .replace(pattern_gt, '&gt;').replace(pattern_para, '&para;<br>');
      switch (op) {
        case DIFF_INSERT:
          html[x] = `<ins>${text}</ins>`;
          break;
        case DIFF_DELETE:
          html[x] = `<del>${text}</del>`;
          break;
        case DIFF_EQUAL:
          html[x] = `<span>${text}</span>`;
          break;
      }
    }
    return html.join('');
  };

  // Exports
  return Handlebars;
});
