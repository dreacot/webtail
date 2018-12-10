# Webtail

### Browser based remote log viewer written in Golang

Webtail is a web-socket based server that streams log files onto your browser. Written in [Golang](https://golang.org), this application provides a basic and clean material UI to view the logs as well. This implementation is a simplified minimal version of the upstream code from https://github.com/prateeknischal/webtail .

### Usage

```
$ go run main.go --help
usage: main [<flags>] [<dir>...]

Flags:
      --help               Show context-sensitive help (also try --help-long and --help-man).
  -p, --port=8080          Port number to host the server

Args:
  [<dir>]  Directory path(s) to look for files

```

To view the UI, navigate to *http(s)://server_ip:port* and you will be presented with a UI to view the logs.
