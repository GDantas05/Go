package main

import (
	"fmt"
	"os"
	"strconv"
) //IMPORTANDO AS BLIBIOTECAS DO GO

func main() {

	/*
	 * SE A QUANTIDADE DE PARÂMETROS PASSADOS FOR MENOR QUE 3 UMA MENSAGEM DE ERRO É MOSTRADA
	 * E O PROGRAMA É ABORTADO.
	 * OS PARÂMETROS SÃO: 1º - NOME DO PROGRAMA(SEMPRE SERÁ O PRIMEIRO PARÂMETRO). 2º - VALOR A
	 * SER CONVERTIDO. 3° - A UNIDADE DE ORIGEM
	 */
	if len(os.Args) < 3 {
		fmt.Println("Uso: conversor <valores> <unidade>")
		os.Exit(1)
	}

	unidadeOrigem := os.Args[len(os.Args)-1]
	valoresOrigem := os.Args[1 : len(os.Args)-1]

	//QUANDO UMA VARIÁVEL É DECLARADA SEM SER INICIALIZADA SEU TIPO DEVE SER INFORMADO
	var unidadeDestino string

	if unidadeOrigem == "celsius" {
		unidadeDestino = "fahrenheit"
	} else if unidadeOrigem == "quilometros" {
		unidadeDestino = "milhas"
	} else {
		fmt.Printf("%s não é unidade conhecida!", unidadeDestino)
		os.Exit(1)
	}

	/*
	 * A FUNÇÃO RANGE TRANSFORMA O SLICE VALORESDEORIGEM EM DUAS PARTES ATRIBUINDO O
	 * INDICE A LETRA I E O VALOR A LETRA V
	 */
	for i, v := range valoresOrigem {

		//STRCONV RETORNA DOIS VALORES: O NUMERO CONVERTIDO E UMA MENSAGEM DE ERRO, CASO
		//NÃO TENHA NENHUMA ERRO NA CONVERSÃO RETORNA NIL
		valorOrigem, err := strconv.ParseFloat(v, 64)

		if err != nil {
			fmt.Printf("O valor %s na posição %d não é um número válido!\n", v, i)
			os.Exit(1)
		}

		var valorDestino float64

		if unidadeOrigem == "celsius" {
			valorDestino = valorOrigem*1.8 + 32
		} else {
			valorDestino = valorOrigem / 1.60934
		}

		//O .2F ARREDONDA UM VALOR E FORMATA PARA DUAS CASAS DECIMAIS
		fmt.Printf("%.2f %s = %.2f %s\n", valorOrigem, unidadeOrigem, valorDestino, unidadeDestino)
	}

}
