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

// TestRestoreSessDisconnTimeout tests autoconnect timeout
// when the server is unreachable
func TestRestoreSessDisconnTimeout(t *testing.T) {
	// Initialize client
	client, err := newClient(
		wwrclt.Options{
			ReconnectionInterval:  5 * time.Millisecond,
			DefaultRequestTimeout: 50 * time.Millisecond,
		},
		&memchan.ClientTransport{},
		clientHooks{},
	)
	require.NoError(t, err)

	// Send request and await reply
	err = client.Connection.RestoreSession(
		context.Background(),
		[]byte("inexistent_key"),
	)
	require.Error(t, err)
	require.IsType(t, wwr.ErrTimeout{}, err)
	require.True(t, wwr.IsErrTimeout(err))
}
