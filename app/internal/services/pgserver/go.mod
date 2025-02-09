module github.com/93mmm/chatting-app/app/internal/services/pgserver

go 1.23.5

replace github.com/93mmm/chatting-app/app/pkg => ../../../pkg

require (
	github.com/93mmm/chatting-app/app/pkg v0.0.0-00010101000000-000000000000
	github.com/jackc/pgx/v5 v5.7.2
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)
