package main // É a primeira coisa que deve vir no arquivo principal da aplicação, e mostra para o Go que este arquivo é um executável

import "fmt" //Lista de imports: É onde referenciado os pacotes que serão utilizados na aplicação
//     parametros 
func main () { //Como definir uma função. Mas está função é o ponto inicial da aplicação, onde tudo será executado.
	fmt.Println("Hello, World!")
}
