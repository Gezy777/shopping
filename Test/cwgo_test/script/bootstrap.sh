#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/cwgo_test"
exec "$CURDIR/bin/cwgo_test"