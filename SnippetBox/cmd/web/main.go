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

  log.Println("Inicializando o servidor na porta: 4000")
  
  //Caso ocorra erro: 
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}
