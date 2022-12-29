#!/bin/sh

# TODO setup lambda rie
if [ -z "${AWS_LAMBDA_RUNTIME_API}" ]; then
  echo "Emulated: $@" >&2
  exec "$@"
else
  echo "Real: $@" >&2
  exec "$@"
fi

