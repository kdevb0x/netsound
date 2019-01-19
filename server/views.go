// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package server

import (
	"html/template"
	"net/http"
)

type ViewController struct {
	// RouteMap maps handler routes (url paths) to View objects
	RouteMap map[string]*View
}

type View func(route string, tmpl *template.Template) http.Handler

func Start() chan error {
	errchan := make(chan error, 10)
	return errchan
}
