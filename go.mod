module go-mysql-react

go 1.15

require (
	github.com/gorilla/mux v1.8.0
	// github.com/joho/godotenv v1.3.0
	go-mysql-react/helpers v0.1.0
	go-mysql-react/models v0.1.0
)

replace (
	go-mysql-react/helpers => ./helpers
	go-mysql-react/models => ./models
)