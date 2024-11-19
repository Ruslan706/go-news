package repository

import (
	"context"
	"time"

	"github.com/Ruslan706/go-news/commit/86c72d46a63efd17c0da10c052fdbd455e194421/internal/domain/models"
	"github.com/jackc/pgx/v5"
)

type DBStorage struct {
	conn *pgx.Conn
}

func NewDB(conn *pgx.Conn) DBStorage {
	return DBStorage{
		conn: conn,
	}
}

func (db *DBStorage) GeatAllUser() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, err := db.conn.Query(ctx, "SELECT name, login, password  FROM users")
	if err != nil {
		return nil, err
	}
	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Name, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
