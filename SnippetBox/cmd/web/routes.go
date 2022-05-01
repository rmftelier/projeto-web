package main

import "net/http"

func(app *application) routes() *http.ServeMux{

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

    return mux
}