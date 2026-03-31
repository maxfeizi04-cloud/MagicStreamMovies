package utils

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/maxfeizi04-cloud/MagicStreamMovies/Server/MagicStreamMoviesServer/database"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var SecretKey = os.Getenv("SECRET_KEY")
var SecretRefreshKey = os.Getenv("SECRET_REFRESH_KEY")
var userCollection = database.OpenCollection("users")

type SignedDetails struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
	UserID    string `json:"user_id"`
	jwt.RegisteredClaims
}

func generateToken(email, firstName, lastName, role, userID string, secret string, duration time.Duration) (string, error) {
	claims := &SignedDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Role:      role,
		UserID:    userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "MagicStream",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}

	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return Token.SignedString([]byte(secret))
}

func GenerateAllTokens(email, firstName, lastName, role, userID string) (string, string, error) {

	accessToken, err := generateToken(
		email, firstName, lastName, role, userID,
		SecretKey,
		24*time.Hour,
	)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := generateToken(
		email, firstName, lastName, role, userID,
		SecretRefreshKey,
		7*24*time.Hour, // 建议更长
	)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func UpdateAllTokens(userID, accessToken, refreshToken string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	updateAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	updateData := bson.M{
		"$set": bson.M{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
			"updated_at":    updateAt,
		},
	}
	_, err := userCollection.UpdateOne(ctx, bson.M{"user_id": userID}, updateData)
	if err != nil {
		return err
	}
	return nil
}

func GetAccessToken(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("authorization")
	if authHeader == "" {
		return "", errors.New("missing authorization header")
	}
	tokenString := authHeader[len("Bearer "):]
	if tokenString == "" {
		return "", errors.New("missing bearer token")
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*SignedDetails, error) {
	claims := &SignedDetails{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, err
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	return claims, nil
}

func GetUserIdFromContext(c *gin.Context) (string, error) {
	userID, exists := c.Get("userID")
	if !exists {
		return "", errors.New("userID does not exists in this context")
	}
	id, ok := userID.(string)
	if !ok {
		return "", errors.New("unable to retrieve userID")
	}
	return id, nil
}
