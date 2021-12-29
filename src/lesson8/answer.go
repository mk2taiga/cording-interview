package lesson8

// 8.1

func countWays(n int) int {
	memo := make([]int, 0, n+1)
	for i := range memo {
		memo[i] = -1
	}
	return countWaysByMemo(n, memo)
}

// NOTE: メモ化は再帰関数など、時間がかかる処理で得られた結果をメモに残す。
func countWaysByMemo(n int, memo []int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	} else if memo[n] > -1 {
		return memo[n]
	}

	memo[n] = countWaysByMemo(n-1, memo) + countWaysByMemo(n-2, memo) + countWaysByMemo(n-3, memo)
	return memo[n]
}

// 8.2
type Point struct {
	x, y int
}

func getPath(maze [][]bool) []Point {
	if len(maze) == 0 {
		return nil
	}

	path := []Point{}
	failed := map[Point]bool{}
	if getPathByMemo(maze, len(maze)-1, len(maze[0])-1, path, failed) {
		return path
	}
	return nil
}

// 道のりが見つかればいい(深さ優先探索すればいい)だけなので、成功はキャッシュしなくていい。
func getPathByMemo(maze [][]bool, row, col int, path []Point, failed map[Point]bool) bool {
	if row < 0 || col < 0 || !maze[row][col] {
		return false
	}

	p := Point{x: col, y: row}

	if failed[p] {
		return false
	}

	isOrigin := col == 0 && row == 0
	if isOrigin || getPathByMemo(maze, row-1, col, path, failed) || getPathByMemo(maze, row, col-1, path, failed) {
		path = append(path, p)
		return true
	}

	failed[p] = true
	return false
}

// 動的計画法のいろはを学べればよかったので、後の問題は一旦飛ばす。
// 再帰関数は、トップダウン、ボトムアップ、半々法がある。ボトムアップが一番シンプルなので、これを最初に考えるのが良い。
