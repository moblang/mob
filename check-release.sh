#!/bin/bash

echo "Checking moblang v0.0.1 release status..."
echo ""

# Check if release exists
echo "📦 Checking GitHub Release..."
curl -s -o /dev/null -w "  Status: %{http_code}\n" \
  https://github.com/moblang/mob/releases/tag/v0.0.1

if [ $? -eq 0 ]; then
  echo "  ✅ Release page exists"
else
  echo "  ❌ Release page not found yet"
fi

echo ""

# Check binaries
echo "📥 Checking AMD64 binary..."
curl -s -o /dev/null -w "  Status: %{http_code}\n" \
  https://github.com/moblang/mob/releases/download/v0.0.1/mob-linux-amd64

echo ""
echo "📥 Checking ARM64 binary..."
curl -s -o /dev/null -w "  Status: %{http_code}\n" \
  https://github.com/moblang/mob/releases/download/v0.0.1/mob-linux-arm64

echo ""
echo "🔗 Links:"
echo "  Release: https://github.com/moblang/mob/releases/tag/v0.0.1"
echo "  Actions: https://github.com/moblang/mob/actions"
