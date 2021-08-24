#!/usr/bin/env bash

set -e

go build ./cmd/one-liner

cd test-data

inputs='small.txt medium.txt large.txt huge.txt'
for input in $inputs; do
    INPUT_FILE="$input"

    echo "Input file: $INPUT_FILE"

    echo "----------- wc ------------"
    wc "$INPUT_FILE"
    echo "---------------------------"

    echo "----------- tr test ------------"
    tr -s '\n' ' ' < "$INPUT_FILE" | wc

    time tr -s '\n' ' ' < "$INPUT_FILE" > /dev/null
    echo "--------------------------------"

    echo "----------- awk test ------------"
    awk '{printf "%s ",$0}' < "$INPUT_FILE" | wc

    time awk '{printf "%s ",$0}' < "$INPUT_FILE" > /dev/null
    echo "--------------------------------"

    echo "----------- one-liner test ------------"
    ../one-liner < "$INPUT_FILE" | wc

    time ../one-liner < "$INPUT_FILE" > /dev/null
    echo "---------------------------------------"
done


