# 🧩 Padrões de Projeto no CORE

Esses padrões não são opcionais.
Eles são essenciais para manter uma aplicação CORE saudável e sustentável.

### 1. State Projection Pattern (Padrão de Projeção de Estado)

A UI é **sempre** uma projeção do estado que vive no servidor.

`Estado → Visão → HTML`

Nunca o contrário (`UI → Lógica → Estado`). A UI não origina mudanças de estado, ela apenas envia eventos para o servidor, que é o único que pode alterar o estado.

### 2. Thin Handler Pattern (Padrão de Handler Fino)

Os handlers HTTP (controladores) devem ser mínimos.

**Responsabilidades:**

1.  Receber o contexto da requisição.
2.  Chamar serviços ou lógica de negócio.
3.  Retornar o estado resultante para o renderer.

Eles não devem conter lógica de negócio pesada.

### 3. Explicit Data Flow (Padrão de Fluxo de Dados Explícito)

Nada de mágica. Nada de dados implícitos.

O estado que uma página ou componente usa deve ser passado explicitamente.

**Exemplo:**
`ctx.State.Set("currentUser", user)`

Isso torna o fluxo de dados fácil de rastrear e depurar.

### 4. No Client Logic Pattern (Padrão de Lógica Zero no Cliente)

O cliente (navegador) não decide.
Ele apenas renderiza o HTML que recebe e envia eventos (cliques, submissões de formulário) para o servidor.

### 5. Layout Composition Pattern (Padrão de Composição de Layout)

Layouts contêm e compõem páginas.
Páginas não controlam o layout.

`{{define "base"}}...{{template "content" .}}...{{end}}`

Isso garante consistência visual e separa a estrutura global do conteúdo específico da página.

### 6. Failure-as-State Pattern (Padrão de Falha como Estado)

Um erro não é uma exceção que quebra a renderização.
Um erro é apenas mais uma parte do estado.

```go
// Exemplo de estado com erro
State.Set("error", "Credenciais inválidas")
```

A UI então projeta esse estado, mostrando a mensagem de erro.

### 7. Progressive Enhancement (Padrão de Melhoria Progressiva)

Toda a funcionalidade central deve funcionar perfeitamente sem JavaScript.
JavaScript é usado apenas para *melhorar* a experiência (ex: atualizações parciais, validações em tempo real), não para habilitá-la.

### 8. Stability First Rule (Regra de Estabilidade Primeiro)

Se uma nova solução, padrão ou dependência:

- aumenta a complexidade desnecessariamente
- dificulta a depuração
- cria estado oculto ou implícito

...ela não deve ser adotada no CORE.

---

## 📌 Conclusão Geral

O CORE não é apenas um framework.
Ele é um **sistema coerente de decisões arquiteturais**.

Cada documento e cada padrão existem para garantir:

- **Previsibilidade**
- **Segurança**
- **Longevidade**
- **Clareza**
