package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 1},
	}
	islands := countIslands(matrix)
	fmt.Println("no:of islands:", islands)
}

func countIslands(matrix [][]int) int {
	islands := 0
	visited := make([][]bool, len(matrix))

	for i := range matrix {
		visited[i] = make([]bool, len(matrix[i]))
	}

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 1 && !visited[i][j] {
				dfs(matrix, &visited, i, j)
				islands++
			}
		}
	}
	return islands
}

func dfs(matrix [][]int, visited *[][]bool, i, j int) {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[i]) || (*visited)[i][j] || matrix[i][j] == 0 {
		return
	}
	(*visited)[i][j] = true
	dfs(matrix, visited, i+1, j)
	dfs(matrix, visited, i-1, j)
	dfs(matrix, visited, i, j+1)
	dfs(matrix, visited, i, j-1)

}
