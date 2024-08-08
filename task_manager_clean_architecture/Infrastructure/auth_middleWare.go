package Infrastructure

import (

)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		jwtService := NewJWTService(jwtSecret)
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")

		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized , gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		token, err := jwtService.VerifyToken(authParts[1])

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		username := claims["username"].(string)
		r := claims["role"].(string)

		user := &domain.User{
			UserName: username,
			Role:     r,
        	}
        	c.Set("user", user)

		c.Next()
	}
}

func RoleBasedMiddleware(requiredRole string) gin.HandlerFunc {

	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		u, ok := user.(*domain.User)

		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if u.Role != requiredRole {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
