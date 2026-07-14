/*
   Copyright (C) 2026 Gleb Obitotsky

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program. If not, see <https://www.gnu.org/licenses/>.

   --- COMMERCIAL USE NOTICE ---
   This open-source license (AGPLv3) STRICTLY REQUIRES you to open your source
   code if you run this application on a server for commercial services.

   If you want to use this software for your business WITHOUT opening your
   source code, you MUST purchase a commercial license from the author.

   To purchase a commercial license, customize the system, or get tech support,
   please contact the author directly via:
   - Telegram: https://t.me/scip1o
   - Email 1: glebobitockij3@gmail.com
   - Email 2: glebobitotsky@yandex.com
   - VK: https://vk.me/imperoxxx
*/

package handler

import (
	"encoding/json"
	"net/http"

	"backend/internal/user/usecase"
)

type UserHandler struct {
	svc *usecase.UserService
}

func NewUserHandler(svc *usecase.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// Get возвращает пользователя по telegram_id (query-параметр)
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	telegramID := r.URL.Query().Get("telegram_id")
	if telegramID == "" {
		http.Error(w, "telegram_id required", http.StatusBadRequest)
		return
	}

	user, err := h.svc.GetOrCreate(r.Context(), telegramID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Create регистрирует нового пользователя
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		TelegramID string          `json:"telegram_id"`
		Phone      string          `json:"phone,omitempty"`
		Metadata   json.RawMessage `json:"metadata,omitempty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if req.TelegramID == "" {
		http.Error(w, "telegram_id required", http.StatusBadRequest)
		return
	}

	user, err := h.svc.Create(r.Context(), req.TelegramID, req.Phone, req.Metadata)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
