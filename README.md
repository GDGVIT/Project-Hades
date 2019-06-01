# Project Hades

The last event coordinator you will ever need

[![CodeFactor](https://www.codefactor.io/repository/github/l04db4l4nc3r/project-hades/badge)](https://www.codefactor.io/repository/github/l04db4l4nc3r/project-hades) [![view docs](https://img.shields.io/badge/docs-view%20API%20documentation-orange.svg)](https://l04db4l4nc3r.github.io/Project-Hades/) [![view mobile app](https://img.shields.io/badge/app-view%20mobile%20app-blue.svg)](https://github.com/GDGVIT/Hades_Admin_App.git) [![view web frontend](https://img.shields.io/badge/frontend-view%20web%20frontend-yellow.svg)](https://github.com/GDGVIT/Project-Hades-Frontend.git)

<br />

## Directions to run
---

<br/>

### Quick setup
Quick setup uses performat dockerfiles made from scratch. It directly runs the compiled binaries.

To `compile` binaries run

```bash
$ make
$ docker-compose -f docker-compose-light.yml build
$ docker-compose -f docker-compose-light.yml up -d neo4j
$ docker-compose -f docker-compose-light.yml up --d
```
<br/>

### Stable setup
Stable setup compiles binaries on a heavier container runs them using `watcher`. 

To run the setup simply do the following 

```bash
$ docker-compose build 
$ docker-compose up -d neo4j
$ docker-compose up -d
```

Go to `localhost:7474` to see if the *database* is running. You would have to create a new username and password the first time

<br/>

### Testing and documentation
To generate docs and testing, simply do the following

```bash
$ make docs
$ make test
```

<br/>
<br/>

### WomenTechies'19 database
<img src="./static/images/neo_event.png" width=100%>


<br />
<br />

### High level flow

![functional architecture](./static/images/HADESxml.png)
