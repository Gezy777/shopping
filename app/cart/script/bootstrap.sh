#!/bin/sh
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/cart"
exec "$CURDIR/bin/cart"
