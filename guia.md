# Guia de Arquitetura do CORE

**Server-First Web Framework**

## 1. Visão Arquitetural

O CORE adota uma arquitetura Server-First, onde o servidor é o centro do sistema.

Não existe uma “aplicação no navegador”.
Existe apenas um cliente de renderização que recebe HTML produzido pelo servidor.

> A interface não decide. Ela reflete.

## 2. Princípio Fundamental

O princípio que governa toda a arquitetura do CORE é:

**Estado único no servidor.**

Tudo deriva disso:

- fluxo
- UI
- permissões
- conteúdo
- comportamento

Se o estado está correto, a interface também estará.

## 3. Camadas do CORE

A arquitetura do CORE é intencionalmente pouca profunda, para reduzir acoplamento e facilitar entendimento.

```
┌────────────┐
│  Browser   │
└─────┬──────┘
      │ HTTP / Event
┌─────▼──────┐
│   Router   │
└─────┬──────┘
      │ Context
┌─────▼─────────┐
│  Middleware   │
└─────┬─────────┘
      │ State
┌─────▼─────────┐
│   Application │
│   Logic (Go)  │
└─────┬─────────┘
      │ View Model
┌─────▼─────────┐
│   Renderer    │
│   (HTML)      │
└─────┬─────────┘
      │ HTML
┌─────▼─────────┐
│   Browser     │
└───────────────┘
```

Cada camada tem responsabilidade única.

## 4. Router

O Router é responsável apenas por:

- mapear URLs
- identificar método HTTP
- extrair parâmetros
- criar o contexto inicial

Ele não contém lógica de negócio.
Ele não renderiza UI.

## 5. Contexto de Execução

Cada requisição cria um `Context`.

O `Context` contém:

- request
- response
- sessão
- estado
- parâmetros
- usuário (se autenticado)

Esse `Context` é imutável por padrão, com alterações explícitas.

## 6. Middleware

Middlewares no CORE formam uma cadeia previsível.

Responsabilidades típicas:

- logging
- autenticação
- autorização
- carregamento de sessão
- proteção CSRF
- métricas

**Regra importante:** Middleware não decide UI.

Ele apenas:

- permite
- bloqueia
- modifica o contexto

## 7. Gerenciamento de Estado

O estado no CORE:

- vive no servidor
- é isolado por sessão
- é serializável
- é observável
- é auditável

Exemplo conceitual:

```go
type State struct {
  User    *User
  Flash   []Message
  Data    map[string]any
}
```

Não existe:

- estado no DOM
- estado em JavaScript
- sincronização frágil

## 8. Application Logic

A lógica da aplicação:

- é escrita em Go
- é testável
- não depende de UI
- não conhece HTML diretamente

> Ela produz estado, não markup.

## 9. Renderer

O Renderer:

- recebe estado
- escolhe página
- aplica layout
- renderiza componentes
- gera HTML final

Renderização sempre acontece:

- no servidor
- de forma determinística

## 10. Pages, Layouts e Components

- **Pages:** Ponto de entrada, associadas a rotas, consomem estado.
- **Layouts:** Estrutura global, composição de páginas.
- **Components:** Fragmentos reutilizáveis, sem estado próprio, recebem dados explicitamente.

Tudo é renderizado no servidor.

## 11. Atualizações Dinâmicas

O CORE suporta:

- requisições HTTP normais
- atualizações parciais
- eventos opcionais (SSE / WebSocket)

Mesmo com tempo real:

- o estado continua no servidor
- o HTML continua sendo a saída

## 12. Erros e Falhas

No CORE:

- erros são tratados no servidor
- falhas não quebram o cliente
- mensagens são parte do estado
- o HTML sempre representa a verdade atual

Sem “estado quebrado” no navegador.

## 13. Segurança

A arquitetura do CORE reduz a superfície de ataque:

- menos JS
- menos lógica no cliente
- sessões server-side
- validação centralizada
- políticas claras

> Segurança não é plugin, é consequência do design.

## 14. SEO e Acessibilidade

Como o HTML é:

- completo
- semântico
- entregue no primeiro request

O CORE:

- é naturalmente indexável
- funciona bem com leitores de tela
- respeita padrões web
- não depende de execução de scripts

## 15. Escalabilidade

A escalabilidade do CORE é horizontal:

- estado isolado por sessão
- renderização determinística
- fácil cache
- balanceamento simples

Sem estado distribuído no cliente.

## 16. Por que não SPA

SPAs:

- duplicam estado
- exigem hidratação
- aumentam complexidade
- criam bugs sutis
- envelhecem mal

O CORE remove essas classes de problemas por design.

## 17. Evolução Arquitetural

O CORE evolui:

- adicionando capacidades
- sem quebrar fundamentos
- mantendo simplicidade
- respeitando contratos

Mudanças grandes exigem justificativa arquitetural clara.

## 18. Princípio Final

> Se não pode ser explicado facilmente, não pertence ao CORE.

A arquitetura existe para reduzir a carga cognitiva, não aumentá-la.

### Conclusão

O CORE não tenta ser tudo.
Ele tenta ser **correto, previsível e durável**.

Sua arquitetura reflete uma escolha consciente:

- menos magia, mais controle
- menos hype, mais engenharia
