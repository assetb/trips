package middleware

import "fmt"

//Error treats errors of application
func Error(e ...interface{}) {
	fmt.Println(e...)
}
