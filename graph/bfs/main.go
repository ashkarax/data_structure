package main

import (
	"fmt"
)

type vertex struct {
	data     int
	adjacent []*vertex
}

type Graph struct {
	vertices []*vertex
}

type Queue struct {
	arr []int
}

func (g *Graph) bfs(vertex int) {
	queue := Queue{}
	visited := make(map[int]bool)
	queue.arr = append(queue.arr, vertex)
	visited[vertex] = true
	for len(queue.arr) != 0 {
		vertex = queue.arr[0]
		queue.arr = queue.arr[1:]
		fmt.Println(vertex, "")
		for _, v := range g.getVertex(vertex).adjacent {
			if !visited[v.data] {
				visited[v.data] = true
				queue.arr = append(queue.arr, v.data)
			}
		}
	}
}

func (g *Graph) InsertVertex(data int) {
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
		fmt.Println("errrrrrrrrrrrrrrrr")
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
	for _, v := range g.vertices {
		fmt.Print("\nvertex ", v.data, ":")
		for _, vv := range v.adjacent {
			fmt.Print(" ", vv.data)
		}
	}
}

func main() {
	var graph Graph

	graph.InsertVertex(4)
	graph.InsertVertex(5)
	graph.InsertVertex(6)
	graph.InsertVertex(7)
	graph.InsertVertex(8)

	graph.addEdge(4, 5)
	graph.addEdge(5, 6)
	graph.addEdge(6, 7)
	graph.addEdge(7, 8)
	graph.addEdge(5, 8)

	graph.Display()
	fmt.Println("-------------------------")
	graph.bfs(8)

	fmt.Println(" ")

}
