# Atlas Corp

## Running
To run this program, run this from the project root in the terminal:
go run main.go

## API
API runs on port 3000
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