package test

import (
	"testing"
	"time"

	wwr "github.com/qbeon/webwire-go"
	wwrclt "github.com/qbeon/webwire-go-client"
	"github.com/stretchr/testify/require"
)

// TestStatus tests the Client.Status method
func TestStatus(t *testing.T) {
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
			Autoconnect:           wwr.Disabled,
		},
		clientHooks{},
	)

	require.NotEqual(t,
		wwrclt.StatusConnected, client.Connection.Status(),
		"Expected client to be disconnected "+
			"before the connection establishment",
	)

	// Connect to the server
	require.NoError(t, client.Connection.Connect())

	require.Equal(t,
		wwrclt.StatusConnected, client.Connection.Status(),
		"Expected client to be connected after the connection establishment",
	)

	// Disconnect the client
	client.Connection.Close()

	require.NotEqual(t,
		wwrclt.StatusConnected, client.Connection.Status(),
		"Expected client to be disconnected after closure",
	)
}
