// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package netsound

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type StreamServer struct {
	*http.Server
	Conn         *websocket.Conn
	StreamSource *soundfile
}

func NewStreamServer() *StreamServer {
	srv := &StreamServer{}
	return srv
}
