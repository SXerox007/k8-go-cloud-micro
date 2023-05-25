package samples

func MaxDepths(s string) int {
	max := 0
	depth := 0

	for _, c := range s {
		if c == '(' {
			depth++
			if depth > max {
				max = depth
			}
		} else if c == ')' {
			depth--
		}
	}

	return max
}

func MaxDepthAdvances(s string) int {
	var stack []rune
	var depth, maxDepth int

	for _, ch := range s {
		if ch == '(' {
			stack = append(stack, ch)
			depth++
			if depth > maxDepth {
				maxDepth = depth
			}
		} else if ch == ')' {
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
				depth--
			} else {
				return -1 // invalid expression
			}
		}
	}

	if len(stack) > 0 {
		return -1 // invalid expression
	}

	return maxDepth
}
