# note: call scripts from /scripts
gobuild:
	go build -o build/main.exe cmd/main/main.go
gotest:
	go test -v ./...

# ./build/main.exe --data_start 2021-08-01 --data_end 2021-09-01 --city Москва --region Московская --limit 100 --offset 0