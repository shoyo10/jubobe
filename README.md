# jubobe

A RESTful jubo backend API application

## Makefile

### Local run

First start PostgreSQL server

`$ make local-deploy-up`

Second migrate database

`$ make pgmigrate`

Third run API server

`$ make jubobe`

### Run unit test

*Befer running test, make sure you already do the first step in local run description*

`$ make unit-test`

## API Usage

After running API server, a quick way to play the APIs is to open the swagger UI: http://127.0.0.1:9090/swagger/index.html

## Project layout

* cmd: for CLI commands
* configs: config files
* docs: swagger documents
* internal: application specific features/funcitons
* internal/delivery: network application layer
* internal/model: the place for core business logic data models
* internal/repository: data storage layer
* internal/service: core business logic functions
* pkg: general purpose features/funcitons
* pkg/config: read config
* pkg/echorouter: initialize echo router, including series middlerwares, such as set request id to header, set request id to ctx, set logger with ctx, recover. Beside, setting a centralized error handle function to process api returned error and unify error control. Also, registering pprof and swagger UI to echo router.
* pkg/errors: define series of http errors
* pkg/postgres: initialize postgres connection
* pkg/zerolog: initialize zerolog
* build/docker: docker build application image file
* deployment: setting up some services needed for local development, such as PostgreSQL Server.

## Core concept

* internal/model is the place for core business data model, and we use the models to exchange data between each layer. 
* we can notice that delivery <-> service <-> repository, when they receive input or return data, all of them use the structures defined in internal/model
* and each layer they should communicate each other by interface

## Packages

Here are the mainly used packages to build the jubo backend server

* cobra: CLI application
* viper: read configuration
* echo: web framework
* zerolog: log
* fx: dependency injection
* postgres: data storge
* gorm: orm
* echo-swagger: generate swagger document & UI
* go-gormigrate: migrate DB table
