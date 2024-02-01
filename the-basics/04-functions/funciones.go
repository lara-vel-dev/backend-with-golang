package main

import "fmt"

// podemos retornar mas de un valor
// si queremos retornar mas de un valor, debemos de especificar el tipo de dato que vamos a retornar
// y ponerlo entre parentesis
// en este caso estamos retornando un string y un int64
func mensajeFunction(message string, gravity int64) (string, int64) {
	return message, gravity
}

func main() {
	fmt.Println(mensajeFunction("Hola soy una función en GO!", 9))
}
