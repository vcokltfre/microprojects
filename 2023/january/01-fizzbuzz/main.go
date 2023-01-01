package main

import "fmt"

type divisorMapping struct {
	divisor int
	word    string
}

var defaults = []divisorMapping{
	{3, "Fizz"},
	{5, "Buzz"},
}

func Fizzbuzz(n int, custom ...[]divisorMapping) {
	mapping := defaults
	if len(custom) > 0 {
		mapping = custom[0]
	}

	for i := 1; i <= n; i++ {
		output := ""
		for _, m := range mapping {
			if i%m.divisor == 0 {
				output += m.word
			}
		}
		if output == "" {
			output = fmt.Sprintf("%d", i)
		}
		fmt.Println(output)
	}
}

func main() {
	fmt.Println("Default Fizzbuzz:")
	Fizzbuzz(15)

	fmt.Println("\nCustom Fizzbuzz:")
	Fizzbuzz(15, []divisorMapping{
		{2, "Fizz"},
		{3, "Buzz"},
		{4, "Bazz"},
	})
}
