FROM golang:latest

WORKDIR /go/app

#instalando biblioteca para o kafka funcionar com o go
RUN apt-get update && apt-get install -y librdkafka-dev

CMD ["tail", "-f", "/dev/null"]

#Comandos bash:

#subir o container:
# docker-compose up -d

#listar serviços em execução
# docker ps

#excutar e trabalhar com o banco de dados
#docker-compose exec mysql bash

#criar database products
#mysql -uroot -p products

#criar tabela -> código do arquivo bd_query

#executar e trabalhar com o kafka
#docker-compose exec kafka bash

#criar topico product
#kafka-topics --bootstrap-server=localhost:9092 --topic=product --create

#listar topicos
#kafka-topics --list --bootstrap-server=localhost:9092

#subir a aplicação e executar o app
#docker-compose exec goapp bash
#go run cmd/app/main.go

#testar kafka -> enviar mensagem
#kafka-console-producer --bootstrap-server=localhost:9092 --topic=product {"name": "product 2", "price": 20}

#ler mensagens gravadas no topico product
#kafka-console-consumer --bootstrap-server localhost:9092 --topic product --from-beginning


