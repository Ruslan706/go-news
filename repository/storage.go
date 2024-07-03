package repository

import "github.com/jackc/pgx/v5"

type DBStorage struct {
	conn *pgx.conn
}

func NewDB(conn *pgx.Conn) DBStorage {
	return DBStorage{
		conn: conn,
	}

}
