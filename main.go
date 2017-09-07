package main

import(
	_ "huanggh.site/learning/api-gateway/http"
	"huanggh.site/learning/api-gateway/conf"
	"fmt"
)
func main(){
	fmt.Printf("api-gateway had stop,reason:%s\n.", <-conf.StopChan)

}
