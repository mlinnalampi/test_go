package main
import "strconv"
import "fmt"
func ordinal(n int) string {
	suffix := "th"
	switch n {
	case 1:
		suffix = "st"
	case 2:
		suffix = "nd"
	case 3:
		suffix = "rd"
}
return strconv.Itoa(n) + suffix
}

func main() {
	fmt.Println(ordinal(5))
}
