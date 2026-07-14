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

package sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"backend/internal/config"
)

type TelegramSender struct {
	token string
}

func NewTelegramSender() *TelegramSender {
	cfg := config.Get()
	return &TelegramSender{token: cfg.Telegram.BotToken}
}

type sendMessageReq struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func (s *TelegramSender) Send(chatID string, text string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", s.token)
	body, _ := json.Marshal(sendMessageReq{ChatID: chatID, Text: text})
	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("telegram api returned status %d", resp.StatusCode)
	}
	return nil
}
