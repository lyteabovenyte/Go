package main

import "fmt"

func main() {

	product := "amir"
	another := "another"

	i := 0
	for _, char := range another {
		i += int(char)
	}
	fmt.Println(i)

	for index, ch := range product { // treats strings as a rune character.
		fmt.Println("index:", index, "ch:", int(ch))
	}

	fmt.Println("--------")

	products := []string{"amir", "mahkam", "ala"}
	fmt.Printf("type of the string: %T\n", products)
	for index, item := range products {
		fmt.Printf("index: %d, item: %s\n", index, item)
	}

	fmt.Println("---------")

	val := "amir"

	for index, ch := range val {
		switch (ch) {
		case 'a':
			fmt.Printf("A at position %d\n", index)
		case 'r', 'R':
			if ch == 'R' {
				fmt.Println("that a typo")
				break
			}
			fmt.Println("the end of the name is ", string(ch))
		}
	}

	fmt.Println("---------")

	
}
