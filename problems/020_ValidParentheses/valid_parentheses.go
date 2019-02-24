package validparentheses

// https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {
	brackets := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	var stack []rune
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
			continue
		}
		if ch == ')' || ch == ']' || ch == '}' {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if brackets[top] != ch {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
