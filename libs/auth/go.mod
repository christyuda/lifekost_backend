module lifekost/libs/auth

go 1.24.4

require (
	lifekost/auth-services/pkg/domain v0.0.0
	github.com/golang-jwt/jwt/v5 v5.2.2
)

replace lifekost/auth-services/pkg/domain => ../../services/auth-services/pkg/domain
