package main

import "fmt"

/*
	@@ Aula - Funções
	Funções são blocos de texto reutilizaveis para realizar tarefas especificas.
	Em go para declarar uma função utilizamos a palavra reservada "func" e em sua estrutura podemos
	receber parametros, tipos de retorno e o codigo que gere a logica que será utilizada.
	Exemplos:
*/

// 1 - Essa função recebe dois paramentros do tipo int8 e seu retorno também e do tipo int8:
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}
/* 2 - É possivel receber mais de dois tipos de retorno, precisando especificar cada uma delas 
	entre parenteses. 
*/
func calculosMatematicos(n1, n2 int8) (int8, int8){
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	// 1 - aqui ele é chamado e o valor final fica armazenado na variavel
	soma := somar(8, 6)
	fmt.Println(soma)
	/* É possivel contruir uma função da seguinte maneira. Pois a função também é um tipo */
	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)
}

