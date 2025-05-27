protoc --go_out=. --go-grpc_out=. api/protos/*

Para criar um novo schema
go run -mod=mod entgo.io/ent/cmd/ent --target internal/app/ent/schema new Products

Para gerar os modelos
go generate ./internal/app/ent/schema

para excluir
git ls-files -i -c --exclude-from=.gitignore | xargs -d '\n' git rm --cached

