package admin

const (
	GetAllAdmins = `SELECT id, username, email, name, mobile FROM users WHERE type = 1 AND deleted_at IS NULL ORDER BY id DESC`
	GetAdminByID = `SELECT id, username, email, name, mobile FROM users WHERE id = ? AND type = 1 AND deleted_at IS NULL`
	UpdateAdmin  = `UPDATE users SET username = ?, email = ?, name = ?, mobile = ? WHERE id = ? AND type = 1`
	DeleteAdmin  = `UPDATE users SET deleted_at = NOW() WHERE id = ? AND type = 1`
	InsertAdmin  = `INSERT INTO users (username, email, password, name, mobile, type) VALUES (?, ?, ?, ?, ?, 1)`
)

// GetAdminByUsernameOrEmail is used to fetch an admin by username or email
const GetAdminByUsernameOrEmail = `
	SELECT id, username, email, name, mobile 
	FROM users 
	WHERE (username = ? OR email = ?) AND type = 1 AND deleted_at IS NULL
`

// GetAdminByUsernameOrEmailAndPassword is used to fetch an admin by username or email and password
const GetAdminByUsernameOrEmailAndPassword = `
	SELECT id, username, email, name, mobile, password
	FROM users 
	WHERE (username = ? OR email = ?) AND type = 1 AND deleted_at IS NULL
`

// GetAdminByIDAndPassword is used to fetch an admin by ID and password
const GetAdminByIDAndPassword = `
	SELECT id, username, email, name, mobile, password	
	FROM users 
	WHERE id = ? AND type = 1 AND deleted_at IS NULL
`
