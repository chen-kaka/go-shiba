export RUN_MODE=test

cd $GOPATH/src/go-shiba
go clean -i
go build

./go-shiba