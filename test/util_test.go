package test

import (
	"fmt"
	"testing"

	wwr "github.com/qbeon/webwire-go"
	wwrclt "github.com/qbeon/webwire-go-client"
	"github.com/qbeon/webwire-go/transport/memchan"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type serverSetup struct {
	Transport wwr.Transport
	Server    wwr.Server
}

type serverSetupTest struct {
	t *testing.T
	serverSetup
}

// setupServer helps setting up and launching the server together with the
// underlying transport
func setupServer(
	impl *ServerImpl,
	opts wwr.ServerOptions,
	trans wwr.Transport,
) (serverSetup, error) {
	// Use default session manager if no specific one is defined
	if opts.SessionManager == nil {
		opts.SessionManager = newInMemSessManager()
	}

	// Use the transport layer implementation specified by the CLI arguments
	if trans == nil {
		// Use default configuration
		trans = &memchan.Transport{}
	}

	// Initialize webwire server
	server, err := wwr.NewServer(impl, opts, trans)
	if err != nil {
		return serverSetup{}, fmt.Errorf(
			"failed setting up server instance: %s",
			err,
		)
	}

	// Run server in a separate goroutine
	go func() {
		if err := server.Run(); err != nil {
			panic(fmt.Errorf("server failed: %s", err))
		}
	}()

	// Return reference to the server and the address its bound to
	return serverSetup{
		Server:    server,
		Transport: trans,
	}, nil
}

// setupTestServer creates a new server setup failing the test immediately if
// the anything went wrong
func setupTestServer(
	t *testing.T,
	impl *ServerImpl,
	opts wwr.ServerOptions,
	trans wwr.Transport,
) serverSetupTest {
	setup, err := setupServer(impl, opts, trans)
	require.NoError(t, err)
	return serverSetupTest{t, setup}
}

// newClient sets up a new test client instance
func (setup *serverSetup) newClient(
	options wwrclt.Options,
	transport wwr.ClientTransport,
	hooks clientHooks,
) (*client, error) {
	return newClient(
		options,
		transport,
		hooks,
	)
}

// newClient sets up a new test client instance
func (setup *serverSetupTest) newClient(
	options wwrclt.Options,
	hooks clientHooks,
) *client {
	clt, err := newClient(
		options,
		&memchan.ClientTransport{
			Server: setup.Transport.(*memchan.Transport),
		},
		hooks,
	)
	require.NoError(setup.t, err)
	return clt
}

// newClient sets up a new test client instance
func newClient(
	options wwrclt.Options,
	clientTransport wwr.ClientTransport,
	hooks clientHooks,
) (*client, error) {
	newClt := &client{Hooks: hooks}

	// Initialize connection
	conn, err := wwrclt.NewClient(newClt, options, clientTransport)
	if err != nil {
		return nil, fmt.Errorf("failed setting up client instance: %s", err)
	}

	newClt.Connection = conn

	return newClt, nil
}

// CompareSessions compares a webwire session
func CompareSessions(t *testing.T, expected, actual *wwr.Session) {
	if actual == nil && expected == nil {
		return
	}

	assert.NotNil(t, expected)
	assert.NotNil(t, actual)
	assert.Equal(t, expected.Key, actual.Key)
	assert.Equal(t, expected.Creation.Unix(), actual.Creation.Unix())
}
