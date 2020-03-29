#/bin/bash

RESULT=`go run ./1/1.go`
rm ./1/1.go

if [ $RESULT = "vim-go" ]; then
    exit 0
fi

exit 1
