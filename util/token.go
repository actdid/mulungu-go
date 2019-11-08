package util

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/actdid/mulungu/constant"
)

//TokenSecret FIXME:obtain from application confiugrations logic, should be stored in datastore as part of application setup sequence
const TokenSecret = "Eo0Ac6AxohJ9aithu6Seom1ree2Huojohna2oJohT9oqueeCiseer4AengahyohMee9aeFoz1Xea8bei4liXoerieDaicahch4eecoh9aayaichin9ahmahboh7baseich5thahz9oochei1thoo6Waig4eJeemie7rou5as9AeNietiojah5Eig0iropuijeiJohraero3ooXu4amaengiori6Lie6uu5jaabeek7iew0usaNg3johchohch9liv9thi1queivahru8eengoecaidaph0ahsohr5Oo9iehaev4foos7ingaefuweiDo9Pei4aelai4ahhooXai4ahshahJ6tae8Eoyoog0ohpe5oa0tiel6ahph5ief9Ao7phai5om3pefae2ohquafeerah7liet2aedielomo8ul0Apo7Gijiu3eiyeePoogh6aiThii9Aedek7miu0gaith9thu1shohs9taiya7aehochixuquaibeirae3eibe"

//Token obtains *jwt.Token from request
func Token(r *http.Request) (*jwt.Token, error) {
	tokenPrefix := "Bearer "
	jwtToken := r.Header.Get(constant.HeaderAuthorization)
	if jwtToken != "" && strings.Contains(jwtToken, tokenPrefix) {
		//we want to get the token, so trim away the tokenPrefix
		tokenString := strings.TrimPrefix(jwtToken, tokenPrefix)
		token, tokenParseError := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			//FIXME: validation of token signature required
			// if reflect.TypeOf(token.Method) == *jwt.SigningMethodHS512 {
			// 	logger.Criticalf(ctx, "middleware authorization", "invalid token received, token not signed by us, consider rejecting request")
			// }
			return []byte(TokenSecret), nil
		})

		return token, tokenParseError
	}
	return nil, nil
}

//Claims returns claims from request
func Claims(r *http.Request) (jwt.Claims, error) {
	token, tokenError := Token(r)
	if tokenError != nil {
		return nil, tokenError
	}

	claims := token.Claims.(jwt.MapClaims)

	return claims, nil
}

//ClaimExists checks if request has claim
func ClaimExists(key string, r *http.Request) (bool, error) {
	claims, claimsError := Claims(r)
	if claimsError != nil {
		return false, claimsError
	}

	if _, ok := claims.(jwt.MapClaims)[key]; ok {
		return true, nil
	}

	return false, nil
}

//Claim returns claim
func Claim(key string, r *http.Request) (interface{}, error) {
	claims, claimsError := Claims(r)
	if claimsError != nil {
		return false, claimsError
	}

	if claim, ok := claims.(jwt.MapClaims)[key]; ok {
		return claim, nil
	}

	return nil, fmt.Errorf("can't find claim %s", key)
}

//GenerateJWTToken generates and returns a signed JWTToken string use with header Authorization Bearer generatedTokenString
func GenerateJWTToken(signingMethod jwt.SigningMethod, claims jwt.Claims, signingKey string) string {
	if signingMethod == nil {
		signingMethod = jwt.SigningMethodHS256
	}

	JWTToken := jwt.NewWithClaims(signingMethod, claims)
	signedJWTToken, signingError := JWTToken.SignedString([]byte(signingKey))

	if signingError != nil {
		//log.Error("failed to sign JWTToken with provided signing key")
		return ""
	}

	return signedJWTToken
}

// Audience  string `json:"aud,omitempty"`
// ExpiresAt int64  `json:"exp,omitempty"`
// Id        string `json:"jti,omitempty"`
// IssuedAt  int64  `json:"iat,omitempty"`
// Issuer    string `json:"iss,omitempty"`
// NotBefore int64  `json:"nbf,omitempty"`
// Subject   string `json:"sub,omitempty"`

//Claims structure for customized claims
// type Claims struct {
// 	Roles    []string
// 	TenantID int64
// 	UserID   int64
// 	jwt.StandardClaims
// }
//
// //NewClaim generates a new security claim
// func NewClaim(roles []string, tenantID int64, UserID int64, jti string, aud string, exp int64, iss string, sub string) Claims {
// 	claims := Claims{roles, tenantID, UserID,
// 		jwt.StandardClaims{
// 			Id:        jti,
// 			Audience:  aud,
// 			ExpiresAt: exp,
// 			IssuedAt:  time.Now().Unix(),
// 			NotBefore: time.Now().Unix(),
// 			Issuer:    iss,
// 			Subject:   sub}}
//
// 	return claims
// }
