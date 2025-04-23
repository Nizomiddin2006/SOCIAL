package main

import (
	"SOCIAL/internal/db"
	"SOCIAL/internal/store"
	"log"

	_ "github.com/lib/pq"
	
	"SOCIAL/internal/env"
)

func main() {
	// CONFIG: Atrof-muhit o'zgaruvchilaridan kerakli sozlamalarni olish
	cfg := config{
		addr: env.GetString("ADDR", ":8081"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost:5433/social_db?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database connection pool established")

	store := store.NewStorage(db)

	// APPLICATION: Application obyekti yaratish
	app := &application{
		config: cfg,
		store:  store,
	}

	// MUX: HTTP marshrutizatorni olish
	mux := app.mount()

	// SERVER: HTTP serverni ishga tushirish
	log.Fatal(app.run(mux))
}
