package data

import (
	"fmt"
)

func init() {
	fmt.Println("package data init function invoked!!")
}

func GetData() []string {
	return []string{"amir", "arman", "seyfi", "others"}
}
