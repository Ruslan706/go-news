package config

import "flag" 

type Config struct{
	Addr string
	dbaddr string
}
func Readconfig() Config{
	var addr string
	var dbaddr string
	flag.StringVar (&. "addr","8080","Server, Address")
	flag.StringVar(&, "dbaddr","", "postgres://postgres:64866550localhost:5432/posgres )
	flag.Parse()
	return Config{
		addr: addr,
		Dbaddr: dbaddr,
	}
}