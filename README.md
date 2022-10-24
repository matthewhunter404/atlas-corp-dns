# Atlas Corp

## Running
To run this program, run this from the project root in the terminal:
```console
$ go run main.go
```

To run the project in a docker container:
```console
$ docker-compose up -d
```

Config can be set via the following environment variables:

```console
atlas_dns_sector_id
atlas_dns_port
```

If not set the program will revert to using default values of SectorID = 1 and Port = 3000

## API
API runs on port 3000 by default
### POST /location

Request:
```json
{
    "x": "123.12",
    "y": "456.56",
    "z": "789.89",
    "vel": "20.0",
}
```

Response:
```json
{
    "loc": 1389.57
}
```