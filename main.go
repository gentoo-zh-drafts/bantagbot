package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token := os.Getenv("tgbot_token")
	if token == "" {
		log.Fatal("tgbot_token environment variable is not set")
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false // 启用调试模式
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue // 忽略非消息更新
		}

		// 检查消息是否以 #GentooZh 开头
		if strings.HasPrefix(update.Message.Text, "#GentooZh") {
			// 构建删除请求
			deleteConfig := tgbotapi.DeleteMessageConfig{
				ChatID:    update.Message.Chat.ID,
				MessageID: update.Message.MessageID,
			}

			// 执行删除操作
			_, err := bot.Request(deleteConfig)
			if err != nil {
				log.Printf("删除消息失败: %v", err)
				// 可选：发送错误提示给群组
				// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "删除消息失败: "+err.Error())
				// bot.Send(msg)
			} else {
				log.Printf("已删除消息 ID %d", update.Message.MessageID)
			}
		}
	}
}
