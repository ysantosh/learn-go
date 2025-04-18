package main

import (
	"context"
	"flag"
	"log"
	"time"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Script for connecting to postgre DB and test the connect time and execution time.
func main() {
	dbURI := flag.String("h", "postgres://user:${pass}@ip-address:5432/dbname?sslmode=require", "DB URI")
  // replace $ by %24, @ by %40, # by %23 when setting value for pass
	flag.Parse()
  // for i := range 10 is failing with error "cannot range over 5 (untyped int constant)"
	for i := 1; i <= 10; i++ {
		sqlxConnect(*dbURI, i)
	}
}

func sqlxConnect(dbURI string, iteration int) {
	pctx := context.Background()

	st := time.Now()
	db, err := sqlx.Connect("postgres", dbURI)
	if err != nil {
		log.Fatalf("[FATAL] failed to connect: %v", err)
	}
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)
	defer db.Close()
	connElapsed := time.Since(st).Milliseconds()

	ctx, cancel := context.WithTimeout(pctx, 1*time.Second)
	defer cancel()

	st = time.Now()
	if err := db.PingContext(ctx); err != nil {
		log.Printf("[ERROR] %d) Connect time: %d ms, Ping failed: %v", iteration, connElapsed, err)
		return
	}
	log.Printf("[INFO] %d) Connect time: %d ms, Ping time: %d ms", iteration, connElapsed, time.Since(st).Milliseconds())
}
