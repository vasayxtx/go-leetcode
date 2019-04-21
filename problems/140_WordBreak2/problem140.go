package problem140

import (
	"strings"
)

type Node struct {
	word      string
	children  []*Node
	sentences [][]string
}

func buildTree(node *Node, s string, wordDict []string, impossibleStrs map[string]bool) bool {
	if s == "" {
		return true
	}
	if impossibleStrs[s] {
		return false
	}

	for _, w := range wordDict {
		if strings.HasPrefix(s, w) {
			node.children = append(node.children, &Node{word: w})
		}
	}

	if len(node.children) == 0 {
		impossibleStrs[s] = true
		return false
	}

	result := false
	for _, child := range node.children {
		if buildTree(child, s[len(child.word):], wordDict, impossibleStrs) {
			result = true
			if len(child.sentences) > 0 {
				for _, s := range child.sentences {
					s := append(s, child.word)
					node.sentences = append(node.sentences, s)
				}
			} else {
				node.sentences = append(node.sentences, []string{child.word})
			}
		}
	}
	node.children = nil
	if result == false {
		impossibleStrs[s] = true
	}
	return result
}

func reverse(a []string) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func wordBreak(s string, wordDict []string) []string {
	root := &Node{word: ""}
	var result []string
	cache := make(map[string]bool)
	buildTree(root, s, wordDict, cache)
	for _, sen := range root.sentences {
		reverse(sen)
		result = append(result, strings.Join(sen, " "))
	}
	return result
}
