// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/freestix/libstix/indicator"
	"github.com/freestix/libstix/stix"
)

func main() {
	s := stix.Create()
	i1 := indicator.Create()
	i1.SetTimestampToNow()
	i1.AddTitle("Attack 2015-02")
	i1.AddType("IP Watchlist")
	i1.AddType("URL Watchlist")

	h := stix.CreateHandling()
	h.AddControlledStructure("Foobar")
	h.AddTLPMarking("Red")

	i1.AddHandling(h)

	s.AddIndicator(i1)

	var data []byte
	data, _ = json.MarshalIndent(s, "", "    ")

	fmt.Println(string(data))
}
