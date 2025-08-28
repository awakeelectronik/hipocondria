#!/bin/bash

# Script para desarrollo - recarga automática

echo "🔄 Iniciando servidor de desarrollo con recarga automática..."

# Verificar si air está instalado
if ! command -v air &> /dev/null; then
    echo "📦 Instalando Air para recarga automática..."
    go install github.com/cosmtrek/air@latest
fi

# Configurar archivo de Air si no existe
if [ ! -f .air.toml ]; then
    echo "📝 Creando configuración de Air..."
    cat > .air.toml << EOF
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = []
  bin = "./tmp/main"
  cmd = "go build -o ./tmp/main ."
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor", "testdata"]
  exclude_file = []
  exclude_regex = ["_test.go"]
  exclude_unchanged = false
  follow_symlink = false
  full_bin = ""
  include_dir = []
  include_ext = ["go", "tpl", "tmpl", "html"]
  include_file = []
  kill_delay = "0s"
  log = "build-errors.log"
  poll = false
  poll_interval = 0
  rerun = false
  rerun_delay = 500
  send_interrupt = false
  stop_on_root = false

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  main_only = false
  time = false

[misc]
  clean_on_exit = false

[screen]
  clear_on_rebuild = false
  keep_scroll = true
EOF
fi

# Iniciar con Air
air
