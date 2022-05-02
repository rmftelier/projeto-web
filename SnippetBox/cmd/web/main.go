//Comando para rodar o site: cd SnippetBox e logo dps go run cmd/web/*

package main

import (
  "flag"
	"log"
	"net/http"
  "os"
  "database/sql"

  _ "github.com/go-sql-driver/mysql"
)

//Criação do struct
type application struct{
    errorLog *log.Logger
    infoLog *log.Logger
    
}

func main() {
   //Flag 
  addr := flag.String("addr", ":4000", "Porta da Rede")
  //CONNECTION STRING
  dsn := flag.String("dsn", "QH7S9cJJ7J:c4mr1j47e4@tcp(remotemysql.com)/QH7S9cJJ7J?parseTime=true", "MySQL DSN")
  
  flag.Parse()

  infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
  errorLog := log.New(os.Stderr, "ERRO:\t", log.Ldate|log.Ltime|log.Lshortfile)


  db, err := openDB(*dsn)
  if err != nil{
     errorLog.Fatal(err)
  }
  
  defer db.Close()
  
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
   err = srv.ListenAndServe()
   errorLog.Fatal(err)
}

//Função responsável por conectar o BD 
func openDB(dsn string)(*sql.DB, error){
    db, err := sql.Open("mysql", dsn)
    if err != nil{
       return nil, err
    }
    //Teste para verificar o funcionamento do BD  
    if err = db.Ping(); err != nil{
       return nil, err 
    }

    return db, nil
}