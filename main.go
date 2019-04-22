package main

import (
	"fmt"

	"github.com/ofonimefrancis/cs-fundamentals/foundations/sorting"
)

func main() {
	data := []int{5, 10, 1, 3, 9, 2, 6}
	fmt.Println(sorting.Selection(data))
}
