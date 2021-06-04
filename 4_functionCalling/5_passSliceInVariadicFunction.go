package main

import (
	"fmt"
	"strings"
)

func joinStrs(element ...string) string {
	return strings.Join(element, "-")
}

func main() {
	element := []string{"geek", "for", "geek"}
	fmt.Println(joinStrs(element...))

}
