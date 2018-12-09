package test

import (
	wwr "github.com/qbeon/webwire-go"
	wwrclt "github.com/qbeon/webwire-go-client"
)

type clientHooks struct {
	OnSessionCreated func(*wwr.Session)
	OnSessionClosed  func()
	OnDisconnected   func()
	OnSignal         func(wwr.Message)
}

// client implements the wwrclt.Implementation interface
type client struct {
	Connection wwrclt.Client
	Hooks      clientHooks
}

// OnSessionCreated implements the wwrclt.Implementation interface
func (clt *client) OnSessionCreated(newSession *wwr.Session) {
	if clt.Hooks.OnSessionCreated != nil {
		clt.Hooks.OnSessionCreated(newSession)
	}
}

// OnSessionClosed implements the wwrclt.Implementation interface
func (clt *client) OnSessionClosed() {
	if clt.Hooks.OnSessionClosed != nil {
		clt.Hooks.OnSessionClosed()
	}
}

// OnDisconnected implements the wwrclt.Implementation interface
func (clt *client) OnDisconnected() {
	if clt.Hooks.OnDisconnected != nil {
		clt.Hooks.OnDisconnected()
	}
}

// OnSignal implements the wwrclt.Implementation interface
func (clt *client) OnSignal(msg wwr.Message) {
	if clt.Hooks.OnSignal != nil {
		clt.Hooks.OnSignal(msg)
	}
}
