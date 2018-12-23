package test

import (
	"context"
	"sync"
	"testing"
	"time"

	wwr "github.com/qbeon/webwire-go"
	wwrclt "github.com/qbeon/webwire-go-client"
	"github.com/stretchr/testify/require"
)

// TestRestoreSession tests manual session restoration by key
func TestRestoreSession(t *testing.T) {
	test := func(t *testing.T, clientOptions wwrclt.Options) {
		lookupTriggered := sync.WaitGroup{}
		lookupTriggered.Add(1)
		clientHookTriggered := sync.WaitGroup{}
		clientHookTriggered.Add(1)

		var sessionKey = "testsessionkey"
		sessionCreation := time.Now()

		// Initialize server
		setup := setupTestServer(
			t,
			&ServerImpl{},
			wwr.ServerOptions{
				SessionManager: &SessionManager{
					SessionLookup: func(key string) (
						wwr.SessionLookupResult,
						error,
					) {
						defer lookupTriggered.Done()
						if key != sessionKey {
							// Session not found
							return nil, nil
						}
						return wwr.NewSessionLookupResult(
							sessionCreation, // Creation
							time.Now(),      // LastLookup
							nil,             // Info
						), nil
					},
				},
			},
			nil, // Use the default transport implementation
		)

		// Initialize client
		client := setup.newClient(
			clientOptions,
			clientHooks{
				OnSessionCreated: func(*wwr.Session) {
					clientHookTriggered.Done()
				},
			},
		)

		if clientOptions.Autoconnect == wwr.Disabled {
			require.NoError(t, client.Connection.Connect())
		}

		// Send restoration request and await reply
		err := client.Connection.RestoreSession(
			context.Background(),
			[]byte(sessionKey),
		)
		require.NoError(t, err)

		lookupTriggered.Wait()
		clientHookTriggered.Wait()
	}

	t.Run("Autoconn", func(t *testing.T) {
		test(t, wwrclt.Options{
			ReconnectionInterval:  5 * time.Millisecond,
			DefaultRequestTimeout: 50 * time.Millisecond,
		})
	})

	t.Run("NoAutoconn", func(t *testing.T) {
		test(t, wwrclt.Options{
			Autoconnect: wwr.Disabled,
		})
	})
}
