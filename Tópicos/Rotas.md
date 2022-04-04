# Anotações 

### Os componentes de uma URL
> Fonte: 🔗[IBM](https://www.ibm.com/docs/en/cics-ts/5.2?topic=concepts-components-url)

A URL possuí normalmente três ou quatro componentes:

1. **Scheme**: Identifica o protocolo a ser usado para acessar o recurso na internet. Pode ser HTTP (sem SSL) ou HTTPS (com SSL).
2. **Host**: O nome do host identifica o host que contém o recurso. Por exemplo, www.exemplo.com 
3. **Path**: O path, caminho em tradução literal, identifica o recurso específico no host que o cliente da Web deseja acessar. Por exemplo, /software/htp/ciscs/index.html.
4. **Query string:** Se uma string de consulta for usada, ela seguirá o componente de path, e fornecerá uma string de informações que o recurso pode usar para alguma finalidade (por exemplo, como parâmetros 
para uma pesquisa ou como dados a serem processados. A query string geralmente é uma string de pares de nome e valor; por exemplo, termo=azul. Os pares de nome e valor
são separados um do outro por um &; por exemplo, term=bluebird&source=browser-search. 

> Na URL: "uvv.com.br/alunos" 
> 
> "uvv.com.br" é o host 
> 
> "/alunos" é o path

---

### Rotas 

+ **Fixed Paths/Caminhos Absolutos**: Fixed Paths, ou Rotas Fixas em tradução literal, são rotas onde o handler só é 
chamado quando o caminho dela é exatamente igual ao que foi estabelecido anteriormente.
Exemplo: "/snippet" | "/snippet/create"

+ **Subtree Pathes/Caminhos Relativos**: Subtree Paths, ou Caminhos em Subarvores em tradução literal, são rotas onde o handler
será chamado se o início da url corresponder ao padrão, por exemplo o nosso "/".
É como se existisse um wildcard após o "/", exemplo *"/*", e isso significa que qualquer URL
que contenha "/" e que não seja igual as definidas como caminho absoluto fará com que o usuário seja levado de volta para
o home que está indicado como a função chamada no projeto. 

