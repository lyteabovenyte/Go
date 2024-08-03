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

	val := "AmirAlaeifar"

	for index, ch := range val {
		switch ch {
		case 'A', 'a':
			if ch == 'A' {
				fmt.Printf("uppercase 'A' at position %d\n", index)
				break // going to the next loop. instaed of else block.
			}
			fmt.Printf("lowecase 'a' at position %d\n", index)
		case 'm':
			fmt.Printf("letter 'm' at position %d\n", index)
		}
	}

	fmt.Println("---------")

	exept_a := 0

	for index, ch := range val {
		switch ch {
		case 'a':
			fmt.Printf("loop1 --> ch --> %s\n", string(ch))
			fmt.Printf("got the letter %d at position %d\n", ch, index)
			fallthrough
		case 'A':
			fmt.Printf("loop2 --> ch --> %s\n", string(ch))
			fmt.Printf("uppercase %d letter at position %d\n", ch, index)
		default:
			exept_a++
			fmt.Println("not wanted")
		}
	}

	fmt.Printf("not A letters --> %d\n", exept_a)

	fmt.Println("-----------")

	day := "Thursday"

	switch day {
	case "Monday", "Tuesday", "Wednesday":
		fmt.Println("Weekday")
	case "Thursday":
		fmt.Println("Weekday")
		fallthrough // transfer the execution to the next block of code, whether the condition are met or not.
	case "Friday":
		fmt.Println("Start of the weekend")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	}

	fmt.Println("---------")

	for counter := 0; counter <= 20; counter++ {
		switch val := counter / 2; val {
		case 2, 3, 5, 7:
			fmt.Printf("prime value %d\n", val)
		default:
			fmt.Printf("not prime value %d\n", val)
		}
	}

	fmt.Println("------labels-------")

	counter := 0
target:
	fmt.Printf("target --> counter = %d\n", counter)
	counter++
	if counter < 5 {
		goto target
	}

	fmt.Println("------labels Use Cases-------")

	for i := 0; i < 10; i++ {
		if i == 5 {
			goto loop
		}
		fmt.Println(i)
	}
loop:
	fmt.Println("reached the loop label")

	
}
