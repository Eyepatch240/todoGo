package security

import (
	"github.com/gin-gonic/gin"
	"time"
	"todoGo/models"

	"github.com/golang-jwt/jwt/v5"
)

// here for ease of set up, will be moved to env variables on vps
var (
	accessTokenSecret  = []byte("c031b0649589c76611e6e4830eafb3b33f9e111295ca11ea430831dc6ed12e63f8b88de55c2e6d790cd6aeac444c5e6f0f22c027470151dedef6cbcf6e0e5caf55f4679adf4f85e6f7e48d4e6295cae3d897563c872975df3a69bfc147dccde41423ec8c28dee4541839a68a57f51b7b07f276953fa357426094dfb9158e46e89bbc0b52b428f073fd05ca11b081586be7c8de401a3744a783dfd79acad02c4a75dd7a02ac6fc202e1fa83ddcd0f91ec250ab9ae938d91584ef9b1228f474f0a42aef2fe8442d24ac82f8a6efb322b9d9639dffbe0c247690431fc687a3744dbb34c3cd0f98fea93fe121b7f42e5490679d3ae9630956e8b02a359f422441cb7")
	refreshTokenSecret = []byte("bc08f8ad8ea2a054e30a247dbc6fdca8e0566827dbc01d7407a9a89a6396dda5a58bf5751ec34f0ee0a66cae7254972885edc540db9bb80a346c383547fcee796383342b9ab157b79d01d91fecf1799c095b1e66cac7ccc1ed04012622e387aa52e8162239861b886cb1063e4b9711da3d37046316e5d3409d829cb4965ce9fc23b8475420e0cf6e524b98b7035679549d1f0c39e16d35f5316c3b159d93d23e8e0cb751cb195a1328ef9b100bf47016a5dac3f90e652af40da14ad024948327deef5b46e271eacbce48cb1924299093801f22a1e2d947d5e0b9bfc0d82a2f43e5aeb52cdf64ebcee5aa2222b59f572481aa1fd0ac7c7c22e2954b6655e6652a")
)

func GenerateToken(user models.User, tokenType string) (string, error) {
	var tokenString string
	var err error

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"iss":     "todoGo",
		"iat":     time.Now().Unix(),
	}

	switch tokenType {
	case "access":
		claims["exp"] = time.Now().Add(5 * time.Minute).Unix() // 5 minutes
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err = token.SignedString(accessTokenSecret)

	case "refresh":
		claims["exp"] = time.Now().Add(15 * 24 * time.Hour).Unix() // 15 days
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err = token.SignedString(refreshTokenSecret)

	default:
		return "", jwt.ErrInvalidKey
	}

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		// grab access token from cookie (accessToken)
		// if(present and valid) c.continue
		// else error, will call /refresh endpoint outside.
	}
}
