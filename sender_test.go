package main

import (
	"net/http/httptest"
	"net/http/httputil"
	"testing"
)

func TestRequestListener(t *testing.T) {
	client := httptest.NewRecorder()
	httputil.NewClientConn()
}
