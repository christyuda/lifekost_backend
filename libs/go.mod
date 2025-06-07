module github.com/lifekost/libs

go 1.24.4

require (
	github.com/christyuda/lifekost_backend/services/auth-services v0.0.0
	github.com/golang-jwt/jwt/v5 v5.2.2
)

replace github.com/christyuda/lifekost_backend/services/auth-services => ../services/auth-services

replace github.com/christyuda/lifekost_backend/libs/auth => ../../libs/auth
