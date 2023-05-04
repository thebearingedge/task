# task

A random name/joke API in Go.

## Endpoint

### `/v1/random-joke`

Retrieves a random `firstName` and `lastName` from `NAMES_SERVICE_BASE_URL` and feeds them into a request to `JOKES_SERVICE_BASE_URL`. The endpoint's response body is the text of the resulting joke, using the `firstName` and `lastName`.

#### Example

```shell
## GET /v1/random-joke
## Hasina Tanweer’s OSI network model has only one layer - Physical.
```

## Environment Variables

See `.env.example`. Make a copy to use them in `make` recipes.

```shell
cat .env.example

# #!/bin/sh
#
# #shellcheck disable=SC2034
#
# LISTEN_ADDRESS=:8080
# NAMES_SERVICE_BASE_URL=https://names.mcquay.me/api/v0/
# JOKES_SERVICE_BASE_URL=http://joke.loc8u.com:8888/joke/
```

#### `LISTEN_ADDRESS`

Listen address to bind the API server to. [See `engine.Run()` in the `gin-gonic` docs](https://pkg.go.dev/github.com/gin-gonic/gin#Engine.Run).

#### `NAMES_SERVICE_BASE_URL`

The remote address of the random name API.

#### `JOKES_SERVICE_BASE_URL`

The remote address of the random joke API. The `firstName` and `lastName` received from `NAMES_SERVICE_BASE_URL` will be added as query string parameters at runtime.

## Development

### Cloning

Clone the repository from [https://github.com/thebearingedge/task](https://github.com/thebearingedge/task) and copy the `.env.example` file to `.env`.

```shell
git clone https://github.com/thebearingedge/task thebearingedge-task
cd thebearingedge-task
cp .env.example .env
```

### Testing

Run unit tests with **`make test`**.

Generate a test coverage report with **`make cover`**.

### Building

Build the application into `.bin/task` with **`make build`**. **NOTE:**, running the binary directly requires exported environment variables in the current shell.

Build a container image tagged as `thebearingedge/task` with **`make image`**. **NOTE** [running a container](https://docs.docker.com/engine/reference/commandline/run/) requires environment variables and a published port.

### Running

**After copying `.env.example` to `.env`**, you should be able to run the application locally with **`make run`**. Then send a request to the server.

```shell
curl localhost:8080/v1/random-joke
# Boubeker Kunst can dereference NULL.
```

## Room for Improvement

- More complete error handling in external service (names, jokes) gateways.
- Better leveraging of `gin-gonic` capabilities, e.g. logging.
- More robust / centralized configuration management.
- More consistent / idiomatic package organization.
- Some kind of retry, circuit breaking, health monitoring.
- 3rd-party dependency auditing and automated security patching.
- CI automation.
- Container image versioning and publishing automation.
