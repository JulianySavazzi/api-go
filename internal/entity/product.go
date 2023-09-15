// go é diferente das linguagens orientadas a objetos, ela é uma linguagem hexagonal
package entity

import "github.com/google/uuid"

// a interface define o contrato
// repositorio para acessar o banco de dados
// error é um tipo de dados que transforma um erro em string
type ProductRepository interface {
	//recebe o produto ou mostra o erro
	Create(Product *Product) error
	//lista os produtos ou mostra o erro
	FindAll() ([]*Product, error)
	//atualizar produto
	//excluir produto
}

// vamos criar uma struct que funcionara como uma classe
type Product struct {
	ID    string
	Name  string
	Price float64
}

// funcao para adicionar novo produto
func NewProduct(name string, price float64) *Product {
	//essa funcao recebe os atributos da struct
	//por parametro e retorna um ponteiro para o novo produto
	return &Product{
		//id usando pacote uuid do google
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

//funcao para alterar produto

//funcao para remover produto
