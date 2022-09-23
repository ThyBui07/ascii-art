# ascii-art
---
Solved during studying in Gritlab coding school on Åland, September 2022

---

## Task description and audit questions

https://github.com/01-edu/public/tree/master/subjects/ascii-art

## Usage

### `go run . [STRING_TO_BE_ARTED]`

### Example: `go run . "Hello World!"`

## Unit testing

### `go test ascii-art/drawing -v`

## Project structure


The `main.go` file contains only argument checking


All other functions for ASCII art generation are located in `drawing` package. There are also unit tests for this package.
