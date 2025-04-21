# Rede Social Baseada em Texto 🌐

Uma rede social minimalista desenvolvida com **Go** (backend) e **Next.js** (frontend), focada em postagens de texto e interações sociais. Projeto construído como desafio técnico.

## 🚀 Funcionalidades Principais

### 👤 Autenticação e Perfil

- Cadastro e login com e-mail/senha
- Atualização de perfil (bio, informações)
- Exclusão de conta

### 📝 Postagens

- Criação/edição de posts com editor rich-text
- Definição de privacidade (público/privado)
- Feed público e feed pessoal com paginação infinita
- Curtidas em posts de usuários

### 🛠️ Gestão Técnica

- Cache de consultas via Redis
- Telemetria com OpenTelemetry e Grafana Tempo
- Documentação Swagger integrada (API)
- UI responsiva com sidebar colapsável

## 🛠️ Stack Tecnológica

### **Backend (API Go)**

- **Linguagem**: Go 1.20+
- **Banco de Dados**: PostgreSQL 16+
- **Cache**: Redis (cache de leituras frequentes)
- **Autenticação**: JWT + Argon2id para hashing de senhas
- **Arquitetura**: Hexagonal, DDD, CQRS
- **Ferramentas**: OpenTelemetry, Echo (Web), Sqlx (SQL)

### **Frontend (Next.js)**

- **Framework**: Next.js 15 (App Router)
- **Estilização**: Tailwind CSS + Shadcn/ui
- **Gerenciamento de Estado e consultas**: TanStack Query
- **Autenticação**: Iron Session (cookies criptografados)
- **Validação**: Zod + React Hook Form
- **Editor de Texto**: TipTap com suporte a Markdown básico

### **Outros**

- Docker Compose (PostgreSQL, Redis, Grafana Tempo)
- Paginação baseada em cursores (UUIDv7 ordenável)
- Sanitização HTML para prevenção de XSS

## 🏗️ Arquitetura e Decisões Técnicas

### **Backend**

- **Arquitetura**: Separação entre lógica de negócio, infraestrutura e interfaces, com conceitos de Arquitetura Hexagonal, CQRS e DDD.
- **Cache Estratégico**: Redis para consultas frequentes (ex: feed público)
- **Segurança**:
  - Senhas armazenadas com Argon2id ([recomendação OWASP](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html))
  - Tokens JWT assinados com chave HMAC
  - Sanitização de HTML em posts

### **Frontend**

- **Server Actions**: Para requisições à API do backend
- **Otimistic UI**: Atualizações instantâneas em likes/exclusões de likes
- **Componentização Modular**:
  - `shadcn/ui` para componentes acessíveis
  - Separação entre componentes genéricos (`ui/`) e específicos (`app/`)
- **Performance**:
  - Paginação infinita com Intersection Observer

## ⚡ Como Executar

### **Via Docker Compose (Recomendado)**

```bash
docker compose up
```

**Observação**: Ao executar o projeto via docker compose, a API do backend não está exposta para acesso fora dos containers, já que todas as interações com ela são feitas pelo servidor do Next.js.

Caso queira acessar a referência de endpoints via Swagger, é necessário adicionar o mapeamento de portas no o arquivo `docker-compose.yaml` na raiz do projeto da seguinte forma:

```yaml
backend: # definição de serviço do backend
  depends_on:
    - postgres
    - redis
  build: ./packages/api-go
  ports:
    - "8080:8080" # Adicionar o mapeamento de portas para permitir o acesso externo
```

**Contêineres Iniciados**:

- PostgreSQL (banco de dados)
- Redis (cache)
- API Go
- Frontend Next.js (`:3000`)
- Grafana Tempo (telemetria, somente traces)
- Grafana (visualização de traces)

O projeto será executado localmente e poderá ser acessado através da seguinte URL:

- Aplicação: [http://localhost:3000](http://localhost:3000)
- Painel do Grafana: [http://localhost:3001](http://localhost:3001)
  - **Usuário:** `admin`
  - **Senha:** `admin`

**Visualizando os traces da aplicação no Grafana**:

Para visualizar os traces no Grafana, é necessário configurar um **data source** do **Tempo** com a URL `http://tempo:3200`:

### **Execução Manual**

O código da API se encontra na pasta [packages/api-go](/packages/api-go/), e o código do frontend se encontra na pasta [packages/frontend-nextjs](/packages/frontend-nextjs/).

1. **API Go**:

   ```bash
   cd packages/api-go
   cp config.example.yaml config.yaml # Configure credenciais
   go mod tidy
   go build -o api && ./api
   ```

⚠️ **Importante**: É necessário ter uma instância do PostgreSQL em execução, com as tabelas definidas no arquivo [init.sql](/packages/api-go/db/init.sql) criadas no schema `public`.

As credenciais de acesso ao PostgreSQL e ao Redis devem ser configuradas no arquivo `config.yaml` criado.

2. **Frontend Next.js**:
   ```bash
   cd packages/frontend-nextjs
   npm install
   echo "API_BASE_URL=http://localhost:8080\nCOOKIE_PASSWORD=chave_secreta" > .env
   npm run dev
   ```

🔍 **Variáveis de Ambiente Cruciais**:

- `API_BASE_URL`: Endpoint da API Go
- `COOKIE_PASSWORD`: Chave para criptografia de sessões

## 📊 Monitoramento

### **Telemetria (Somente Traces)**

- Configure um endpoint OTLP no `config.yaml` da API para exportar traces ao Grafana Tempo ou a um coletor OTLP.

Por padrão, ao executar o projeto pelo Docker Compose conforme definido no arquivo `docker-compose.yaml`, os traces do projeto são exportados para o Grafana Tempo e podem ser visualizados no Grafana.

## 📚 Documentação

- **API**: Acesse `http://localhost:8080/docs/index.html` em um navegador de internet para acessar a UI do Swagger.

Se você optou por executar o projeto com Docker Compose, **certifique-se de seguir os passos descritos [nesta seção](#via-docker-compose-recomendado)** antes de iniciar os containers. Isso é necessário para garantir o acesso à documentação dos endpoints da API.

## ℹ️ Informações Adicionais

Informações mais detalhadas sobre o Backend e o Frontend do projeto estão disponíveis nos arquivos README localizados nas respectivas pastas: [API](/packages/api-go/README.md) e [Frontend](/packages/frontend-nextjs/README.md). Neles, você encontrará uma visão geral da estrutura dos arquivos de cada parte do projeto, bem como instruções específicas de execução e desenvolvimento.
