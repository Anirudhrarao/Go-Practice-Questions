package main

import (
	"fmt"
)

type Person struct {
	name    string
	age     int
	friends []string // Slice of string to store friend's names
}

func (p *Person) addFriend(friendName string) {
	p.friends = append(p.friends, friendName)
}

func main() {
	// Creating struct
	p1 := Person{
		name:    "Anirudhra",
		age:     22,
		friends: []string{},
	}
	// Adding friends
	p1.addFriend("yash")
	p1.addFriend("ankit")
	// Displaying the person's details
	fmt.Printf("Name: %s, Age: %d\n", p1.name, p1.age)
	fmt.Println("Friends:", p1.friends)
}
