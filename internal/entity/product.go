// go é diferente das linguagens orientadas a objetos
package entity

import "github.com/google/uuid"

// a interface define o contrato
// repositorio para acessar o banco de dados
type ProductRepository interface {
	//recebe o produto ou mostra o erro
	Create(Product *Product) error
	//lista os produtos ou mostra o erro
	FindAll() ([]*Product, error)
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
