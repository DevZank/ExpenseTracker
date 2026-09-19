# 💸 ExpenseTracker API

Uma API REST simples e eficiente para gerenciamento de despesas pessoais, construída com **Go** e **PostgreSQL**.

---

## 🚀 Tecnologias

| Tecnologia | Descrição |
|---|---|
| [Go](https://golang.org/) | Linguagem principal |
| [Gorilla Mux](https://github.com/gorilla/mux) | Roteamento HTTP |
| [PostgreSQL](https://www.postgresql.org/) | Banco de dados relacional |
| [lib/pq](https://github.com/lib/pq) | Driver PostgreSQL para Go |
| [godotenv](https://github.com/joho/godotenv) | Carregamento de variáveis de ambiente |

---

## 📁 Estrutura do Projeto

```
ExpenseTracker/
├── config/
│   └── db.go            # Configuração e conexão com o banco de dados
├── handlers/
│   └── expenseTracker.go # Handlers HTTP (lógica dos endpoints)
├── models/
│   └── expense.go       # Model da despesa + SQL de criação da tabela
├── main.go              # Entry point da aplicação
├── go.mod
├── go.sum
└── .env                 # Variáveis de ambiente (não versionado)
```

---

## ⚙️ Configuração

### Pré-requisitos

- [Go 1.21+](https://golang.org/dl/)
- [PostgreSQL](https://www.postgresql.org/download/) rodando localmente ou em um servidor

### 1. Clone o repositório

```bash
git clone https://github.com/DevZank/ExpenseTracker.git
cd ExpenseTracker
```

### 2. Configure as variáveis de ambiente

Crie um arquivo `.env` na raiz do projeto com as suas credenciais do PostgreSQL:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_NAME=nome_do_banco
```

### 3. Instale as dependências

```bash
go mod tidy
```

### 4. Execute a aplicação

```bash
go run main.go
```

A API estará disponível em: `http://localhost:8080`

> A tabela `expenses` é criada automaticamente no banco de dados na primeira execução.

---

## 📌 Endpoints

### Base URL: `http://localhost:8080`

| Método | Rota | Descrição |
|--------|------|-----------|
| `GET` | `/expenses` | Lista todas as despesas |
| `POST` | `/expenses` | Cria uma nova despesa |
| `PUT` | `/expenses/{id}` | Atualiza uma despesa existente |
| `DELETE` | `/expenses/{id}` | Remove uma despesa |

---

### 📋 Modelo de Despesa

```json
{
  "id": 1,
  "description": "Mercado",
  "category": "Alimentação",
  "value": 150.75,
  "date": "2026-09-19"
}
```

| Campo | Tipo | Obrigatório | Descrição |
|-------|------|-------------|-----------|
| `description` | `string` | ✅ | Descrição da despesa (máx. 150 chars) |
| `category` | `string` | ✅ | Categoria da despesa (máx. 40 chars) |
| `value` | `float` | ✅ | Valor da despesa |
| `date` | `string` | ❌ | Data da despesa |

---

### 📖 Exemplos de Requisição

#### GET /expenses
```bash
curl -X GET http://localhost:8080/expenses
```

**Resposta (200 OK):**
```json
[
  {
    "id": 1,
    "description": "Mercado",
    "category": "Alimentação",
    "value": 150.75,
    "date": "2026-09-19"
  }
]
```

---

#### POST /expenses
```bash
curl -X POST http://localhost:8080/expenses \
  -H "Content-Type: application/json" \
  -d '{
    "description": "Netflix",
    "category": "Assinatura",
    "value": 39.90,
    "date": "2026-09-19"
  }'
```

**Resposta (201 Created):**
```json
{
  "id": 2,
  "description": "Netflix",
  "category": "Assinatura",
  "value": 39.90,
  "date": "2026-09-19"
}
```

---

#### PUT /expenses/{id}
```bash
curl -X PUT http://localhost:8080/expenses/2 \
  -H "Content-Type: application/json" \
  -d '{
    "description": "Netflix Premium",
    "category": "Assinatura",
    "value": 55.90,
    "date": "2026-09-19"
  }'
```

**Resposta (200 OK):**
```json
{
  "id": 2,
  "description": "Netflix Premium",
  "category": "Assinatura",
  "value": 55.90,
  "date": "2026-09-19"
}
```

---

#### DELETE /expenses/{id}
```bash
curl -X DELETE http://localhost:8080/expenses/2
```

**Resposta (204 No Content)**

---

## 🗄️ Schema do Banco de Dados

```sql
CREATE TABLE IF NOT EXISTS expenses (
    id          SERIAL PRIMARY KEY,
    description VARCHAR(150) NOT NULL,
    category    VARCHAR(40)  NOT NULL,
    value       FLOAT        NOT NULL,
    date        VARCHAR(60)
);
```

