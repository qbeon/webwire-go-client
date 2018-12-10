<!-- HEADER -->
<h1 align="center">
	<br>
	<a href="https://github.com/qbeon/webwire-go"><img src="https://cdn.rawgit.com/qbeon/webwire-go/c7c2c74e/docs/img/webwire_logo.svg" alt="WebWire" width="256"></a>
	<br>
	<br>
	WebWire for <a href="https://golang.org/">Go</a>
	<br>
	<sub>An asynchronous duplex messaging library</sub>
</h1>
<p align="center">
	<a href="https://travis-ci.org/qbeon/webwire-go-client">
		<img src="https://travis-ci.org/qbeon/webwire-go-client.svg?branch=master" alt="Travis CI: build status">
	</a>
	<a href="https://coveralls.io/github/qbeon/webwire-go-client?branch=master">
		<img src="https://coveralls.io/repos/github/qbeon/webwire-go-client/badge.svg?branch=master" alt="Coveralls: Test Coverage">
	</a>
	<a href="https://goreportcard.com/report/github.com/qbeon/webwire-go-client">
		<img src="https://goreportcard.com/badge/github.com/qbeon/webwire-go-client" alt="GoReportCard">
	</a>
	<a href="https://codebeat.co/projects/github-com-qbeon-webwire-go-client-master">
		<img src="https://codebeat.co/badges/809181da-797c-4cdd-bb23-d0324935f3b0" alt="CodeBeat: Status">
	</a>
	<a href="https://codeclimate.com/github/qbeon/webwire-go-client/maintainability">
		<img src="https://api.codeclimate.com/v1/badges/243a45cacec7d850c64d/maintainability" alt="CodeClimate: Maintainability">
	</a>
	<br>
	<a href="https://opensource.org/licenses/MIT">
		<img src="https://img.shields.io/badge/License-MIT-green.svg" alt="Licence: MIT">
	</a>
	<a href="https://app.fossa.io/projects/git%2Bgithub.com%2Fqbeon%2Fwebwire-go-client?ref=badge_shield" alt="FOSSA Status">
		<img src="https://app.fossa.io/api/projects/git%2Bgithub.com%2Fqbeon%2Fwebwire-go-client.svg?type=shield"/>
	</a>
	<a href="https://godoc.org/github.com/qbeon/webwire-go-client">
		<img src="https://godoc.org/github.com/qbeon/webwire-go-client?status.svg" alt="GoDoc">
	</a>
</p>
<p align="center">
	<a href="https://opencollective.com/webwire">
		<img src="https://opencollective.com/webwire/tiers/backer.svg?avatarHeight=64" alt="OpenCollective">
	</a>
</p>
<br>

