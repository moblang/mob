# Moblang Compiler

Um compilador moderno para a linguagem de programação moblang.

## Instalação

### Via Make (Recomendado para Desenvolvimento)

```bash
make install
```

### Via Script (Instalação Simples)

```bash
curl -fsSL https://mob.dev/install | bash
```

### Via Script (Instalação Simples)

```bash
curl -fsSL https://raw.githubusercontent.com/moblang/mob/main/install.sh | bash
```

## Uso

### Executar um programa mob

```bash
mob run main.mob
```

### Compilar para binário nativo

```bash
mob build main.mob
./main
```

### Iniciar servidor HTTP

```bash
mob serve main.mob
```

### Executar linter

```bash
mob lint .
```

## Desenvolvimento

### Instalação Global

```bash
make install     # Instala em ~/.local/bin/mob
make reinstall   # Reinstala após modificações
```

### Atualização Após Modificações

**Opção 1 - Script rápido:**
```bash
./update-mob.sh
```

**Opção 2 - Make:**
```bash
make reinstall
```

**Opção 3 - Automático (Watcher):**
```bash
./watch-and-rebuild.sh
# Monitora mudanças e recompila automaticamente
```

### Build

```bash
make build       # Build do binário local
make build-all   # Build para todas as plataformas
```

### Testes

```bash
make test
```

### Lint

```bash
make lint
```

### Limpar

```bash
make clean
```

### Formatar

```bash
make fmt
```

### Compilar para binário nativo

```bash
mob build main.mob
./main
```

### Iniciar servidor HTTP

```bash
mob serve main.mob
```

### Executar linter

```bash
mob lint .
```

## Desenvolvimento

### Build

```bash
make build
```

### Build para todas as plataformas

```bash
make build-all
```

### Testes

```bash
make test
```

### Lint

```bash
make lint
```

### Limpar

```bash
make clean
```

## Estrutura do Projeto

```
moblang/
├── cmd/
│   └── mob/          # CLI principal
├── pkg/
│   └── compiler/     # Lexer, Parser, CodeGen
├── main.mob          # Exemplo Hello World
├── Makefile          # Build automatizado
└── install.sh        # Script de instalação
```

## Exemplo Hello World

Crie um arquivo `main.mob`:

```mob
print("Hello World!")
```

Execute:

```bash
mob run main.mob
```

Saída:
```
Hello World!
```

## Licença

MIT
