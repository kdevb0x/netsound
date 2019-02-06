// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package main

import (
	"os"

	"github.com/kdevb0x/netsound"
	"github.com/spf13/pflag"
)

func main() {
	_ := netsound.LoadSoundFile("")
	pflag.Parse()
	_ := os.Args()
}
