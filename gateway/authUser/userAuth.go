package authUser

import (
	"cesargdd/grpc-gateway/pb"
	"cesargdd/grpc-gateway/servers"
	"context"
	"net/http"
)

var UserCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

var u, _ = servers.UserServer()

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			username, err := ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// create user and check if user exists in db
			user := pb.User{Username: username}
			res, err := u.GetUserByUsername(context.Background(), &pb.GetUserByUsernameRequest{Username: username})
			id := res.User.Id
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			user.Id = id
			// put it in context
			ctx := context.WithValue(r.Context(), UserCtxKey, &user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *pb.User {
	raw, _ := ctx.Value(UserCtxKey).(*pb.User)
	return raw
}
