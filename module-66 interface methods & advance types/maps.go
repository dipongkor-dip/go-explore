package main

import "fmt"

type Vehicle struct {
	name string
	sits int
}

func maps_func() {
	// var mp map[string]int // thats nil map

	// var mp = make(map[string]int)
	// mp["fScore"] = 20
	// mp["sScore"] = 30
	// mp["tScore"] = 40

	// var mp = map[string]int{"fScore": 20, "sScore": 30, "tScore": 40}
	// delete(mp, "sScore")

	var mp = map[string]Vehicle{"bus": {name: "Samoli Bus", sits: 20}}

	fmt.Printf("%v\n", mp)
}
