# 🔐 Documentação de Segurança do CORE

## Princípio-base

> Menos lógica no cliente = menos superfície de ataque.

## Superfície de ataque reduzida

Por design, o CORE elimina ou reduz drasticamente:

- estado sensível armazenado no browser
- lógica de negócio crítica em JavaScript
- APIs de dados complexas expostas para a UI

## Sessões

- **Server-side:** O estado da sessão vive no servidor, não em tokens JWT decodificáveis no cliente.
- **Cookies seguros:** Use flags `HttpOnly`, `Secure`, e `SameSite=Lax` ou `Strict`.
- **Tokens curtos:** O cookie contém apenas um ID de sessão de alta entropia.
- **Renovação controlada:** O servidor gerencia o ciclo de vida da sessão.

## Autenticação

- **Centralizada:** A lógica de autenticação é um middleware no servidor.
- **Validada antes da renderização:** O estado de autenticação é verificado antes de qualquer lógica de página ser executada.

> Estado autorizado gera UI autorizada.

## CSRF (Cross-Site Request Forgery)

- **Tokens por sessão:** Gere tokens anti-CSRF e armazene-os no estado da sessão.
- **Validação em middleware:** Verifique o token para todas as requisições que alteram estado (`POST`, `PUT`, `DELETE`).
- **Formulários protegidos:** Inclua o token em um campo oculto nos formulários.

## XSS (Cross-Site Scripting)

- **HTML renderizado no servidor:** A principal fonte de HTML é o motor de templates do Go.
- **Escape automático:** Use `html/template` para garantir que toda a saída de dados seja escapada por padrão.
- **Sem template dinâmico no cliente:** Evite renderizar HTML a partir de strings no JavaScript.

## Headers de Segurança

O CORE recomenda e facilita o uso de headers de segurança padrão:

- `Content-Security-Policy` (CSP)
- `Strict-Transport-Security` (HSTS)
- `X-Frame-Options`
- `X-Content-Type-Options`

## Falhas

- **Falha nunca quebra a UI:** Erros no servidor são tratados e podem ser renderizados como parte do estado da página (ex: uma mensagem de erro).
- **Respostas sempre válidas:** O cliente sempre recebe um documento HTML válido, mesmo em caso de erro.
