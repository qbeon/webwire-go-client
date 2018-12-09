package test

import (
	"context"
	"testing"
	"time"

	wwr "github.com/qbeon/webwire-go"
	wwrclt "github.com/qbeon/webwire-go-client"
	"github.com/stretchr/testify/require"
)

// TestSignalDisconnNoAutoconn tests client.Signal expecting it to return a
// DisconnectedErr when autoconn is disabled and the client is disconnected
func TestSignalDisconnNoAutoconn(t *testing.T) {
	// Initialize webwire server
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
		nil, // Use the default transport implementation
		clientHooks{},
	)

	// Try to send a signal and expect a DisconnectedErr error
	err := client.Connection.Signal(
		context.Background(),
		nil,
		wwr.Payload{Data: []byte("test")},
	)
	require.Error(t, err)
	require.IsType(t, wwr.DisconnectedErr{}, err)
}
