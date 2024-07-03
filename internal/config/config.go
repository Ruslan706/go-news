package config

import "flag"

type Config struct {
	Addr   string
	DbAddr string
}

func ReadConfig() Config {
	var addr string
	var dbaddr string
	flag.StringVar(&addr, "addr", "8080", "Server, Address")
	flag.StringVar(&dbaddr, "dbaddr", "", "postgres://postgres:64866550localhost:5432/posgres")
	flag.Parse()
	return Config{
		Addr:   addr,
		DbAddr: dbaddr,
	}
}
