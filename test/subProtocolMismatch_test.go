package test

import (
	"context"
	"testing"

	wwr "github.com/qbeon/webwire-go"
	wwrclt "github.com/qbeon/webwire-go-client"
	"github.com/stretchr/testify/require"
)

// TestSubProtocolMismatch tests dial failure on sub-protocol mismatch
func TestSubProtocolMismatch(t *testing.T) {

	t.Run("AB", func(t *testing.T) {
		/* SERVER: A; CLIENT: B */

		// Initialize server
		setupMismatch := setupTestServer(
			t,
			&ServerImpl{},
			wwr.ServerOptions{
				SubProtocolName: []byte("serverprotocol"),
			},
			nil, // Use the default transport implementation
		)

		// Initialize client
		clientMismatch := setupMismatch.newClient(
			wwrclt.Options{
				Autoconnect:     wwr.Disabled,
				SubProtocolName: []byte("clientprotocol"),
			},
			clientHooks{},
		)

		err := clientMismatch.Connection.Connect(context.Background())
		require.Error(t, err)
	})

	t.Run("NB", func(t *testing.T) {
		/* SERVER: nil; CLIENT: B */

		// Initialize server (no sub-protocol)
		setupNoSubProto := setupTestServer(
			t,
			&ServerImpl{},
			wwr.ServerOptions{},
			nil, // Use the default transport implementation
		)

		// Initialize client
		clientNoSubProto := setupNoSubProto.newClient(
			wwrclt.Options{
				Autoconnect:     wwr.Disabled,
				SubProtocolName: []byte("clientprotocol"),
			},
			clientHooks{},
		)

		err := clientNoSubProto.Connection.Connect(context.Background())
		require.Error(t, err)
	})

	t.Run("AN", func(t *testing.T) {
		/* SERVER: A; CLIENT: nil */

		// Initialize server (no sub-protocol)
		setupNoCltSubProto := setupTestServer(
			t,
			&ServerImpl{},
			wwr.ServerOptions{
				SubProtocolName: []byte("serverprotocol"),
			},
			nil, // Use the default transport implementation
		)

		// Initialize client
		clientNoCltSubProto := setupNoCltSubProto.newClient(
			wwrclt.Options{
				Autoconnect: wwr.Disabled,
			},
			clientHooks{},
		)

		err := clientNoCltSubProto.Connection.Connect(context.Background())
		require.Error(t, err)
	})
}
