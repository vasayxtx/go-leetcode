package problem22

func generateParenthesis(n int) []string {
	buffer := make([]byte, n*2)
	var res []string
	var generate func(int, int, int)
	generate = func(index, opensCount, closesCount int) {
		if index == n*2 {
			res = append(res, string(buffer))
			return
		}
		if opensCount < n {
			buffer[index] = '('
			generate(index+1, opensCount+1, closesCount)
		}
		if opensCount > closesCount {
			buffer[index] = ')'
			generate(index+1, opensCount, closesCount+1)
		}
	}
	generate(0, 0, 0)
	return res
}
