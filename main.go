package main

import (
	"fmt"
)

type person struct {
	firstName      string
	lastName       string
	favoriteFlavor []string
}

func main() {
	var p person

	p = person{
		firstName:      "Aarjap",
		lastName:       "Piya",
		favoriteFlavor: []string{"blackcurrent", "vanilla", "Strawberry"}}

	var a map[string]person
	a = make(map[string]person)

	a[p.lastName] = p
	// _ is connect to p.lastName and value is connect to p where p has firstName, lastName and favoriteFlavor. p is store in value of for loop which is used to print stuff
	for _, value := range a {
		fmt.Println(value.firstName)
		fmt.Println(value.lastName)
		for _, flavor := range value.favoriteFlavor {
			fmt.Println(flavor)
		}
	}

}
