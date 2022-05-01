package main

//go run cmd/web/*
import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

//A parte (app *application) significa que a função home pertence a application
func(app *application) home(rw http.ResponseWriter, r *http.Request){
  //Teste para verificar se um caminho está ou não no código, caso não esteja será enviado ao usuário um erro
  if r.URL.Path != "/"{
    app.notFound(rw) //Chamamos o módulo criado no helper.go para esse erro 
    return
  } 
 
  files := []string{
    "./ui/html/home.page.tmpl.html",
    "./ui/html/base.layout.tmpl.html",
    "./ui/html/footer.partial.tmpl.html",
  }
  
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  
  err = ts.Execute(rw, nil)
  if err != nil{
    app.serverError(rw, err)
    return
  }  
}

//http://localhost:4000/snippet?id=123
func(app *application) showSnippet(rw http.ResponseWriter, r *http.Request){

  //Pegamos o id da URL e convertemos de string para int e já tratamos possíveis erros no back-end
  id,err := strconv.Atoi(r.URL.Query().Get("id"))
  //Tratamos o erro para quando o id não puder ser convertido para inteiro por não ser um número e quando o id for uma string de um número negativo
  if err != nil || id < 1 {
    app.notFound(rw)
    return
  }
  fmt.Fprintf(rw, "Exibir o Snippet de ID: %d", id)
}

func(app *application) createSnippet(rw http.ResponseWriter, r *http.Request){
  if r.Method != "POST"{
    //Tem que ser colocado acima dos códigos WriteHeader e http.Error
    rw.Header().Set("Allow","POST")
    app.clientError(rw, http.StatusMethodNotAllowed)
    return
  }
  
  rw.Write([]byte("Criar novo snippet"))
}