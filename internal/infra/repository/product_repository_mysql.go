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

// metodo create da struct produto -> inserir dados no banco de dados
func (r *ProductRepositoryMysql) Create(product *entity.Product) error {
	// o Exec está passando o comando SQL insert
	// o _ está ignorando os detalhes dos erros, se quisessemos mostrar, teriamos que colocar uma variavel
	_, err := r.DB.Exec("Insert into products (id, name, price) values(?,?,?)",
		product.ID, product.Name, product.Price)
	//se acontecer algum erro, retorna o erro, se não retorna vazio
	if err != nil {
		return err
	}
	return nil
}

// metodo findAll da struct produto
func (r *ProductRepositoryMysql) FindAll() ([]*entity.Product, error) {
	//será mostrado o erro e a linha -> rows, err
	//usamos parametro de erro no go, pq ele nao tem try catch
	rows, err := r.DB.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close() //é executado depois de tudo que está no escopo (dentro desse metodo)
	//bom para liberar memoria

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}
