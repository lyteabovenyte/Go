package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	product := "kayak"

	fmt.Println(strings.IndexRune(product, 'a'))
	fmt.Println(strings.IndexAny(product, "abc"))

	fmt.Println(strings.IndexAny(product, "abc"))

	for _, r := range product {
		fmt.Printf("rune %c\n", r)
		fmt.Println("actual r-->", r)
	}

	rp := "你好，世界!"

	fmt.Println("using len func-->", len(rp)) // it's a rune.

	fmt.Println("using rune len func-->", utf8.RuneCountInString(rp))

	fmt.Println("-----------------------")

	for _, ch := range product {
		fmt.Println(string(ch), "--> to upper with rune -->", string(unicode.ToUpper(ch)))
	}

	fmt.Println("---------------")

	s := "  Amir Alaeifar  "

	fmt.Println(strings.IndexByte(s, 'M'))

	num := "۱۲۳۴"

	dict := map[rune]int{
		'۱': 1,
		'۲': 2,
		'۳': 3,
		'۴': 4,
	}
	for _, n := range num {
		fmt.Println(dict[n])
	}

	fmt.Println("-------XXX--------")

	num1 := []rune(num)

	fmt.Println(string(num1))
	fmt.Printf("type of num1 --> %T", num1) // rune is an alias of int32

	var convertor = map[rune]rune{}

	for _, runeValue := range num1 {
		convertor[runeValue] = runeValue
	}

	fmt.Printf("value --> %c\n", convertor)

	fmt.Println("-----------------")

	result := strings.Map(func(r rune) rune {
		if r == '۱' {
			return '1'
		}
		if r == '۲' {
			return '2'
		}
		if r == '۳' {
			return '3'
		}
		if r == '۴' {
			return '4'
		}
		return r
	}, num)

	fmt.Println(result)

	fmt.Println("-----------------")

	str := "  Amir Alaeifar is here :)  "

	fmt.Println(strings.Fields(str))

	splits := strings.Split(str, " ")
	for _, ch := range splits {
		fmt.Println("split >>", ch, "<<")
	}

	fmt.Println("--------------------------")

	words := map[int]string{}

	i := 1

	fields := strings.Fields(str)
	fmt.Println("fields -->", fields)

	for _, x := range fields {
		words[i] = x
		i++
	}

	fmt.Println(words)
	fmt.Println("---------------")

	jomle := "this  is  double  space"

	field := strings.Fields(jomle)

	for _, x := range field {
		fmt.Println(x)
	}

	fmt.Println("----------------")

	split := strings.SplitN(jomle, "  ", 4)
	for _, y := range split {
		fmt.Println(y)
	}

	fmt.Println("----------------")

	fmt.Println("s is -->", s)

	s = "Amir Alaeifar"

	fmt.Println(".", strings.TrimSpace(s), ".")

	sufFunc := func(r rune) bool {
		return r == 'r'
	}

	fmt.Println(strings.TrimFunc(s, sufFunc))
	fmt.Println(strings.TrimRight(s, "ar"))
	fmt.Println(strings.TrimSuffix(s, "ar"))
	fmt.Println(strings.Trim(s, "Aar"))

	fmt.Println("---------------")

	replacer := strings.NewReplacer("Amir", "amir", "Alaeifar", "alaeifar")

	newS := replacer.Replace(s)
	fmt.Println(newS)

	fmt.Println("--------------")

	mapper := func(r rune) rune {
		if r == '۱' {
			return '1'
		}
		return r
	}

	mapString := "amir۱"
	fmt.Println(strings.Map(mapper, mapString))

	fmt.Println("--------------")

	persian_nums := "۱۲۳۴"
	mp := map[rune]rune{
		'۱': '1',
		'۲': '2',
		'۳': '3',
		'۴': '4',
		'۵': '5',
		'۶': '6',
		'۷': '7',
		'۸': '8',
		'۹': '9',
		'۰': '0',
	}

	converted := strings.Map(func(r rune) rune {
		if c, ok := mp[r]; ok {
			return c
		} else {
			return r
		}
	}, persian_nums)

	fmt.Println(converted)

	fmt.Println("-------------")

	name := "Amir Alaeifar"

	// this can be efficient in using CSV type speadsheets
	split = strings.Fields(name)
	fmt.Printf("the type of split is %T\n", split) // []string
	joined := strings.Join(split, ",")
	fmt.Println(joined)

	fmt.Println("--------------")

	t := "mahkam is the most beautiful one"
	var builder strings.Builder
	fmt.Println(ChanginStrings(t, &builder))

	fmt.Println("-----------regexp----------")

	text := "amir alaeifar namir ta vaghti naresidi :)"

	match, err := regexp.MatchString("[a-z]mir", text)
	if err != nil {
		fmt.Println("err:", err)
	}
	if err == nil {
		fmt.Println(match)
	}

	pattern, err := regexp.Compile("[A-z]+mir")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("-->", pattern.MatchString(text))

	fmt.Println("pattern FindAllStrings -->", pattern.FindAllString(text, -1))
	splitted := pattern.Split(text, -1)

	for _, x := range splitted {
		fmt.Println(x)
	}

	fmt.Println("--------------------")

	t3 := "Kayak. A boat for everyone"

	patent := regexp.MustCompile("K[A-z]{4}|[A-z]oat")
	firstIndex := patent.FindStringIndex(t3) // first occurance
	allIndices := patent.FindAllIndex([]byte(t3), -1)
	fmt.Println("first index:", getSubString(t3, firstIndex))

	// all indices
	for x, idx := range allIndices {
		fmt.Printf("%dth word: %s\n", x, getSubString(t3, idx))
	}

	fmt.Println("------------")

	t4 := "hey, a Book came from skybooks"
	p := regexp.MustCompile("[A-z]*(B|b)[A-z]{2}k[A-z]{0,2}")
	all := p.FindAllIndex([]byte(t4), -1)
	one := p.FindStringIndex(t4)
	fmt.Println(one)
	fmt.Println(getSubString(t4, one))

	for _, idx := range all {
		fmt.Println(getSubString(t4, idx))
	}

	fmt.Println("FindString and FindAllString methods")
	fmt.Println(p.FindString(t4))
	fmt.Println(p.FindAllString(t4, -1))

	fmt.Println("-------------")

	sp := regexp.MustCompile(`,`)
	sps := sp.Split(t4, -1)
	for _, s := range sps {
		fmt.Println("subString:", strings.TrimSpace(s))
	}

	fmt.Println("==========XXXXX==========")

	t5 := `ddia from skybooks
	pipeline from amazon
	`

	// TODO!!
	pattern1 := regexp.MustCompile(`(?P<Name>[A-z]*) from (?P<Source>[A-z]*)`)
	subs := pattern1.FindAllStringSubmatch(t5, -1)

	for _, sub := range subs {
		for _, cat := range []string{"Name", "Source"} {
			fmt.Println(cat, "=", sub[pattern1.SubexpIndex(cat)])
		}
	}

	// for _, sub := range subs {
	// 	fmt.Println("sub:", sub)
	// }

}

// pattern:
// first create a builder
// second Write within the builder as many as you want with arbitary character and strings
// third get the string out of builder with String() method
func ChanginStrings(text string, b *strings.Builder) string {
	replacer := strings.NewReplacer("is", "are", "one", "girls")
	new_text := replacer.Replace(text)
	for _, sub := range strings.Fields(new_text) {
		if sub == "mahkam" {
			b.WriteString(sub)
			b.WriteString(" and ala")
		} else {
			b.WriteString(sub)
		}
		b.WriteRune(' ')
	}
	return b.String()
}

// get acutal string from the indices given by regex pattern
func getSubString(s string, indices []int) string {
	return string(s[indices[0]:indices[1]]) // getting actual word
}
