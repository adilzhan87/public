# sqrt

## Intructions

Écrire une fonction qui renvoie la racine carré d'un `int` passé en paramètre as parameter si cette racine carré est un nombre entier. Autrement elle renvoie `0`.

## Fonction attendue

```go
func Sqrt(int nb) int {

}
```

## Utilisation

Voici un éventuel [programme](TODO-LINK) pour tester votre fonction :

```go
package main

import (
        "fmt"
        piscine ".."
)

func main() {
	arg1 := 4
	arg2 := 3
	fmt.Println(piscine.Sqrt(arg1))
	fmt.Println(piscine.Sqrt(arg2))

}
```

Et son résultat :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
2
0
student@ubuntu:~/piscine/test$
```