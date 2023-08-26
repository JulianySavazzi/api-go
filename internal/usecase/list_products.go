package usecase

import "github.com/JulianySavazzi/api-go/internal/entity"

//listagem simples sem dado de entrada -> precisamos apenas de dado de saída
//se fosse uma pesquisa com dado de entrada, precisaríamos de um dado de entrada e um de saida

type ListProductsOutputDto struct {
	ID    string
	Name  string
	Price float64
}

type ListProductsUseCase struct {
	//interface do repositório
	ProductRepository entity.ProductRepository
}

// funcao construtora
func NewListProductsUseCase(productRepository entity.ProductRepository) *ListProductsUseCase {
	return &ListProductsUseCase{ProductRepository: productRepository}
}

// metodo listar
func (u *ListProductsUseCase) Execute() ([]*ListProductsOutputDto, error) {
	products, err := u.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var productsOutput []*ListProductsOutputDto
	for _, product := range products {
		productsOutput = append(productsOutput, &ListProductsOutputDto{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}
	return productsOutput, nil
}
