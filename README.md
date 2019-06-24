<p align="center">
    <img src="https://user-images.githubusercontent.com/30529572/60032255-cd613600-96c3-11e9-891f-c2e69c3ce96e.png" width=150px />
</p>
<h1 align = "center">Project Hades</h1>

The last event coordinator you will ever need

[![CodeFactor](https://www.codefactor.io/repository/github/l04db4l4nc3r/project-hades/badge)](https://www.codefactor.io/repository/github/l04db4l4nc3r/project-hades) [![view docs](https://img.shields.io/badge/docs-view%20API%20documentation-orange.svg)](https://l04db4l4nc3r.github.io/Project-Hades/) [![view mobile app](https://img.shields.io/badge/app-view%20mobile%20app-blue.svg)](https://github.com/GDGVIT/Hades_Admin_App.git) [![view web frontend](https://img.shields.io/badge/frontend-view%20web%20frontend-yellow.svg)](https://github.com/GDGVIT/Project-Hades-Frontend.git) [![Mailer service](https://img.shields.io/badge/service-view%20mailer%20service-green.svg)](https://github.com/GDGVIT/Mailer-Microservice.git) [![Permission generator](https://img.shields.io/badge/service-view%20permission--generator%20service-red.svg)](https://github.com/GDGVIT/Permission-Generator-Microservice.git)

<br />

## Directions to run
Get started with setting up **Project Hades** locally. This project uses `GO111MODULE`. To install dependencies run a `go mod vendor` after cloning.

<br/>

### Environment variables
These variables should reside as key value pairs in a file called `.env`.

| Variable Name | Description | Data type |
|:-------------:|:-----------:|:------:|
| SALT | Salt for hashing | Integer |
| JWT_PASSWORD | Password for JWT | String |
| PROD_URI | Neo4j URI | String |
| MAIL_FROM | Mailer email for analytics | String |
| MAIL_TO | Destination email for analytics| String | 
| MAIL_PASSWORD | Password for MAIL_FROM | String |

<br/>


### Production setup
Production setup uses performat dockerfiles made from scratch. It directly runs the compiled binaries.

To `build` binaries run

```bash
$ make # needs go installed
$ docker-compose -f performant-compose.yaml build
$ docker-compose -f performant-compose.yaml up -d neo4j
$ docker-compose -f performant-compose.yaml up -d
```

### Dev setup
Dev setup compiles binaries on a heavier container runs them using `watcher`. 

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
