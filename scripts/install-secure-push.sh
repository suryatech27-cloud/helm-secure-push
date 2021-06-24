#!/bin/bash

UNAME=$( command -v uname)
OS=""

case $( "${UNAME}" | tr '[:upper:]' '[:lower:]') in
  linux*)
    OS='linux'
    ;;
  darwin*)
    OS='darwin'
    ;;
  msys*|cygwin*|mingw*)
    # or possible 'bash on windows'
    OS='windows'
    ;;
  nt|win*)
    OS='windows'
    ;;
  *)
    OS='unknown'
    ;;
esac


DIR=$(helm env HELM_PLUGINS)

if [ $OS == "windows" ]; then
        #echo $HELM_PLUGIN_DIR
        cd $DIR'/helm-secure-push/bin'; rm 'securepush';
        #echo $PUSH_BIN_PATH

        cd $DIR'/helm-secure-push/bin'; env GOOS=windows GOARCH=amd64 go build ../cmd/securepush
        cd 'C:/Windows/System32'; setx path "$DIR\helm-secure-push\bin"
else
  echo "Installed Successfully Done."
fi

#finish
    
