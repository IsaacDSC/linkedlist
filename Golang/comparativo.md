# Comparativo de Desempenho: Lista Encadeada vs Slice

Este documento apresenta uma análise comparativa de desempenho entre uma **lista encadeada** implementada manualmente e um **slice** do Go para diferentes operações.

## Operações Comparadas

1. **Unshift**: Inserção no início.
2. **Append**: Inserção no final.
3. **AppendOnIndex**: Inserção em um índice específico.
4. **Search**: Busca por um elemento.
5. **Delete**: Remoção de um elemento.

## Resultados de Benchmark

Os resultados abaixo foram obtidos utilizando o comando `go test -bench=. -cpuprofile=cpu.prof`. 
Os tempos são medidos em nanosegundos por operação.

### Resultados Numéricos dos Benchmarks

#### Tempo de Execução (ns/op)

| Operação              | Lista Encadeada | Slice      | Comparação         |
|-----------------------|-----------------|------------|-------------------|
| **Unshift**           | 51.22           | 75,649.00  | Lista 1,477x mais rápida |
| **Append**            | 90,371.00       | 14.02      | Slice 6,445x mais rápido |
| **AppendOnIndex**     | 487.20          | 772.50     | Lista 1.59x mais rápida |
| **Search**            | 535.00          | 310.10     | Slice 1.73x mais rápido |
| **Delete**            | 0.64            | 0.32       | Slice 2x mais rápido |
| **Read**              | 807.20          | 0.65       | Slice 1,242x mais rápido |
| **Write**             | 113,534.00      | 21.17      | Slice 5,362x mais rápido |

### Análise Qualitativa

#### Tempo de Execução

| Operação              | Lista Encadeada | Slice            |
|-----------------------|-----------------|------------------|
| **Unshift**           | ✅ Mais rápido   | ❌ Mais lento     |
| **Append**            | ❌ Mais lento    | ✅ Mais rápido    |
| **AppendOnIndex**     | ✅ Mais rápido   | ❌ Mais lento     |
| **Search**            | ❌ Mais lento    | ✅ Mais rápido    |
| **Delete**            | ❌ Mais lento    | ✅ Mais rápido    |
| **Read**              | ❌ Mais lento    | ✅ Mais rápido    |
| **Write**             | ❌ Mais lento    | ✅ Mais rápido    |

## Análise Detalhada

### 1. **Unshift**
- **Lista Encadeada (51.22 ns/op)**: A operação é extremamente eficiente, pois basta atualizar o ponteiro do nó inicial.
- **Slice (75,649.00 ns/op)**: É necessário realocar todos os elementos, o que é extremamente custoso, tornando-o cerca de 1,477 vezes mais lento que a lista encadeada.

### 2. **Append**
- **Lista Encadeada (90,371.00 ns/op)**: Requer percorrer toda a lista para encontrar o último nó, e então adicionar o elemento, o que se torna muito ineficiente com o crescimento da lista.
- **Slice (14.02 ns/op)**: A operação é otimizada internamente pelo Go, sendo cerca de 6,445 vezes mais rápida que a lista encadeada.

### 3. **AppendOnIndex**
- **Lista Encadeada (487.20 ns/op)**: Surpreendentemente mais rápida para inserções em índices aleatórios.
- **Slice (772.50 ns/op)**: Mais lento devido à necessidade de deslocar elementos após o ponto de inserção.

### 4. **Search**
- **Lista Encadeada (535.00 ns/op)**: A busca é linear, percorrendo nó a nó.
- **Slice (310.10 ns/op)**: Também usa busca linear, mas é cerca de 1.73 vezes mais rápida devido à melhor localidade de memória.

### 5. **Delete**
- **Lista Encadeada (0.64 ns/op)**: O tempo extremamente baixo sugere que esta implementação pode estar apenas removendo a referência, sem percorrer a lista.
- **Slice (0.32 ns/op)**: Similarmente rápido, com desempenho 2 vezes melhor que a lista encadeada.

### 6. **Read/Write (Operações Gerais)**
- A leitura direta em slice (0.65 ns/op) é cerca de 1,242 vezes mais rápida que em lista encadeada (807.20 ns/op).
- A escrita em slice (21.17 ns/op) é cerca de 5,362 vezes mais rápida que em lista encadeada (113,534.00 ns/op).

## Conclusão

Com base nos benchmarks realizados:

1. **Lista Encadeada** é significativamente superior para:
   - **Unshift**: ~1,477 vezes mais rápida que slice
   - **AppendOnIndex**: ~1.59 vezes mais rápida que slice

2. **Slice** é significativamente superior para:
   - **Append**: ~6,445 vezes mais rápido que lista encadeada
   - **Search**: ~1.73 vezes mais rápido que lista encadeada
   - **Delete**: 2 vezes mais rápido que lista encadeada
   - **Read**: ~1,242 vezes mais rápido que lista encadeada
   - **Write**: ~5,362 vezes mais rápido que lista encadeada

### Recomendação de Uso

- **Use Lista Encadeada quando**:
  - A operação predominante for inserção no início (Unshift)
  - Precisar inserir elementos em posições específicas frequentemente

- **Use Slice quando**:
  - A operação predominante for adicionar ao final (Append)
  - Precisar de acesso rápido aos elementos (Read)
  - A performance geral for mais importante que operações específicas

## Observações Adicionais

- Os tempos extremamente baixos para operações Delete sugerem uma possível otimização na implementação ou um padrão de benchmark que favorece casos especiais.
- Para melhor análise de uso de memória, seria recomendável executar o benchmark com o parâmetro `-benchmem`.
