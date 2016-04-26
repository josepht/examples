package main

import (
	"fmt"
	"strings"
)

func comma_reverse(s string) string {
	n := len(s)

	if n <= 3 {
		return s
	}

	return s[:3] + "," + comma_reverse(s[3:])
}

func comma(s string) string {
	if strings.Contains(s, ".") {
		i := strings.Index(s, ".")
		return comma(s[:i]) + "." + comma_reverse(s[i+1:])
	}
	n := len(s)

	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	fmt.Println(comma("1234567890"))
	fmt.Println(comma("123456789"))
	fmt.Println(comma("12345678"))
	fmt.Println(comma("1234567"))
	fmt.Println(comma("123456"))

	fmt.Println(comma_reverse("1234567890"))
	fmt.Println(comma_reverse("123456789"))
	fmt.Println(comma_reverse("12345678"))
	fmt.Println(comma_reverse("1234567"))
	fmt.Println(comma_reverse("123456"))

	fmt.Println(comma("12345.67890"))
	fmt.Println(comma("12345.6789"))
	fmt.Println(comma("1234.5678"))
	fmt.Println(comma("12.34567"))

	fmt.Println(comma("1230.4560"))
}
