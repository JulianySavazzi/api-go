FROM golang:latest

WORKDIR /go/app

#instalando biblioteca para o kafka funcionar com o go
RUN apt-get update && apt-get install -y librdkafka-dev

CMD ["tail", "-f", "/dev/null"]


#subir o container:
# docker-compose up -d
