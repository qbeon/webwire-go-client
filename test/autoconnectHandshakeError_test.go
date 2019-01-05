package test

import (
	"context"
	"testing"

	wwr "github.com/qbeon/webwire-go"
	wwrclt "github.com/qbeon/webwire-go-client"
	"github.com/stretchr/testify/require"
)

// TestAutoconnectHandshakeError tests dial failure on sub-protocol mismatch
func TestAutoconnectHandshakeError(t *testing.T) {
	t.Run("SubProtocolMismatch_AB", func(t *testing.T) {
		/* SERVER: A; CLIENT: B */
		setupMismatch := setupTestServer(
			t,
			&ServerImpl{},
			wwr.ServerOptions{
				SubProtocolName: []byte("serverprotocol"),
			},
			nil, // Use the default transport implementation
		)

		// Initialize client
		newClient := setupMismatch.newClient(
			wwrclt.Options{
				SubProtocolName: []byte("clientprotocol"),
			},
			clientHooks{},
		)
		err := newClient.Connection.Connect(context.Background())
		require.Error(t, err)
	})

	t.Run("SubProtocolMismatch_NB", func(t *testing.T) {
		/* SERVER: nil; CLIENT: B */
		setupMismatch := setupTestServer(
			t,
			&ServerImpl{},
			wwr.ServerOptions{},
			nil, // Use the default transport implementation
		)

		// Initialize client
		newClient := setupMismatch.newClient(
			wwrclt.Options{
				SubProtocolName: []byte("clientprotocol"),
			},
			clientHooks{},
		)
		err := newClient.Connection.Connect(context.Background())
		require.Error(t, err)
	})

	t.Run("SubProtocolMismatch_AN", func(t *testing.T) {
		/* SERVER: A; CLIENT: nil */
		setupMismatch := setupTestServer(
			t,
			&ServerImpl{},
			wwr.ServerOptions{
				SubProtocolName: []byte("serverprotocol"),
			},
			nil, // Use the default transport implementation
		)

		// Initialize client
		newClient := setupMismatch.newClient(wwrclt.Options{}, clientHooks{})
		err := newClient.Connection.Connect(context.Background())
		require.Error(t, err)
	})
}
