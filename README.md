[![Build Status](https://img.shields.io/travis/walle/fval.svg?style=flat)](https://travis-ci.org/walle/fval)
[![Coverage](https://img.shields.io/codecov/c/github/walle/fval.svg?style=flat)](https://codecov.io/github/walle/fval)
[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/walle/fval)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/walle/fval/master/LICENSE)
[![Go Report Card](http://goreportcard.com/badge/walle/fval?t=3)](http:/goreportcard.com/report/walle/fval)

# fval

Simple package for validating file and directory existence, eg. input to CLI applications.
Originally built for [go-arg](https://github.com/alexflint/go-arg) - [Pull
request](https://github.com/alexflint/go-arg/pull/17). The pull request was
not accepted as go-arg is to be kept minimal, so this is extracted to a
separate package.
The usage is of course not limited to CLI application input, it is general
purpose, but was built with CLI input in mind.

## Installation

```shell
$ go get github.com/walle/fval
```

## Usage
```go
// CLI application with usage: example INPUTFILE OUTPUTDIR
if !fval.FileExists(os.Args[1]) {
        fmt.Fprintf(os.Stderr, "Usage error: %s is not a valid file", os.Args[1])
        os.Exit(1)
}

fval.DirExistsOrCreate(os.Args[2], 0766)
```

## Testing

Use the `go test` tool.

```shell
$ go test -cover
```

## Contributing

All contributions are welcome! See [CONTRIBUTING](CONTRIBUTING.md) for more
info.

## License

The code is under the MIT license. See [LICENSE](LICENSE) for more
information.
