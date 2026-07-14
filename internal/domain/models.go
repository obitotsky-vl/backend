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

package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID       `json:"id"`
	TelegramID string          `json:"telegram_id"`
	Phone      string          `json:"phone,omitempty"`
	Metadata   json.RawMessage `json:"metadata"`
}

type Slot struct {
	ID        uuid.UUID       `json:"id"`
	BranchID  uuid.UUID       `json:"branch_id"`
	StartTime time.Time       `json:"start_time"`
	EndTime   time.Time       `json:"end_time"`
	IsBooked  bool            `json:"is_booked"`
	Metadata  json.RawMessage `json:"metadata"`
}

type Booking struct {
	ID       uuid.UUID       `json:"id"`
	SlotID   uuid.UUID       `json:"slot_id"`
	UserID   uuid.UUID       `json:"user_id"`
	BookedAt time.Time       `json:"booked_at"`
	Status   string          `json:"status"`
	Metadata json.RawMessage `json:"metadata"`
}

type Service struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	DurationMinutes int             `json:"duration_minutes"`
	Price           float64         `json:"price"`
	Metadata        json.RawMessage `json:"metadata"`
}
