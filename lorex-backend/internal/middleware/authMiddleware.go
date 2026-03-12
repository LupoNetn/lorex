package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luponetn/lorex/utils"
)


func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenParts := c.GetHeader("Authorization")
		if tokenParts == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization token required"})
			slog.Error("authorization token required")
			return
		}

		// get token key
		tokenString := utils.ExtractToken(tokenParts)
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization token required"})
			slog.Error("authorization token required")
			return
		}

		// validate token
		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			slog.Error("invalid or expired token", "error", err)
			return
		}

		user := struct {
			ID    string
			Email string
			Role  string
		}{
			ID:    claims.ID,
			Email: claims.Email,
			Role:  claims.Role,
		}

		// set user and claims to context
		c.Set("user", user)
		c.Set("claims", claims)
		c.Next()
	}
}

// RoleMiddleware checks if the user has one of the required roles
func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		val, exists := c.Get("claims")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: no claims found"})
			return
		}

		claims, ok := val.(*utils.CustomClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to parse claims"})
			return
		}

		roleMatch := false
		for _, role := range roles {
			if claims.Role == role {
				roleMatch = true
				break
			}
		}

		if !roleMatch {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden: insufficient permissions"})
			return
		}

		c.Next()
	}
}

// CompanyOnly restricts access to only users with the 'company' role
func CompanyOnly() gin.HandlerFunc {
	return RoleMiddleware("company")
}

// DriverOnly restricts access to only users with the 'driver' role
func DriverOnly() gin.HandlerFunc {
	return RoleMiddleware("driver")
}

// CustomerOnly restricts access to only users with the 'customer' role
func CustomerOnly() gin.HandlerFunc {
	return RoleMiddleware("customer")
}

