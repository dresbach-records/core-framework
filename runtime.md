# 📐 Especificação do Runtime CORE

## Objetivo do Runtime

O runtime CORE é o motor de execução responsável por:

- receber requisições
- criar contexto
- carregar estado
- executar lógica
- renderizar HTML
- responder ao cliente

Ele não é um framework de frontend.
Ele é o núcleo operacional da aplicação.

## Componentes do Runtime

### 1. HTTP Server

- Baseado em HTTP padrão
- Sem abstrações excessivas
- Compatível com proxies e load balancers
- Responsável apenas por:
    - aceitar conexões
    - delegar para o router

### 2. Router

- Mapeia rota + método → handler
- Suporta parâmetros dinâmicos
- Não contém lógica de negócio
- **Contrato:** `Handle(method, path, Handler)`

### 3. Context

- Criado por requisição.
- Contém:
    - Request
    - Response
    - Params
    - Session
    - State
    - User (opcional)

> O Context é o fio condutor do runtime.

### 4. Middleware Engine

- Pipeline previsível
- Ordem explícita
- Execução síncrona
- **Regra:** Middleware não renderiza HTML.

### 5. State Manager

- Estado único por sessão
- Serializável
- Isolado
- Recuperável

> Nenhum estado vive no cliente.

### 6. Renderer

- Recebe estado
- Resolve página
- Aplica layout
- Renderiza componentes
- Produz HTML final

### 7. Response

- HTML completo ou parcial
- Status HTTP correto
- Headers explícitos
