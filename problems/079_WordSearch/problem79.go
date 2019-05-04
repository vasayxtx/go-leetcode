package problem79

type cell struct {
	x, y int
}

func checkIfCellValid(c cell, board [][]byte) bool {
	if len(board) == 0 {
		return false
	}
	return c.x >= 0 && c.y >= 0 && c.x < len(board) && c.y < len(board[0])
}

func tryBuildWord(word string, board [][]byte, curCell cell) bool {
	if word == "" {
		return true
	}
	if !checkIfCellValid(curCell, board) || board[curCell.x][curCell.y] != word[0] {
		return false
	}

	word = word[1:]
	ch := board[curCell.x][curCell.y]
	board[curCell.x][curCell.y] = '-' // By condition word cat contain only letters
	ok := tryBuildWord(word, board, cell{curCell.x - 1, curCell.y}) ||
		tryBuildWord(word, board, cell{curCell.x, curCell.y - 1}) ||
		tryBuildWord(word, board, cell{curCell.x + 1, curCell.y}) ||
		tryBuildWord(word, board, cell{curCell.x, curCell.y + 1})
	board[curCell.x][curCell.y] = ch
	return ok
}

func exist(board [][]byte, word string) bool {
	for i, line := range board {
		for j := range line {
			if tryBuildWord(word, board, cell{i, j}) {
				return true
			}
		}
	}
	return false
}
