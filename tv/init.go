package tv

import (
	"sync"
)

var Outputmap map[string]string
var Wg sync.WaitGroup
var ch chan bool
var First bool

func init() {
	Outputmap = make(map[string]string)
	First = true
	flag = false
}
