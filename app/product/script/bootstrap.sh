#!/bin/sh
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/product"
exec "$CURDIR/bin/product"
