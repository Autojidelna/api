package auth

import (
	"context"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

var FirebaseClient *firebase.App
var authClient *auth.Client

func InitFirebase() {
	// [START initialize_app]
	client, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		if gin.DebugMode == "debug" {
			log.Println("error initializing Firebase: %v\n", err)
		}
		log.Fatalf("error initializing app: %v\n", err)
	}
	authyClient, err := client.Auth(context.Background())
	if err != nil {
		if gin.DebugMode == "debug" {
			log.Println("Failed to get Firebase Auth client: %v\n", err)
		} else {
			log.Fatalf("Failed to get Firebase Auth client: %v", err)
		}
	}
	authClient = authyClient
	FirebaseClient = client
}

// Is user logged in to firebase?
func FirebaseAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Check if the token starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		// Extract the token from the header
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Verify the token using Firebase Auth
		decodedToken, err := authClient.VerifyIDToken(context.Background(), token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Add the decoded token to the context so other handlers can use it
		c.Set("firebaseUser", decodedToken)

		// Continue to the next middleware/handler
		c.Next()
	}
}
