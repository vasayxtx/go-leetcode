package problem207

func checkPossibilityToFinish(cDeps [][]int, c1, c2 int) bool {
	for _, c := range cDeps[c2] {
		if c == c1 || !checkPossibilityToFinish(cDeps, c1, c) {
			return false
		}
	}
	return true
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	cDeps := make([][]int, numCourses)
	for _, p := range prerequisites {
		c1, c2 := p[0], p[1]
		if !checkPossibilityToFinish(cDeps, c1, c2) {
			return false
		}
		cDeps[c1] = append(cDeps[c1], c2)
	}
	return true
}
