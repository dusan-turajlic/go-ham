#!/bin/sh

OUT_DIR="${APPLICATION_DIR}/dist/$(tr -dc A-Za-z0-9 </dev/urandom | head -c 13 ; echo '')"
/usr/local/go/bin/go build -o "$OUT_DIR" go-ham

sh -c "${OUT_DIR}"

