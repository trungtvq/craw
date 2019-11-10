#!/bin/bash

HOME=$(pwd)
APPNAME=$(basename $HOME)
SERVER=trungtvq.ddns.net
SERVER_DIR=/server/go/$APPNAME
SERVER_ADDR=root@$SERVER:$SERVER_DIR
echo $SERVER_DIR

ssh root@$SERVER sh $SERVER_DIR/runserver stop /
rsync -av --delete $HOME/$APPNAME $SERVER_ADDR/
rsync -av --delete $HOME/*.toml $SERVER_ADDR/
rsync -av --delete $HOME/runserver $SERVER_ADDR/

# runserver
ssh root@$SERVER sh $SERVER_DIR/runserver start