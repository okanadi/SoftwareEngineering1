// internal/web/middleware.go
package web

import (
	"backend/internal/domain"
	"backend/internal/service"
	"context"
	"net/http"
)

type AuthMiddleware struct {
	userService *service.UserService
}

func NewAuthMiddleware(us *service.UserService) *AuthMiddleware {
	return &AuthMiddleware{userService: us}
}

// RoleMiddleware prüft, ob der User eine der benötigten Rollen hat
func (m *AuthMiddleware) RoleMiddleware(allowedRoles ...domain.UserRole) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			customerLastName := r.Header.Get("lastname")
			userID := r.Header.Get("userID")

			// Prüfen, ob überhaupt eine Identifikation mitgeschickt wurde
			if userID == "" && customerLastName == "" {
				http.Error(w, "Nicht autorisiert: Keine Identifikation bereitgestellt", http.StatusUnauthorized)
				return
			}

			var authUser domain.UserDB
			isAllowed := false

			// --- FALL 1: Interner User (Handwerker, Innendienst, Admin) ---
			if userID != "" {
				user, err := m.userService.GetUserByID(r.Context(), userID)
				if err != nil {
					http.Error(w, "User nicht gefunden", http.StatusForbidden)
					return
				}
				authUser = user

				for _, role := range allowedRoles {
					if user.Role == role {
						isAllowed = true
						break
					}
				}
			} else if customerLastName != "" {
				// Prüfen, ob die aufgerufene Route für Kunden erlaubt ist
				for _, role := range allowedRoles {
					if role == domain.RoleCustomer {
						isAllowed = true
						break
					}
				}

				if isAllowed {
					// Wir erstellen ein temporäres User-Objekt für den Context,
					// damit nachfolgende Handler den Namen/Rolle kennen.
					authUser = domain.UserDB{
						Name: customerLastName,
						Role: domain.RoleCustomer,
					}
				}
			}

			if !isAllowed {
				http.Error(w, "Keine Berechtigung für diese Aktion", http.StatusForbidden)
				return
			}

			// User-Info (entweder aus DB oder virtuell für Kunde) in den Context legen
			ctx := context.WithValue(r.Context(), "user", authUser)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
