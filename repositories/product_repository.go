package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetProducts(name string) ([]models.Product, error) {
	query := "SELECT p.id, p.category_id, c.name as category_name, p.name, p.price, p.stock FROM products p INNER JOIN categories c ON p.category_id = c.id"

	args := []interface{}{}
	if name != "" {
		query += " WHERE p.name ILIKE $1"
		args = append(args, "%"+name+"%")
	}

	query += " ORDER BY p.id"

	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]models.Product, 0)
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.CategoryID, &product.CategoryName, &product.Name, &product.Price, &product.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (repo *ProductRepository) CreateProduct(product *models.Product) error {
	query := "INSERT INTO products (category_id, name, price, stock) VALUES ($1, $2, $3, $4) RETURNING id"
	err := repo.db.QueryRow(query, product.CategoryID, product.Name, product.Price, product.Stock).Scan(&product.ID)
	return err
}

// GetByID - ambil produk by ID
func (repo *ProductRepository) GetProductByID(id int) (*models.Product, error) {
	query := "SELECT p.id, p.category_id, c.name as category_name, p.name, p.price, p.stock FROM products p INNER JOIN categories c ON p.category_id = c.id WHERE p.id = $1"

	var product models.Product
	err := repo.db.QueryRow(query, id).Scan(&product.ID, &product.CategoryID, &product.CategoryName, &product.Name, &product.Price, &product.Stock)
	if err == sql.ErrNoRows {
		return nil, errors.New("Product not found")
	}
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *ProductRepository) UpdateProduct(product *models.Product) error {
	query := "UPDATE products SET category_id = $1, name = $2, price = $3, stock = $4 WHERE id = $5"
	result, err := repo.db.Exec(query, product.CategoryID, product.Name, product.Price, product.Stock, product.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("Product not found")
	}

	return nil
}

func (repo *ProductRepository) DeleteProduct(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("Product not found")
	}

	return err
}
