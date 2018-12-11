package test

import (
	"testing"

	wwr "github.com/qbeon/webwire-go"
	wwrclt "github.com/qbeon/webwire-go-client"
	"github.com/stretchr/testify/require"
)

// TestSubProtocolMatch tests dial success on sub-protocol match
func TestSubProtocolMatch(t *testing.T) {
	// Initialize server
	setup := setupTestServer(
		t,
		&ServerImpl{},
		wwr.ServerOptions{
			SubProtocolName: []byte("sharedprotocol"),
		},
		nil, // Use the default transport implementation
	)

	// Initialize client
	client := setup.newClient(
		wwrclt.Options{
			Autoconnect:     wwr.Disabled,
			SubProtocolName: []byte("sharedprotocol"),
		},
		nil, // Use the default transport implementation
		clientHooks{},
	)

	require.NoError(t, client.Connection.Connect())
}
