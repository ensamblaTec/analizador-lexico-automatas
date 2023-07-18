package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/g-code99/arboles/analizador-lexico-automatas/utils"
)

// Estructura del nodo
type Node struct {
	Value string
	Child []*Node
}

// Estructura del arbol
type Tree struct {
	start, top *Node
}

var valid bool = true
var cnt int
var input []byte

// Recorre el árbol de forma infija (in-order) (modificada)
func inorderTraversal(node *Node) {
	if node == nil {
		return
	}

	if len(node.Child) > 0 {
		for _, child := range node.Child {
			inorderTraversal(child) // Recorre los otros hijos del nodo actual
		}
	} else {
		// Verificar si la entrada es correcta
		if cnt < len(input) && input[cnt] != []byte(node.Value)[0] {
			valid = false
		}
		fmt.Printf("%s ", node.Value)
		cnt++
	}
}

// Método para crear el árbol
func createLexTree(data string) *Node {
	lines := strings.Split(data, "\n")
	nodesMap := make(map[string]*Node)

	// Crear todos los nodos sin hijos
	for _, line := range lines {
		parts := strings.Split(line, "|")
		value := parts[0]
		nodesMap[value] = &Node{Value: value}
	}

	// Agregar los hijos a los nodos
	for _, line := range lines {
		parts := strings.Split(line, "|")
		value := parts[0]
		childrenData := parts[1]

		if len(childrenData) > 0 {
			children := strings.Split(childrenData, ",")
			for _, childValue := range children {
				childNode, exists := nodesMap[childValue]
				if exists {
					nodesMap[value].Child = append(nodesMap[value].Child, childNode)
				} else {
					// Crear un nuevo nodo hoja y agregarlo como hijo
					newChildNode := &Node{Value: childValue}
					nodesMap[value].Child = append(nodesMap[value].Child, newChildNode)
				}
			}
		}
	}

	return nodesMap["S"]
}

// Método para imprimir el árbol en forma identada
func printLex(node *Node, depth int) {
	if node == nil {
		return
	}

	fmt.Printf("%sValue: %s\n", strings.Repeat("  ", depth), node.Value)
	for _, child := range node.Child {
		printLex(child, depth+1)
	}
}

func main() {
	// Obtener los datos del documento
	data := strings.Trim(utils.ReadFile("ejer1.txt"), "\n")
	// Crear el arbol
	lex := createLexTree(data)

	// Imprimir el arbol
	fmt.Println("Arbol léxico creado:")
	printLex(lex, 0)

	// pedir datos
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Ingresa la cadena a reconocer: ")
	scanner.Scan()
	input = scanner.Bytes()

	// Realizar la búsqueda en el árbol
	fmt.Print("Búsqueda: ")
	inorderTraversal(lex)
	fmt.Println()
	fmt.Printf("Expresión válida: %v\n", valid)
}

// // Raiz del arbol
// root := &Node{
// 	Value: "S",
// 	Child: []*Node{
// 		{
// 			Value: "M",
// 			Child: []*Node{
// 				{
// 					Value: "x",
// 					Child: nil,
// 				},
// 				{
// 					Value: "Z",
// 					Child: []*Node{
// 						{
// 							Value: "b",
// 							Child: nil,
// 						},
// 						{
// 							Value: "a",
// 							Child: nil,
// 						},
// 					},
// 				},
// 				{
// 					Value: "b",
// 					Child: nil,
// 				},
// 			},
// 		},
// 		{
// 			Value: "N",
// 			Child: []*Node{
// 				{
// 					Value: "b",
// 					Child: nil,
// 				},
// 			},
// 		},
// 		{
// 			Value: "Z",
// 			Child: []*Node{
// 				{
// 					Value: "X",
// 					Child: []*Node{
// 						{
// 							Value: "a",
// 							Child: nil,
// 						},
// 					},
// 				},
// 				{
// 					Value: "N",
// 					Child: []*Node{
// 						{
// 							Value: "a",
// 							Child: nil,
// 						},
// 						{
// 							Value: "b",
// 							Child: nil,
// 						},
// 						{
// 							Value: "x",
// 							Child: nil,
// 						},
// 					},
// 				},
// 			},
// 		},
// 	},
// }

// // fmt.Println(root)
