# Bicycle store api

Bicycle store api returns a list of bicycle stores around a certain area.

## Installation


```bash
docker build -t storesservice .
```

## Usage

```bash
docker run --env API_KEY=<api_key> -p 8080:8080 storesservice
```

or
```bash
go run cmd/storesservice/main.go -key=<api_key>
```

The api can now be used at http://127.0.0.1:8080/bicyclestores.

Api documentation is found at http://127.0.0.1:8080/docs

Prometheus metrics is found at http://127.0.0.1:8080/metrics


## License
Copyright Timo Truyts
