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

   //Guardamos todas as configurações do servidor nessa variável
   srv := &http.Server{
      Addr: *addr, 
      ErrorLog: errorLog, 
      Handler: app.routes(), 
   }
  
   infoLog.Printf("Inicializando o servidor na porta %s\n", *addr)
   err := srv.ListenAndServe()
   errorLog.Fatal(err)
}