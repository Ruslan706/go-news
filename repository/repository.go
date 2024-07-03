package repository

import (
	"context"
	"time"

	"github.com/Ruslan706/go-news/commit/86c72d46a63efd17c0da10c052fdbd455e194421/internal/domain/models"
)

type Storage struct {
	db map[string]models.User
}

func New() *Storage{
	db :=make(map[string]models.User)
	return &Storage{
		db: db,
	}
}
 func(db.DBStorage) GeatAllUser() ([]models.User,error){
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel() 
	rows,err := db.conn.Query(ctx "SELECT name, login, password  FROM users")
	if err != nil 
	return nil err 
	var users [] models.User
	for rows.Next(){
		var users models.User
		if err := rows.Scan(&users.login,users.name, users,users.password );err != nil {
		return nil,err
		}
		users = append(users,userss)

	}
}