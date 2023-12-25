#!/bin/bash

GOOS=js GOARCH=wasm go build -o output/calc.wasm

cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" output/
