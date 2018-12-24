package test

import (
	"context"
	"sync"
	"testing"
	"time"

	wwr "github.com/qbeon/webwire-go"
	wwrclt "github.com/qbeon/webwire-go-client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAutoSessionRestoration tests automatic session restoration on connection
// establishment
func TestAutoSessionRestoration(t *testing.T) {
	lookupTriggered := sync.WaitGroup{}
	lookupTriggered.Add(1)
	hookTriggered := sync.WaitGroup{}
	hookTriggered.Add(1)

	// Initialize webwire server
	setup := setupTestServer(
		t,
		&ServerImpl{
			Signal: func(
				_ context.Context,
				conn wwr.Connection,
				_ wwr.Message,
			) {
				assert.NoError(t, conn.CreateSession(nil))
			},
		},
		wwr.ServerOptions{
			SessionManager: &SessionManager{
				// Finds session by key
				SessionLookup: func(key string) (
					wwr.SessionLookupResult,
					error,
				) {
					lookupTriggered.Done()
					return wwr.NewSessionLookupResult(
						time.Now(), // Creation
						time.Now(), // LastLookup
						nil,
					), nil
				},
			},
		},
		nil, // Use the default transport implementation
	)

	// Initialize client
	client := setup.newClient(
		wwrclt.Options{
			Autoconnect: wwr.Disabled,
		},
		clientHooks{
			OnSessionCreated: func(*wwr.Session) {
				hookTriggered.Done()
			},
		},
	)

	require.NoError(t, client.Connection.Connect())

	// Create session
	require.NoError(t, client.Connection.Signal(
		context.Background(),
		[]byte("create_session"),
		wwr.Payload{},
	))

	hookTriggered.Wait()
	require.NotNil(t, client.Connection.Session())

	// Disconnect client without closing the session
	hookTriggered.Add(1)
	client.Connection.Close()

	// Reconnect
	require.NoError(t, client.Connection.Connect())

	lookupTriggered.Wait()
	hookTriggered.Wait()
	require.NotNil(t, client.Connection.Session())
}
