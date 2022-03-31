/* curl -i -X POST http://localhost:4000/snippet/create */
package main

import (
	"log"
  "net/http"
)

func home(rw http.ResponseWriter, r *http.Request){

  //Teste para verificar se um caminho está ou não no código, caso não esteja será enviado ao usuário um erro. 
  if r.URL.Path != "/"{
     http.NotFound(rw, r)
    return
  }
  rw.Write([]byte("Bem vindo ao SnippetBox"))
}

func showSnippet(rw http.ResponseWriter, r *http.Request){

  rw.Write([]byte("Mostrar um Snippet específico"))
}

func createSnippet(rw http.ResponseWriter, r *http.Request){

  if r.Method != "POST"{
    //Tem que ser colocado acima dos códigos WriteHeader e http.Error.
    rw.Header().Set("Allow", "POST")
     
    rw.WriteHeader(405) //Erro: Não aceito esse método
    
    //rw.Write([]byte("Método não permitido"))
    http.Error(rw, "Método não permitido", http.StatusMethodNotAllowed)
    return
  }
  
  rw.Write([]byte("Criar novo Snippet"))
}

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
