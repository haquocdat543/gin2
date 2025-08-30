package share

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"gin/src/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

var jwtSecret = []byte("your_super_secret_key") // Use a strong, secure key

func LoadPrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key.(*rsa.PrivateKey), nil
}

func LoadPublicKey(publicKey string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key.(*rsa.PublicKey), nil
}

func CreateJWKSet() jwk.Set {

	publicKey, err := LoadPublicKey(config.ENV.JWTPublicKey)
	if err != nil {
		panic(config.LOAD_JWT_PUBLIC_KEY_ERROR)
	}

	jwkKey, err := jwk.FromRaw(publicKey)
	if err != nil {
		panic(err)
	}

	jwkKey.Set(jwk.KeyIDKey, "my-key-id")

	jwkSet := jwk.NewSet()
	jwkSet.AddKey(jwkKey)

	return jwkSet
}

func GenerateToken(username string, ip string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"ip":       ip,
		"exp":      time.Now().Add(config.GlobalJWTTimeToLive).Unix(), // Token expires in 24 hours
		"iat":      time.Now().Unix(),
	}

	privateKey, err := LoadPrivateKey(config.ENV.JWTPrivateKey)
	if err != nil {
		panic(config.LOAD_JWT_PRIVATE_KEY_ERROR)
	}

	return jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(privateKey)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix if present
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		publicKey, err := LoadPublicKey(config.ENV.JWTPublicKey)
		if err != nil {
			panic(config.LOAD_JWT_PUBLIC_KEY_ERROR)
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return publicKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		username, ok := claims["username"].(string)
		if !ok || username == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Username missing in token"})
			c.Abort()
			return
		}

		// IP binding check
		tokenIP, ok := claims["ip"].(string)
		if !ok || tokenIP != c.ClientIP() {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "IP address mismatch — possible token misuse"})
			c.Abort()
			return
		}

		c.Set("username", username)
		c.Next()
	}
}

// GetUsername extracts the "username" claim from the Gin context.
// This assumes your AuthMiddleware has already stored it.
func GetUsername(c *gin.Context) (string, error) {
	value, exists := c.Get("username")
	if !exists {
		return "", errors.New("username not found in context")
	}

	username, ok := value.(string)
	if !ok || username == "" {
		return "", errors.New("username in context is invalid")
	}

	return username, nil
}
