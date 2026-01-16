# CORE Framework

**O Framework Web Server-First para Construção de Aplicações Robustas em Go**

O CORE é um framework web opinativo, construído em Go, que implementa uma arquitetura **Server-First**. Ele foi projetado para reduzir a complexidade, eliminar a duplicação de estado e construir aplicações web duráveis, previsíveis e seguras, priorizando a engenharia sólida em vez de tendências passageiras.

## Manifesto: A Filosofia do CORE

O CORE nasce como uma resposta à complexidade excessiva dos frameworks modernos. Acreditamos que o servidor deve ser a fonte da verdade, e a interface do usuário, uma projeção direta desse estado.

### Princípios Fundamentais

1.  **Server-First:** Toda lógica de negócio, validação e gerenciamento de estado acontece no servidor. O cliente não governa o sistema; ele apenas reflete o que o servidor determina.
2.  **HTML como Contrato:** O HTML é a interface final, entregue pronta e completa. A aplicação funciona perfeitamente sem JavaScript.
3.  **Estado Único:** Existe uma única fonte de verdade para o estado, e ela vive no servidor. Isso elimina bugs complexos de sincronização.
4.  **JavaScript é Opcional:** JavaScript deve ser usado para melhorias progressivas, não como um pilar estrutural. Se o JS falhar, a aplicação continua 100% funcional.
5.  **Previsibilidade Acima de Efeitos:** Priorizamos clareza e comportamento consistente em vez de "mágica" e efeitos complexos. Um sistema previsível é mais fácil de manter e escalar.
6.  **Longevidade como Meta:** O CORE é projetado para rodar por anos, facilitar a manutenção e envelhecer bem, reduzindo a dívida técnica.

---

## Primeiros Passos

Construa sua primeira aplicação com o CORE em minutos.

### Pré-requisitos

- Go 1.21+

### Instalação e Execução

Para criar e rodar um novo projeto:

```bash
# (Assumindo que a CLI do CORE está instalada e no seu PATH)
core new app meu-projeto
cd meu-projeto
core dev
```

Acesse [http://localhost:8080](http://localhost:8080) para ver sua aplicação no ar.

### Fluxo de Desenvolvimento

O fluxo mental do CORE é simples:

1.  Uma requisição HTTP chega ao servidor.
2.  O estado da aplicação é carregado ou criado.
3.  A lógica de negócio, escrita em Go, é executada.
4.  O estado resultante é passado para um renderizador.
5.  O renderizador gera o HTML final.
6.  Uma resposta HTML completa é enviada ao navegador.

> Você não constrói páginas. Você constrói um fluxo **Estado → Renderização**.

---

## Visão Geral da Arquitetura

O CORE possui uma arquitetura com camadas claras e responsabilidades únicas, projetada para reduzir o acoplamento e aumentar a clareza.

```
┌────────────┐
│  Browser   │
└─────┬──────┘
      │ HTTP
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

-   **Router:** Mapeia URLs para handlers, sem lógica de negócio.
-   **Middleware:** Modifica o contexto da requisição (ex: autenticação, logging).
-   **Application Logic:** O coração da sua aplicação, escrito em Go. Produz estado, não markup.
-   **Renderer:** Recebe o estado e gera o HTML final, usando layouts e componentes.

Essa arquitetura garante que a segurança, a performance (SEO) e a acessibilidade sejam consequências naturais do design, não preocupações tardias.

## Conclusão

O CORE não tenta ser tudo para todos. Ele existe para equipes que constroem produtos reais, valorizam a estabilidade e preferem clareza à complexidade. Ele é a base para construir sistemas que duram.

**CORE: O núcleo antes do excesso. A base antes do brilho. A engenharia antes do hype.**
