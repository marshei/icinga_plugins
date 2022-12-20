# Contributing

There are many ways to contribute, you can fork the project and enhance the code base, submitting bug reports or feature
requests and you can help by testing open pull requests. Any help is much appreciated.

## Development

This Icinga plugin support module is written in Go. 

It is recommendable to use an editor with helpers for Go e.g. auto-completion. For instance [Visual Studio Code](https://code.visualstudio.com/) provides
such helper features with the [Go](https://code.visualstudio.com/docs/languages/go) extension. 


## Testing

The `Makefile` contains the target `test` which will run `go test` on the project directory.


The `Makefile` is also used in our CI using [GitHub Actions](https://github.com/mcktr/check_fritz/actions).