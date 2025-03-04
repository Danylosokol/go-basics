package data

import "fmt"

var Countries [10]string
var Slice []int
var Codes map[int]string

func init() {
	Countries[0] = "Argentina"
	Countries[1] = "Brazil"
	Countries[8] = "UK"
	qty := len(Countries)
	fmt.Println(Countries, qty)
}
