#!/bin/sh


if [ $USE_SECRET_MANAGER ] ;then
  echo "use secret-manager"
else
  echo "doesn't use secret-manager"
fi

/usr/local/bin/subsidy
