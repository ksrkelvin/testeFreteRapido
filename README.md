# testeFreteRapido
Teste técnico para empresa Frete Rápido
# FreteRápido 🚛

Este repositório contém a resolução do **Desafio Back-end #2** proposto pela FreteRápido. A aplicação implementa uma API REST para cotações fictícias e métricas de fretes.

## 🧠 Funcionalidades

1. **Cotação de frete** `[POST] /quote`
   Recebe dados de entrada e realiza uma cotação fictícia utilizando a API da FreteRápido.

   * Armazena os resultados no banco de dados para consultas posteriores.
* Entrada esperada:

```json
{
	"recipient":{
		"address":{
				"zipcode":"01311000"
			}
		},
   "volumes":[
		 {
			"category":7,
			"amount":1,
			"unitary_weight":5,
			"price":349,
			"sku":"abc-teste-123",
			"height":0.2,
			"width":0.2,
			"length":0.2,
			"unitary_price": 10
		},
		{
			"category":7,
			"amount":2,
			"unitary_weight":4,
			"price":556,
			"sku":"abc-teste-527",
			"height":0.4,
			"width":0.6,
			"length":0.15,
			"unitary_price": 10
			}
		]
}
```
* Retorno esperado:

```json
{
	"carrier": [
		{
			"name": "CORREIOS - SEDEX",
			"service": "SEDEX",
			"deadline": "1",
			"price": 92.37
		},
		{
			"name": "CORREIOS",
			"service": "SEDEX",
			"deadline": "1",
			"price": 92.37
		},
		{
			"name": "CORREIOS - SEDEX",
			"service": "PAC",
			"deadline": "5",
			"price": 112.74
		},
		{
			"name": "CORREIOS",
			"service": "PAC",
			"deadline": "5",
			"price": 112.74
		}
	]
}
```
---

2. **Métricas de cotações** `[GET] /metrics?last_quotes={?}`
   Consulta métricas das cotações armazenadas no banco de dados.

   * Parâmetro opcional: `last_quotes` (quantidade de cotações recentes).
   * Métricas retornadas:
     * Quantidade de resultados por transportadora
     * Total de preço por transportadora
     * Média de preço por transportadora
     * Frete mais barato geral
     * Frete mais caro geral

* Entrada esperada:

```json
	{}
```
* Retorno esperado:
```json
{
	"CarrierMetrics": [
		{
			"CarrierName": "CORREIOS - SEDEX",
			"TotalQuotes": 2,
			"TotalPrice": 205.11,
			"AveragePrice": 102.555
		},
		{
			"CarrierName": "CORREIOS",
			"TotalQuotes": 2,
			"TotalPrice": 205.11,
			"AveragePrice": 102.555
		}
	],
	"CheapestCarrier": {
		"CarrierName": "CORREIOS - SEDEX",
		"Price": 92.37
	},
	"ExpensiveCarrier": {
		"CarrierName": "CORREIOS - SEDEX",
		"Price": 112.74
	}
}
```
---

## ▶️ Como executar a aplicação

A aplicação é **containerizada** usando Docker. Todos os comandos abaixo devem ser executados no diretório raiz do projeto.

### 1️⃣ Variáveis de ambiente

Esse projeto precisa das seguintes variaveis de ambiente 

```env
FRETE_RAPIDO_API_AUTH_TOKEN= (Usar Token de autenticação da API frete Rápido)
FRETE_RAPIDO_API_PLATAFORM_CODE= (Usar Código da Plataforma da API frete Rápido)

```
---
### 2️⃣ Build & Run

Primeiramente é necessário atribuir as variaveis de ambiente no arquivo docker/docker-compose.yml:
      FRETE_RAPIDO_API_AUTH_TOKEN
      FRETE_RAPIDO_API_PLATAFORM_CODE
Após isso executar o comando abaixo a partir da pasta raiz do projeto 

```bash
docker-compose -f docker/docker-compose.yml up --build
```

A API estará disponível em: `http://localhost:8080`

---

## 📁 Estrutura do projeto

```
C:.
│   .env
│   .gitignore
│   go.mod
│   go.sum
│   README.md
│   
├───.vscode
│       launch.json
│       
├───cmd
│   └───server
│           main.go
│           __debug_bin.exe
│
├───docker
│       docker-compose.yml
│       Dockerfile
│
├───internal
│   ├───adapter
│   │   ├───gateways
│   │   │   └───freterapido
│   │   │           freterapido_adapter.go
│   │   │           freterapido_constants.go
│   │   │           freterapido_gateway.go
│   │   │
│   │   ├───http
│   │   │   │   router.go
│   │   │   │
│   │   │   ├───dtos
│   │   │   │       metrics_response.go
│   │   │   │       quote_request.go
│   │   │   │       quote_response.go
│   │   │   │
│   │   │   ├───handlers
│   │   │   │       metrics_handler.go
│   │   │   │       quote_handler.go
│   │   │   │
│   │   │   └───mappers
│   │   │           metrics_http_mapper.go
│   │   │           quote_http_mapper.go
│   │   │
│   │   └───repositories
│   │       │   repository.go
│   │       │
│   │       ├───mappers
│   │       │       metrics_repository_mapper.go
│   │       │       quote_repository_mapper.go
│   │       │
│   │       ├───migrations
│   │       │       quote_migrations.go
│   │       │
│   │       ├───models
│   │       │       quote_model.go
│   │       │
│   │       └───postgres
│   │               metrics_repository.go
│   │               quote_repository.go
│   │
│   ├───config
│   │       config.go
│   │       contstants.go
│   │       database.go
│   │       env.go
│   │
│   ├───domain
│   │   ├───entity
│   │   │       metrics_entity.go
│   │   │       quote_entity.go
│   │   │
│   │   └───interfaces
│   │           metrics_repository.go
│   │           quote_gateway.go
│   │           quote_repository.go
│   │
│   └───usecase
│       ├───metrics
│       │       metrics_output.go
│       │       metrics_usecase.go
│       │
│       └───quote
│               quote_input.go
│               quote_output.go
│               quote_usecase.go
│
└───pkg
    ├───freteRapidoApi
    │       api_error.go
    │       client.go
    │       constants.go
    │       cotacao_v3.go
    │       http_client.go
    │       models.go
    │
    └───httpx
            app_error.go
            factory.go
            response.go
            result.go
```

---

## ✅ Tecnologias utilizadas

* Golang
* Docker & Docker Compose
* PostgreSQL
* Boas práticas de Clean Code e TDD

---

## 📌 Observações

* Todas as entradas são validadas antes de consumir a API externa.
* Mensagens de erro são claras e padronizadas.
* Dados sensíveis, como token e CNPJ, são configuráveis via `.env`.
* Rotas seguem o padrão REST com JSON de entrada e saída.
