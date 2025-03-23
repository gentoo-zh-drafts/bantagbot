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

	bot.Debug = false
	log.Printf("Bot started: @%s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil || update.Message.Text == "" {
			continue // 忽略非文本消息
		}

		// 转换为小写进行统一比较
		lowerText := strings.ToLower(update.Message.Text)
		if strings.HasPrefix(lowerText, "#gentoozh") {
			// 构建删除请求
			deleteConfig := tgbotapi.DeleteMessageConfig{
				ChatID:    update.Message.Chat.ID,
				MessageID: update.Message.MessageID,
			}

			// 执行删除并添加日志
			if _, err := bot.Request(deleteConfig); err != nil {
				log.Printf("删除失败 [%d]: %v", update.Message.MessageID, err)

				// 可选：发送错误提示
				// warnMsg := tgbotapi.NewMessage(
				//	 update.Message.Chat.ID,
				//	 "⚠️ 删除失败: "+err.Error(),
				// )
				// bot.Send(warnMsg)
			} else {
				log.Printf("已删除消息 [%d]: %s",
					update.Message.MessageID,
					truncateText(update.Message.Text, 20), // 截断长文本
				)
				// 可选：发送删除提示
				deleteMsg := tgbotapi.NewMessage(
					update.Message.Chat.ID, "tag 已删除",
				)
				bot.Send(deleteMsg)

			}
		}
	}
}

// 辅助函数：截断过长的消息文本
func truncateText(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
