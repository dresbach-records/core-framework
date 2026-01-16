# CORE — Server-First Web Framework

**CORE é um framework web server-first, focado em HTML semântico, estado único no servidor e estabilidade de longo prazo.**

O CORE remove a complexidade desnecessária do frontend, preservando uma experiência web moderna, rápida e acessível.

> Estabilidade é uma feature.
> O servidor possui a verdade.

## O que é o CORE?

O CORE **não é** um framework de frontend.

O CORE é o **núcleo** de uma aplicação web.

Em vez de empurrar a lógica e o estado para o navegador, o CORE mantém:

- o **estado** no servidor
- o **HTML** como o contrato da interface
- o **JavaScript** opcional

O navegador renderiza o que o servidor decide.

## Filosofia

O CORE é construído sobre alguns princípios rígidos:

1.  **Server-First**: O servidor controla o estado, o fluxo e as decisões.
2.  **HTML como Contrato**: O HTML é a interface primária entre o servidor e o navegador.
3.  **Fonte Única de Verdade**: Não há estado duplicado entre cliente e servidor.
4.  **JavaScript é Opcional**: O JS pode existir, mas nunca como o núcleo do sistema.
5.  **Previsibilidade sobre Efeitos**: Comportamento claro é mais importante que truques visuais.
6.  **Longevidade sobre Hype**: O CORE foi projetado para durar anos.

### O que o CORE NÃO é

- ❌ Não é um framework SPA
- ❌ Não é React, Vue ou Next
- ❌ Não é uma solução pesada no frontend
- ❌ Não é um projeto guiado por tendências

O CORE abraça a web como ela é, não como uma gambiarra.

## Quando usar o CORE

O CORE é ideal para:

- Sites institucionais
- Áreas de cliente
- Dashboards
- Painéis administrativos
- Sistemas de faturamento
- Plataformas de hospedagem
- Ferramentas internas
- Aplicações B2B
- Produtos de longa duração

Se a sua prioridade é **estabilidade, SEO, acessibilidade e manutenibilidade**, o CORE se encaixa naturalmente.

## Arquitetura do CORE

```
Navegador
  ↓ (requisição / evento)
Servidor CORE
  ↓ (estado + lógica)
HTML (completo ou parcial)
  ↓
Atualização do DOM
```

Não há hidratação, nem DOM virtual, nem lógica duplicada.

## Estrutura do Projeto

Um projeto CORE típico se parece com isto:

```
core-app/
├─ cmd/server/main.go
├─ internal/
│  ├─ router/
│  ├─ middleware/
│  ├─ state/
│  ├─ render/
│  ├─ session/
│  └─ ws/
├─ ui/
│  ├─ layouts/
│  │  └─ main.layout.core
│  ├─ pages/
│  │  ├─ home.page.core
│  │  └─ login.page.core
│  └─ components/
│     └─ button.comp.core
├─ static/
│  └─ css/
├─ core.config.toml
└─ README.md
```

## Páginas, Layouts e Componentes

### Páginas

Páginas representam rotas e projeções de estado.

- `home.page.core`
- `dashboard.page.core`

### Layouts

Layouts definem a estrutura global.

- `main.layout.core`

### Componentes

Componentes são fragmentos de UI reutilizáveis renderizados no servidor.

- `button.comp.core`
- `alert.comp.core`

Toda a renderização acontece no servidor.

## Exemplo (Conceitual)

```go
func Home(ctx *core.Context) core.View {
  return ctx.Render("home", core.State{
    "services": services.List(),
  })
}
```

A visão é um resultado do estado, não o contrário.

## CLI (Conceito)

```bash
core new app mysite
core new page dashboard
core new component alert
core dev
core run
```

Simples, previsível e consistente.

## Acessibilidade e SEO

O CORE fornece naturalmente:

- HTML totalmente renderizado
- Saída amigável para leitores de tela
- Navegação por teclado
- Conteúdo rastreável
- Sem dependência de JS para o conteúdo

Isso se alinha de perto com os padrões e melhores práticas da web.

## Comparação

| Aspecto      | CORE                | Frameworks SPA      |
|--------------|---------------------|---------------------|
| Estado       | Apenas no servidor  | Cliente + Servidor  |
| HTML         | Primário            | Secundário          |
| SEO          | Nativo              | Condicional         |
| Complexidade | Baixa               | Alta                |
| Longevidade  | Alta                | Baixa–Média         |
| Depuração    | Previsível          | Complexa            |

## Origem

O CORE foi originalmente desenvolvido na **Dresbach Hosting do Brasil** como parte de sistemas reais em produção.

O framework existe para resolver problemas reais, não para seguir tendências.

## Status

- 🟢 Conceito validado
- 🟡 Runtime do núcleo em evolução
- 🟡 Documentação em progresso
- 🔵 Extensões planejadas

O CORE é desenvolvido intencionalmente com cautela e clareza.

## Licença

Licença MIT.

---

### Nota Final

O CORE não tenta substituir tudo.

Ele foca em fazer uma coisa bem:

**construir aplicações web estáveis, manuteníveis e dirigidas pelo servidor.**
