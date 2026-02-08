package main

import (
	"errors"
	"fmt"
)

func main() {
	/* @@ Aula - Tipo de Dados

		*Tipos basicos em Go*

		Tipos Numéricos
			- Numéros Inteiro (int):
					Os numeros tipo int8 só suportam numeros até 8bits e sucessivamente até o 64bits 
					Inteiros: int, int8, int16, int32, int64
					O tipo *int* ele vai usar a arquitetura da máquina que está sendo utilizada para
					definir o tamanho do tipo. Por exemplo, um PC 32 bits ele vai ter o tamanho de int32.
			- Numeros Inteiros Sem Sinal "unsigned" (uint):
					Os numeros tipo uint8 só suportam numeros até 8bits positvos e sucessivamente até o 64bits
					Unsigned: uint, uint8, uint16, uint32, uint64
					O tipo *uint* ele vai usar a arquitetura de máquina que esta sendo utilizada para
					definir o tamanho do tipo. Por exemplo, um PC 64 bits ele vai ter o tamanho de uint64
			- Numeros Reais (float):
					São numeros com ponto flutuante, porém só suportam dois tipos: 32 e 64bits
					Reais: float32 e float64
					O tipo aceita numeros com pontos flutantes, por exemplo: 720.30 não aceitam virgula,
					apenas ponto. (.)
					Também não é possivel definir o tipo float, apenas os seus tipos.
			- Cadeia de Caracteres (string):
					É o tipo de texto da linguagem, deve ser escrito entras aspas duplas ("").
					Para o tipo char(unico caracter), ele utilizar o tipo string, mas o caracter precisa 
					está entre aspas simples (''), e ele irá retornar o valor daquele caracter da tabela
					ASCII e será do tipo int.
			- Booleano(bool):
					É o tipo que possui apenas 2 valores, verdadeiro(true) e falso (false).
			- Erro (error):
					É um tipo que o possui para lidar com erros, junto com o pacote padrão errors.
	*/ 

	// Numeros inteiros

	var numero int = -10000000
	fmt.Println(numero)
	var numero2 uint32 = 100000
	fmt.Println(numero2)

	//Alias
	//Alias de int32
	var numero3 rune = 552056
	fmt.Println(numero3)
	//Alias de uint8
	var numero4 byte = 123
	fmt.Println(numero4)


	//Numeros Reais
	var numeroReal1 float32 = 123.44
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 7256.32
	fmt.Println(numeroReal2)

	numeroReal3 := 54448.23
	fmt.Println(numeroReal3)

	//Strings (Cadeia de Caracteres)
	var str string = "Texto"
	fmt.Println(str)
	char := "A"
	fmt.Println(char)

	//Booleano

	var boolean bool = true
	fmt.Println(boolean)

	// Error

	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)

	/* Todos os tipos possui um valor padrão. Quando uma variavel é inicializada, porém sem ter
	valor, ela já possui um valor proprio para seu tipo.
	Ex: 
		# var str2 string 
		retorna um string vazio
		# var numb int
		retorna um 0

		# var bol bool
		retorna o valor false
	*/
	var str2 string
	fmt.Println(str2)
}
