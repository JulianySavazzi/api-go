FROM golang:latest

WORKDIR /go/app

#instalando biblioteca para o kafka funcionar com o go
RUN apt-get update && apt-get install -y librdkafka-dev

CMD ["tail", "-f", "/dev/null"]

#subir o container:
# docker-compose up -d

#excutar e trabalhar com o banco de dados
#docker-compose exec mysql bash

#criar database products
#mysql -uroot -p products

#criar tabela -> código do arquivo bd_query

#executar e trabalhar com o kafka
#docker-compose exec kafka bash

#criar topico product
#kafka-topics --bootstrap-server=localhost:9092 --topic=product --create

#subir a aplicação e executar o app
#docker-compose exec goapp bash
#go run cmd/app/main.go

#testar kafka
#kafka-console-producer --bootstrap-server=localhost:9092 --topic=product {"name": "product 2", "price": 20}
#kafka-console-producer --bootstrap-server=localhost:9094 --topic=product {"name": "product 2", "price": 200}