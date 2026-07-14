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

package usecase

import (
	"context"
	"encoding/json"

	"backend/internal/config"
	"backend/internal/domain"

	"github.com/google/uuid"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetOrCreate находит пользователя по telegram_id.
// Если auto_register == true и пользователь не найден — создаёт нового.
func (s *UserService) GetOrCreate(ctx context.Context, telegramID string) (*domain.User, error) {
	user, err := s.repo.GetByTelegramID(ctx, telegramID)
	if err == nil {
		return user, nil
	}
	if err != domain.ErrUserNotFound {
		return nil, err
	}

	cfg := config.Get()
	if !cfg.Features.AutoRegister {
		return nil, domain.ErrUserNotFound
	}

	newUser := &domain.User{
		ID:         uuid.New(),
		TelegramID: telegramID,
		Metadata:   json.RawMessage(`{}`),
	}
	if err := s.repo.Create(ctx, newUser); err != nil {
		return nil, err
	}
	return newUser, nil
}

func (s *UserService) Create(ctx context.Context, telegramID, phone string, meta json.RawMessage) (*domain.User, error) {
	// Проверка на существование
	_, err := s.repo.GetByTelegramID(ctx, telegramID)
	if err == nil {
		return nil, domain.ErrUserAlreadyExists
	}
	if err != domain.ErrUserNotFound {
		return nil, err
	}

	u := &domain.User{
		ID:         uuid.New(),
		TelegramID: telegramID,
		Phone:      phone,
		Metadata:   meta,
	}
	if err := s.repo.Create(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}
