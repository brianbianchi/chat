## TCP

> The Transmission Control Protocol (TCP) is one of the main protocols of the Internet protocol suite. It originated in the initial network implementation in which it complemented the Internet Protocol (IP). Therefore, the entire suite is commonly referred to as TCP/IP. TCP provides reliable, ordered, and error-checked delivery of a stream of octets (bytes) between applications running on hosts communicating via an IP network. Major internet applications such as the World Wide Web, email, remote administration, and file transfer rely on TCP, which is part of the Transport Layer of the TCP/IP suite. SSL/TLS often runs on top of TCP. -[wiki](https://en.wikipedia.org/wiki/Transmission_Control_Protocol)

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
