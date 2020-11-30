# transaction-challenge

Este projeto é a implementação de uma API REST de integração para cadastro de transações na API da Pagar.me.

### Ferramentas
- Linguagem: [Go](https://golang.org/ "Go")

### Dependências
- Gin (go get github.com/gin-gonic/gin)
- Viper (go get github.com/spf13/viper)

### Arquivo de configuração: *config.yml*
Exemplo:
```yaml
pagarmeapi:
  baseUrl: https://api.pagar.me/1
  key: SUA_API_KEY

server:
  address: :8080
```
- **pagarmeapi**: configurações da API da Pagar.me
	- **baseUrl**: endereço base para acesso à API da Pagar.me
  - **key**: key para acesso à API da Pagar.me
- **server**: configurações do servidor da API
	- **address**: endereço e porta de acesso à API

#### Adicionar um planeta (com nome, clima e terreno)

> Método: POST
Endpoint: /v1/transactions

- **Campos do corpo**:
  - **auth_type**: tipo de autenticação na API da Pagar.me [obrigatório]
    - **basic**: a api_key é enviada como Basic Auth
    - **body**: a api_key é enviada no corpo da requisição
    - **url**: a api_key é enviada na url (query param)
  - **transaction**: objeto "Transação" da API da Pagar.me descrito [aqui](https://docs.pagar.me/reference#criar-transacao "aqui") [obrigatório]

##### Exemplo requisição:
> POST /v1/transactions
```json
{
    "auth_type": "basic",
    "transaction": { ... }
}
```

------------

#### Usando localmente:
Para rodar a aplicação localmente é necessário executar os seguintes passos:
1. Instalar as ferramentas abaixo na máquina local:
	- Go v1.14.6+
	- Docker
2. Baixar a imagem do golang no Docker: **docker pull golang:1.14-alpine3.12**
3. Clonar esse repositório em qualquer diretório
4. Alterar o arquivo *config.yml* com as configurações desejadas
5. No diretório clonado, buildar a imagem da aplicação usando: **docker build --rm -t desafiogo-ana .**
5. Rodar a aplicação usando: **docker run -p 8080:8080 --rm desafiogo-ana**