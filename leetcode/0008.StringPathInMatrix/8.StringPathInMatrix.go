package leetcode

func existStringPathInMatrix(board [][]byte, word string) bool {
	if len(board) == 0 || len(word) == 0 {
		return false
	}

	checked := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		checked[i] = make([]bool, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if existPath(board, checked, i, j, word, 0) {
				return true
			}
		}
	}

	return false
}

func existPath(board [][]byte, checked [][]bool, r int, c int, word string, w int) bool {
	if board[r][c] == word[w] {
		if w == len(word)-1 {
			return true
		}
		checked[r][c] = true

		if r-1 >= 0 && !checked[r-1][c] && existPath(board, checked, r-1, c, word, w+1) {
			return true
		}
		if r+1 < len(board) && !checked[r+1][c] && existPath(board, checked, r+1, c, word, w+1) {
			return true
		}
		if c-1 >= 0 && !checked[r][c-1] && existPath(board, checked, r, c-1, word, w+1) {
			return true
		}
		if c+1 < len(board[0]) && !checked[r][c+1] && existPath(board, checked, r, c+1, word, w+1) {
			return true
		}

		checked[r][c] = false
	}

	return false
}
