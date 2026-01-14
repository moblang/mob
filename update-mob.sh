#!/bin/bash

# Script para atualizar o mob globalmente após modificações

set -e

REPO_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BINARY_DIR="${HOME}/.local/bin"

echo "🔄 Atualizando mob..."
echo ""

# Entrar no diretório do repositório
cd "$REPO_DIR"

# Reinstalar
echo "📦 Rebuildando..."
make reinstall

echo ""
echo "✅ Mob atualizado com sucesso!"
echo ""
echo "📍 Localização: ${BINARY_DIR}/mob"
echo ""
echo "Para testar:"
echo "  mob version"
echo "  mob help"
