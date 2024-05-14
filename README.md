# assessment

Requires Go 1.22 or higher.

Please follow the [install instructions](https://golang.org/doc/install) for relevant version.

## Build & Test

### Build

```bash
make
```

### Test

```bash
make test
```

## Run locally

### Requirements

- [Docker](https://docs.docker.com/engine/install/)
- [Helm](https://helm.sh/docs/intro/install/)
- A local kubernetes cluster if you intend to deploy locally. [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation) could be used for this.

### Run binary

```bash
make run
```

Access the service from `http://localhost:8080`

### Configuration

The service is primarily configured through a set of environment variables.

| Environment Variable    | Description                                                                             | Required | Default |
| ----------------------- | --------------------------------------------------------------------------------------- | -------- | ------- |
| SERVER_PORT             | This configures the port the service listens on.                                        | `false`  | `8080`  |
| SCRAPER_REQUEST_TIMEOUT | This configures the timeout setting for http requests made by the scraper               | `false`  | `30s`   |
| SERVER_IDLE_TIMEOUT     | [IdleTimeout](https://pkg.go.dev/net/http#Server.IdleTimeout) setting for http server   | `false`  | `30s`   |
| SERVER_READ_TIMEOUT     | [ReadTimeout](https://pkg.go.dev/net/http#Server.ReadTimeout) setting for http server   | `false`  | `15s`   |
| SERVER_WRITE_TIMEOUT    | [WriteTimeout](https://pkg.go.dev/net/http#Server.WriteTimeout) setting for http server | `false`  | `15s`   |

## CI/CD

The app is built, tested and released using GitHub Actions workflows which can be found in the [.github/workflows](.github/workflows) folder.
