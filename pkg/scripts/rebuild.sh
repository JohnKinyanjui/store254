
go build -o ./bin/eccommerce/main ./cmd/eccommerce/main.go
go build -o ./bin/public/main ./cmd/public/main.go

sudo supervisorctl restart all