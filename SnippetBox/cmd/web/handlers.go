package main

import(
	"net/http"
	"strconv"
  "fmt"
)

func home(rw http.ResponseWriter, r *http.Request){

  //Teste para verificar se um caminho está ou não no código, caso não esteja será enviado ao usuário um erro 
  if r.URL.Path != "/"{
     http.NotFound(rw, r)
    return
  }
  rw.Write([]byte("Bem vindo ao SnippetBox"))
}

func showSnippet(rw http.ResponseWriter, r *http.Request){
  
  //Pegamos o id da URL e convertemos de string para int e já tratamos possíveis erros no back-end
  id, err := strconv.Atoi(r.URL.Query().Get("id"))
  //Tratamos o erro para quando o id não puder ser convertido para inteiro e quando o id for negativo
  if err != nil || id < 1{
      http.NotFound(rw, r)
      return
  }
  fmt.Fprintf(rw, "Exibir o Snippet de ID: %d", id)
}

func createSnippet(rw http.ResponseWriter, r *http.Request){

  if r.Method != "POST"{
    //Tem que ser colocado acima dos códigos WriteHeader e http.Error
    rw.Header().Set("Allow", "POST")
     
    rw.WriteHeader(405) //Erro: Não aceito esse método

    http.Error(rw, "Método não permitido", http.StatusMethodNotAllowed)
    return
  }
  
  rw.Write([]byte("Criar novo Snippet"))
}
