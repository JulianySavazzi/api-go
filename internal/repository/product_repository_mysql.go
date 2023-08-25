package repository

//conexao com o banco de dados
import (
	"database/sql"

	"github.com/JulianySavazzi/api-go/internal/entity"
)

type ProductRepositoryMysql struct {
	DB *sql.DB
}

// funcao construtora
func NewProductRepositoryMysql(db *sql.DB) *ProductRepositoryMysql {
	return &ProductRepositoryMysql{DB: db}
}

// metodo para inserir dados no banco de dados
func (r *ProductRepositoryMysql) Create(product *entity.Product) error {
	_, err := r.DB.Exec("Insert into products (id, name, price) values(?,?,?)",
		product.ID, product.Name, product.Price)

	if err != nil {
		return err
	}
	return nil
}
