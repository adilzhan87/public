## sortwordarr

### Instructions

Écrire une fonction `SortWordArr` qui trie par ordre `ascii` (dans l'ordre croissant) un tableau de `string`.

### Fonction attendue

```go
func SortWordArr(array []string) {

}
```

### Utilisation

Voici un éventuel programme pour tester votre fonction :

```go
package main

import (
	"fmt"
	piscine ".."
)

func main() {

	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	piscine.SortWordArr(result)

	fmt.Println(result)
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
[1 2 3 A B C a b c]
student@ubuntu:~/[[ROOT]]/test$
```
