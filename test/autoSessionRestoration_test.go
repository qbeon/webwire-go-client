package test

import (
	"context"
	"testing"

	wwr "github.com/qbeon/webwire-go"
	wwrclt "github.com/qbeon/webwire-go-client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAutoSessionRestoration tests automatic session restoration on connection
// establishment
func TestAutoSessionRestoration(t *testing.T) {
	sessionStorage := make(map[string]*wwr.Session)

	currentStep := 1
	var createdSession *wwr.Session

	// Initialize webwire server
	setup := setupTestServer(
		t,
		&ServerImpl{
			Request: func(
				_ context.Context,
				conn wwr.Connection,
				msg wwr.Message,
			) (wwr.Payload, error) {
				if currentStep == 2 {
					// Expect the session to have been automatically restored
					CompareSessions(t, createdSession, conn.Session())
					return wwr.Payload{}, nil
				}

				// Try to create a new session
				err := conn.CreateSession(nil)
				assert.NoError(t, err)
				return wwr.Payload{}, err
			},
		},
		wwr.ServerOptions{
			SessionManager: &SessionManager{
				// Saves the session
				SessionCreated: func(conn wwr.Connection) error {
					session := conn.Session()
					sessionStorage[session.Key] = session
					return nil
				},
				// Finds session by key
				SessionLookup: func(key string) (
					wwr.SessionLookupResult,
					error,
				) {
					// Expect the key of the created session to be looked up
					assert.Equal(t, createdSession.Key, key)

					assert.Contains(t, sessionStorage, key)
					session := sessionStorage[key]
					// Session found
					return wwr.NewSessionLookupResult(
						session.Creation,                      // Creation
						session.LastLookup,                    // LastLookup
						wwr.SessionInfoToVarMap(session.Info), // Info
					), nil
				},
			},
		},
		nil, // Use the default transport implementation
	)

	// Initialize client
	client := setup.newClient(wwrclt.Options{}, clientHooks{})

	require.NoError(t, client.Connection.Connect())

	/*****************************************************************\
		Step 1 - Create session and disconnect
	\*****************************************************************/

	// Create a new session
	_, err := client.Connection.Request(
		context.Background(),
		[]byte("login"),
		wwr.Payload{
			Encoding: wwr.EncodingBinary,
			Data:     []byte("auth"),
		},
	)
	require.NoError(t, err)

	createdSession = client.Connection.Session()

	// Disconnect client without closing the session
	client.Connection.Close()

	// Ensure the session isn't lost
	require.NotEqual(t,
		wwrclt.StatusConnected, client.Connection.Status(),
		"Client is expected to be disconnected",
	)
	require.NotEqual(t,
		"", client.Connection.Session().Key,
		"Session lost after disconnection",
	)

	/*****************************************************************\
		Step 2 - Reconnect, restore and verify authentication
	\*****************************************************************/
	currentStep = 2

	// Reconnect (this should automatically try to restore the session)
	require.NoError(t, client.Connection.Connect())

	// Verify whether the previous session was restored automatically
	// and the server authenticates the user
	_, err = client.Connection.Request(
		context.Background(),
		[]byte("verify"),
		wwr.Payload{
			Encoding: wwr.EncodingBinary,
			Data:     []byte("is_restored?"),
		},
	)
	require.NoError(t, err)
}
