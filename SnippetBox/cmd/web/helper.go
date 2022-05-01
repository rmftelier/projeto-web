//Helper -> Métodos para nossa aplicação
package main

import(
  "net/http"
  "fmt"
  "runtime/debug"
) 

//Função/Método que agrupa os erros de servidor e faz com que não seja necessária a repetição de código, ou seja, o programa respeita o princípio DRY agora.
func(app *application) serverError(w http.ResponseWriter, err error){
    //Sprintf -> retorna uma string
    trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

    //Vinculação do erro ao lugar onde o erro aconteceu realmente, nesse caso: dois níveis acima do atual. Quando utilizamos o print ele mostra o erro como se acontecendo na linha em que o print está declarado o que não é o que realmente ocorre. 
    app.errorLog.Output(2, trace)
  
    //Erro não mostrado para o usuário, informação útil para o programador
    app.errorLog.Println(trace)

    //Erro mostrado ao usuário
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func(app *application) clientError(w http.ResponseWriter, status int){
  http.Error(w, http.StatusText(status), status)
}

func(app *application) notFound(w http.ResponseWriter){
   app.clientError(w, http.StatusNotFound) 
}

//Com base nesses módulos, nós podemos criar erros personalizados. 