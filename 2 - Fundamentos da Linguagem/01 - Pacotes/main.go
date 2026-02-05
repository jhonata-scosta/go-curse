/*
	Os pacotes são uma coleção de códigos-fontes que estão no mesmo diretorio e são compilados juntos
	Eles fornecem uma estrutura para organizar o código em módulos lógicos.

	Módulo é um conjuto de pacotes que compõe o projeto. Esse pacotes podem ser criados ou ser utilizado
	os externos de terceiros (github, etc...) e os pacotes padrões do go.

	Para criar um módulo, na pasta raiz do projeto:

	$ go mod init <nome-do-modulo>

	É semelhante com o package.json do Node, onde fica armazenado todas as dependências do projeto
	com versionamento e outro detalhes.

	É possivel compilar o projeto para gerar um arquivo executável, que pega todas as informações,
	e gera apenas um unico arquivo, onde estará o seu projeto.

	$ go build

	Para criar um pacote, necessita-se de criar uma pasta para cada um modulo diferente.
	O nome da pasta e o nome do arquivo não importam para criar, só é necessario criar um arquivo
	.go e colocar o nome do pacote assim dentro do arquivo.
	
	# package <nome-do-pacote>

	Como Go, não é uma Linguagem orientada a objetos, não existe Public, Private, Protected
	No Go, para definir se algo como função, variavel,constante, struct é Public ou Private
	é pela primeira letra delas.
	Por exemplo, uma função começar com letra minuscula:

	# func escrever(){
	# }
	
	Ela só será visivel dentro do seu proprio pacote. Se ela começa com letra maiuscula:
	
	# func Escrever(){	
	#	}
	
	Ela será visivel para outros pacotes.

	O comando $ go install , ele gerar um arquivo binário e adiciona ele na raiz da instalação do go.
*/

package main

import (
	"auxiliar/auxiliar"
	"fmt"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
}
