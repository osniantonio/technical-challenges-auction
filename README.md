# Abertura e fechamento do Leilão - Go Routines

## Objetivo: Adicionar uma nova funcionalidade ao projeto já existente para o leilão fechar automaticamente a partir de um tempo definido.

Clone o seguinte repositório: clique para acessar o repositório (https://github.com/devfullcycle/labs-auction-goexpert).

Toda rotina de criação do leilão e lances já está desenvolvida, entretanto, o projeto clonado necessita de melhoria: adicionar a rotina de fechamento automático a partir de um tempo.

Para essa tarefa, você utilizará o go routines e deverá se concentrar no processo de criação de leilão (auction). A validação do leilão (auction) estar fechado ou aberto na rotina de novos lançes (bid) já está implementado.

## Você deverá desenvolver:

Uma função que irá calcular o tempo do leilão, baseado em parâmetros previamente definidos em variáveis de ambiente;
Uma nova go routine que validará a existência de um leilão (auction) vencido (que o tempo já se esgotou) e que deverá realizar o update, fechando o leilão (auction);
Um teste para validar se o fechamento está acontecendo de forma automatizada;

## Dicas:

Concentre-se na no arquivo internal/infra/database/auction/create_auction.go, você deverá implementar a solução nesse arquivo;
Lembre-se que estamos trabalhando com concorrência, implemente uma solução que solucione isso:
Verifique como o cálculo de intervalo para checar se o leilão (auction) ainda é válido está sendo realizado na rotina de criação de bid;
Para mais informações de como funciona uma goroutine, clique aqui e acesse nosso módulo de Multithreading no curso Go Expert;

## Entrega:

O código-fonte completo da implementação.
Documentação explicando como rodar o projeto em ambiente dev.
Utilize docker/docker-compose para podermos realizar os testes de sua aplicação.

## Como acessar o MongoDB

Utilizando o mongo shell para connectar o mongoDB

    Com o MongoDB Shell Instalado localmente (https://www.mongodb.com/try/download/shell)

    Executar o comando para conectar ao mongodb
    	mongosh "mongodb://admin:admin@localhost:27017"

    Selecionar a base de dados admin
    	use admin

    Para ver as collections dessa base
    	show collections
    		system.users
    		system.version

    Consultando os dados da collection users
    	db.system.users.find()

## Como executar o projeto:

Para subir o app:

```sh
make up
```

Para parar o contêiner:

```sh
make down
```

Para ver os logs:

```sh
make logs
```

## Como testar o projeto:

Primeiro, para facilitar a crição do leilão e dar os lances, utilizar o arquivo ./test/auction.http
