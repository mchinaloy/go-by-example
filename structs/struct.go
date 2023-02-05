package main

import "fmt"

type Person struct {
	first string
	last  string
}

type SecretAgent struct {
	person Person
	ltk    bool
}

func main() {
	person_1 := Person{
		first: "Michael",
		last:  "Chin",
	}

	secret_agent_1 := SecretAgent{
		person: person_1,
		ltk:    true,
	}

	fmt.Println(person_1)
	fmt.Println(secret_agent_1)

	secret_agent_1.print_ltk()
}

func (s SecretAgent) print_ltk() {
	fmt.Println(s.ltk)
}
