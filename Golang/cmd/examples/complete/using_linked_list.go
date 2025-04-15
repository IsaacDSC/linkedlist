package main

import (
	"fmt"
	"linkedlist/Golang/algoritms"
)

func main() {
	// Criar uma lista de inteiros
	lista := algoritms.LinkedList[int]{}
	fmt.Println("Lista criada:")
	lista.Debug()

	// Adicionar elementos
	fmt.Println("\nAdicionando elementos:")
	lista.Append(10)
	lista.Append(20)
	lista.Append(30)
	lista.Debug()

	// Adicionar no início
	fmt.Println("\nAdicionando 5 no início:")
	lista.Unshift(5)
	lista.Debug()

	// Adicionar em uma posição específica
	fmt.Println("\nAdicionando 15 na posição 2:")
	lista.AppendOnIndex(15, 2)
	lista.Debug()

	// Buscar elementos
	fmt.Println("\nBuscando elementos:")
	if index, found := lista.Search(15); found {
		fmt.Printf("Elemento 15 encontrado na posição %d\n", index)
	}

	// Buscar por índice
	fmt.Println("\nBuscando por índice:")
	if val, err := lista.FindByIndex(1); err == nil {
		fmt.Printf("Elemento no índice 1: %d\n", val)
	}

	// Remover elementos
	fmt.Println("\nRemovendo elemento 15:")
	lista.Delete(15)
	lista.Debug()

	// Testar remoção no início e no fim
	fmt.Println("\nRemovendo primeiro e último elementos:")
	lista.Delete(5)  // Remove o primeiro
	lista.Delete(30) // Remove o último
	lista.Debug()
}
