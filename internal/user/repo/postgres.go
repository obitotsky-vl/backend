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

package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"backend/internal/domain"

	"github.com/google/uuid"
)

type UserPostgresRepo struct {
	db *sql.DB
}

func NewUserPostgresRepo(db *sql.DB) *UserPostgresRepo {
	return &UserPostgresRepo{db: db}
}

func (r *UserPostgresRepo) GetByTelegramID(ctx context.Context, telegramID string) (*domain.User, error) {
	row := r.db.QueryRowContext(ctx,
		`SELECT id, telegram_id, phone, metadata FROM users WHERE telegram_id = $1`,
		telegramID,
	)
	var u domain.User
	err := row.Scan(&u.ID, &u.TelegramID, &u.Phone, &u.Metadata)
	if err == sql.ErrNoRows {
		return nil, domain.ErrUserNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("get user by telegram: %w", err)
	}
	return &u, nil
}

func (r *UserPostgresRepo) Create(ctx context.Context, user *domain.User) error {
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}
	if user.Metadata == nil {
		user.Metadata = json.RawMessage(`{}`)
	}
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO users (id, telegram_id, phone, metadata)
         VALUES ($1, $2, $3, $4)`,
		user.ID, user.TelegramID, user.Phone, user.Metadata,
	)
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}
	return nil
}
