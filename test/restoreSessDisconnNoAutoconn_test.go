package test

import (
	"context"
	"testing"
	"time"

	wwr "github.com/qbeon/webwire-go"
	wwrclt "github.com/qbeon/webwire-go-client"
	"github.com/qbeon/webwire-go/transport/memchan"
	"github.com/stretchr/testify/require"
)

// TestRestoreSessDisconnNoAutoconn tests disconnected error
// when trying to manually restore the session
// while the server is unreachable and autoconn is disabled
func TestRestoreSessDisconnNoAutoconn(t *testing.T) {
	// Initialize client
	client, err := newClient(
		wwrclt.Options{
			Autoconnect:           wwr.Disabled,
			ReconnectionInterval:  5 * time.Millisecond,
			DefaultRequestTimeout: 50 * time.Millisecond,
		},
		&memchan.ClientTransport{},
		clientHooks{},
	)
	require.NoError(t, err)

	// Try to restore a session and expect a ErrDisconnected error
	err = client.Connection.RestoreSession(
		context.Background(),
		[]byte("inexistent_key"),
	)
	require.Error(t, err)
	require.IsType(t, wwr.ErrDisconnected{}, err)
}
