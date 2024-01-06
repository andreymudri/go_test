# Golang Challenge

Simple Golang challenge which needs 2 endpoints:
One endpoint sending 2 arrays and another to merge in order those two arrays

Also using a ListNode struct.

Project port as of right now is hard-coded on 8080(first commit or so).

## Local Deployment

Clone the project

```bash
  git clone https://github.com/andreymudri/go_test
```

Go to the project directory

```bash
  cd go-test
```

To deploy this project run:
open your terminal inside the project folder

```bash
  go run main.go
```

or you can use Docker.

```bash
  docker compose up --build
```

##Running Tests

To run tests, run the following command

```bash
go test -v ./...
```

````


## API Usage

#### POST Lists

#### POST /saveLists
example body:
```bash
{
  "List1":[11,2,5,4],
  "List2":[666,999,111,1]
}
````

#####GET /merge

Should returns a sorted array provided that you given 2 arrays in the previous endpoint
