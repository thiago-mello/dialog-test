# Rede Social Simples - Next.js

Uma rede social minimalista baseada em texto, desenvolvida com Next.js (App Router) e TypeScript. Permite que usuários criem postagens, interajam com conteúdo e gerenciem seus perfis.

## 🚀 Funcionalidades Principais

- **Autenticação**: Login e registro de usuários
- **Postagens**:
  - Criação/edição de textos com editor rich-text (TipTap)
  - Likes e exclusão de posts
  - Feed público e feed dos próprios posts
- **Perfil**:
  - Atualização de informações
  - Exclusão de conta
- UI Responsiva com sidebar colapsável
- Paginação infinita no feed
- Toasts para feedback de ações

## 🛠️ Tecnologias

- **Frontend**:
  - Next.js 15 (App Router)
  - TypeScript
  - Tailwind CSS + Shadcn/ui
  - React Hook Form + Zod (validação)
  - TanStack Query (gerenciamento de estado e consultas)
- **Autenticação**: Iron Session (sessões criptografadas)
- **Editor de Texto**: TipTap (com extensões para markdown básico)
- **Ferramentas**:
  - date-fns (formatação de datas)
  - Lucide React (ícones)

## 📁 Estrutura de Arquivos e Pastas

```plaintext
.
├── actions/            → Lógica de API e Server Actions
│   ├── api/            → Endpoints específicos (posts, users)
│   └── login.ts        → Autenticação
├── app/                → Páginas e componentes do Next.js
│   ├── posts/          → Funcionalidades de postagem
│   ├── profile/        → Gerenciamento de perfil
│   └── timeline/       → Feed de postagens
├── components/         → Componentes reutilizáveis
│   ├── ui/             → Componentes Shadcn/ui personalizados
│   └── app-sidebar.tsx → Sidebar principal
├── constants/          → Constantes globais
├── hooks/              → Custom hooks
├── lib/                → Utilitários de baixo nível
├── providers/          → Context providers
└── utils/              → Funções utilitárias
```

## 🔧 Decisões Técnicas

### 1. Arquitetura de API

- **Server Actions**: Utilizadas para operações CRUD diretamente do frontend
- **Tipagem Forte**: Interfaces TypeScript para as respostas de API
- **Tratamento de Erros**: Padronização com objetos `ApiResponse<T>`

### 2. Gerenciamento de Estado

- **TanStack Query**: Para caching e sincronização de dados do feed
- **Otimistic Updates**: Atualização imediata da UI em likes/exclusões de likes
- **Session Management**: Iron Session com cookies criptografados

### 3. Componentização

- **Shadcn/ui**: Componentes acessíveis e personalizáveis
- **Separação Clara**:
  - `components/ui/`: Building blocks genéricos
  - `app/components/`: Componentes específicos de páginas

### 4. Performance

- **Loading States**: Skeletons durante carregamentos
- **Infinite Scroll**: Paginação otimizada com Intersection Observer
- **Compressão CSS**: Tailwind com purge de classes não utilizadas

## ▶️ Como Executar

1. Instale as dependências:

```bash
npm install
```

2. Configure variáveis de ambiente (.env):

```env
API_BASE_URL=urlDoBackend:porta
SECRET_COOKIE_PASSWORD=suachavesupersecreta
```

3. Inicie o servidor de desenvolvimento:

```bash
npm run dev
```

Certifique-se de que a API que alimenta o projeto esteja em execução.
