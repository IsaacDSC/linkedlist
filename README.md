# linkedlist

## Lista Encadeada

### ✅ Vantagens:

| Vantagem | Explicação |
|----------|------------|
| ✅ Tamanho dinâmico | Cresce conforme necessário, sem precisar alocar espaço fixo. |
| ✅ Inserção/remoção eficientes | Inserir/remover elementos no início ou meio é rápido (O(1) ou O(n)). |
| ✅ Uso eficiente da memória | Não precisa de blocos contíguos de memória, como arrays. |

### ❌ Desvantagens:

| Desvantagem | Explicação |
|-------------|------------|
| ❌ Acesso lento | Achar um elemento por índice exige percorrer a lista (O(n)). |
| ❌ Maior uso de memória por nó | Cada nó precisa de espaço extra pro ponteiro/referência. |
| ❌ Cache performance ruim | Como os elementos não são contíguos, o acesso é mais lento pra CPU. |

## Lista Comum (Array / Slice)

### ✅ Vantagens:

| Vantagem | Explicação |
|----------|------------|
| ✅ Acesso rápido | Acesso direto por índice (O(1)) é instantâneo. |
| ✅ Melhor uso de cache | Elementos contíguos na memória = mais rápido pra CPU. |
| ✅ Estrutura simples | Mais fácil de implementar e entender. |

### ❌ Desvantagens:

| Desvantagem | Explicação |
|-------------|------------|
| ❌ Tamanho fixo (em arrays puros) | Precisa saber o tamanho antecipadamente (exceto com slices em Go). |
| ❌ Inserções/remoções custosas | Inserir/remover do meio exige deslocar elementos (O(n)). |
| ❌ Realoque em crescimento | Em slices dinâmicos (como em Go), pode haver cópia ao crescer demais. |

## 📊 Comparativo rápido:

| Operação | Lista Encadeada | Lista Comum (Array/Slice) |
|----------|----------------|--------------------------|
| Acesso por índice | ❌ O(n) | ✅ O(1) |
| Inserção no início | ✅ O(1) | ❌ O(n) (precisa mover) |
| Inserção no fim | ✅/❌ O(n)/O(1)* | ✅ O(1) amortizado |
| Remoção no início | ✅ O(1) | ❌ O(n) |
| Remoção no meio | ✅ O(n) | ❌ O(n) |
| Uso de memória | ❌ Maior (ponteiros) | ✅ Mais eficiente |

\* Inserção no fim em lista encadeada é O(1) se tiver ponteiro pro último nó, senão é O(n).
