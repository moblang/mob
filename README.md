# Moblang Programming Language

**moblang** é uma linguagem de programação moderna, compilada e fortemente tipada, criada para desenvolvedores que buscam **produtividade extrema**, **alto desempenho** e uma **experiência de desenvolvimento limpa**, sem gambiarras.

Arquivos da linguagem utilizam a extensão **`.mob`**.

## 📚 Comandos

### Comandos Básicos

```bash
mob run <file.mob>      # Compila e executa
mob build <file.mob>    # Compila para binário nativo
mob serve <file.mob>    # Inicia servidor HTTP
mob lint <path>         # Executa linter
mob version             # Mostra versão
mob help                # Mostra ajuda
mob info                # Mostra informações do sistema
```

### Comandos de Ajuda

```bash
mob help                # Mostra todos os comandos disponíveis
mob help run            # Ajuda detalhada do comando run
mob help build          # Ajuda detalhada do comando build
mob --help              # Mostra ajuda (alternativa)
mob -h                  # Mostra ajuda (alternativa curta)
```

### Atalhos

```bash
mob <file.mob>          # Atalho para "mob run <file.mob>"
mob -v                  # Mostra versão
mob --version           # Mostra versão
```

### Opções do Build

```bash
mob build -o <nome> <file>    # Especifica nome do binário
mob build --output <nome>       # Especifica nome do binário
```

## 🚀 Instalação Rápida

### Via Script (Recomendado)

```bash
curl -fsSL https://mob.dev/install | bash
```

### Via GitHub

```bash
curl -fsSL https://raw.githubusercontent.com/moblang/mob/main/install.sh | bash
```

### Manual

```bash
git clone https://github.com/moblang/mob.git
cd mob
make build
sudo make install
```

## 👋 Hello World

Crie `main.mob`:

```mob
print("Hello World!")
```

Execute:

```bash
mob run main.mob
```

Ou compile para binário nativo:

```bash
mob build main.mob
./main
```

## 📚 Comandos

```bash
mob run <file.mob>      # Compila e executa
mob build <file.mob>    # Compila para binário nativo
mob serve <file.mob>    # Inicia servidor HTTP
mob lint <path>         # Executa linter
mob version             # Mostra versão
```

## 🧱 Exemplos

Ver a pasta `examples/` para mais exemplos.

### Hello World
```mob
print("Hello World!")
```

### Orientação a Objetos (em breve)
```mob
class User extends Model:
    public name: string

    public function hello():
        print("Hello " + this.name)
```

## 🛠️ Desenvolvimento

```bash
make build       # Build do binário
make test        # Executa testes
make lint        # Executa linter
make clean       # Limpa build
make install     # Instala em ~/.local/bin
```

## 📦 Features

- ✅ Linguagem compilada (performance igual/superior a Go)
- ✅ Sintaxe Python-like com indentação
- ✅ Tipagem forte verificada em compilação
- ✅ Orientação a objetos com classes
- ✅ Import de módulos
- ⚡ Assincronismo e concorrência (em breve)
- 🧹 Linter nativo (em breve)
- 🌐 Servidor HTTP embutido (em breve)
- 🐞 Debugger nativo (em breve)

## 📁 Estrutura do Projeto

```
moblang/
├── cmd/mob/              # CLI principal
├── pkg/compiler/         # Lexer, Parser, CodeGen
├── examples/             # Exemplos de código
├── main.mob              # Hello World exemplo
├── Makefile              # Automatização de build
└── install.sh            # Script de instalação
```

## 🧠 Filosofia

- Código deve ser legível
- Performance não é opcional
- Ferramentas devem ser nativas
- Menos configuração, mais execução
- O desenvolvedor vem antes do framework

## 📝 Status do Projeto

🚧 moblang está em desenvolvimento ativo

Este repositório contém o compilador inicial da linguagem.

## 🤝 Contribuindo

Contribuições são bem-vindas! Veja CONTRIBUTING.md para mais detalhes.

## 📄 Licença

MIT License - veja LICENSE para mais detalhes.

---

"Uma linguagem criada por quem já programou em quase todas — e decidiu criar a sua."
