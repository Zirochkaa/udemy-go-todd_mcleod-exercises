package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByLast []user
type ByAge []user

func (l ByLast) Len() int           { return len(l) }
func (l ByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}
	fmt.Println("All users:")
	for _, u := range users {
		fmt.Println("\t", u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, s := range u.Sayings {
			fmt.Printf("\t\t%v\n", s)
		}
	}

	sort.Sort(ByLast(users))
	fmt.Println("\nUsers sorted by last name:")
	for _, u := range users {
		fmt.Println("\t", u.First, u.Last, u.Age)
		for _, s := range u.Sayings {
			fmt.Printf("\t\t%v\n", s)
		}
	}

	sort.Sort(ByAge(users))
	fmt.Println("\nUsers sorted by age:")
	for _, u := range users {
		fmt.Println("\t", u.First, u.Last, u.Age)
		for _, s := range u.Sayings {
			fmt.Printf("\t\t%v\n", s)
		}
	}
}
