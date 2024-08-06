#Parking Management

Api de gerenciamento de estacionamento com controle de entrada e saida dos carros do estabelecimento

## Tecnologias Usadas

- golang
- docker

## Rodando o projeto 

- clone o repostiorio e entre na pasta do projeto
```bash
  git clone https://github.com/j0n4t45d3v/parking_management.git
  cd parking_management/ 
```

- depois suba o banco de dados
```bash
  docker-compose up -d database
```

- Apos subir o banco de dados rode as migrations e as seeds
```bash
  go run cmd/cli/main.go migrate
  go run cmd/cli/main.go seed
```

- E por fim é so rodar a api
```bash
  go run cmd/api/main.go
``` 
