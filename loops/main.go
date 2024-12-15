package main

import (
	"fmt"
	"slices"
)

func main() {
	animals := [2]string{}
	animals[0] = "human"
	animals[1] = "me"
	fmt.Println(animals)
	/* slice */
	new_animals := []string{}
	new_animals = append(new_animals, "you")
	fmt.Println(new_animals)
	new_animals = slices.Delete(new_animals, 0, 1)
	fmt.Println(new_animals)

	for i := 0; i < len(animals); i++ {
		fmt.Printf("%s", animals[i])
	}
	for value := range 10 {
		fmt.Println(value)
	}
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}
