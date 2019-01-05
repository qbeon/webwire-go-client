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

// TestReqDisconn(tests Client.Request expecting it to return a
// disconnected-error when trying to send a request while the client is
// disconnected and autoconn is disabled
func TestReqDisconn(t *testing.T) {
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

	// Try to send a request and expect a ErrDisconnected error
	_, err = client.Connection.Request(
		context.Background(),
		nil,
		wwr.Payload{Data: []byte("testdata")},
	)
	require.Error(t, err)
	require.IsType(t, wwr.ErrDisconnected{}, err)
	require.False(t, wwr.IsErrCanceled(err))
	require.False(t, wwr.IsErrTimeout(err))
}
