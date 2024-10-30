package config

import (
	"context"
	"fmt"
	"log"

	//	"secposretriever/config"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBstruct struct {
	ConPool    *pgxpool.Pool
	Connection *pgxpool.Conn
	Config     *pgxpool.Config
}

// info sur la library : https://github.com/jackc/pgx
// tuto : https://www.youtube.com/watch?v=sXMSWhcHCf8
// https://github.com/jackc/pgx

// tout ce qui est au dessus est très complet, ptet trop pour le petit besoin. permet d'aller loin.
// pour le petit besoin, cet article semble suffisant :
// https://medium.com/@neelkanthsingh.jr/understanding-database-connection-pools-and-the-pgx-library-in-go-3087f3c5a0c

func DBConfig(dburl string) (dbConfig *pgxpool.Config) {
	const defaultMaxConns = int32(4)
	const defaultMinConns = int32(0)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5

	// Your own Database URL
	DATABASE_URL := dburl
	//	const DATABASE_URL string = "postgres://postgres:12345678@localhost:5432/postgres?"

	dbConfig, err := pgxpool.ParseConfig(DATABASE_URL)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		log.Println("Before acquiring the connection pool to the database!!")
		return true
	}

	dbConfig.AfterRelease = func(c *pgx.Conn) bool {
		log.Println("After releasing the connection pool to the database!!")
		return true
	}

	dbConfig.BeforeClose = func(c *pgx.Conn) {
		log.Println("Closed the connection pool to the database!!")
	}

	return dbConfig
}

// check avec cela: https://medium.com/@neelkanthsingh.jr/understanding-database-connection-pools-and-the-pgx-library-in-go-3087f3c5a0c
// est t'on vraiment obligé de mettre dans le main ? nooon, je fais le blaireau
func (DB *DBstruct) DBConnect() (err error) {

	DB.Config = DBConfig(Envs.DBCONN)

	DB.ConPool, err = pgxpool.NewWithConfig(context.Background(), DB.Config)
	if err != nil {
		log.Fatal("Error while creating connection to the database!!")
	}

	DB.Connection, err = DB.ConPool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error while acquiring connection from the database pool!!")
	}

	defer DB.Connection.Release()

	err = DB.Connection.Ping(context.Background())
	if err != nil {
		log.Fatal("Could not ping database")
	}

	fmt.Println("Connected to the database!!")

	return err

}
