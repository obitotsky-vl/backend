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

package config

import (
	"encoding/json"
	"os"
	"sync/atomic"
	"time"
)

type FeatureFlags struct {
	HasDepositPayment   bool `json:"has_deposit_payment"`
	HasMultipleBranches bool `json:"has_multiple_branches"`
	AutoRegister        bool `json:"auto_register"`
}

type ServiceConfig struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	DurationMinutes int             `json:"duration_minutes"`
	Price           float64         `json:"price"`
	Metadata        json.RawMessage `json:"metadata"`
}

type Config struct {
	ServerPort         string       `json:"server_port"`
	BusinessID         string       `json:"business_id"`
	BusinessName       string       `json:"business_name"`
	Timezone           string       `json:"timezone"`
	Features           FeatureFlags `json:"features"`
	NotificationMethod string       `json:"notification_method"`
	Telegram           struct {
		BotToken    string `json:"bot_token"`
		AdminChatID string `json:"admin_chat_id"`
	} `json:"telegram"`
	Services []ServiceConfig `json:"services"`
	DSN      string          `json:"-"`
}

var globalConfig atomic.Value

func Load(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}
	cfg.DSN = os.Getenv("DB_DSN")
	if cfg.ServerPort == "" {
		cfg.ServerPort = "8080"
	}
	loc, err := time.LoadLocation(cfg.Timezone)
	if err != nil {
		return nil, err
	}
	time.Local = loc

	globalConfig.Store(&cfg)
	return &cfg, nil
}

func Get() *Config {
	return globalConfig.Load().(*Config)
}
