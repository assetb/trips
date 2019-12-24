package middleware

import "fmt"

//Error treats errors of application
//current is a dummy just for future use
func Error(e ...interface{}) {
	fmt.Println(e...)
}
