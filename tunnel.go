package main

import (
	"github.com/gofrs/uuid"
	gw "go-web-utilities"
	"time"
)

type Tunnel struct {
	*gw.Heartbeat
	ID string
	cache map[string]interface{}
	toHostMessages *[]interface{}
	toClientMessages *[]interface{}
	toHostChan chan interface{}
	toClientChan chan interface{}
}

func NewTunnel(msToDie time.Duration) *Tunnel {
	id, _ := uuid.NewV4()
	return &Tunnel{
		gw.NewHeartbeat(msToDie),
		id.String(),
		map[string]interface{}{},
		&[]interface{}{},
		&[]interface{}{},
		nil,
		nil,
	}
}

func (tu *Tunnel) SetCache (key string, value interface{}) {
	tu.cache[key] = value
}

func (tu *Tunnel) GetCache (key string) interface{}  {
	if value, ok := tu.cache[key]; ok {
		return value
	} else {
		return nil
	}
}

func (tu *Tunnel) SetHostChan(c chan interface{})  {
	tu.toHostChan = c
}

func (tu *Tunnel) SetClientChan(c chan interface{})  {
	tu.toClientChan = c
}

func (tu *Tunnel) SendToClient(data interface{})  {
	sendMessageTo(tu.toClientChan, tu.toClientMessages, data)
	tu.toClientChan = nil
}

func (tu *Tunnel) SendToHost(data interface{})  {
	sendMessageTo(tu.toHostChan, tu.toHostMessages, data)
	tu.toHostChan = nil
}

func sendMessageTo(c chan interface{}, msgs *[]interface{}, data interface{})  {
	*msgs = append(*msgs, data)
	if c != nil && len(*msgs) > 0 {
		c <- *msgs
		*msgs = []interface{}{}
	}
}
