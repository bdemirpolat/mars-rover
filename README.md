# Code Review: Mars Rover

[![Go](https://github.com/bdemirpolat/mars-rover/actions/workflows/go-actions.yml/badge.svg?branch=main)](https://github.com/bdemirpolat/mars-rover/actions/workflows/go-actions.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/bdemirpolat/mars-rover)](https://goreportcard.com/report/github.com/bdemirpolat/mars-rover)

### Warn
Two rover's can position same place while it's moving, its easy to avoid but i did not implement.

## Example
#### Input
```
5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM
```
#### Output<br>
Result of First Rover: "1 3 N"<br>
Result of Second Rover: "5 1 E"<br><br>

### Run

```
go run .
```

### Test

```
go test ./...
```

### Benchmark

```
go test -bench=.
```

### Docker
```
docker build . -t mars-rover
```
```
docker run mars-rover
```

# License

This repository released under the Apache 2.0 license. See [LICENSE.txt](https://github.com/bdemirpolat/mars-rover/blob/main/LICENSE)
