// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/pflag"
)

var usageString = `Usage: netsound_server [OPTION]... [FILE]...

Stream audio to netsound_client from FILE(s) (or current directory by default).

Options:
  -d, --debug=[info, all] 	   print debugging info to stdout.




  -i, --inbound=[address]:port 	  address to listen for incomming http requests
  -o, --outbound=[address]:port   address to use for outbound audio stream(tcp)
				   (by default, this is chosen at random)

  -h, --help 			  display help and exit

DEBUGGING OPTIONS:
  'info' only prints basic info such as starting or stopping an audio stream,
  'all'  prints finer-grained information, and is more useful for debugging.

`

var (
	flags struct {
		debug    *bool   // use pflag.BoolP
		inbound  *string // use pflag.StringP
		outbound *string // ^^
		serveDir []string
	}
)

func getflags() {
	flags.debug = pflag.BoolP(flags.debug, "debug", "d", false,
		"print debug info")

	flags.inbound = pflag.StringP("inbound", "i", ":80",
		"address to listen for incomming HTTP requests")
	flags.outbound = pflag.StringP("outbound", "o", "",
		"address to use for outbound audio stream(tcp)")
	flags.serveDir = append(flags.serveDir, pflag.Args()[len(os.Args()-len(pflag.Arg(len(pflag.Args()-1)))):])

	pflag.Usage = func() { fmt.Sprint(usageString) }
	pflag.Parse()
}

func main() {
	getflags()
	for {
		for e := range server.Start() {
			log.Println(e.Error())
		}
	}

}
