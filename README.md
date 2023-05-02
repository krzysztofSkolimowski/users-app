# users-app

## Instructions how to start the application

### Prerequisites

Regardless of the machine you are using, you will need to have the following tools installed:

- docker
- docker-compose
- make

In order to generate gRPC code and HTTP code, you will need to have the following tools installed:

- protoc
- oapi-codegen (https://github.com/deepmap/oapi-codegen)

In order to run tests, you will need to have the following tools installed:

- gotestsum (https://github.com/gotestyourself/gotestsum) - altough it is optional, it is needed to run make test
  command, but `go test ./...` will work as well

You will need any grpc, http client of your choice, but I recommend using evans and importing swagger spec to postman.
The same goes with redis and postgres clients. I was using pgcli for postgres and my small program for redis (it's
located in a `redis-client/main.go`)

In order to interact with redis I suggest just using:
`docker compose run redis redis-cli -h redis `

### Running the application

application uses `.env` file. Docker compose will automatically load it. You can find example file in `.env.example`.

The applications lets configure if desired interface is grpc, http or both. By default it is both. You can change it by
setting `RUN_HTTP` and `RUN_GRPC` variables in `.env` file.

In order to start the application, you need to run `make up` command. It will build the application and start it.

In order to check if the application is up, the easiest way is to query health endpoint:
`curl -v http://localhost:8080/health`

`make down` will stop the application and remove containers.

Application produces 2 special logs:

- events.log - it contains all events
- postgresql.log - it contains all db logs

## Explanations of technical choices and assumptions

### Chosen architecture

- Even though the application is small, I decided to follow hexagonal architecture and DDD. It is of course a little bit
  overkill for and app with just 4 endpoints, but I wanted to display that I am familiar with those concepts.
- There are few places which break it a little bit (like uuid library imported in the domain), mostly because of the
  time constraints.
- I also aimed for presenting finished application, which might be close to production ready, so, there might be some
  small shortcuts

Layers description:

- domain - This layer is responsible for representing the domain concepts, entities, and logic of the system.
- application - This layer coordinates the application activity.
- adapters - This layer is responsible for interacting with infrastructure
- ports - This layer is responsible for converting data from the external world to the internal format and receiving
  requests

When I've read the assignment, I've decide to follow CQRS pattern and split the application into 2 parts:

- queries - responsible for reading data
- commands - responsible for writing data

In order to fulfill the requirement of notifying other services about changes, I've decided to use events. This is
probably most elegant way of doing it.

This is much simpler, when all the mutating operations are separated from the reading ones. Also, the typical readmodel
and write model for user are different, so it was a natural choice.

## testing approach

Usually my goto testing approach is to write four layers of tests:

- [x] unit tests - testing go package
- [x] integration test - testing integration between a go package (e.g. repository) and outside service (e.g. postgres)
- [ ] component level tests - end to end tests for a single component (e.g. this users app)
- [ ] bdd test - testing the app in real environment from users perspective

In this case, with limited time constraint I've decided to skip component level tests and bdd tests.
The most important part of the application is the domain, so I've decided to focus on unit tests.

For the ease of development I've also decided to go for integration tests for postgres adapter.
All the tests in `adapters` layer work against running postgres.

So in order to run tests, you need to have running postgres instance. You can use the one that gets started
by `docker compose up`. Remember to set `DB_TEST_HOST` variable in `.env` file.

In order to run tests, you need to run `make test` command.

## Possible extensions or improvements to the service

- [ ] At the moment, the serialization of events does not exist at all. It should be added, so that the events can be
  stored in the database and replayed in case of failure. I would suggest just marshalling events to proto
- [ ] Publishing events at the moment has no retry mechanism. It should be added, so that the events are not lost in
  case of failure - also events should be buffered and published by a separate process
- [ ] Add missing layers of tests
- [ ] The logging is very basic, it should be improved - It should add the proper configurable, structured logging.
  Also, loggers should log corresponding request id (which gets added to the context) with every log entry.
- [ ] The error handling is also very basic, it should be improved, there is no good translation between domain errors
  and http errors - most of the time any error is translated to 500. Also, there is no good way of adding additional
  context

### Directories

- [api](api/) OpenAPI and gRPC definitions
- [docker](docker/) Dockerfiles
- [fixtures](fixtures/) It contains few example users, useful for testing
- [internal](internal/) application code
- [logs](logs/) logs directory
- [redis-client](redis-client/) redis client
- [sql](sql/) sql migrations
