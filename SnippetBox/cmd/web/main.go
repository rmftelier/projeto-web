//Comando para rodar o site: cd SnippetBox e logo dps go run cmd/web/*
package main

import (
  "flag"
	"log"
	"net/http"
  "os"
)

//Criação do struct
type application struct{
    errorLog *log.Logger
    infoLog *log.Logger
}

func main() {
   //Flag 
  addr := flag.String("addr", ":4000", "Porta da Rede")
  flag.Parse()

  infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
  errorLog := log.New(os.Stderr, "ERRO:\t", log.Ldate|log.Ltime|log.Lshortfile)

   app := &application{
        errorLog: errorLog,
        infoLog: infoLog,
   }

  
	//Conexão ao servidor:
  mux := http.NewServeMux()

  //Caminho relativo:
  mux.HandleFunc("/", app.home)

  //Caminhos absolutos:
  mux.HandleFunc("/snippet", app.showSnippet)
  mux.HandleFunc("/snippet/create", app.createSnippet)

  //Criação de um manipulador (handler) de servidor de arquivos (file server)
  fileServer := http.FileServer(http.Dir("./ui/static/"))
  mux.Handle("/static/",http.StripPrefix("/static",fileServer))
  
   infoLog.Printf("Inicializando o servidor na porta %s\n", *addr)
   err := http.ListenAndServe(*addr, mux)
   errorLog.Fatal(err)
}