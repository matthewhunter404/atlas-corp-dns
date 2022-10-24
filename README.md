# Atlas Corp

## Running
To run this program, run this from the project root in the terminal:
```console
go run main.go
```

To run the project in a docker container:
```console
docker-compose up -d
```

## API
API runs on port 3000 by default
### POST /location

Request:
{
    "x": "123.12",
    "y": "456.56",
    "z": "789.89",
    "vel": "20.0",
}

Response:
{
    "loc": 1389.57
}