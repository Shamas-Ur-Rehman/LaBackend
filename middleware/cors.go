package middleware

import "github.com/gin-gonic/gin"

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, X-Requested-With, X-HTTP-Method-Override, Origin, Cache-Control, Pragma")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Max-Age", "86400")
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}

// package middleware

// import "github.com/gin-gonic/gin"

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Set CORS headers for all requests
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

// 		// FIXED: Add missing headers that your frontend needs
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, X-Requested-With, X-HTTP-Method-Override")

// 		// FIXED: Ensure all methods are properly listed
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

// 		// FIXED: Add expose headers for better debugging
// 		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type")

// 		// Handle preflight OPTIONS requests properly
// 		if c.Request.Method == "OPTIONS" {
// 			// FIXED: Return 200 instead of 204 for better compatibility
// 			c.AbortWithStatus(200)
// 			return
// 		}

// 		c.Next()
// 	}
// }

// // Alternative: If you want to be more specific about origins (recommended for production)
// func CORSMiddlewareSecure() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		origin := c.Request.Header.Get("Origin")

// 		// Allow your specific domains
// 		allowedOrigins := []string{
// 			"https://www.laorganics310.com",
// 			"https://laorganics310.com",
// 			"http://localhost:3000",
// 			"http://localhost:3001",
// 			"https://laorganics.vercel.app",
// 		}

// 		isAllowed := false
// 		for _, allowedOrigin := range allowedOrigins {
// 			if origin == allowedOrigin {
// 				isAllowed = true
// 				break
// 			}
// 		}

// 		if isAllowed {
// 			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
// 		}

// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, X-Requested-With, X-HTTP-Method-Override")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
// 		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(200)
// 			return
// 		}

// 		c.Next()
// 	}
// }
