package conf

import "sync"

var StopChan chan string = make(chan string,2)

var AccessCount int64 = 0
var ACMute sync.Mutex


func init (){

}