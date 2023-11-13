package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/srivardhanreddy01/webapplication_go/api/handlers"
)

type AuthMiddlewareDependencies struct {
	AuthHandlerDependencies handlers.AuthHandlerDependencies
}

func BasicAuthMiddleware(deps AuthMiddlewareDependencies, next http.Handler) http.Handler {
	// fmt.Print("hello")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		email, password, ok := r.BasicAuth()
		// fmt.Println(email)
		// fmt.Println(password)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// fmt.Println(email)
		// fmt.Println(password)

		// Authenticate the user based on the provided username and password
		user, err := handlers.AuthenticateUser(deps.AuthHandlerDependencies, email, password)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if user == nil {
			// User not found or authentication failed
			log.Printf("User not found or authentication failed for email: %s", email)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If authentication is successful, you can proceed to the next handler
		// and attach the authenticated user to the request context.
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
