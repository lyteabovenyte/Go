package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	r := regexp.MustCompile(`/([0-9]+)`)

	s := "amir/5fateme/44/33another/1"

	g := r.FindAllStringSubmatch(s, -1)

	fmt.Println(g)
	fmt.Println(len(g))
	fmt.Println(g[0])
	fmt.Println(len(g[0]))

	idString := g[0][1]
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println("cannot convert the id")
	}
	fmt.Println("id -->", id)

}
