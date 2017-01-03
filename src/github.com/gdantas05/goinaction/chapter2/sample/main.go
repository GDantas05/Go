package main

import (
	//O IDENTIFICADOR BLANK ANTES DA IMPORTAÇÃO DO PACOTE PERMITE QUE O MESMO SEJA INICIALIZADO MESMO QUE NENHUM IDENTIFICADOR DESSE PACOTE
	//SEJA USADO DIRETAMENTE
	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
	"log"
	"os"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	//DEFINE QUE O LOGGER DA BIBLIOTECA PADRÃO DEVE ESCREVER NO DISPOSITIVO STDOUT, AO INVÉS DE ESCREVER NO PADRÃO STDERR
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	//CHAMA A FUNÇÃO RUN DO PACOTE SEARCH
	search.Run("president")
}
