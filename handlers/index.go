package handlers

import (
	"big-integers-calculator/operations/numbers"
	"big-integers-calculator/operations/polynomials"
	"html/template"
	"net/http"
	"regexp"
	"strings"
)

const MULTIPLY_NUMBERS string = "on"

type Data struct {
	Input  string
	Result []int
}

func IndexGetHandler(writer http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("html/index.html"))
	template.Execute(writer, nil)
}

func IndexPostHandler(writer http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("html/index.html"))
	request.ParseForm()
	input := request.FormValue("expression")

	var res []int
	validateInput(input)
	left, right := parse(input)
	poly1, poly2 := createPolys(left, right)
	if request.FormValue("multiplyNumbers") == MULTIPLY_NUMBERS {
		fillPolys(poly1, poly2, left, right, true)
		res = numbers.Multiply(poly1, poly2)
	} else {
		fillPolys(poly1, poly2, left, right, false)
		res = polynomials.Multiply(poly1, poly2)
	}

	data := Data{
		Input:  request.FormValue("expression"),
		Result: res,
	}
	template.Execute(writer, data)
}

func validateInput(input string) {
	pattern := `^\d+\*\d+$`
	correctInput, _ := regexp.Match(pattern, []byte(input))
	if !correctInput {
		panic("Incorrect input!")
	}
}

func parse(input string) (left, right string) {
	data := strings.Split(input, "*")
	left, right = data[0], data[1]
	return left, right
}

func createPolys(left, right string) (poly, otherPoly []complex128) {
	size := getSize(len(left), len(right))
	poly = make([]complex128, size)
	otherPoly = make([]complex128, size)
	return poly, otherPoly
}

func getGreaterLen(len1, len2 int) int {
	if len1 > len2 {
		return len1
	}
	return len2
}

func getSize(len1, len2 int) int {
	greaterLen := getGreaterLen(len1, len2)
	size := 1
	for size < greaterLen+1 {
		size <<= 1
	}
	size <<= 1
	return size
}

func fillPolys(poly, otherPoly []complex128, left, right string, mulNumbers bool) {
	var index int
	for i := 0; i < len(left); i++ {
		if mulNumbers {
			index = len(left) - 1 - i
		} else {
			index = i
		}
		poly[i] = complex(float64(rune(left[index])-'0'), 0)
	}

	for i := 0; i < len(right); i++ {
		if mulNumbers {
			index = len(right) - 1 - i
		} else {
			index = i
		}
		otherPoly[i] = complex(float64(rune(right[index])-'0'), 0)
	}
}
