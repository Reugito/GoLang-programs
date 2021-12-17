package main

import "fmt"

func main() {
	var obj1 = emp{10, "rao", "dhotre"}
	var obj2 = emp{lastName: "om"}
	fmt.Println(obj1, obj2)
}

type emp struct {
	id        int
	firstName string
	lastName  string
}
