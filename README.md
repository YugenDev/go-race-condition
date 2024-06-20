## Condición de Carrera y Solución con `sync.Mutex` en Go

### Problema de Condición de Carrera

En programación concurrente, una condición de carrera (race condition) ocurre cuando varias goroutines acceden y modifican simultáneamente una variable compartida sin la sincronización adecuada. Esto puede llevar a resultados inesperados e incorrectos.

**Escenario de Ejemplo:**

Imaginemos una cuenta bancaria con un saldo almacenado en la variable `balance`. Si dos goroutines (representando depósitos) intentan agregar dinero al saldo al mismo tiempo, podrían surgir problemas:

- **Lecturas Inconsistentes:** Una goroutine podría leer el saldo actual (`b = balance`) antes de que la otra lo actualice. Luego, ambas goroutines agregan sus depósitos al valor `b` desactualizado, lo que resulta en un saldo final incorrecto.
- **Actualizaciones Perdidas:** Si una goroutine incrementa `balance` mientras otra está leyendo su valor, la lectura podría perder la actualización, lo que resultaría en un saldo final incorrecto.

### Solución con `sync.Mutex`

El código Go proporcionado previene efectivamente las condiciones de carrera utilizando un bloqueo `sync.Mutex` (exclusión mutua). Veamos cómo funciona:

#### Inicialización de Mutex

```go
var lock sync.Mutex
```

#### Función `Depositar`

```go
func Depositar(amount int, wg *sync.WaitGroup, lock *sync.Mutex) {
  defer wg.Done()
  lock.Lock()
  b := balance
  balance = b + amount
  lock.Unlock()
}
```

- `lock.Lock()`: Antes de acceder a la variable compartida `balance`, llamamos a la función lock de la libreria mutex, la gorutine lee esto y bloquea la sección del codigo que se encuentra antes del desbloqueo. Esto asegura que solo una goroutine pueda proceder a la vez.
- `lock.Unlock()`: Después de la sección crítica (actualizando `balance`), se libera el bloqueo, lo que permite que otras goroutines lo adquieran.

### Prevención de Condiciones de Carrera

Al poner el bloqueo mutex antes de acceder a `balance`, la función `Depositar` crea una sección crítica. Solo una goroutine puede tener el bloqueo a la vez, lo que garantiza que:

- Ninguna otra goroutine puede modificar `balance` mientras la actual lo está leyendo o actualizando.
- El valor leído de `balance` (almacenado en `b`) permanece constante durante la operación de depósito.
- La actualización final de `balance` refleja con precisión los depósitos combinados de todas las goroutines.

### En Resumen

El `sync.Mutex` en este código Go actúa como un bloqueo, asegurando que solo una goroutine pueda acceder y modificar la variable compartida `balance` a la vez. Esto evita las condiciones de carrera y garantiza actualizaciones consistentes del saldo, incluso en escenarios concurrentes.
