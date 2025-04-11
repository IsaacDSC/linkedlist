# Comparativo de Rendimiento: Lista Enlazada vs Slice

Este documento presenta un análisis comparativo de rendimiento entre una **lista enlazada** implementada manualmente y un **slice** de Go para diferentes operaciones.

## Operaciones Comparadas

1. **Unshift**: Inserción al inicio.
2. **Append**: Inserción al final.
3. **AppendOnIndex**: Inserción en un índice específico.
4. **Search**: Búsqueda de un elemento.
5. **Delete**: Eliminación de un elemento.

## Resultados del Benchmark

Los resultados a continuación fueron obtenidos utilizando el comando `go test -bench=. -cpuprofile=cpu.prof`. 
Los tiempos se miden en nanosegundos por operación.

### Resultados Numéricos de los Benchmarks

#### Tiempo de Ejecución (ns/op)

| Operación             | Lista Enlazada  | Slice      | Comparación         |
|-----------------------|-----------------|------------|-------------------|
| **Unshift**           | 51.22           | 75,649.00  | Lista 1,477x más rápida |
| **Append**            | 90,371.00       | 14.02      | Slice 6,445x más rápido |
| **AppendOnIndex**     | 487.20          | 772.50     | Lista 1.59x más rápida |
| **Search**            | 535.00          | 310.10     | Slice 1.73x más rápido |
| **Delete**            | 0.64            | 0.32       | Slice 2x más rápido |
| **Read**              | 807.20          | 0.65       | Slice 1,242x más rápido |
| **Write**             | 113,534.00      | 21.17      | Slice 5,362x más rápido |

### Análisis Cualitativo

#### Tiempo de Ejecución

| Operación             | Lista Enlazada  | Slice            |
|-----------------------|-----------------|------------------|
| **Unshift**           | ✅ Más rápido    | ❌ Más lento     |
| **Append**            | ❌ Más lento     | ✅ Más rápido    |
| **AppendOnIndex**     | ✅ Más rápido    | ❌ Más lento     |
| **Search**            | ❌ Más lento     | ✅ Más rápido    |
| **Delete**            | ❌ Más lento     | ✅ Más rápido    |
| **Read**              | ❌ Más lento     | ✅ Más rápido    |
| **Write**             | ❌ Más lento     | ✅ Más rápido    |

## Análisis Detallado

### 1. **Unshift**
- **Lista Enlazada (51.22 ns/op)**: La operación es extremadamente eficiente, pues solo requiere actualizar el puntero del nodo inicial.
- **Slice (75,649.00 ns/op)**: Es necesario reubicar todos los elementos, lo que es extremadamente costoso, haciéndolo aproximadamente 1,477 veces más lento que la lista enlazada.

### 2. **Append**
- **Lista Enlazada (90,371.00 ns/op)**: Requiere recorrer toda la lista para encontrar el último nodo, y luego añadir el elemento, lo que se vuelve muy ineficiente con el crecimiento de la lista.
- **Slice (14.02 ns/op)**: La operación está optimizada internamente por Go, siendo aproximadamente 6,445 veces más rápida que la lista enlazada.

### 3. **AppendOnIndex**
- **Lista Enlazada (487.20 ns/op)**: Sorprendentemente más rápida para inserciones en índices aleatorios.
- **Slice (772.50 ns/op)**: Más lento debido a la necesidad de desplazar elementos después del punto de inserción.

### 4. **Search**
- **Lista Enlazada (535.00 ns/op)**: La búsqueda es lineal, recorriendo nodo por nodo.
- **Slice (310.10 ns/op)**: También utiliza búsqueda lineal, pero es aproximadamente 1.73 veces más rápida debido a la mejor localidad de memoria.

### 5. **Delete**
- **Lista Enlazada (0.64 ns/op)**: El tiempo extremadamente bajo sugiere que esta implementación podría estar simplemente eliminando la referencia, sin recorrer la lista.
- **Slice (0.32 ns/op)**: Similarmente rápido, con un rendimiento 2 veces mejor que la lista enlazada.

### 6. **Read/Write (Operaciones Generales)**
- La lectura directa en slice (0.65 ns/op) es aproximadamente 1,242 veces más rápida que en lista enlazada (807.20 ns/op).
- La escritura en slice (21.17 ns/op) es aproximadamente 5,362 veces más rápida que en lista enlazada (113,534.00 ns/op).

## Conclusión

Basado en los benchmarks realizados:

1. **Lista Enlazada** es significativamente superior para:
   - **Unshift**: ~1,477 veces más rápida que slice
   - **AppendOnIndex**: ~1.59 veces más rápida que slice

2. **Slice** es significativamente superior para:
   - **Append**: ~6,445 veces más rápido que lista enlazada
   - **Search**: ~1.73 veces más rápido que lista enlazada
   - **Delete**: 2 veces más rápido que lista enlazada
   - **Read**: ~1,242 veces más rápido que lista enlazada
   - **Write**: ~5,362 veces más rápido que lista enlazada

### Recomendación de Uso

- **Use Lista Enlazada cuando**:
  - La operación predominante sea inserción al inicio (Unshift)
  - Necesite insertar elementos en posiciones específicas frecuentemente

- **Use Slice cuando**:
  - La operación predominante sea añadir al final (Append)
  - Necesite acceso rápido a los elementos (Read)
  - El rendimiento general sea más importante que operaciones específicas

## Observaciones Adicionales

- Los tiempos extremadamente bajos para operaciones Delete sugieren una posible optimización en la implementación o un patrón de benchmark que favorece casos especiales.
- Para un mejor análisis del uso de memoria, sería recomendable ejecutar el benchmark con el parámetro `-benchmem`.
