package main

import "fmt"

type vertex struct {
	data     int
	adjacent []*vertex
}

type Graph struct {
	vertices []*vertex
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
		fmt.Println("error ,nil found")
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

func (g *Graph) print() {
	for _, vs := range g.vertices {
		fmt.Print("\n vertex:", vs.data, ":")
		for _, v := range vs.adjacent {
			fmt.Print(" ", v.data)
		}
	}
}

func (g *Graph) deleteVertex(data int) {
	for i, vertex := range g.vertices {
		if vertex.data == data {
			g.vertices = append(g.vertices[:i], g.vertices[i+1:]...)
			for _, vertexx := range g.vertices {
				vertexx.deleteAdjacent(data)
			}
		}
	}
}

func (v *vertex) deleteAdjacent(data int) {
	for i, adj := range v.adjacent {
		fmt.Println("----------------")
		fmt.Println("adj-data:", adj.data, "  and data is:", data)

		if adj.data == data {
			v.adjacent = append(v.adjacent[:i], v.adjacent[i+1:]...)
		}
	}
}

func (g *Graph) deleteEdge(u, v int) {
	from := g.getVertex(u)
	to := g.getVertex(v)
	if from == nil || to == nil {
		fmt.Println("errrrrrrrrrrrrrrrrrrrrrr")
		return
	}
	from.deleteAdjacent(v)
	to.deleteAdjacent(u)

}

func main() {
	var graph Graph
	graph.addVertex(1)
	graph.addVertex(2)
	graph.addVertex(3)
	graph.addVertex(4)
	graph.addVertex(5)

	graph.addEdge(1, 2)
	graph.addEdge(3, 2)
	graph.addEdge(3, 4)
	graph.addEdge(4, 2)
	graph.addEdge(5, 2)
	graph.addEdge(3, 5)

	graph.print()

	graph.deleteVertex(4)

	graph.print()

	graph.deleteEdge(5, 2)
	graph.print()

	fmt.Println("")

}
