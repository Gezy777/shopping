#!/bin/sh
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/email"
exec "$CURDIR/bin/email"
