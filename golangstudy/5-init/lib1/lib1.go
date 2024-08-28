package lib1

import "fmt"

// lib1提供的api
func TestLib1() {
	fmt.Println("lib1Test()....")
}

func init() {
	fmt.Println("lib1.5-init() starting.....")
}
