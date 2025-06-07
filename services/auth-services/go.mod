module github.com/christyuda/lifekost_backend/services/auth-service

go 1.24.4

require (
	github.com/golang-jwt/jwt/v5 v5.2.2
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
	golang.org/x/crypto v0.39.0
	github.com/christyuda/lifekost_backend/libs/auth v0.0.0
	github.com/christyuda/lifekost_backend/services/auth-services v0.0.0
)

replace github.com/christyuda/lifekost_backend/services/auth-services => ../services/auth-services

replace github.com/christyuda/lifekost_backend/libs/auth => ../libs/auth
