#!/bin/sh
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/checkout"
exec "$CURDIR/bin/checkout"
