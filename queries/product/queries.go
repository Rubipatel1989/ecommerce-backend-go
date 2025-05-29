package product

const (
	GetAllProducts = `SELECT id, name, price, stock FROM products where deleted_at is null ORDER BY id DESC`
	GetProductByID = `SELECT id, name, price, stock FROM products WHERE id = ? AND deleted_at IS NULL`
	UpdateProduct  = `UPDATE products SET name = ?, price = ?, stock = ? WHERE id = ?`
	DeleteProduct  = `UPDATE products SET deleted_at = NOW() WHERE id = ?`
	InsertProduct  = `INSERT INTO products (name, price, stock) VALUES (?, ?, ?)`
)
