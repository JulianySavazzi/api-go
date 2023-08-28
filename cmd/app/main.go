package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/JulianySavazzi/api-go/internal/infra/akafka"
	"github.com/JulianySavazzi/api-go/internal/infra/repository"
	"github.com/JulianySavazzi/api-go/internal/infra/web"
	"github.com/JulianySavazzi/api-go/internal/usecase"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
	//"github.com/go-sql-driver/mysql"
)

func main() {
	//criar conexao com banco de dados
	//também vamos usar docker
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/products")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//criar obejto do repositorio
	repository := repository.NewProductRepositoryMysql(db)
	createProductUseCase := usecase.NewCreateProductUseCase(repository)
	listProductsUseCase := usecase.NewListProductsUseCase(repository)

	productHandlers := web.NewProductHandlers(createProductUseCase, listProductsUseCase)

	//roteador -> mapear rotas  do servidor web
	r := chi.NewRouter()
	r.Post("/products", productHandlers.CreateProductHandler)
	r.Get("/products", productHandlers.ListProductsHandler)

	//servidor http -> executa o servidor web
	//go rotines -> para não interromper a execucao do servidor antes de rodar td o codigo
	go http.ListenAndServe(":8000", r)

	//canal de comunicacao
	//criar canal que consome mensagens kafka
	msgChan := make(chan *kafka.Message)
	//go rotines -> para usar a funcao que consome as mensagens (coloca ela para executar em segundo plano enquanto o resto do codigo é executado)
	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}
		//pega os dados que recebeu do kafka em JSON e joga para o DTO
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			//logar o erro
		}
		_, err = createProductUseCase.Execute(dto)
	}
}

/*

--- testar no kafka ---

docker-compose exec kafka bash
kafka-topics --botstrap-server=localhost:9092 --topic=product --create
kafka-console-producer --botstrap-server=localhost:9092 --topic=product
{"neme": "My Product2", "price": 200}

*/
