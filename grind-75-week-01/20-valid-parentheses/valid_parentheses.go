package valid_parentheses

type Stack struct {
	content []string
}

func (stack *Stack) push(s string) {
	stack.content = append(stack.content, s)
}

func (stack *Stack) pop() {
	n := len(stack.content) - 1
	stack.content = stack.content[:n]
}

func (stack *Stack) len() int {
	return len(stack.content)
}

func (stack *Stack) top() string {
	n := len(stack.content) - 1
	lastParentheses := stack.content[n]
	return lastParentheses
}

func isValid(s string) bool {
	var stack = Stack{}
	for i := 0; i < len(s); i++ {
		parentheses := string(s[i])
		if stack.len() != 0 {
			lastParentheses := stack.top()
			if (lastParentheses == "(" && parentheses == ")") ||
				(lastParentheses == "[" && parentheses == "]") ||
				(lastParentheses == "{" && parentheses == "}") {
				stack.pop()
				continue
			}
		}

		stack.push(parentheses)
	}

	return 0 == stack.len()
}
