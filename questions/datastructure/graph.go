package datastructure

import "fmt"

// Graph interface
type Graph interface {
	AddVertex(interface{})
	AddEdge(interface{}, interface{})
	RemoveVertex(interface{})
	RemoveEdge(interface{}, interface{})
}

// Vertex holds vertex value and neighbor adjacency list
type Vertex struct {
	value     interface{}
	neighbors map[interface{}]*Vertex
}

// UndirectedGraph is a map of vertices
type UndirectedGraph struct {
	vertices map[interface{}]*Vertex
}

// NewUndirectedGraph returns a new undericted graph
func NewUndirectedGraph() *UndirectedGraph {
	return &UndirectedGraph{make(map[interface{}]*Vertex)}
}

// AddVertex inserts a new vertex
func (g *UndirectedGraph) AddVertex(val interface{}) {
	g.vertices[val] = &Vertex{val, make(map[interface{}]*Vertex)}
}

// RemoveVertex deletes a vertex
func (g *UndirectedGraph) RemoveVertex(val interface{}) {
	if _, ok := g.vertices[val]; ok {
		delete(g.vertices, val)
		for _, vertex := range g.vertices {
			if _, ok := vertex.neighbors[val]; ok {
				delete(vertex.neighbors, val)
			}
		}
	}
}

// AddEdge adds an edge for two vertices
func (g *UndirectedGraph) AddEdge(src, dest interface{}) {
	if _, ok := g.vertices[src]; ok {
		if _, ok := g.vertices[dest]; ok {
			g.vertices[src].neighbors[dest] = g.vertices[dest]
			g.vertices[dest].neighbors[src] = g.vertices[src]
		}
	}
}

// RemoveEdge removes an edge of two vertices
func (g *UndirectedGraph) RemoveEdge(src, dest interface{}) {
	if _, ok := g.vertices[src]; ok {
		if _, ok := g.vertices[dest]; ok {
			delete(g.vertices[src].neighbors, dest)
			delete(g.vertices[dest].neighbors, src)
		}
	}
}

func (g *UndirectedGraph) String() string {
	str := "[ "

	for k, v := range g.vertices {
		str += fmt.Sprintf("'%c:", k)

		for e := range v.neighbors {
			str += fmt.Sprintf(" %c", e)
		}

		str += "', "
	}

	return str[:len(str)-2] + " ]"
}
