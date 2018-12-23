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

// TestConcurrentSignal tests concurrent calling of Client.Signal
func TestConcurrentSignal(t *testing.T) {
	concurrentAccessors := 16
	finished := sync.WaitGroup{}
	finished.Add(concurrentAccessors * 2)

	// Initialize webwire server
	setup := setupTestServer(
		t,
		&ServerImpl{
			Signal: func(
				_ context.Context,
				_ wwr.Connection,
				_ wwr.Message,
			) {
				finished.Done()
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
		clientHooks{},
	)
	defer client.Connection.Close()

	require.NoError(t, client.Connection.Connect())

	sendSignal := func() {
		defer finished.Done()
		assert.NoError(t, client.Connection.Signal(
			context.Background(),
			[]byte("sample"),
			wwr.Payload{
				Encoding: wwr.EncodingBinary,
				Data:     []byte("samplepayload"),
			},
		))
	}

	for i := 0; i < concurrentAccessors; i++ {
		go sendSignal()
	}

	finished.Wait()
}
