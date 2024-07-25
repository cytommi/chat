package service

import (
	"chat/internal/config"
	"context"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var ErrNoUsername = errors.New("no auth info in context")

type authKey string

const UserIdKey authKey = "username"

type ChatClaims struct {
	UserId string
	jwt.RegisteredClaims
}

type authHelper struct {
	jwtPublicKey  *rsa.PublicKey
	jwtPrivateKey *rsa.PrivateKey
}

func newAuthHelper() (*authHelper, error) {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(config.JwtPublicPEM)
	if err != nil {
		return nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(config.JwtPrivateKeyPEM)
	if err != nil {
		return nil, err
	}

	return &authHelper{
		jwtPublicKey:  publicKey,
		jwtPrivateKey: privateKey,
	}, nil
}

func UserFromContext(ctx context.Context) (string, error) {
	userId, ok := ctx.Value(UserIdKey).(string)
	if !ok {
		return "", ErrNoUsername
	}
	return userId, nil
}

type AuthenticateRequest struct {
	UserId string `json:"userId"`
}

type AuthenticateResponse struct {
	Token string `json:"token"`
}

// TODO: impl actaul auth
func (a *authHelper) Authenticate(w http.ResponseWriter, r *http.Request) {
	var reqBody JoinRoomRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	tk := jwt.NewWithClaims(JwtSigningMethod, &ChatClaims{
		UserId: reqBody.UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:       uuid.New().String(),
			Subject:  JwtSubject,
			Issuer:   JwtIssuer,
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	})

	tkSigned, err := tk.SignedString(a.jwtPrivateKey)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	jsonRes, err := json.Marshal(AuthenticateResponse{
		Token: tkSigned,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	w.Write(jsonRes)
}

func (a *authHelper) WithAuthCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get auth header
		bearer := r.Header.Get("Authorization")
		if !strings.HasPrefix(bearer, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		tkStr := strings.TrimPrefix(bearer, "Bearer ")
		if len(tkStr) <= 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Decode
		tk, err := jwt.ParseWithClaims(tkStr, &ChatClaims{}, func(t *jwt.Token) (any, error) {
			return a.jwtPublicKey, nil
		})
		if err != nil {
			http.Error(w, err.Error(), 401)
		}

		claims, ok := tk.Claims.(*ChatClaims)
		if !ok || !tk.Valid {
			http.Error(w, "invalid claims", 401)
		}

		ctx := context.WithValue(r.Context(), UserIdKey, claims.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
