package notify

import "sync"

var LoadDownConfig = make(chan bool, 1)
var MainWait = new(sync.WaitGroup)
