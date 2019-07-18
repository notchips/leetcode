/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 */
func solveSudoku(board [][]byte) {
	rows, columns, squares := getSet(), getSet(), getSet()
	// 标记所有存在的点
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			c := board[i][j]
			if c >= '0' && c <= '9' {
				rows[i][c-'0'] = true            // 标记第i行有c                0  1  2
				columns[j][c-'0'] = true         // 标记第j列有c                3  4  5 
				squares[i-i%3+j/3][c-'0'] = true // 标记第i-i%3+j/3个矩形有c:   6  7  8
			}
		}
	}
	exit := make(chan struct{})
	go fillNums(board, rows, columns, squares, exit)
	<-exit
}

type point struct {
	validNums []int
	i, j      int
}

func fillNums(board [][]byte, rows, columns, squares [][]bool, exit chan struct{}) {
	// 找到可以填最少数字的点
	p := point{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			c := board[i][j]
			if c == '.' {
				validNums := getValidNums(rows[i], columns[j], squares[i-i%3+j/3])
				if p.validNums == nil || len(validNums) < len(p.validNums) {
					p.validNums = validNums
					p.i, p.j = i, j
				}
			}
		}
	}

	if p.validNums == nil { // 表已填完，中止程序
		close(exit)
		time.Sleep(time.Millisecond * 1000)
	}

	for _, validNum := range p.validNums { // 回溯填表
		board[p.i][p.j] = byte(validNum) + '0'
		rows[p.i][validNum] = true
		columns[p.j][validNum] = true
		squares[p.i-p.i%3+p.j/3][validNum] = true

		fillNums(board, rows, columns, squares, exit)

		board[p.i][p.j] = '.'
		rows[p.i][validNum] = false
		columns[p.j][validNum] = false
		squares[p.i-p.i%3+p.j/3][validNum] = false
	}
}

// 获取空位能填的数字
func getValidNums(row, column, square []bool) []int {
	ret := make([]int, 0, 9)
	for i := 1; i <= 9; i++ {
		if !row[i] && !column[i] && !square[i] {
			ret = append(ret, i)
		}
	}
	return ret
}

func getSet() [][]bool {
	set := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		set[i] = make([]bool, 10)
	}
	return set
}
