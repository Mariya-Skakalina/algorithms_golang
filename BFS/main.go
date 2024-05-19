package main

import "fmt"

type Person struct {
	Name    string
	Seller  bool
	Friends []Person
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func search(arr []Person) {

	for _, v := range arr {
		var sellers []string
		if contains(sellers, v.Name) == true {
			continue
		}
		sellers = append(sellers, v.Name)
		if v.Seller {
			fmt.Printf("Продавец манго найден %v\n", v.Name)
			break
		}
	}
	for _, v := range arr {
		search(v.Friends)
	}

}

func main() {
	var Jom = Person{Name: "Jom", Seller: false}
	var Kol = Person{Name: "Kol", Seller: false}
	var Bob = Person{Name: "Bob", Seller: false, Friends: []Person{Jom, Kol}}

	var Kate = Person{Name: "Kate", Seller: true}
	var Sveta = Person{Name: "Sveta", Seller: false}
	var Alisa = Person{Name: "Alice", Seller: false, Friends: []Person{Kol, Kate, Sveta}}
	people := []Person{Bob, Alisa}
	search(people)
}
