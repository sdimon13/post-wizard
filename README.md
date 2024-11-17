# Post Wizard

**Post Wizard** — это инструмент для автоматического поиска и публикации постов из различных источников (сайтов, Telegram-каналов и т.д.) в ваши социальные сети или каналы.

## 🌟 Возможности

- **Автопостинг:** автоматически публикует интересные посты, найденные в указанных источниках.
- **Подключение к Telegram API:** позволяет взаимодействовать с Telegram-каналами и обрабатывать сообщения.
- **Гибкая настройка:** хранение конфигурации в `.env` файле для удобного управления.

## 🚀 Установка

1. **Склонируйте репозиторий:**
   ```bash
    git clone https://github.com/ваш-пользователь/post-wizard.git
    cd post-wizard

2. **Установите зависимости:**
   ```bash
    go mod tidy

3. **Создайте файл .env: В корне проекта создайте файл .env и добавьте следующие переменные:**
   ```bash
    APP_ID=ваш_app_id
    APP_HASH=ваш_app_hash
    PHONE_NUMBER=ваш_номер_телефона

4. **Запустите приложение:**
   ```bash
    go run main.go

## ⚙️ Использование
На текущем этапе приложение:

- Подключается к Telegram через API.
- Логинится с использованием вашего номера телефона.
- Обрабатывает входящие сообщения в приватных чатах и отправляет ответ "**Hello from Gogram!**".

## 📄 Пример кода

### main.go
```go

package main

import (
	"github.com/amarnathcjd/gogram/telegram"
	"log"
	"post-wizard/config"
)

func main() {
	cfg := config.LoadConfig()

	client, err := telegram.NewClient(telegram.ClientConfig{
		AppID:   cfg.AppID,
		AppHash: cfg.AppHash,
	})

	if err != nil {
		log.Fatal(err)
	}

	client.Conn()

	client.Login(cfg.PhoneNumber)

	client.On(telegram.OnMessage, func(message *telegram.NewMessage) error {
		message.Reply("Hello from Gogram!")
		return nil
	}, telegram.FilterPrivate)

	client.Idle()
}
```

### config.go
```go
package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppID       int32
	AppHash     string
	PhoneNumber string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Не удалось загрузить .env файл, используем переменные окружения из системы")
	}

	appIDStr := os.Getenv("APP_ID")
	appHash := os.Getenv("APP_HASH")
	phoneNumber := os.Getenv("PHONE_NUMBER")

	if appIDStr == "" || appHash == "" || phoneNumber == "" {
		log.Fatal("Не удалось получить APP_ID, APP_HASH или PHONE_NUMBER из переменных окружения")
	}

	// Преобразуем APP_ID из строки в int32
	appID, err := strconv.Atoi(appIDStr)
	if err != nil {
		log.Fatalf("Ошибка преобразования APP_ID в число: %v", err)
	}

	return &Config{
		AppID:       int32(appID),
		AppHash:     appHash,
		PhoneNumber: phoneNumber,
	}
}
```


## 🛠 Конфигурация
Конфигурация приложения хранится в файле .env. Используемые переменные окружения:

- `APP_ID` — ID вашего приложения Telegram API.
- `APP_HASH` — хэш вашего приложения Telegram API.
- `PHONE_NUMBER` — номер телефона, привязанный к Telegram.

## 🌱 Развитие
**Планируемые функции:**

- Поддержка граббинга постов с веб-сайтов.
- Настройка фильтров для отбора интересного контента.
- Подключение дополнительных соцсетей для автопостинга.

## 📜 Лицензия
Проект распространяется под лицензией **MIT**. Подробности см. в файле [LICENSE](LICENSE).

---
💡 **Вклад в проект**:

Если у вас есть идеи для улучшения, создайте новый issue или отправьте pull request. Мы будем рады вашему участию!