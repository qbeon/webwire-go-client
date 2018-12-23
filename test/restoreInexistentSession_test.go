package test

import (
	"context"
	"testing"
	"time"

	wwr "github.com/qbeon/webwire-go"
	wwrclt "github.com/qbeon/webwire-go-client"
	"github.com/stretchr/testify/require"
)

// TestRestoreInexistentSession tests the restoration of an inexistent session
func TestRestoreInexistentSession(t *testing.T) {
	// Initialize server
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
		},
		clientHooks{},
	)

	// Try to restore the session and expect it to fail
	// due to the session being inexistent
	sessionRestorationError := client.Connection.RestoreSession(
		context.Background(),
		[]byte("lalala"),
	)
	require.Error(t, sessionRestorationError)
	require.IsType(t, wwr.ErrSessionNotFound{}, sessionRestorationError)
}
