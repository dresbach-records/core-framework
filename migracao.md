# 🔄 Guia de Migração para CORE

**Migração não é reescrita cega.**

Migrar para CORE significa mudar o modelo mental, não apenas trocar tecnologia.

## De SPA → CORE

### Antes (SPA)

- Estado no cliente
- API separada
- Hidratação
- Sincronização frágil

### Depois (CORE)

- Estado no servidor
- HTML como resposta
- Navegação real
- Sem hidratação

## Estratégia recomendada

### Fase 1 — Backend primeiro

1.  Mova regras de negócio para o servidor.
2.  Centralize o estado.
3.  Simplifique a API (idealmente, elimine-a em favor de renderização direta).

### Fase 2 — HTML server-side

1.  Renderize páginas completas no backend.
2.  Elimine a necessidade de JSON para construir a UI.
3.  Use rotas e links (`<a>`) reais.

### Fase 3 — Remoção gradual do JS

1.  JavaScript se torna opcional, para melhorias.
2.  Remova lógica de estado ou de renderização duplicada no cliente.
3.  Mantenha apenas interações pontuais e específicas.

## Migração incremental

CORE não exige um *big bang*.
Você pode migrar:

- página por página
- rota por rota
- módulo por módulo
