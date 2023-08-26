package usecase

import "github.com/JulianySavazzi/api-go/internal/entity"

//precisamos de um dado de entrada e um dado de saida
//DTO -> data transfer object
//dados passarem pelas camadas da aplicação sem gerar acoplamento do domínio

type CreateProductInputDto struct {
	//tag JSON para informar qual variavel ele esta se referindo
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateProductOutputDto struct {
	ID    string
	Name  string
	Price float64
}

type CreateProductUseCase struct {
	// passar a interface do repositório, em vez de passar o repositório inteiro
	ProductRepository entity.ProductRepository
}

// funcao construtora do tipo criado
func NewCreateProductUseCase(productRepository entity.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{ProductRepository: productRepository}
}

func (u *CreateProductUseCase) Execute(input CreateProductInputDto) (*CreateProductOutputDto, error) {
	// criar produto
	product := entity.NewProduct(input.Name, input.Price)
	err := u.ProductRepository.Create(product)
	if err != nil {
		return nil, err
	}
	//dado retornado sem estar acoplado ao dominio
	return &CreateProductOutputDto{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}
