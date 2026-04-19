#!/bin/sh
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/payment"
exec "$CURDIR/bin/payment"
