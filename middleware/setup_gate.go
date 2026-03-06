package middleware

import (
	"app/constant"
	"app/gen/setup/v1/setupv1connect"
	"context"
	"encoding/json"
	"net/http"
)

type SetupStatusChecker interface {
	IsSetupCompleted(ctx context.Context) (bool, error)
}

func WithSetupGate(next http.Handler, checker SetupStatusChecker) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isCompleted, err := checker.IsSetupCompleted(r.Context())
		if err != nil {
			writeSetupJSON(w, http.StatusInternalServerError, map[string]any{
				"success": false,
				"message": constant.ErrInternalServer.Error(),
			})
			return
		}

		path := r.URL.Path
		if !isCompleted {
			if !isAllowedDuringSetup(path) {
				writeSetupJSON(w, http.StatusForbidden, map[string]any{
					"success": false,
					"message": constant.ErrSetupRequired.Error(),
				})
				return
			}
		}

		if isCompleted && isSetupOnlyPath(path) {
			writeSetupJSON(w, http.StatusConflict, map[string]any{
				"success": false,
				"message": constant.ErrSetupAlreadyCompleted.Error(),
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isAllowedDuringSetup(path string) bool {
	return path == setupv1connect.SetupGetStatusProcedure || path == setupv1connect.SetupRunSetupProcedure
}

func isSetupOnlyPath(path string) bool {
	return path == setupv1connect.SetupRunSetupProcedure
}

func writeSetupJSON(w http.ResponseWriter, statusCode int, payload map[string]any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(payload)
}
