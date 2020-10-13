# Bicycle store api

Bicycle store api returns a list of bicycle stores around a certain area.

## Installation

Use the package manager [pip](https://pip.pypa.io/en/stable/) to install foobar.

```bash
docker build -t storesservice .
```

## Usage

```bash
docker run --env API_KEY=<api_key> -p 8080:8080 -p 8081:8081 storesservice
```

or
```bash
go run cmd/storesservice/main.go -key=<api_key>
```

The api can now be used at http://127.0.0.1:8080/bicyclestores.

Api documentation is found at http://127.0.0.1:8080/docs

Prometheus metrics is found at http://127.0.0.1:8081/metrics


## License
Copyright Timo Truyts
