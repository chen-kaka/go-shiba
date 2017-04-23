export RUN_MODE=prod

cd $GOPATH/src/go-shiba
go clean -i
go build

./go-shiba