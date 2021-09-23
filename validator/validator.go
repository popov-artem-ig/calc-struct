package validator

import (
	"strconv"
)

func Validate() {
	

}

func Is(s string) bool {

	//regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

    
    return true
}

func isCheckingBrackets(str string) bool {

	var count = 0
    for i := 0; i < len(str); i++ {
        if str[i] == '(' {
            count++
		} else if str[i] == ')' {
            if count < 0 {
                return false
			}
            count--
		}
	}

    if count == 0 {
        return true
	}
    return false
}

func isArithmeticSymbol(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}


func isInt(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}