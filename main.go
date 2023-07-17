package main

import "fmt"

type Node struct {
	Value string
	Child []*Node
}

type Tree struct {
	start, top *Node
}

// Recorre el árbol de forma infija (in-order)
func inorderTraversal(node *Node) {
	if node == nil {
		return
	}

	if len(node.Child) > 0 {
		for _, child := range node.Child {
			inorderTraversal(child) // Recorre los otros hijos del nodo actual
		}
	} else {
		fmt.Printf("%s ", node.Value)
	}
}

func main() {
	// Raiz del arbol
	root := &Node{
		Value: "S",
		Child: []*Node{
			{
				Value: "M",
				Child: []*Node{
					{
						Value: "x",
						Child: nil,
					},
					{
						Value: "Z",
						Child: []*Node{
							{
								Value: "b",
								Child: nil,
							},
							{
								Value: "a",
								Child: nil,
							},
						},
					},
					{
						Value: "b",
						Child: nil,
					},
				},
			},
			{
				Value: "N",
				Child: []*Node{
					{
						Value: "b",
						Child: nil,
					},
				},
			},
			{
				Value: "Z",
				Child: []*Node{
					{
						Value: "X",
						Child: []*Node{
							{
								Value: "a",
								Child: nil,
							},
						},
					},
					{
						Value: "N",
						Child: []*Node{
							{
								Value: "a",
								Child: nil,
							},
							{
								Value: "b",
								Child: nil,
							},
							{
								Value: "x",
								Child: nil,
							},
						},
					},
				},
			},
		},
	}

	// Realizar la búsqueda infija en el árbol
	fmt.Print("Búsqueda")
	inorderTraversal(root)
	fmt.Println()
}
