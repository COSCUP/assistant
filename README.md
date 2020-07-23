# coscup assistant

This action is for COSCUP assistant, a Google action.

Design files is in /designs folder

Backend (fulfillment) is in /fulfillment folder, wrote on Golang.




## Set env
```
cd fulfillment
export GOPATH=`pwd`

cd src/github.com/COSCUP/assistant
dep ensure
```


## build and run

```
cd fulfillment/src/github.com/COSCUP/assistant/serverlet
go build
./serverlet


```

## Problem

dep: command not found

see https://github.com/golang/dep
