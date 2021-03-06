package handlers

import (
	"big-integers-calculator/cmd/operations/numbers"
	"big-integers-calculator/cmd/operations/polynomials"
	"big-integers-calculator/cmd/types"
	"errors"
	"html/template"
	"net/http"
	"regexp"
	"strings"
)

const (
	MULTIPLY_NUMBERS    string = "on"
	INDEX_PATH          string = "assets/html/index.html"
	HTML_INPUT_NAME     string = "expression"
	HTML_MULTIPLY       string = "multiplyNumbers"
	INCORRECT_INPUT_MSG string = "incorrect input"
)

func IndexGetHandler(writer http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles(INDEX_PATH))
	template.Execute(writer, nil)
}

func IndexPostHandler(writer http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles(INDEX_PATH))
	request.ParseForm()
	var data types.Data = types.Data{
		Input: strings.TrimSpace(request.FormValue(HTML_INPUT_NAME)),
	}

	if ValidateInput(data.Input) {
		Multiply(&data, request)
	} else {
		data.Error = errors.New(INCORRECT_INPUT_MSG)
	}
	template.Execute(writer, data)
}

func ValidateInput(input string) (valid bool) {
	pattern := `^\d+\*\d+$`
	valid, _ = regexp.Match(pattern, []byte(input))
	return valid
}

func Multiply(data *types.Data, request *http.Request) {
	left, right := parse(data.Input)
	poly1, poly2 := createPolys(left, right)

	if request.FormValue(HTML_MULTIPLY) == MULTIPLY_NUMBERS {
		fillNumber(poly1, left)
		fillNumber(poly2, right)
		res := numbers.Multiply(poly1, poly2)
		data.Result = res.Trim().String()
	} else {
		fillPoly(poly1, left)
		fillPoly(poly2, right)
		res := polynomials.Multiply(poly1, poly2)
		data.Result = res.Trim().String()
	}
}

func parse(input string) (left, right string) {
	delimeter := "*"
	data := strings.Split(input, delimeter)
	left, right = data[0], data[1]
	return left, right
}

func createPolys(left, right string) (poly, otherPoly []complex128) {
	size := getSize(len(left), len(right))
	poly = make([]complex128, size)
	otherPoly = make([]complex128, size)
	return poly, otherPoly
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

func getGreaterLen(len1, len2 int) int {
	if len1 > len2 {
		return len1
	}
	return len2
}

func fillPoly(poly []complex128, data string) {
	dataSize := len(data)
	for i := 0; i < dataSize; i++ {
		poly[i] = complex(float64(rune(data[i])-'0'), 0)
	}
}

func fillNumber(number []complex128, data string) {
	dataSize := len(data)
	for i := 0; i < dataSize; i++ {
		number[i] = complex(float64(rune(data[dataSize-1-i])-'0'), 0)
	}
}
