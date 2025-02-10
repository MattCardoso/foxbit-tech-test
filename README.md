# foxbit-tech-test
Teste Tecnico da Foxbit

## Para buildar o projeto:
Para gerar o executavel basta rodar:

`CGO_ENABLED=0 GOOS=linux go build -o server`

Em seguida só executar: `./server`


## Especificações do teste:
Para rodar os testes utilize o comando:

`go test -v ./...`

Para rodar os testes com o intuito de criar um perfil de cobertura utilize os seguinte comandos:
```bash
go test -v -coverprofile=./coverage/coverage.out ./...

go tool cover -html ./coverage/coverage.out -o ./coverage/coverage.html
```

Depois só abrir o arquivo `coverage/coverage.html` no navegador 

## Para buildar a imagem Docker
Execute o seguinte comando:

```bash
# para buildar a imagem
docker build -t server .

#para rodar a imagem
docker run --rm -p 8000:8000 server:latest
```


## Salvando imagem no GHCR

```bash
docker tag server:latest ghcr.io/mattcardoso/foxbit-tech-test:latest
docker push ghcr.io/mattcardoso/foxbit-tech-test:latest

```