#!/usr/bin/env bash

set -e

INPUT_FILE='./test-data/large-file.txt'

go build ./cmd/one-liner

echo "Input file: $INPUT_FILE"

echo "----------- wc ------------"
wc "$INPUT_FILE"
echo "---------------------------"

echo "----------- tr test ------------"
tr -s '\n' ' ' < "$INPUT_FILE" | wc

time tr -s '\n' ' ' < "$INPUT_FILE" > /dev/null
echo "--------------------------------"

echo "----------- one-liner test ------------"
./one-liner < "$INPUT_FILE" | wc

time ./one-liner < "$INPUT_FILE" > /dev/null
echo "---------------------------------------"

