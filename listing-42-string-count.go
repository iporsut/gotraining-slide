package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "สวัสดีครับ"
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))

	fmt.Println(utf8.RuneCountInString("👨‍👩‍👦‍👦"))
}
