package test

import (
	"context"
	"sync"
	"testing"
	"time"

	wwr "github.com/qbeon/webwire-go"
	wwrclt "github.com/qbeon/webwire-go-client"
	"github.com/stretchr/testify/assert"
)

// TestConcurrentConnect tests concurrent calling of Client.Connect
func TestConcurrentConnect(t *testing.T) {
	concurrentAccessors := 16
	finished := sync.WaitGroup{}
	finished.Add(concurrentAccessors)

	// Initialize webwire server
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

	connect := func() {
		defer finished.Done()
		assert.NoError(t, client.Connection.Connect(context.Background()))
	}

	for i := 0; i < concurrentAccessors; i++ {
		go connect()
	}

	finished.Wait()
}
