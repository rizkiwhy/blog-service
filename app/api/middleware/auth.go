package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"rizkiwhy-blog-service/api/presenter"
	pkgUser "rizkiwhy-blog-service/package/user"
	mUser "rizkiwhy-blog-service/package/user/model"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type AuthMiddleware struct {
	UserService     pkgUser.Service
	CacheRepository pkgUser.CacheRepository
}

func NewAuthMiddleware(userService pkgUser.Service, cacheRepository pkgUser.CacheRepository) *AuthMiddleware {
	return &AuthMiddleware{
		UserService:     userService,
		CacheRepository: cacheRepository,
	}
}

func (am *AuthMiddleware) AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Error().Msg("[AuthMiddleware][AuthJWT] Missing authorization header")
			c.AbortWithStatusJSON(http.StatusUnauthorized, presenter.FailureResponse(MissingAuthHeaderTitleMessage, MissingAuthHeaderErrorMessage))
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Error().Msg("[AuthMiddleware][AuthJWT] Invalid signing method")
				return nil, errors.New(ErrInvalidSigningMethodMessage)
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			log.Error().Err(err).Msg("[AuthMiddleware][AuthJWT] Failed to parse token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, presenter.FailureResponse(ErrInvalidAuthHeaderTitleMessage, err.Error()))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			log.Error().Msg("[AuthMiddleware][AuthJWT] Invalid token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, presenter.FailureResponse(ErrInvalidAuthHeaderTitleMessage, ErrInvalidTokenMessage))
			return
		}

		valueJWTPayload, err := am.CacheRepository.GetJWTPayload(mUser.GetJWTPayloadRequest{JIT: uuid.MustParse(fmt.Sprintf("%v", claims["jit"]))})
		if err != nil {
			log.Error().Err(err).Msg("[AuthMiddleware][AuthJWT] Failed to get JWT payload from cache")
			c.AbortWithStatusJSON(http.StatusUnauthorized, presenter.FailureResponse(ErrInvalidAuthHeaderTitleMessage, err.Error()))
			return
		}

		err = valueJWTPayload.ValidateTokenClaims(claims)
		if err != nil {
			log.Error().Err(err).Msg("[AuthMiddleware][AuthJWT] Invalid token claims")
			c.AbortWithStatusJSON(http.StatusUnauthorized, presenter.FailureResponse(ErrInvalidAuthHeaderTitleMessage, err.Error()))
			return
		}

		c.Set("user_id", valueJWTPayload.UserID)
		c.Set("email", valueJWTPayload.Email)

		c.Next()
	}
}
