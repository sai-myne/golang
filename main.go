package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// loop maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		23242343: "mario",
		12332434: "luigi",
		56463565: "peach",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[23242343])

	phonebook[12332434] = "bowser"
	fmt.Println(phonebook)

	phonebook[56463565] = "yoshi"
	fmt.Println(phonebook)

}
