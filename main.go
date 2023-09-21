package main

import (
	"fmt"

	"github.com/satria-technology/gommon/utils"
)

func main() {
	needle := "coba"
	haystack := []string{""}
	result := utils.StringInSliceOfStrings(needle, haystack)
	fmt.Println(result)
}
