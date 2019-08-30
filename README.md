# nosh
Run programs without a shell.

Windows often causes a shell to pop up when running certain programs. `nosh`
simply suppresses this window.

## Installation

1. [Install Go](https://golang.org/doc/install)
2. [Install Git](https://git-scm.com/downloads)
3. Using a shell with Git (such as Git Bash), run the following command:

```
go get -ldflags="-H windowsgui" github.com/anaminus/nosh
```

If you configured Go correctly, this will install `nosh` to `$GOPATH/bin`, which
will allow you run it directly from a shell.

## Usage

```shell
nosh COMMAND ARGS...
```
