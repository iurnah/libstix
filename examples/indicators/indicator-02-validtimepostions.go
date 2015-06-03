package main

import (
	"encoding/json"
	"fmt"
	"github.com/jordan2175/freestix/libstix/indicator"
	"github.com/jordan2175/freestix/libstix/stix"
)

func main() {

	s := stix.New()

	i1 := indicator.New()
	i1.AddValidTimePosition("2015-01-01T00:00:00-0700", "2015-02-02T23:59:59-0700")
	i1.AddValidTimePosition("2014-01-01T00:00:00-0700", "2014-02-02T23:59:59-0700")

	s.AddIndicator(i1)

	fmt.Println("====================================")
	var data []byte
	data, _ = json.MarshalIndent(s, "", "    ")

	fmt.Println(string(data))

}
