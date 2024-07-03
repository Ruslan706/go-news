package server

type Repository interface{

}

type Server struct{
 db Repository
 func New(db Repository ) *Server{
	return &Server
	db:db,
 }
}