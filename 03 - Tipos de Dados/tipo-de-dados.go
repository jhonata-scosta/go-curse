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
			- 
	*/
	var numero int = 10000000
	fmt.Println(numero)
	
}
