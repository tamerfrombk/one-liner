#!/usr/bin/env bash

set -e

INPUT_FILE='./test-data/large-file.txt'

go build ./cmd/one-liner

echo "Input file: $INPUT_FILE"

echo "----------- tr test ------------"
time tr -d '\n' < "$INPUT_FILE" > /dev/null
echo "--------------------------------"

echo "----------- one-liner test ------------"
time ./one-liner < "$INPUT_FILE" > /dev/null
echo "---------------------------------------"

