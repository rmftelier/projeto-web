package mysql

import(
   "database/sql"

  "github.com/rmftelier/projeto-web/pkg/models"
)


type SnippetModel struct{
    DB *sql.DB
}


//Função para Inserir um Snippet no BD
func(m *SnippetModel)Insert(title, content, expires string) (int, error){

   stmt := `INSERT INTO snippets (title, content, created, expires)
            VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

  result, err := m.DB.Exec(stmt, title, content, expires)
  if err != nil{
     return 0, err
  }

  id, err := result.LastInsertId()
  if err != nil{
     return 0, err
  }
  return int(id), nil 
}


//Função para
func(m *SnippetModel) Get(id int)(*models.Snippet, error){
   return nil, nil
}

//Função para
func(m *SnippetModel) Latest()([]*models.Snippet, error){
   return nil, nil
}