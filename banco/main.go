package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	//Passando o nome dos parâmetros
	contaDoBruno := ContaCorrente{
		titular: "Bruno",
		numeroAgencia: 7413,
		numeroConta: 1234,
		saldo: 75.7,
	}

	//Sem nome dos parâmetros
	contaDoMi := ContaCorrente{
		"Mi",
		7413,
		1234,
		75.7,
	}

	fmt.Println(contaDoBruno)
	fmt.Println(contaDoMi)
}
