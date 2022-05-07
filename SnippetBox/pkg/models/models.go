package models

import(
  "errors"
  "time"
)

var ErrNoRecord = errors.New("models: no matching record Found")

//Responsável por guardar as informações do Banco de Dados
type Snippet struct{
     ID int
     Title string 
     Content string 
     Created time.Time
     Expires time.Time
}