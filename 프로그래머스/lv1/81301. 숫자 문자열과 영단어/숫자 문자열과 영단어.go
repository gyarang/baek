import (
	"strconv"
	"unicode"
)

var textVal = [...]string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func isDigitText(s string) (bool, int) {
	for i, v := range textVal {
		if v == s {
			return true, i
		}
	}

	return false, 0
}

func isNum(c rune) bool {
	return unicode.IsNumber(c)
}

func solution(s string) int {
	resultStr := ""
	tempStr := ""
	strArr := []rune(s)

	for i := 0; i < len(s); i++ {
		if isNum(strArr[i]) {
			resultStr += string(strArr[i])
			continue
		}

		tempStr += string(strArr[i])
		isDigit, val := isDigitText(tempStr)
		if isDigit {
			resultStr += strconv.Itoa(val)
			tempStr = ""
		}
	}

	result, _ := strconv.Atoi(resultStr)

	return result
}