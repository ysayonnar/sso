package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"jwt-go/internal/database"
	"jwt-go/internal/logger"
	"jwt-go/pkg/password"
	"jwt-go/pkg/token"
	"log/slog"
	"net/http"
	"net/mail"
)

type RegistrationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JwtResponse struct {
	JwtToken string `json:"token"`
}

func (r *RegistrationRequest) Validate() error {
	if len(r.Password) < 8 || len(r.Password) > 30 {
		return fmt.Errorf("passwod is too short or long")
	}

	_, err := mail.ParseAddress(r.Email)
	if err != nil {
		return fmt.Errorf("invalid email")
	}

	return nil
}

func Registration(log *slog.Logger, storage *database.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.Registration"

		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Error("body reading error", logger.Error(fmt.Errorf("op: %s, err: %w", op, err)))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var req RegistrationRequest
		err = json.Unmarshal(body, &req)
		if err != nil {
			log.Error("json parsing error", logger.Error(fmt.Errorf("op: %s, err: %w", op, err)))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = req.Validate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%s", err.Error())
		}

		var count int
		err = storage.Db.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1;", req.Email).Scan(&count)
		if err != nil || count == 1 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "user with such email already exists")
			return
		}

		password_hash, err := password.HashPassword(req.Password)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Error("hashing password error", logger.Error(fmt.Errorf("op: %s, err: %w", op, err)))
			return
		}

		var userId int
		err = storage.Db.QueryRow(`INSERT INTO users(email, password_hash) VALUES($1, $2) RETURNING user_id;`, req.Email, password_hash).Scan(&userId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Error("insert error", logger.Error(fmt.Errorf("op: %s, err: %w", op, err)))
			return
		}

		jwtToken, err := token.New(userId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error("jwt-token signing error", logger.Error(err))
			return
		}

		response, err := json.Marshal(JwtResponse{JwtToken: jwtToken})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error("error while marshiling json", logger.Error(err))
			return
		}

		w.Header().Add("Content-type", "application/json")
		fmt.Fprint(w, string(response))
	}
}
