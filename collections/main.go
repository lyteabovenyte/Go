package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var names [3]string

	names[0] = "amir"
	names[1] = "mahkam"
	names[2] = "ala"
	fmt.Printf("the type is %T and the array is %v", names, names)

	fmt.Println("------reference to an fixed-length array------")
	otherArray := &names
	fmt.Printf("the type of the otherArray is %T and the values is %v\n", otherArray, *otherArray)

	fmt.Println("------array of pointers-------")

	val1 := "amir"
	val2 := "mahkam"
	val3 := "ala"

	thirdArray := [3]*string{&val1, &val2, &val3}
	for _, value := range thirdArray {
		fmt.Println(*value)
	}
	val1 = "changed"
	for _, value := range thirdArray {
		fmt.Println(*value)
	}

	fmt.Println("---------------")

	array := [3]string{"amir", "mahkam", "ala"}
	arraySlice := make([]string, 3)

	for _, element := range array {
		arraySlice = append(arraySlice, element)
	}

	for _, element := range arraySlice {
		fmt.Println(element)
	}

	fmt.Println("--------------")

	namesSlice := make([]string, 3, 6)
	for index, element := range namesSlice {
		fmt.Printf("index is %d and element is %v\n", index, element)
	}

	anotherSlice := append(namesSlice, "amir", "another-amir")
	fmt.Println(cap(namesSlice))
	fmt.Println(cap(anotherSlice))
	fmt.Println(len(anotherSlice))
	// fmt.Println(namesSlice[5]) // can't be reached, cause runtime error cause the value is empty not the zero value.

	fmt.Println("---------------")
	notInitSlice := make([]string, 3)

	notInitSlice = append(notInitSlice, anotherSlice...) // TODO: RESERACH IT.
	fmt.Println(notInitSlice)

	fmt.Println("---------------")

	white := make([]string, 1, 6)

	black := append(white, "red", "blue")

	white[0] = "another-color"
	fmt.Println("white:")
	for i, ele := range white {
		fmt.Printf("%d) %s\n", i, ele)
	}

	fmt.Printf("white is %v\n", white)
	fmt.Printf("black is %v\n", black)
	fmt.Printf("the type of the black %T\n", black)

	fmt.Println("--------variadic function---------")

	Hello("amir", "hastam")

	fmt.Println("-------------")

	variadic := []string{"amir", "alaeifar", "happy"}

	appended := append(white, variadic...)
	for i, ele := range appended {
		fmt.Printf("%d) %s\n", i, ele)
	}

	fmt.Println("-----------")

	arr := []string{"amir", "mahkam", "ala"}
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	appendedArr := append(arr, "aria", "alan")

	fmt.Printf("arr --> %v\n", arr)
	fmt.Printf("appendedArr --> %v\n", appendedArr)

	fmt.Println("changing...")
	arr[0] = "senior.Amir"

	fmt.Printf("arr --> %v\n", arr)
	fmt.Printf("appendedArr --> %v\n", appendedArr)

	fmt.Println("-------try again for same underlying array--------")

	test := make([]string, 3, 6)
	test[0] = "amir"
	test[1] = "mahkam"
	test[2] = "ala"

	appendedTest := append(test, "aria", "alan") // backed by the same array.

	fmt.Println(test)
	fmt.Println(appendedTest)

	fmt.Println("changing...")
	test[0] = "senior.Amir"

	fmt.Println(test)
	fmt.Println(appendedTest)

	fmt.Println("---------focus--------")

	family := [4]string{"amir", "mahkam", "ala", "aria"}

	girls := family[1:3] // slicing a slice, makes the underlying array to be shared

	fmt.Printf("family --> %v\n", family)
	fmt.Printf("girls --> %v\n", girls)

	fmt.Println("changing...")
	family[1] = "Ms.mahkam"
	fmt.Printf("family --> %v\n", family)
	fmt.Printf("girls --> %v\n", girls)

	fmt.Println("----------------")

	fmt.Println(family)

	familySlice := family[1:3:3]
	fmt.Printf("first pointer %p\n", &familySlice)

	fmt.Println(cap(familySlice))
	fmt.Println(familySlice)
	fmt.Println("appending...")

	familySlice = append(familySlice, "another-baby", "another?")

	fmt.Println(familySlice)
	fmt.Printf("second pointer %p\n", &familySlice)

	fmt.Println("---------")

	oups := []string{"amir", "mahkam", "ala"}

	someOups := make([]string, 2)
	fmt.Printf("len --> %v and cap --> %v\n", len(someOups), cap(someOups))
	fmt.Println(oups)

	copy(someOups, oups)
	fmt.Println(someOups)

	fmt.Println("------------")

	var nilSlice []string
	fmt.Printf("the cap of nil --> %d\n", cap(nilSlice))

	emptySlice := make([]string, 0)
	fmt.Printf("the cap of empty --> %d\n", cap(emptySlice))

	fmt.Println(nilSlice)
	fmt.Println(emptySlice)

	fmt.Printf("nil --> %p\n", &nilSlice)
	fmt.Printf("empty --> %p\n", &emptySlice)

	fmt.Println("--------------")

	list := []string{"macbook", "mouse", "monitor", "keyboard"}
	fmt.Printf("init list memory address --> %p\n", list)
	littleList := []string{"book", "source"}

	copy(list, littleList)
	fmt.Printf("after copy. list --> %v\n", list)
	fmt.Printf("the memory address of list after copy --> %p\n", list)

	list = append(list, littleList...)
	fmt.Printf("after append. list --> %v\n", list)
	fmt.Printf("the memory address of list after append --> %p\n", list)

	fmt.Println("-----important------")

	num1 := 1
	num2 := 2
	num3 := 3

	var collection []*int
	collection = append(collection, &num1, &num2, &num3)

	for i, val := range collection {
		fmt.Printf("index %d --> value %d\n", i, *val)
	}

	fmt.Println("-----------")

	p1 := []string{"amir", "mahkam", "ala"}

	// p1ptr := (*[3]string)(p1)
	// for _, val := range p1ptr {
	// 	fmt.Println(val)
	// }
	arraye := [2]string(p1)
	fmt.Printf("the type is %T and the val is %v\n", arraye, arraye)

	fmt.Println("--------------")

	m1 := map[string]string{
		"husband": "amir",
		"wife":    "mahkam",
		"girl":    "ala",
		"boy":     "aria",
	}

	// comma ok approach for maps.
	value, ok := m1["boy"]
	if ok {
		fmt.Printf("Mr.%s exists\n", value)
	} else {
		fmt.Println("do not born")
	}

	fmt.Println("------------")

	fmt.Println(m1)
	delete(m1, "boy")
	fmt.Println(m1)

	if value, ok := m1["boy"]; ok {
		fmt.Printf("the value %v exists", value)
	} else {
		fmt.Printf("the value do not exists\n")
	}

	fmt.Println("------enumerating maps-------")

	for key, value := range m1 {
		fmt.Printf("key --> %s, value --> %s\n", key, value)
	}

	fmt.Println("-------------")

	var keys []string

	for key, _ := range m1 {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, value := range keys {
		fmt.Printf("key --> %s, value --> %s\n", value, m1[value])
	}

	fmt.Println("--------------")

	// strings are array of bytes --> so be careful when slicing them.
	var price string = "€48.95"

	// var currency byte = price[0]

	// amount, parseErr := strconv.ParseFloat(price[1:], 64)

	// fmt.Println("currency: ", string(currency))
	// if parseErr == nil {
	// 	fmt.Println("amount: ", amount)
	// } else {
	// 	fmt.Printf("not parsable, %v", parseErr)
	// }
	fmt.Println(len(price))

	fmt.Println("-----solution to parsing------")

	var rune_price []rune = []rune(price)

	var currency = rune_price[0]
	fmt.Printf("the type is %T and the value is %v\n", currency, currency)

	amount, parseErr := strconv.ParseFloat(string(rune_price[1:]), 64)

	if parseErr == nil {
		fmt.Printf("amount: %v\n", amount)
	} else {
		fmt.Printf("not parsable: %v\n", parseErr)
	}

	p := "$49.02"

	var curr byte = p[0] // byte is uint8
	fmt.Printf("the type is %T and value is %v\n", curr, string(curr))

	if amount, parseErr := strconv.ParseFloat(p[1:], 64); parseErr == nil {
		fmt.Printf("amount: %v\n", amount)
	} else {
		fmt.Println("not parsable. ", parseErr)
	}

	fmt.Println("----------------------")
	var r_s []rune = []rune("mahkam")
	fmt.Println(r_s)

	fmt.Println("========================================")

	str := "amir"
	fmt.Printf("type --> %T, value --> %v\n", []byte(str), []byte(str))
	fmt.Printf("type --> %T, value --> %v\n", []rune(str), []rune(str))
	var first_one []byte = []byte(str)
	var second_two []rune = []rune(str)

	fmt.Println(len(first_one))
	fmt.Println(len(second_two))

	fmt.Println("--------bytes and runes------------")

	s := "amir€ alaeifarwith euro by a ‹se›uhﬁ‡‹°"
	var b []byte = []byte(s)
	fmt.Printf("actual --> %v, string --> %s\n", b, string(b))

	var r []rune = []rune(s)

	var s_r []rune = []rune(s)
	fmt.Println(string(s_r[4]))
	var euro_r []rune = []rune(string(s_r[4]))
	fmt.Println(string(euro_r))
	fmt.Printf("actual --> %v, string --> %s\n", r, string(r))

	fmt.Println("------mocking--------")

	v := "€49.450"

	var v_r []rune = []rune(v)

	var v_b []byte = []byte(v)
	char_byte := string(v_b[0])

	fmt.Println("char dollar: ", char_byte)

	char := v_r[0]

	fmt.Println("char: ", string(char))
	a, parseErr := strconv.ParseFloat(string(v_r[1:]), 64)
	if parseErr == nil {
		fmt.Println("amount: ", a)
	} else {
		fmt.Println("not parsable: ", parseErr)
	}
}

// variadic function
func Hello(a ...string) string {
	for _, str := range a {
		fmt.Println(str)
	}
	return ""
}
