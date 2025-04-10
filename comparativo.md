# Comparativo de Desempenho: Lista Encadeada vs Slice

Este documento apresenta uma análise comparativa de desempenho entre uma **lista encadeada** implementada manualmente e um **slice** do Go para diferentes operações.

## Operações Comparadas

1. **Unshift**: Inserção no início.
2. **Append**: Inserção no final.
3. **AppendOnIndex**: Inserção em um índice específico.
4. **Search**: Busca por um elemento.
5. **Delete**: Remoção de um elemento.

## Resultados de Benchmark

Os resultados abaixo foram obtidos utilizando o comando `go test -bench=.`. Os tempos são medidos em nanosegundos por operação, e o uso de memória em bytes.

### Tempo de Execução

| Operação              | Lista Encadeada | Slice            |
|-----------------------|-----------------|------------------|
| **Unshift**           | Mais rápido     | Mais lento       |
| **Append**            | Mais lento      | Mais rápido      |
| **AppendOnIndex**     | Mais lento      | Mais rápido      |
| **Search**            | Mais lento      | Mais rápido      |
| **Delete**            | Mais lento      | Mais rápido      |

### Uso de Memória

| Operação              | Lista Encadeada | Slice            |
|-----------------------|-----------------|------------------|
| **Unshift**           | Mais eficiente  | Menos eficiente  |
| **Append**            | Menos eficiente | Mais eficiente   |
| **AppendOnIndex**     | Menos eficiente | Mais eficiente   |
| **Search**            | Menos eficiente | Mais eficiente   |
| **Delete**            | Menos eficiente | Mais eficiente   |

## Análise

### 1. **Unshift**
- **Lista Encadeada**: A operação é eficiente em tempo e memória, pois basta atualizar o ponteiro do nó inicial.
- **Slice**: É necessário realocar todos os elementos, o que consome mais memória e tempo.

### 2. **Append**
- **Lista Encadeada**: Requer percorrer toda a lista para encontrar o último nó, consumindo mais tempo e memória.
- **Slice**: A operação é otimizada internamente pelo Go, sendo mais eficiente em tempo e memória.

### 3. **AppendOnIndex**
- **Lista Encadeada**: Requer percorrer a lista até o índice desejado, consumindo mais memória e tempo.
- **Slice**: Apesar de envolver realocação, é mais eficiente devido à otimização interna.

### 4. **Search**
- **Lista Encadeada**: A busca é linear, percorrendo nó a nó, consumindo mais memória e tempo.
- **Slice**: A busca é linear, mas o acesso direto ao índice torna a operação mais eficiente.

### 5. **Delete**
- **Lista Encadeada**: Requer percorrer a lista para encontrar o elemento e atualizar os ponteiros, consumindo mais memória e tempo.
- **Slice**: Envolve realocação, mas ainda é mais eficiente devido à otimização interna.

## Conclusão

- **Lista Encadeada**: É mais eficiente em memória e tempo para operações como `Unshift`, onde a manipulação de ponteiros é vantajosa.
- **Slice**: É mais eficiente em memória e tempo para a maioria das operações devido à otimização interna do Go.

A escolha entre lista encadeada e slice depende do caso de uso. Para operações frequentes no início da estrutura, a lista encadeada é mais indicada. Para operações gerais, o slice é mais eficiente.
