package main

import "fmt"

type vertex struct {
	data     int
	adjacent []*vertex
}

type Graph struct {
	vertices []*vertex
}

type Stack struct {
	arr []int
}

func (g *Graph) dfs(vertex int) {
	stack := Stack{}
	visited := make(map[int]bool)
	stack.arr = append(stack.arr, vertex)
	visited[vertex] = true
	for len(stack.arr) != 0 {
		vertex = stack.arr[len(stack.arr)-1]
		stack.arr = stack.arr[:len(stack.arr)-1]
		fmt.Print(vertex, " ")
		for _, v := range g.getVertex(vertex).adjacent {
			if !visited[v.data] {
				visited[v.data] = true
				stack.arr = append(stack.arr, v.data)
			}
		}
	}

}

func (g *Graph) addVertex(data int) {
	if !g.contains(data) {
		g.vertices = append(g.vertices, &vertex{data: data})
	}
}

func (g *Graph) contains(data int) bool {
	for _, v := range g.vertices {
		if v.data == data {
			return true
		}
	}
	return false
}

func (g *Graph) addEdge(u, v int) {
	from := g.getVertex(u)
	to := g.getVertex(v)
	if from == nil || to == nil {
		fmt.Println("errrrrrrrrrrrrrrrrrrr")
		return
	}
	from.adjacent = append(from.adjacent, to)
	to.adjacent = append(to.adjacent, from)
}

func (g *Graph) getVertex(data int) *vertex {
	for _, v := range g.vertices {
		if v.data == data {
			return v
		}
	}
	return nil
}

func (g *Graph) Display() {
	for _, vx := range g.vertices {
		fmt.Print("\nvertex", vx.data, ":")
		for _, v := range vx.adjacent {
			fmt.Print(" ", v.data)
		}
	}
}

func main() {
	var graph Graph
	graph.addVertex(1)
	graph.addVertex(2)
	graph.addVertex(3)
	graph.addVertex(4)
	graph.addVertex(5)
	graph.addVertex(6)

	graph.addEdge(1, 2)
	graph.addEdge(2, 3)
	graph.addEdge(3, 4)
	graph.addEdge(4, 5)
	graph.addEdge(5, 6)
	graph.addEdge(1, 6)
	graph.addEdge(1, 5)

	graph.Display()

	fmt.Println("\n---------------\n")
	graph.dfs(1)

	fmt.Println(" ")
}
