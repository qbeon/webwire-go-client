package test

import (
	"context"
	"testing"
	"time"

	wwr "github.com/qbeon/webwire-go"
	wwrclt "github.com/qbeon/webwire-go-client"
	"github.com/stretchr/testify/require"
)

// TestSignalAutoconn tests sending signals on disconnected clients expecting it
// to automatically establish a connection
func TestSignalAutoconn(t *testing.T) {
	// Initialize webwire server given only the request
	setup := setupTestServer(
		t,
		&ServerImpl{},
		wwr.ServerOptions{},
		nil, // Use the default transport implementation
	)

	// Initialize client
	client := setup.newClient(
		wwrclt.Options{
			DefaultRequestTimeout: 2 * time.Second,
			Autoconnect:           wwr.Enabled,
		},
		nil, // Use the default transport implementation
		clientHooks{},
	)

	// Skip manual connection establishment and rely on autoconnect instead
	require.NoError(t, client.Connection.Signal(
		context.Background(),
		nil,
		wwr.Payload{Data: []byte("testdata")},
	))
}
