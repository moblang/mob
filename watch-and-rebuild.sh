#!/bin/bash

# Watcher automático para recompilar e reinstalar o mob quando houver mudanças

set -e

REPO_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo "👀 Mob Auto-Update Watcher"
echo ""
echo "Diretório: $REPO_DIR"
echo ""
echo "Monitorando alterações no código..."
echo "Pressione Ctrl+C para parar"
echo ""

# Entrar no diretório do repositório
cd "$REPO_DIR"

# Usar inotifywait para monitorar mudanças
# Se não estiver instalado, usa polling simples
if command -v inotifywait >/dev/null 2>&1; then
    echo "✅ Usando inotifywait (monitoramento em tempo real)"
    echo ""
    while true; do
        inotifywait -e modify,create,delete -r ./cmd ./pkg 2>/dev/null || true
        echo ""
        echo "🔄 Mudanças detectadas! Atualizando mob..."
        make reinstall
        echo ""
        echo "✅ Atualizado! Continue codando..."
        echo ""
    done
else
    echo "⚠️  inotifywait não encontrado. Instale para monitoramento em tempo real:"
    echo "   sudo apt install inotify-tools"
    echo ""
    echo "Usando modo de polling (verifica a cada 2 segundos)..."
    echo ""
    while true; do
        sleep 2
        make reinstall > /dev/null 2>&1 && echo "✅ Mob atualizado!" || true
    done
fi
