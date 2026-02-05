/*
	No Go há duas maneiras de definir uma variável, onde uma deixa explicito o tipo e
	a outra que deixa ela implicita.

	De maneira Explicita:
	# var <nome-da-variavel> <tipo> = <valor>
	Não necessariamente não precisa inicia uma variavel, ela já vem com um valor "Zerado"

	De maneira implicita:
	# <nome-da-variavel> := <valor>
	Aqui ele vai auto-atribuir um tipo a variavel dependendo do valor.

	O go não deixa você definir variaveis e não utiliza-las, ele dá erro na execução, informando que a
	variavel foi declarada, mas não usada.

*/

package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1" // Explicito
	variavel2 := "Variável 2" //Implicito = Inferência de Tipo
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	//Outros meios de adicionar variaveis
	var (
		variavel3 string = "lalala"
		variavel4 string = "lalala"
	)
	fmt.Println(variavel3, variavel4)

  variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	//Constantes
	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5, variavel6)
}
