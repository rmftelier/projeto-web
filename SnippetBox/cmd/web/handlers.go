package main

//go run cmd/web/*
import (
  "fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/rmftelier/projeto-web/pkg/models"
)

//A parte (app *application) significa que a função home pertence a application
func(app *application) home(rw http.ResponseWriter, r *http.Request){
  //Teste para verificar se um caminho está ou não no código, caso não esteja será enviado ao usuário um erro
  if r.URL.Path != "/"{
    app.notFound(rw) //Chamamos o módulo criado no helper.go para esse erro 
    return
  } 
  
  snippets, err := app.snippets.Latest()
  if err != nil{
      app.serverError(rw, err)
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
  
  err = ts.Execute(rw, snippets)
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
    
    s, err := app.snippets.Get(id)
    if err == models.ErrNoRecord{
        app.notFound(rw)
        return
    }else if err != nil{
        app.serverError(rw, err)
        return
    }
  
    files := []string{
      "./ui/html/show.page.tmpl.html",
      "./ui/html/base.layout.tmpl.html",
      "./ui/html/footer.partial.tmpl.html",
    }
    
    ts, err := template.ParseFiles(files...)
    if err != nil{
      app.serverError(rw, err)
      return
    }
    //s é o snippet que foi lido do banco de dados
    err = ts.Execute(rw, s)
    if err != nil{
      app.serverError(rw, err)
      return
    }  
}

func(app *application) createSnippet(rw http.ResponseWriter, r *http.Request){
  if r.Method != "POST"{
    //Tem que ser colocado acima dos códigos WriteHeader e http.Error
    rw.Header().Set("Allow","POST")
    app.clientError(rw, http.StatusMethodNotAllowed)
    return
  }

	title := "Aula de hoje"
	content := "Tentando lidar com o banco de dados"
	expires := "19"

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(rw, err)
		return
	}
  
  http.Redirect(rw, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}