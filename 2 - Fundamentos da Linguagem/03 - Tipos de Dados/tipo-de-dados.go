package main

import "fmt"

func main() {
	/* @@ Aula - Tipo de Dados

		*Tipos basicos em Go*

		Tipos Numéricos
			- Numéros Inteiro (int):
					Os numeros tipo int8 só suportam numeros até 8bits e sucessivamente até o 64bits 
					Inteiros: int, int8, int16, int32, int64
					O tipo *int* ele vai usar a arquitetura da máquina que está sendo utilizada para
					definir o tamanho do tipo. Por exemplo, um PC 32 bits ele vai ter o tamanho de int32.
			- Numeros Inteiros Sem Sinal "unsigned" (uint)
					Os numeros tipo uint8 só suportam numeros até 8bits positvos e sucessivamente até o 64bits
					Unsigned: uint, uint8, uint16, uint32, uint64
					O tipo *uint* ele vai usar a arquitetura de máquina que esta sendo utilizada para
					definir o tamanho do tipo. Por exemplo, um PC 64 bits ele vai ter o tamanho de uint64
			- Numeros Reais (float)
					São numeros com ponto flutuante, porém só suportam dois tipos: 32 e 64bits
					Reais: float32 e float64
					O tipo aceita numeros com pontos flutantes, por exemplo: 720.30 não aceitam virgula,
					apenas ponto. (.)
					Também não é possivel definir o tipo float, apenas os seus tipos.
	*/
<<<<<<< HEAD:2 - Fundamentos da Linguagem/03 - Tipos de Dados/tipo-de-dados.go

	// Numeros inteiros
=======
>>>>>>> 29e189acc1067d0b59ca1c67431920f6673cb80a:03 - Tipos de Dados/tipo-de-dados.go
	var numero int = -10000000
	fmt.Println(numero)
	var numero2 uint32 = 100000
	fmt.Println(numero2)
<<<<<<< HEAD:2 - Fundamentos da Linguagem/03 - Tipos de Dados/tipo-de-dados.go

	//Alias
	//Alias de int32
	var numero3 rune = 552056
	fmt.Println(numero3)
	//Alias de uint8
	var numero4 byte = 123
	fmt.Println(numero4)
=======
>>>>>>> 29e189acc1067d0b59ca1c67431920f6673cb80a:03 - Tipos de Dados/tipo-de-dados.go
	

	//Numeros Reais
	var numeroReal1 float32 = 123.44
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 7256.32
	fmt.Println(numeroReal2)

	numeroReal3 := 54448.23
	fmt.Println(numeroReal3)

}
