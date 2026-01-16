# Primeiros Passos com o CORE

## O que é o CORE

CORE é um framework web server-first, onde:

- o servidor controla o estado
- o HTML é o contrato
- o cliente apenas renderiza
- JavaScript é opcional

> Você não constrói páginas.
> Você constrói **estado → renderização**.

## Pré-requisitos

- Go 1.21+
- Conhecimento básico de HTTP e HTML
- Nenhuma dependência de frontend

## Criando seu primeiro projeto

```bash
core new app meu-site
cd meu-site
core dev
```

Acesse: `http://localhost:8080`

### Estrutura inicial do projeto

```
meu-site/
├─ cmd/server/main.go
├─ internal/
│  ├─ router/
│  ├─ middleware/
│  ├─ state/
│  ├─ render/
│  └─ session/
├─ ui/
│  ├─ layouts/
│  │  └─ main.layout.core
│  ├─ pages/
│  │  └─ home.page.core
│  └─ components/
├─ static/
│  └─ css/
└─ core.config.toml
```

## Criando uma página

```bash
core new page dashboard
```

Isso cria `ui/pages/dashboard.page.core`.

## Fluxo mental do CORE

1.  Request chega
2.  Estado é carregado
3.  Lógica roda no servidor
4.  HTML é renderizado
5.  Resposta enviada

Não existe lógica de UI no cliente.

---

## 📄 Especificação dos Arquivos .core

### 1. Pages (`*.page.core`)

Representam rotas e estados principais.

**Exemplo:** `home.page.core`

**Responsabilidades:**

- Associar rota
- Definir layout
- Consumir estado
- Renderizar conteúdo

### 2. Layouts (`*.layout.core`)

Estrutura global da aplicação.

**Exemplo:** `main.layout.core`

**Responsabilidades:**

- HTML base
- Slots de conteúdo (`{{template "content" .}}`)
- Header / Footer
- Navegação

### 3. Components (`*.comp.core`)

Fragmentos reutilizáveis de UI.

**Exemplo:** `button.comp.core`, `alert.comp.core`

**Características:**

- Renderizados no servidor
- Recebem estado como parâmetro
- Não possuem estado próprio no cliente
- Reutilizáveis entre páginas

### Regra fundamental

> Nenhum arquivo `.core` contém lógica de negócio complexa.
> A lógica vive no Go, a UI é a projeção.
