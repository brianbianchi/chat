## Usage

```console
$ go run main.go
```

```console
$ telnet 127.0.0.1 8080 #connect as "Client 1"
$ Hello
$ Client 2: Hi there
$ Lorem ipsum dolor sit amet
$ Client 2: Ut enim ad minim veniam
```

```console
$ telnet 127.0.0.1 8080 #connect as "Client 2"
$ Client 1: Hello
$ Hi there
$ Client 1: Lorem ipsum dolor sit amet
$ Ut enim ad minim veniam
```
