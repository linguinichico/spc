package storage

import (
	"database/sql"
	"os"

	t "github.com/linguinichico/spc/api/types"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

// Initialize a new connection to a database.
func NewPostgresStore() (*PostgresStore, error) {

	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_IP := os.Getenv("POSTGRES_IP")
	POSTGRES_PORT := os.Getenv("POSTGRES_PORT")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")

	connectionString := "postgres://" + POSTGRES_USER + ":" + POSTGRES_PASSWORD + "@" + POSTGRES_IP + ":" + POSTGRES_PORT + "/" + POSTGRES_DB + "?sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) ClosePostgresStore() error {
	err := s.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) CreateGroupUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(100) NOT NULL,
			password VARCHAR(100) NOT NULL,
			created timestamp
	)`
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) CreateUser(user *t.User) error {
	query := `INSERT INTO users
			(username, password, created)
			values ($1,$2,$3)
	`
	_, err := s.db.Query(query,
		user.Username,
		user.Password,
		user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) GetUsers() (*t.GroupUsers, error) {
	query := `SELECT * FROM users`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	group_users := &t.GroupUsers{}
	for rows.Next() {
		user := &t.User{}
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.CreatedAt)
		if err != nil {
			return nil, err
		}
		group_users.Users = append(group_users.Users, user)
	}
	return group_users, nil

}
