func isValidSudoku(board [][]byte) bool {

	// check for row
	for i := 0 ; i < 9 ; i++{
		visited := make(map[byte]bool)
		for j := 0 ; j < 9 ; j++{
			if(board[i][j] == '.'){
				continue
			} 
			if(visited[board[i][j]] == true){
				return false
			}
			visited[board[i][j]] = true
		}
	}

	// check for col
	for i := 0 ; i < 9; i++{
		visited := make(map[byte]bool)
		for j := range board{
			if(board[j][i] == '.'){
				continue
			} 
			if(visited[board[j][i]] == true){
				return false
			}
			visited[board[j][i]] = true
		}
	}


	// check for 3x3, specific formula
	for square := 0 ; square < 9 ; square++ { // there are 9 square
		visited := make(map[byte]bool)
		for i := 0 ; i < 3; i++{
			for j := 0 ; j < 3 ; j++{
			row := (square / 3) * 3 + i // traversing
			col := (square % 3) * 3 + j

			if(board[row][col] == '.'){
				continue
			} 
			if(visited[board[row][col]] == true){
				return false
			}
			visited[board[row][col]] = true
		}
		}
	}


	return true
}
