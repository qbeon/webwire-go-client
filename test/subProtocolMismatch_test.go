package test

import (
	"testing"

	wwr "github.com/qbeon/webwire-go"
	wwrclt "github.com/qbeon/webwire-go-client"
	"github.com/stretchr/testify/require"
)

// TestSubProtocolMismatch tests dial failure on sub-protocol mismatch
func TestSubProtocolMismatch(t *testing.T) {
	/* SERVER: B; CLIENT: A */

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
	require.Error(t, clientMismatch.Connection.Connect())

	/* SERVER: nil; CLIENT: A */

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
	require.Error(t, clientNoSubProto.Connection.Connect())

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
	require.Error(t, clientNoCltSubProto.Connection.Connect())
}
