package main

import (
	"fmt"
	"os"
	"strconv"
)

// fornece uma grande variedade de funções para a conversão de strings para outros formatos e vice-versa

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: conversor <valores> <unidades>")
		os.Exit(1)
	}

	// slices são listas indexadas e de tamanho variável em Go
	// Sendo que o primeiro elemento possui indice 0
	// Sendo assim, o indice do ultimo elemento é sempre n - 1
	// sendo n o número total de elementos presentes no slice
	// numero_total_de_argumentos - 1 = Ultimo elemento
	unidadeOrigem := os.Args[len(os.Args)-1]     // extrai o ultimo argumento da linha de comando e armazena na 'unidadeOrigem'
	valoresOrigem := os.Args[1 : len(os.Args)-1] // O 1 : refere-se ao segundo elemento e o len(os.Args)-1 ao ultimo
	// fmt.Println("Ultimo argumento:", unidadeOrigem)
	// fmt.Println("Primeiro argumento(INDICE 1):", valoresOrigem)

	var unidadeDestino string

	if unidadeOrigem == "celsius" {
		unidadeDestino = "fahrenheit"
	} else if unidadeOrigem == "quilometros" {
		unidadeDestino = "milhas"
	} else {
		fmt.Printf("%s não é uma unidade conhecida!\n", unidadeDestino)
		os.Exit(1)
	}

	for i, v := range valoresOrigem {
		var valorDestino float64
		// atribuimos o indice à variavel i e o elemento à v
		valorOrigem, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf("O valor %s na posição %d não é um número válido\n", v, i)
			os.Exit(1)
		}

		if unidadeOrigem == "celsius" {
			valorDestino = valorOrigem*1.8 + 32
		} else {
			valorDestino = valorOrigem / 1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s\n", valorOrigem, unidadeOrigem, valorDestino, unidadeDestino)
	}
}

// exemplo de execução: go run conversor_medida.go 32 27.4 -3 0 celsius
