package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":  "der",
		"blue": "eulb",
	}
	// var colors map[string]string // JUST USE MAKE APPARENTLY?
	// colors := make(map[string]string)
	colors["white"] = "etihw"
	fmt.Println(colors)
	delete(colors, "white")
	describeColorsMap(colors)
}

func describeColorsMap(cm map[string]string) {
	for k, v := range cm {
		fmt.Println(k, " backwards is ", v)
	}
}
