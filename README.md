# Lista Encadeada - Estudo e Comparativo

## O que é uma Lista Encadeada?

Uma lista encadeada é uma estrutura de dados linear composta por uma sequência de elementos chamados "nós", onde cada nó contém:
1. Um valor de dado
2. Um ponteiro (ou referência) para o próximo nó na sequência

![Estrutura de Lista Encadeada](https://upload.wikimedia.org/wikipedia/commons/6/6d/Singly-linked-list.svg)

Diferentemente de arrays ou slices, que armazenam elementos em posições contíguas de memória, as listas encadeadas podem usar memória não-contígua, com cada nó podendo estar em qualquer local da memória.

### Características Principais:
- **Inserção e remoção eficientes**: Especialmente no início da lista (O(1))
- **Acesso sequencial**: Para acessar um elemento, é necessário percorrer a lista desde o início
- **Tamanho dinâmico**: Cresce conforme necessário, limitado apenas pela memória disponível
- **Uso de memória**: Cada nó requer espaço adicional para armazenar o ponteiro

## Sobre este Projeto

Este projeto visa:
1. Demonstrar a implementação de listas encadeadas em diferentes linguagens (Go e Java)
2. Comparar o desempenho entre listas encadeadas e estruturas nativas (slices em Go)
3. Servir como material de estudo para estruturas de dados

## Implementações

### Golang

A implementação em Go usa generics para criar uma lista encadeada que pode armazenar qualquer tipo de dado:

```go
package main

import (
	"fmt"
	"linkedlist/Golang/algoritms"
)

func main() {
	list := algoritms.LinkedList[int]{}

	list.Append(10)
	list.Unshift(5)
	list.AppendOnIndex(7, 1)

	list.Debug() // 5 -> 7 -> 10 -> nil

	index, isFound := list.Search(7)
	fmt.Println("Índice do 7:", isFound, index) // 1

	list.Delete(7)
	list.Debug() // 5 -> 10 -> nil
}


// Métodos implementados:
// - Unshift: Adiciona elemento no início
// - Append: Adiciona elemento no final
// - AppendOnIndex: Adiciona elemento em posição específica
// - Search: Busca por um elemento
// - Delete: Remove um elemento
```

### Java

```java
import algoritms.LinkedList;

class HelloWorld {
  public static void main(String[] args) {
    LinkedList<Integer> list = new LinkedList<>();
    list.append(10);
    list.unshift(5);
    list.appendOnIndex(7, 1);

    list.printList(); // 5 -> 7 -> 10 

   var indexFound = list.search(null);
    if (indexFound != null) {
      System.out.println("Value found at index: " + indexFound);
    } else {
      System.out.println("Value not found");
    }

   list.delete(7);
    list.printList(); // 5 -> 10

  }
}

// Métodos implementados:
// - Unshift: Adiciona elemento no início
// - Append: Adiciona elemento no final
// - AppendOnIndex: Adiciona elemento em posição específica
// - Search: Busca por um elemento
// - Delete: Remove um elemento
```

## Comparativo de Desempenho

Realizamos testes de benchmark comparando a eficiência de listas encadeadas e slices em Go para várias operações comuns:

### Resultados (ns/op)

| Operação              | Lista Encadeada | Slice      | Comparação         |
|-----------------------|-----------------|------------|-------------------|
| **Unshift**           | 51.22           | 75,649.00  | Lista 1,477x mais rápida |
| **Append**            | 90,371.00       | 14.02      | Slice 6,445x mais rápido |
| **AppendOnIndex**     | 487.20          | 772.50     | Lista 1.59x mais rápida |
| **Search**            | 535.00          | 310.10     | Slice 1.73x mais rápido |
| **Delete**            | 0.64            | 0.32       | Slice 2x mais rápido |
| **Read**              | 807.20          | 0.65       | Slice 1,242x mais rápido |
| **Write**             | 113,534.00      | 21.17      | Slice 5,362x mais rápido |

### Conclusões Principais:
- **Lista Encadeada** é superior para inserções no início (Unshift) e inserções em posições específicas
- **Slice** é superior para a maioria das outras operações, especialmente para acesso direto e append
- Para detalhes completos, veja o [comparativo detalhado](comparativo.md)

## Como Executar

### Exemplos em Go:

```bash
go run ./Golang/cmd/example/main.go
```

### Benchmarks:

```bash
cd Golang/cmd/benchmark
go test -bench=.
```

### Benchmarks com perfil de CPU:

```bash
cd Golang/cmd/benchmark
go test -bench=. -cpuprofile=cpu.prof
go tool pprof cpu.prof
```

### Analisador de Perfil:

```bash
go run ./Golang/cmd/profile/profile_analyzer.go ./Golang/cmd/benchmark/cpu.prof
```

## Estrutura do Projeto

```
/linkedlist
├── README.md                     # Este arquivo
├── algoritms/                    # Implementação da lista encadeada em Go
│   └── linked_list.go
├── benchmark_test.go             # Testes de benchmark
├── cmd/                          # Ferramentas auxiliares  
│   └── pprof/
│       └── profile_analyzer.go   # Analisador de perfil de CPU
├── comparativo.md                # Análise detalhada de desempenho
├── comparativo_es.md             # Versão em espanhol do comparativo
├── java/                         # Implementação em Java (a ser adicionada)
└── main.go                       # Exemplos de uso
```

## Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests com melhorias.

## Licença

Este projeto é disponibilizado sob a licença MIT.
