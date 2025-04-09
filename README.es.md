# linkedlist

## Lista Enlazada

### ✅ Ventajas:

| Ventaja | Explicación |
|----------|------------|
| ✅ Tamaño dinámico | Crece según sea necesario, sin necesidad de asignar espacio fijo. |
| ✅ Inserción/eliminación eficientes | Insertar/eliminar elementos al inicio o en el medio es rápido (O(1) u O(n)). |
| ✅ Uso eficiente de la memoria | No necesita bloques contiguos de memoria, como los arrays. |

### ❌ Desventajas:

| Desventaja | Explicación |
|-------------|------------|
| ❌ Acceso lento | Encontrar un elemento por índice requiere recorrer la lista (O(n)). |
| ❌ Mayor uso de memoria por nodo | Cada nodo necesita espacio adicional para el puntero/referencia. |
| ❌ Rendimiento de caché deficiente | Como los elementos no son contiguos, el acceso es más lento para la CPU. |

## Lista Común (Array / Slice)

### ✅ Ventajas:

| Ventaja | Explicación |
|----------|------------|
| ✅ Acceso rápido | Acceso directo por índice (O(1)) es instantáneo. |
| ✅ Mejor uso de caché | Elementos contiguos en memoria = más rápido para la CPU. |
| ✅ Estructura simple | Más fácil de implementar y entender. |

### ❌ Desventajas:

| Desventaja | Explicación |
|-------------|------------|
| ❌ Tamaño fijo (en arrays puros) | Es necesario conocer el tamaño con antelación (excepto con slices en Go). |
| ❌ Inserciones/eliminaciones costosas | Insertar/eliminar en el medio requiere desplazar elementos (O(n)). |
| ❌ Reasignación al crecer | En slices dinámicos (como en Go), puede haber copia al crecer demasiado. |

## 📊 Comparativa rápida:

| Operación | Lista Enlazada | Lista Común (Array/Slice) |
|----------|----------------|--------------------------|
| Acceso por índice | ❌ O(n) | ✅ O(1) |
| Inserción al inicio | ✅ O(1) | ❌ O(n) (requiere mover) |
| Inserción al final | ✅/❌ O(n)/O(1)* | ✅ O(1) amortizado |
| Eliminación al inicio | ✅ O(1) | ❌ O(n) |
| Eliminación en el medio | ✅ O(n) | ❌ O(n) |
| Uso de memoria | ❌ Mayor (punteros) | ✅ Más eficiente |

\* La inserción al final en una lista enlazada es O(1) si tiene un puntero al último nodo, de lo contrario es O(n).
