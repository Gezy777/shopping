#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=a.b.c
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}