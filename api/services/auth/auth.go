package auth

import (
	"github.com/mikunalpha/mog/api/shared/auth"
	"github.com/mikunalpha/mog/api/shared/utils"
)

// Init sets up jwt token key.
func Init() error {
	auth.SetTokenSecretKey(utils.EnvString("AUTH_JWT_TOKEN_SECRET_KEY", "jwttokensecretkey"))

	return nil
}
