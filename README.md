# Desafio Americanas

O desafio consistiu em 3 etapas:

1. Elaborar uma aplicação que fusionasse duas listas ordenadas;

2. Criar uma API para receber duas listas, em um endpoint e; em outro, retornasse uma lista resultante do merge das listas anteriores;

3. Expor alguma forma de rodar e testá-los;

O código foi escrito em Golang, utilizando o framework Gin junto ao pacote `net/http`. 

Foi utilizado um banco de dados Postgres para o armazenamento das listas e posterior consulta. A tecnologia foi escolhida justamente por suportar o armazenamento de arrays de inteiros.

Para os testes, contamos com a ajuda das ferramentas `Testify`  e `Mockery` para ajudar na organização do código e potencialização do pacote `testing` da biblioteca padrão.

Infelizmente, não foi possível orquestrar com Kubernetes a conexão da aplicação com o banco de dados. Porém, caso não se tenha configurado um servidor Postgres para ser utiilizado, foi disponibilizado um compose-file para rodar via container.

Todavia, a aplicação roda normalmente em Kubernetes, podendo fazer uso dos recursos `port-forward` normalmente para se comunicar tanto com a máquina local quanto com um banco de dados local.

## Configuração

A configuração da aplicação deverá seguir os seguintes parâmetros:

---

```js
DB_HOST={HOST}
DB_PORT={PORT}
DB_USER={USER}
DB_PASSWORD={PASSWORD}
DB_NAME={DATABASE}
DB_SSLMODE=disable
PORT={PORT}
```

---

A configuração deverá ser salva em um arquivo `.env` e armazenado na pasta `/config`.

O argumento do modo de execução deverá ser passado no momento da execução.

Caso seja por meio do go run, o comando será: `go run ./cmd/main.go run-mode`.

Caso seja pelo debug do VS-Code, o arquivo deverá obedecer os arquivos na pasta `.vs-code` que está disponível no repositório, de acordo com o sistema operacional.

## Testes

Os testes foram centrados nas funcionalidades principais do programa, que seria mergear, verificar e processar as listas, bem como seus services.

Para rodá-los, basta aplicar o comando `go test ./...`.

## API

### Endpoints

**SAVE_LISTS**

* Endpoint: `{{HOST}}/api/v1/list/save_list`;

* Request Body: 

---

```json
  {
      "type": "node",
      "list1": {
          "val": 1,
          "next": {
              "val": 7,
              "next": {
                  "val": 9
              }
          }
      },
      "list2": {
          "val": 5,
          "next": {
              "val": 5,
              "next": {
                  "val": 11
              }
          }
      }
  }
```

---

* Response:

---

```json
  {
      "code": 201,
      "message": "Lists saved successfully. Follow the id to merge the lists.",
      "data": 1
  }
```

---

**MERGE**

* Endpoints: `{{HOST}}/api/v1/list/merge/:list_id`
  
  * Params: `:list-id`
  
  * Example: `{{HOST}}/api/v1/list/merge/1`

* Response:

---

```json
  {
      "code": 201,
      "message": "Lists had merged succesfully.",
      "data": {
          "merged_list": {
              "Val": 1,
              "Next": {
                  "Val": 5,
                  "Next": {
                      "Val": 5,
                      "Next": {
                          "Val": 7,
                          "Next": {
                              "Val": 9,
                              "Next": {
                                  "Val": 11,
                                  "Next": null
                              }
                          }
                      }
                  }
              }
          }
      }
  }
```

---

## Sugestões de Testes

1. Enviar duas listas ordenadas pela rota `save_lists`. O resultado será o id pelo qual você irá consultar a lista mergeada. Num primeiro momento, essa lista será mergeada, devolvida para o usuário e, depois, salva no banco de dados.

2. Com o mesmo id reconsultar essa lista. Dessa vez, ela será trazida do banco de dados, e a operação de mergeamento não será mais realizada. 

3. Ainda na rota de `merge`, consultar um id inexistente. Irá retornar que não pode consultar.

4. Também poderá ser feita uma tentativa de enviar uma ou mais listas não ordenadas,  o que o programa não irá aceitar.