<!-- CONTENT -->
[webwire-go-client](https://github.com/qbeon/webwire-go-client) provides a client implementation for the open-source [webwire](https://github.com/qbeon/webwire-go) protocol.
<br>

#### Table of Contents
- [Installation](#installation)
  - [Dep](#dep)
  - [Go Get](#go-get)
- [Contribution](#contribution)
  - [Maintainers](#maintainers)
- [WebWire Binary Protocol](#webwire-binary-protocol)
- [Examples](#examples)
- [Features](#features)
  - [Request-Reply](#request-reply)
  - [Client-side Signals](#client-side-signals)
  - [Server-side Signals](#server-side-signals)
  - [Namespaces](#namespaces)
  - [Sessions](#sessions)
  - [Automatic Session Restoration](#automatic-session-restoration)
  - [Automatic Connection Maintenance](#automatic-connection-maintenance)
  - [Graceful Shutdown](#graceful-shutdown)
  - [Security](#security)


## Installation
Choose any stable release from [the available release tags](https://github.com/qbeon/webwire-go-client/releases) and copy the source code into your project's vendor directory: ```$YOURPROJECT/vendor/github.com/qbeon/webwire-go-client```.

### Dep
If you're using [dep](https://github.com/golang/dep), just use [dep ensure](https://golang.github.io/dep/docs/daily-dep.html#adding-a-new-dependency) to add a specific version of webwire-go-client including all its transitive dependencies to your project: ```dep ensure -add github.com/qbeon/webwire-go-client@v1.0.0```.

### Go Get
You can also use [go get](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies): ```go get github.com/qbeon/webwire-go-client``` but beware that this will fetch the latest commit of the [master branch](https://github.com/qbeon/webwire-go-client/commits/master) which is currently **not yet** considered a stable release branch. It's therefore recommended to use [dep](https://github.com/qbeon/webwire-go-client#dep) instead.

## Contribution
Contribution of any kind is always welcome and appreciated, check out our [Contribution Guidelines](https://github.com/qbeon/webwire-go/blob/master/CONTRIBUTING.md) for more information!

### Maintainers
| Maintainer | Role | Specialization |
| :--- | :--- | :--- |
| **[Roman Sharkov](https://github.com/romshark)** | Core Maintainer | Dev (Go, JavaScript) |
| **[Daniil Trishkin](https://github.com/FromZeus)** | CI Maintainer | DevOps |

## WebWire Binary Protocol
WebWire is built for speed and portability implementing an open source binary protocol ([see here](https://github.com/qbeon/webwire-go#webwire-binary-protocol) for further details).

## Examples
- **[Echo](https://github.com/qbeon/webwire-go-examples/tree/master/examples/echo)** - Demonstrates a simple request-reply implementation.

- **[Pub-Sub](https://github.com/qbeon/webwire-go-examples/tree/master/examples/pubsub)** - Demonstrantes a simple publisher-subscriber tolopology implementation.

- **[Chat Room](https://github.com/qbeon/webwire-go-examples/tree/master/examples/chatroom)** - Demonstrates advanced use of the library. The corresponding [JavaScript Chat Room Client](https://github.com/qbeon/webwire-js/tree/master/examples/chatroom-client-vue) is implemented with the [Vue.js framework](https://vuejs.org/).

## Features
### Request-Reply
Clients can initiate multiple simultaneous requests and receive replies asynchronously. Requests are multiplexed through the connection similar to HTTP2 pipelining.

```go
// Send a request to the server,
// this will block the goroutine until either a reply is received
// or the default timeout triggers (if there is one)
reply, err := client.Request(nil, nil, wwr.NewPayload(
  wwr.EncodingBinary,
  []byte("sudo rm -rf /"),
))
if err != nil {
  // Oh oh, request failed for some reason!
}
defer reply.Close()
reply.PayloadUtf8() // Here we go!
 ```

Requests will respect cancelable contexts and provided deadlines

```go
cancelableCtx, cancel := context.WithCancel(context.Background())
defer cancel()
timedCtx, cancelTimed := context.WithTimeout(cancelableCtx, 1*time.Second)
defer cancelTimed()

// Send a cancelable request to the server with a 1 second deadline
// will block the goroutine for 1 second at max
reply, err := client.Request(timedCtx, nil, wwr.Payload(
  wwr.EncodingUtf8,
  []byte("hurry up!"),
))
// Investigate errors manually...
switch err.(type) {
  case wwr.ErrCanceled:
    // Request was prematurely canceled by the sender
  case wwr.ErrDeadlineExceeded:
    // Request timed out, server didn't manage to reply
    // within the user-specified context deadline
  case wwr.TimeoutErr:
    // Request timed out, server didn't manage to reply
    // within the specified default request timeout duration
  case nil:
    // Replied successfully
    defer reply.Close()
}

// ... or check for a timeout error the easier way:
if err != nil {
  if wwr.IsTimeoutErr(err) {
    // Timed out due to deadline excess or default timeout
  } else {
    // Unexpected error
  }
}

reply // Just in time!
```

### Client-side Signals
Individual clients can send signals to the server. Signals are one-way messages guaranteed to arrive, though they're not guaranteed to be processed like requests are. In cases such as when the server is being shut down, incoming signals are ignored by the server and dropped while requests will acknowledge the failure.

```go
// Send signal to server
err := client.Signal(
  "eventA",
  wwr.NewPayload(
    wwr.EncodingUtf8,
    []byte("something"),
  ),
)
```

### Server-side Signals
The server also can send signals to individual connected clients.

```go
func OnRequest(
  _ context.Context,
  conn wwr.Connection,
  _ wwr.Message,
) (wwr.Payload, error) {
  // Send a signal to the client before replying to the request
  conn.Signal(
    nil, // No message name
    wwr.NewPayload(wwr.EncodingUtf8, []byte("example")),
  )

  // Reply nothing
  return nil, nil
}
```

### Namespaces
Different kinds of requests and signals can be differentiated using the builtin namespacing feature.

```go
func OnRequest(
  _ context.Context,
  _ wwr.Connection,
  msg wwr.Message,
) (wwr.Payload, error) {
  switch msg.Name() {
  case "auth":
    // Authentication request
    return wwr.NewPayload(
      wwr.EncodingUtf8,
      []byte("this is an auth request"),
    )
  case "query":
    // Query request
    return wwr.NewPayload(
      wwr.EncodingUtf8,
      []byte("this is a query request"),
    )
  }

  // Otherwise return nothing
  return nil, nil
}
```
```go
func OnSignal(
  _ context.Context,
  _ wwr.Connection,
  msg wwr.Message,
) {
  switch msg.Name() {
  case "event A":
    // handle event A
  case "event B":
    // handle event B
  }
}
```

### Sessions
Individual connections can get sessions assigned to identify them. The state of the session is automagically synchronized between the client and the server. WebWire doesn't enforce any kind of authentication technique though, it just provides a way to authenticate a connection. WebWire also doesn't enforce any kind of session storage, the user could implement a custom session manager implementing the WebWire `SessionManager` interface to use any kind of volatile or persistent session storage, be it a database or a simple in-memory map.

```go
func OnRequest(
  _ context.Context,
  conn wwr.Connection,
  msg wwr.Message,
) (wwr.Payload, error) {
  // Verify credentials
  if string(msg.Payload().Data()) != "secret:pass" {
    return nil, wwr.ReqErr {
      Code: "WRONG_CREDENTIALS",
      Message: "Incorrect username or password, try again"
    }
  }
  // Create session (will automatically synchronize to the client)
  err := conn.CreateSession(/*something that implements wwr.SessionInfo*/)
  if err != nil {
    return nil, fmt.Errorf("Couldn't create session for some reason")
  }

  // Complete request, reply nothing
  return nil, nil
}
```

WebWire provides a basic file-based session manager implementation out of the box used by default when no custom session manager is defined. The default session manager creates a file with a .wwrsess extension for each opened session in the configured directory (which, by default, is the directory of the executable). During the restoration of a session the file is looked up by name using the session key, read and unmarshalled recreating the session object.

### Automatic Session Restoration
The client will automatically try to restore the previously opened session during connection establishment when getting disconnected without explicitly closing the session before.

```go
// Will automatically restore the previous session if there was any
err := client.Connect()
```

The session can also be restored manually given its key assuming the server didn't yet delete it. Session restoration will fail and return an error if the provided key doesn't correspond to any active session on the server or else if there's another active session already assigned to this client.
```go
err := client.RestoreSession([]byte("yoursessionkeygoeshere"))
```

### Automatic Connection Maintenance
The WebWire client maintains the connection fully automatically to guarantee maximum connection uptime. It will automatically reconnect in the background whenever the connection is lost.

The only things to remember are:
- Client API methods such as `client.Request` and `client.RestoreSession` will timeout if the server is unavailable for the entire duration of the specified timeout and thus the client fails to reconnect.
- `client.Signal` will immediately return a `ErrDisconnected` error if there's no connection at the time the signal was sent.

This feature is entirely optional and can be disabled at will which will cause `client.Request` and `client.RestoreSession` to immediately return a `ErrDisconnected` error when there's no connection at the time the request is made.

### Graceful Shutdown
The server will finish processing all ongoing signals and requests before closing when asked to shut down.
```go
// Will block until all handlers have finished
server.Shutdown()
```
While the server is shutting down new connections are refused with `503 Service Unavailable` and incoming new requests from connected clients will be rejected with a special error: `RegErrSrvShutdown`. Any incoming signals from connected clients will be ignored during the shutdown.

Server-side client connections also support graceful shutdown, a connection will be closed when all work on it is done,
while incoming requests and signals are handled similarly to shutting down the server.
```go
// Will block until all work on this connection is done
connection.Close()
```

### Security
Webwire can be hosted by a [TLS](https://en.wikipedia.org/wiki/Transport_Layer_Security) protected HTTPS server to prevent [man-in-the-middle attacks](https://en.wikipedia.org/wiki/Man-in-the-middle_attack) as well as to verify the identity of the server. Setting up a TLS protected server is easy:
```go
// Setup a secure webwire server instance
server, err := wwr.NewServer(
	serverImplementation,
	wwr.ServerOptions{
    Host: "localhost:443",
    // Use a TLS protected transport layer
    Transport: &wwrfasthttp.Transport{
			TLS: &wwrfasthttp.TLS{
        // Provide key and certificate
				CertFilePath:       "path/to/certificate.crt",
        PrivateKeyFilePath: "path/to/private.key",
        // Specify TLS configs
        Config: &tls.Config{
          MinVersion:               tls.VersionTLS12,
		      CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
		      PreferServerCipherSuites: true,
		      CipherSuites: []uint16{
			      tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			      tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			      tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		      },
        }
			},
		},
	},
)
if err != nil {
	panic(fmt.Errorf("failed setting up wwr server: %s", err))
}
// Launch
if err := server.Run(); err != nil {
	panic(fmt.Errorf("wwr server failed: %s", err))
}
```

To connect the client to a TLS protected webwire server `"https"` must be used as the URL scheme:
```go
connection, err := wwrclt.NewClient(
	url.URL{
		Scheme: "https",
		Host: "localhost:443",
	},
	clientImplementation,
	wwrclt.Options{/*...*/},
	nil, // Use default TLS configuration
)
```

In case of a self-signed certificate used for testing purposes the client will fail to connect but TLS can be configured to skip the certificate verification (which **must be disabled in production!**):
```go
connection, err := wwrclt.NewClient(
	url.URL{
		Scheme: "https",
		Host: "localhost:443",
	},
	clientImplementation,
	wwrclt.Options{/*...*/},
	/*
		--------------------------------------------------------------
		WARNING! NEVER DISABLE CERTIFICATE VERIFICATION IN PRODUCTION!
		--------------------------------------------------------------
		InsecureSkipVerify is enabled for testing purposes only
		to allow the use of a self-signed localhost SSL certificate.
		Enabling this option in production is dangerous and irresponsible.
	*/
	&tls.Config{
		InsecureSkipVerify: true,
	},
)
```

An alternative, somewhat safer approach would be to install the root CA certificate on the test system to make clients accept the self-signed server ceretificate (which was signed using the installed root certificate) instead of enabling `InsecureSkipVerify`.

----

© 2018 Roman Sharkov <roman.sharkov@qbeon.com>
