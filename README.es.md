# Lista Enlazada - Estudio y Comparativo

## ¿Qué es una Lista Enlazada?

Una lista enlazada es una estructura de datos lineal compuesta por una secuencia de elementos llamados "nodos", donde cada nodo contiene:
1. Un valor de dato
2. Un puntero (o referencia) al siguiente nodo en la secuencia

![Estructura de Lista Enlazada](https://upload.wikimedia.org/wikipedia/commons/6/6d/Singly-linked-list.svg)

A diferencia de arrays o slices, que almacenan elementos en posiciones contiguas de memoria, las listas enlazadas pueden usar memoria no contigua, con cada nodo pudiendo estar en cualquier ubicación de la memoria.

### Características Principales:
- **Inserción y eliminación eficientes**: Especialmente al inicio de la lista (O(1))
- **Acceso secuencial**: Para acceder a un elemento, es necesario recorrer la lista desde el inicio
- **Tamaño dinámico**: Crece según sea necesario, limitado solo por la memoria disponible
- **Uso de memoria**: Cada nodo requiere espacio adicional para almacenar el puntero

## Sobre este Proyecto

Este proyecto tiene como objetivo:
1. Demostrar la implementación de listas enlazadas en diferentes lenguajes (Go y Java)
2. Comparar el rendimiento entre listas enlazadas y estructuras nativas (slices en Go)
3. Servir como material de estudio para estructuras de datos

## Implementaciones

### Golang

La implementación en Go usa genéricos para crear una lista enlazada que puede almacenar cualquier tipo de dato:

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
// - Unshift: Añade elemento al inicio
// - Append: Añade elemento al final
// - AppendOnIndex: Añade elemento en posición específica
// - Search: Busca un elemento
// - Delete: Elimina un elemento
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
// - Unshift: Añade elemento al inicio
// - Append: Añade elemento al final
// - AppendOnIndex: Añade elemento en posición específica
// - Search: Busca un elemento
// - Delete: Elimina un elemento
```

## Comparativo de Rendimiento

Realizamos pruebas de benchmark comparando la eficiencia de listas enlazadas y slices en Go para varias operaciones comunes:

### Resultados (ns/op)

| Operación             | Lista Enlazada  | Slice      | Comparación         |
|-----------------------|-----------------|------------|-------------------|
| **Unshift**           | 51.22           | 75,649.00  | Lista 1,477x más rápida |
| **Append**            | 90,371.00       | 14.02      | Slice 6,445x más rápido |
| **AppendOnIndex**     | 487.20          | 772.50     | Lista 1.59x más rápida |
| **Search**            | 535.00          | 310.10     | Slice 1.73x más rápido |
| **Delete**            | 0.64            | 0.32       | Slice 2x más rápido |
| **Read**              | 807.20          | 0.65       | Slice 1,242x más rápido |
| **Write**             | 113,534.00      | 21.17      | Slice 5,362x más rápido |

### Conclusiones Principales:
- **Lista Enlazada** es superior para inserciones al inicio (Unshift) e inserciones en posiciones específicas
- **Slice** es superior para la mayoría de las otras operaciones, especialmente para acceso directo y append
- Para detalles completos, vea el [comparativo detallado](./Golang/comparativo_es.md)

## Cómo Ejecutar

### Ejemplos en Go:

```bash
go run ./Golang/cmd/example/main.go
```

### Benchmarks:

```bash
cd Golang/cmd/benchmark
go test -bench=.
```

### Benchmarks con perfil de CPU:

```bash
cd Golang/cmd/benchmark
go test -bench=. -cpuprofile=cpu.prof
go tool pprof cpu.prof
```

### Analizador de Perfil:

```bash
go run ./Golang/cmd/profile/profile_analyzer.go ./Golang/cmd/benchmark/cpu.prof
```

## Estructura del Proyecto

```
/linkedlist
├── README.md                     # Este archivo
├── algoritms/                    # Implementación de la lista enlazada en Go
│   └── linked_list.go
├── benchmark_test.go             # Pruebas de benchmark
├── cmd/                          # Herramientas auxiliares  
│   └── pprof/
│       └── profile_analyzer.go   # Analizador de perfil de CPU
├── comparativo.md                # Análisis detallado de rendimiento
├── comparativo_es.md             # Versión en español del comparativo
├── java/                         # Implementación en Java (por añadir)
└── main.go                       # Ejemplos de uso
```

## Contribuciones

¡Las contribuciones son bienvenidas! Siéntase libre de abrir issues o enviar pull requests con mejoras.

## Licencia

Este proyecto está disponible bajo la licencia MIT.
