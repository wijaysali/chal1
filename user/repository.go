package user

import (
	"database/sql"
)

// UserRepository represents the database repository for user objects.
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new UserRepository with the given database connection.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// CreateUser inserts a new user object into the database and returns its ID.
func (ur *UserRepository) CreateUser(user *User) (int64, error) {
	result, err := ur.db.Exec("INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)",
		user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// GetUserByID retrieves a user object from the database by ID.
func (ur *UserRepository) GetUserByID(id int) (*User, error) {
	user := User{}

	err := ur.db.QueryRow("SELECT * FROM users WHERE id=?", id).
		Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser updates a user object in the database.
func (ur *UserRepository) UpdateUser(user *User) error {
	_, err := ur.db.Exec("UPDATE users SET first_name=?, last_name=?, email=?, password=? WHERE id=?",
		user.FirstName, user.LastName, user.Email, user.Password, user.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser removes a user object from the database by ID.
func (ur *UserRepository) DeleteUser(id int) error {
	_, err := ur.db.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}
