# Desafio Americanas

O desafio consistiu em 3 etapas:

1. Elaborar uma aplicação que fusionasse duas listas ordenadas;

2. Criar uma API para receber duas listas, em um endpoint e; em outro, retornasse uma lista resultante do merge das listas anteriores;

3. Expor alguma forma de rodar e testá-los;

O código foi escrito em Golang, utilizando o framework Gin junto ao pacote net/http. Foi utilizado um banco de dados Postgres para o armazenamento das listas e posterior consulta. A tecnologia foi escolhida justamente por suportar o armazenamento de arrays de inteiros.

Infelizmente, não foi possível orquestrar com Kubernetes a conexão da aplicação com o banco de dados. Porém, caso não se tenha configurado um servidor Postgres para ser utiilizado, foi disponibilizado um compose-file para rodar via container.

Todavia, a aplicação roda normalmente em Kubernetes, podendo fazer uso dos recursos `port-forward` normalmente para se comunicar tanto com a máquina local quanto com um banco de dados local.

## Configuração

A configuração da aplicação deverá seguir os seguintes parâmetros:

```js
DB_HOST={HOST}
DB_PORT={PORT}
DB_USER={USER}
DB_PASSWORD={PASSWORD}
DB_NAME={DATABASE}
DB_SSLMODE=disable
ENV=dev
```

A configuração deverá ser salva em um arquivo `.env` e armazenado na pasta `/config`.

A chave `env` deverá ser fornecida como caso o programa seja rodado via debbug ou localmente. Ela também poderá ser 
