package main

import "fmt"

func main() {
	people := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 25,
	}
	ageGroups := make(map[int][]string)

	//ageGroups
	//25: alice, charlie
	//30: bob

	// Group people by age
	for name := range people {
		age := people[name]

		ageGroups[age] = append(ageGroups[age], name)

	}
	// Collect keys in a slice
	keys := []int{}
	for key := range ageGroups {
		keys = append(keys, key)
	}
	// Iterate over keys and print age groups
	fmt.Println("People grouped by age:")
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		fmt.Printf("Age %d: %v\n", key, ageGroups[key])
	}
}
