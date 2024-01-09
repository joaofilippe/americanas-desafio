# Desafio Americanas

O desafio consistiu em 3 etapas:

1. Elaborar uma aplicação que fusionasse duas listas ordenadas;

2. Criar uma API para receber duas listas, em um endpoint e; em outro, retornasse uma lista resultante do merge das listas anteriores;

3. Expor alguma forma de rodar e testá-los;

O código foi escrito em Golang, utilizando o framework Gin junto ao pacote `net/http`. Foi utilizado um banco de dados Postgres para o armazenamento das listas e posterior consulta. A tecnologia foi escolhida justamente por suportar o armazenamento de arrays de inteiros.

Infelizmente, não foi possível orquestrar com Kubernetes a conexão da aplicação com o banco de dados. Porém, caso não se tenha configurado um servidor Postgres para ser utiilizado, foi disponibilizado um compose-file para rodar via container.

Todavia, a aplicação roda normalmente em Kubernetes, podendo fazer uso dos recursos `port-forward` normalmente para se comunicar tanto com a máquina local quanto com um banco de dados local.

## Configuração

Em primeiro lugar, deve-se assegurar que há um banco de dados Postegres disponível para uso no momento da execução. Caso não possua, é possível executar o comando `docker-compose up`.

A configuração da aplicação deverá seguir os seguintes parâmetros:

```go
DB_HOST={HOST}
DB_PORT={PORT}
DB_USER={USER}
DB_PASSWORD={PASSWORD}
DB_NAME={DATABASE}
DB_SSLMODE=disable
PORT={PORT}
```

* `DB_HOST`: endereço do host do banco de dados;

* `DB_PORT`: porta para acesso do banco de dados;

* `DB_USER`: usuário com acesso ao banco de dados;

* `DB_PASSWORD`: senha do usuário;

* `DB_NAME`: nome do banco de dados(schema) que será criada a tabela

* `PORT`: porta local onde a aplicação será executada;
  
  

A configuração deverá ser salva em um arquivo `.env` e armazenado na pasta `/config`.

O argumento do modo de execução deverá ser passado no momento da execução.

 Caso seja por meio do go run, o comando será: `go run ./cmd/main.go run-mode`.

Caso seja pelo debug do VS-Code, o arquivo deverá obedecer os arquivos na pasta `.vs-code` que está disponível no repositório, de acordo com o sistema operacional.

## Testes

Os testes foram centrados nas funcionalidades principais do programa, que seria mergear, verificar e processar as listas.

Para rodá-los, basta aplicar o comando `go test ./...`
