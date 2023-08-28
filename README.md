# Fibonacci API in Go

This is a simple Go API that generates Fibonacci sequences in a spiral matrix.

## Installation

Make sure you have Go installed on your system. Then you can clone this repository and build the application:

### follow this command to run API:
```
git clone https://github.com/yourusername/fibonacci-spiral-matrix-api.git
```
```
cd fibonacci-spiral-matrix-api
```
```
go build
```
```
./fibonacci-spiral-matrix-api
```
The API will be available at http://localhost:8080. You can use tools like curl or an HTTP client to interact with it.

### Consume API 
[github.com/nictes1/fibonacci-spiral-frontend] (https://github.com/nictes1/fibonacci-spiral-frontend).



## Endpoints

### Generate a Fibonacci Spiral Matrix

URL: /spiral
HTTP Method: POST
Query Parameters:
rows: The number of rows in the matrix (integer).
cols: The number of columns in the matrix (integer).
Example Request:

curl -X POST "http://localhost:8080/spiral?rows=5&cols=5"
### Example Response:
```
{
  "ts": "1623959127000",
  "rows": [
    [0, 1, 1, 2, 3],
    [610, 987, 1597, 2584, 5],
    [377, 28657, 46368, 4181, 8],
    [233, 17711, 10946, 6765, 13],
    [144, 89, 55, 34, 21]
  ]
}
```



