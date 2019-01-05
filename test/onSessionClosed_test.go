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

// TestOnSessionClosed tests the OnSessionClosed hook
func TestOnSessionClosed(t *testing.T) {
	authenticated := sync.WaitGroup{}
	authenticated.Add(1)
	hookCalled := sync.WaitGroup{}
	hookCalled.Add(1)

	// Initialize webwire server
	setup := setupTestServer(
		t,
		&ServerImpl{
			Request: func(
				_ context.Context,
				conn wwr.Connection,
				_ wwr.Message,
			) (wwr.Payload, error) {
				// Try to create a new session
				err := conn.CreateSession(nil)
				assert.NoError(t, err)
				if err != nil {
					return wwr.Payload{}, err
				}

				go func() {
					// Wait until the authentication request is finished
					authenticated.Wait()

					// Close the session
					assert.NoError(t, conn.CloseSession())
				}()

				return wwr.Payload{}, nil
			},
		},
		wwr.ServerOptions{},
		nil, // Use the default transport implementation
	)

	// Initialize client
	client := setup.newClient(
		wwrclt.Options{
			DefaultRequestTimeout: 2 * time.Second,
		},
		clientHooks{
			OnSessionClosed: func() {
				hookCalled.Done()
			},
		},
	)
	require.NoError(t, client.Connection.Connect(context.Background()))

	// Send authentication request and await reply
	_, err := client.Connection.Request(
		context.Background(),
		[]byte("login"),
		wwr.Payload{
			Encoding: wwr.EncodingBinary,
			Data:     []byte("credentials"),
		},
	)
	require.NoError(t, err)
	authenticated.Done()

	// Verify client session
	hookCalled.Wait()
}
