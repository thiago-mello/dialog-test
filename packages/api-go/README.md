# API de Rede Social em Go 🚀

Uma API simples para uma rede social desenvolvida em Go, utilizando conceitos de Arquitetura Hexagonal, DDD e CQRS.

## 📂 Estrutura de Pastas

```yaml
.
├── db/ # Scripts SQL e templates de consultas
│ └── sql/
│ ├── posts/ # Queries relacionadas a posts (likes, posts)
│ └── users/ # Queries relacionadas a usuários
│
├── src/ # Código-fonte principal
│ ├── config/ # Configurações globais (banco, cache, servidor)
│ ├── libs/ # Bibliotecas internas
│ │ ├── application/ # Lógica de aplicação (middlewares, validação)
│ │ ├── cache/ # Implementação de cache (Redis)
│ │ └── utils/ # Utilitários (banco de dados, paginação, etc.)
│ │
│ ├── modules/ # Módulos da aplicação
│ │ ├── posts/ # Funcionalidades de posts (CRUD, likes)
│ │ └── users/ # Funcionalidades de usuários (autenticação, CRUD)
│ │
│ ├── routes/ # Configuração das rotas da API
│ └── tests/ # Testes de integração e scripts auxiliares
│
└── main.go # Ponto de entrada da aplicação
```

## 🚀 Como Construir e Executar

### Pré-requisitos

- Go 1.20+
- PostgreSQL 16+ e Redis (ou containers Docker)
- Arquivo `config.yaml` (copie de `config.example.yaml` na raiz do projeto)

### Passos:

1. **Configuração do Ambiente**:
   ```bash
   cp config.example.yaml config.yaml
   ```

Preencha o `config.yaml` com suas credenciais de banco e Redis e adicione uma string de segredo JWT para assinar os tokens.

2. **Instalar Dependências**:

   ```bash
   go mod tidy
   ```

3. **Executar Migrações**:

   - Certifique-se de que o PostgreSQL está rodando.
   - Execute o script `db/init.sql` para criar tabelas.

4. **Build e Execução**:

   ```bash
   go build -o api && ./api
   ```

5. **Testar Endpoints**:
   - Use ferramentas como [Postman](https://www.postman.com/) ou [curl](https://curl.se/).
   - Exemplo de endpoint: `GET /v1/posts` para listar posts públicos.

---

## ⚙️ Configuração Detalhada

### Arquivo `config.yaml`

```yaml
database:
  relational:
    host: localhost
    port: 5432
    database-name: nome_do_banco
    auth:
      user: seu_usuario
      password: sua_senha
  redis:
    address: localhost:6379
    password:
    db: 0

sql:
  templates:
    path: ./db/sql/**/*.sql

server:
  port: 8080
  debug: true

api:
  jwt:
    secret: "chave_secreta_para_jwt"
    expires-in: 3h
```

## 🧪 Testes

Para executar testes unitários:

```bash
go test -v ./src/...
```

## 📖 Documentação

A referência de endpoints pode ser acessada por meio da requisição `GET /docs/index.html` em um navegador de internet.

## 📝 Notas Adicionais

- **Cache**: As consultas a posts e usuários são cacheadas via Redis para melhor desempenho.
- **Segurança**: Senhas são hasheadas com Argon2id e tokens JWT são assinados com HMAC.
- **DDD**: Módulos seguem alguns princípios de Domain-Driven Design, como _bounded contexts_.
