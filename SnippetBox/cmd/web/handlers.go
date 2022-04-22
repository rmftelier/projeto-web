package main

//go run cmd/web/*
import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(rw http.ResponseWriter, r *http.Request){
  //Teste para verificar se um caminho está ou não no código, caso não esteja será enviado ao usuário um erro
  if r.URL.Path != "/"{
    http.NotFound(rw, r)
    return
  } 

  //Adicionado dia 13/04
  files := []string{
    "./ui/html/home.page.tmpl.html",
    "./ui/html/base.layout.tmpl.html",
    "./ui/html/footer.partial.tmpl.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    log.Println(err.Error())
    http.Error(rw, "Internal Error", 500)
    return
  }
  err = ts.Execute(rw, nil)
  if err != nil{
    log.Println(err.Error())
    http.Error(rw, "Internal Error", 500)
    return
  }
}

//http://localhost:4000/snippet?id=123
func showSnippet(rw http.ResponseWriter, r *http.Request){

  //Pegamos o id da URL e convertemos de string para int e já tratamos possíveis erros no back-end
  id,err := strconv.Atoi(r.URL.Query().Get("id"))
  //Tratamos o erro para quando o id não puder ser convertido para inteiro por não ser um número e quando o id for uma string de um número negativo
  if err != nil || id < 1 {
    http.NotFound(rw, r)
    return
  }
  fmt.Fprintf(rw, "Exibir o Snippet de ID: %d", id)
}

func createSnippet(rw http.ResponseWriter, r *http.Request){
  if r.Method != "POST"{
    //Tem que ser colocado acima dos códigos WriteHeader e http.Error
    rw.Header().Set("Allow","POST")
    http.Error(rw, "Metodo não permitido", http.StatusMethodNotAllowed)
    return
  }
  
  rw.Write([]byte("Criar novo snippet"))
}