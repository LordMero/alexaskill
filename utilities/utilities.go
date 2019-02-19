package utilities

import "log"

func Catch(e error){
	if e != nil { log.Fatal(e)}
}
