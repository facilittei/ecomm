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

To generate a Mock from an interface go to the interface folder and run:

E.g.

```
cd internal/commmunications/http
```

```
mockery --name=HttpClient --output ./../../mocks/ --filename http_client_mock_test.go --structname HttpClientMock
```

This command will generate a mock for our tests.

Usage:

```
httpClient := &mocks.HttpClientMock{}
httpClient.On("Post").Return(nil, nil)
```