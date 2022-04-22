//Comando para rodar o site: cd SnippetBox e logo dps go run cmd/web/*
package main

import (
	"log"
	"net/http"
)

func main() {
	//Conexão ao servidor:
  mux := http.NewServeMux()

  //Caminho relativo:
  mux.HandleFunc("/", home)

  //Caminhos absolutos:
  mux.HandleFunc("/snippet", showSnippet)
  mux.HandleFunc("/snippet/create", createSnippet)

  //Criação de um manipulador (handler) de servidor de arquivos (file server)
  fileServer := http.FileServer(http.Dir("./ui/static/"))
  mux.Handle("/static/",http.StripPrefix("/static",fileServer))
  
  log.Println("Inicializando o servidor na porta: 4000")
  //Caso ocorra erro:
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}