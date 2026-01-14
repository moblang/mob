# Moblang Compiler Architecture

## Overview

O compilador moblang é implementado em Go e segue um pipeline clássico de compilação:

```
Source (.mob) → Lexer → Parser → AST → Code Generator → Go Code → Native Binary
```

## Componentes

### 1. Lexer (`pkg/compiler/lexer.go`)

Responsável por transformar o código fonte em tokens.

**Tokens suportados:**
- `TokenIdentifier`: identificadores (print, User, name)
- `TokenString`: strings literais ("Hello World")
- `TokenNumber`: números literais
- `TokenLeftParen`: `(`
- `TokenRightParen`: `)`
- `TokenColon`: `:`
- `TokenIndent`: início de bloco (indentação)
- `TokenDedent`: fim de bloco (dedentação)
- `TokenNewline`: quebra de linha
- `TokenEOF`: fim de arquivo

**Características:**
- Suporta indentação baseada em espaços (4 espaços)
- Gera tokens Indent/Dedent automaticamente
- Trata strings com escape de caracteres

### 2. Parser (`pkg/compiler/parser.go`)

Responsável por transformar tokens em uma AST (Abstract Syntax Tree).

**Tipos de nós:**
- `NodeProgram`: programa completo
- `NodeCall`: chamada de função
- `NodeString`: string literal
- `NodeIdentifier`: identificador

**Características:**
- Parse recursivo descendente
- Suporta chamadas de função: `print("Hello")`
- Suporta múltiplas declarações
- Tratamento de erro básico

### 3. Code Generator (`pkg/compiler/codegen.go`)

Responsável por gerar código Go a partir da AST.

**Estratégia de compilação:**
1. Gera código Go válido
2. Compila código Go para binário nativo
3. Performance próxima ao Go nativo

**Mapeamentos:**
- `print()` → `fmt.Println()`
- Strings em .mob → Strings em Go
- Identificadores → Identificadores Go

### 4. Compiler (`pkg/compiler/compiler.go`)

Orquestra todo o processo de compilação.

**Operações:**
- `CompileAndRun()`: compila e executa (temporário)
- `Compile()`: compila para binário persistente

**Processo:**
1. Lê arquivo fonte
2. Executa lexer → tokens
3. Executa parser → AST
4. Executa codegen → Go code
5. Compila Go → binário nativo

### 5. CLI (`cmd/mob/main.go`)

Interface de linha de comando.

**Comandos:**
- `run`: compila e executa
- `build`: compila para binário
- `serve`: servidor HTTP (TODO)
- `lint`: linter nativo (TODO)
- `version`: versão do compilador

## Fluxo de Execução

### Comando: `mob run main.mob`

```
1. CLI recebe comando
2. Compiler.CompileAndRun("main.mob")
3. Lexer tokeniza source
4. Parser cria AST
5. CodeGenerator gera Go code
6. Compiler compila Go para temp_binary
7. Executa temp_binary
8. Remove temp_binary
```

### Comando: `mob build main.mob`

```
1. CLI recebe comando
2. Compiler.Compile("main.mob", "main")
3. Lexer tokeniza source
4. Parser cria AST
5. CodeGenerator gera Go code
6. Compiler compila Go para "./main"
7. Binário persiste
```

## Design Decisions

### Por que Go como código intermediário?

1. **Performance**: Go tem performance excelente, similar a C/C++
2. **Simplicidade**: Código Go é fácil de gerar e debugar
3. **Portabilidade**: Go compila para múltiplas plataformas
4. **Tooling**: Ecossistema Go maduro para build/distribute
5. **Manutenibilidade**: Código gerado é legível e debugável

### Por que não LLVM direto?

1. **Complexidade**: LLVM tem curva de aprendizado íngreme
2. **Debugabilidade**: Código Go gerado é mais fácil de debugar
3. **Time-to-market**: Compilação via Go é mais rápida para MVP
4. **Suficiente**: Performance atinge objetivo do README

### Futuro: Compilação nativa

Planejado:
- Gerar código de máquina diretamente (via LLVM ou similar)
- Otimizações específicas da linguagem
- Remove dependência de Go para compilação final

## Próximos Passos

### Curto Prazo
- [ ] Suporte a variáveis
- [ ] Suporte a expressões matemáticas
- [ ] Suporte a classes e métodos
- [ ] Suporte a tipos básicos (int, float, bool)
- [ ] Linter básico

### Médio Prazo
- [ ] Sistema de tipos completo
- [ ] Imports de módulos
- [ ] Funções definidas pelo usuário
- [ ] Controle de fluxo (if, for, while)

### Longo Prazo
- [ ] Compilação nativa (sem Go intermediate)
- [ ] Debugger step-through
- [ ] Tracer/Profiler
- [ ] Assincronismo e green threads
- [ ] Garbage collector customizado

## Convenções de Código

### Go (compiler e CLI)
- Seguir `Effective Go`
- Usar `gofmt` para formatação
- `golangci-lint` para linting
- Testes unitários para cada componente

### Mob (linguagem)
- 4 espaços de indentação
- snake_case para funções e variáveis
- PascalCase para classes
- UPPER_SNAKE_CASE para constantes

## Performance

Benchmark inicial:
- Hello World: ~100ms (incluindo compilação Go)
- Binário final: ~2-3MB
- Runtime: Similar a Go

Otimizações futuras:
- Cache de compilação
- Compilação incremental
- Link-time optimization
- Redução de tamanho de binário
