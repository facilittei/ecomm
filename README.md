# Ecomm

Responsible for managing payments with third-party providers.

## Getting started

Commands used to build and run the application can be found in [Makefile](./Makefile).

E.g. the application can be started by running `make compose/up`. You can check by visiting port 4000 on your browser.

## Development

### Monitoring

Application monitoring is done with [Prometheus](https://prometheus.io/).

[http://localhost:9090](http://localhost:9090)

The endpoint that Prometheus scrapes is [http://localhost:4000/metrics](http://localhost:4000/metrics)

### Visualization

Application metrics can be visualized with [Grafana](https://grafana.com/).

[http://localhost:3000](http://localhost:3000)

Default login: admin/admin

### Testing

We use [Testify](https://github.com/stretchr/testify) as our testing library in conjunction with the standard library
from Go.

This library also provides Mock capabilities and in order to auto-generate interfaces we can
use [Mockery](https://github.com/vektra/mockery)

You can create all mocks by running `go generate ./...` and the files that has the annotation below, they'll be
automatically generated.

```
//go:generate mockery --name=HttpClient --output ./../../mocks/ --filename http_client.go --structname HttpClientMock
```

If you want to generate a specific Mock from an interface go to the interface folder and run:

E.g.

```bash
cd internal/commmunications/http
```

```bash
mockery --name=HttpClient --output ./../../mocks/ --filename http_client.go --structname HttpClientMock
```

This command will generate a mock for our tests.

Usage:

```
httpClient := &mocks.HttpClientMock{}
httpClient.On("Post").Return(nil, nil)
```

## Migrations

Our database version control is done with [go-migrate](https://github.com/golang-migrate/migrate).

There are a couple of ways of installing it, and homebrew is one of them:

```bash
brew install golang-migrate
```

Every change to the database schema must be done through the usage of this tool.

Usage:

```bash
migrate create -seq -ext=.sql -dir=./migrations create_charges_table
```

Applying migration:

```bash
export DSN='postgres://facilittei:4321@localhost/facilittei?sslmode=disable'
migrate -path ./migrations -database ${DSN} up
```