package rpn

import (
	"errors"
	"fmt"
	"strconv"
)

func Calc(expression string) (float64, error) {
	tokens := make([]string, 0)
	currentToken := ""
	for _, char := range expression {
		if char == ' ' {
			continue
		}
		if char == '+' || char == '-' || char == '*' || char == '/' || char == '(' || char == ')' {
			if currentToken != "" {
				tokens = append(tokens, currentToken)
				currentToken = ""
			}
			tokens = append(tokens, string(char))
		} else {
			currentToken += string(char)
		}
	}
	if currentToken != "" {
		tokens = append(tokens, currentToken)
	}

	postfix := make([]string, 0)
	operators := make([]string, 0)

	for _, token := range tokens {
		if isNumber(token) {
			postfix = append(postfix, token)
		} else if token == "(" {
			operators = append(operators, token)
		} else if token == ")" {
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				postfix = append(postfix, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			if len(operators) == 0 {
				return 0, errors.New("Проблема со скобками")
			}
			operators = operators[:len(operators)-1]
		} else if isOperator(token) {
			for len(operators) > 0 && priority(operators[len(operators)-1]) >= priority(token) {
				postfix = append(postfix, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
		} else {
			return 0, fmt.Errorf("Неправильный ввод")
		}
	}

	for len(operators) > 0 {
		if operators[len(operators)-1] == "(" {
			return 0, errors.New("Проблема со скобками")
		}
		postfix = append(postfix, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	stack := make([]float64, 0)

	for _, token := range postfix {
		if isNumber(token) {
			num, _ := strconv.ParseFloat(token, 64)
			stack = append(stack, num)
		} else if isOperator(token) {
			if len(stack) < 2 {
				return 0, errors.New("Неверное выражение")
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					return 0, errors.New("Делить на ноль нельзя")
				}
				stack = append(stack, a/b)
			default:
				return 0, fmt.Errorf("Неизвестный оператор: %s", token)
			}
		} else {
			return 0, fmt.Errorf("Неизвестный оператор: %s", token)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("Неверное выражение")
	}

	return stack[0], nil
}

func tokenize(expr string) []string {
	tokens := make([]string, 0)
	currentToken := ""
	for _, char := range expr {
		if char == ' ' {
			continue
		}
		if char == '+' || char == '-' || char == '*' || char == '/' || char == '(' || char == ')' {
			if currentToken != "" {
				tokens = append(tokens, currentToken)
				currentToken = ""
			}
			tokens = append(tokens, string(char))
		} else {
			currentToken += string(char)
		}
	}
	if currentToken != "" {
		tokens = append(tokens, currentToken)
	}
	return tokens
}

func infixToPostfix(tokens []string) ([]string, error) {
	output := make([]string, 0)
	operators := make([]string, 0)
	for _, token := range tokens {
		if isNumber(token) {
			output = append(output, token)
		} else if token == "(" {
			operators = append(operators, token)
		} else if token == ")" {
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			if len(operators) == 0 {
				return nil, errors.New("Проблема со скобками")
			}
			operators = operators[:len(operators)-1]
		} else if isOperator(token) {
			for len(operators) > 0 && priority(operators[len(operators)-1]) >= priority(token) {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
		} else {
			return nil, fmt.Errorf("Неверное значение")
		}
	}
	for len(operators) > 0 {
		if operators[len(operators)-1] == "(" {
			return nil, errors.New("Проблема со скобками")
		}
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}
	return output, nil
}

func evaluatePostfix(postfix []string) (float64, error) {
	stack := make([]float64, 0)
	for _, token := range postfix {
		if isNumber(token) {
			num, _ := strconv.ParseFloat(token, 64)
			stack = append(stack, num)
		} else if isOperator(token) {
			if len(stack) < 2 {
				return 0, errors.New("Неверное выражение")
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					return 0, errors.New("Делить на ноль нельзя")
				}
				stack = append(stack, a/b)
			default:
				return 0, fmt.Errorf("Неизвестный оператор: %s", token)
			}
		} else {
			return 0, fmt.Errorf("Неизвестный оператор: %s", token)
		}
	}
	if len(stack) != 1 {
		return 0, errors.New("Неверное выражение")
	}
	return stack[0], nil
}

func isNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

func priority(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}
