package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-ddd-hexagonal/domain/model"
)

type SQLProductRepository struct {
	db *sql.DB
}

func NewSQLProductRepository(db *sql.DB) *SQLProductRepository {
	return &SQLProductRepository{db: db}
}

func (r *SQLProductRepository) Create(ctx context.Context, product *model.Product) error {
	query := `INSERT INTO products (id, name, price) VALUES (?, ?, ?)`

	// Executar a query com os valores do produto
	_, err := r.db.ExecContext(ctx, query, product.ID, product.Name, product.Price)
	if err != nil {
		// Tratamento de erro, caso algo dê errado com a inserção
		return fmt.Errorf("error inserting product: %v", err)
	}

	// Retornar nil em caso de sucesso
	return nil
}

func (r *SQLProductRepository) Update(ctx context.Context, product *model.Product) error {
	query := `UPDATE products SET name = ?, price = ? where id = ?`

	// Executar a query com os valores do produto
	_, err := r.db.ExecContext(ctx, query, product.Name, product.Price, product.ID)
	if err != nil {
		// Tratamento de erro, caso algo dê errado com a inserção
		return fmt.Errorf("error updating product: %v", err)
	}

	// Retornar nil em caso de sucesso
	return nil
}

func (r *SQLProductRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM products where id = ?`

	// Executar a query para deletar o produto
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		// Tratamento de erro, caso algo dê errado com a inserção
		return fmt.Errorf("error delete product: %v", err)
	}

	// Retornar nil em caso de sucesso
	return nil
}

func (r *SQLProductRepository) FindByID(ctx context.Context, id int64) (*model.Product, error) {
	query := `SELECT id, name, price FROM products WHERE id = ?`
	var product model.Product

	// Usar QueryRowContext para executar a query esperando por uma única linha como resultado
	row := r.db.QueryRowContext(ctx, query, id)
	// Escanear o resultado e atribuir aos campos do produto
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			// Se nenhum produto foi encontrado, retornar nil e um erro
			return nil, fmt.Errorf("no product found with id %s", id)
		}
		// Tratamento de outros erros
		return nil, fmt.Errorf("error querying product with id %s: %v", id, err)
	}

	// Retornar o produto encontrado
	return &product, nil
}

func (r *SQLProductRepository) FindAll(ctx context.Context) ([]*model.Product, error) {
	query := `SELECT id, name, price FROM products`
	var products []*model.Product

	// Usar QueryContext para executar a query
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying all products: %v", err)
	}
	defer rows.Close()

	// Iterar sobre todas as linhas retornadas
	for rows.Next() {
		var product model.Product
		// Escanear o resultado e atribuir aos campos do produto atual
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return nil, fmt.Errorf("error scanning product: %v", err)
		}
		// Adicionar o produto à lista de produtos
		products = append(products, &product)
	}
	// Verificar se houve erros durante a iteração
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over products: %v", err)
	}

	// Retornar a lista de produtos
	return products, nil
}
