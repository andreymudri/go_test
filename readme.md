
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




## Running Tests

To run tests, run the following command

```bash
go test -v ./...
```

```
