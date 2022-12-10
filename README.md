# aoc2022

[Advent of Code 2022](https://adventofcode.com/2022/) Solutions 🎄

Solutions in [Golang](https://go.dev) 1.19

## Run Solution

```sh
go run cmd/day<x>/main.go 
```

Solutions are run using inputs stored in `./inputs/` folder. 
- `/inputs` folder has been .gitignored to prevent github scraping of inputs.
- Add input to folder in the format `/inputs/day<x>/part<y>.txt`

## Run Tests

Run a specific test
```sh
go test  ./internal/day<x>/day<x>_test.go
```


Run all tests
```sh
go test  ./internal/... 
```


