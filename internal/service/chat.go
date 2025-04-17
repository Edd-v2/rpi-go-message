package service

import (
	"time"

	"github.com/Edd-v2/rpi-go-message/dto"
	db_model "github.com/Edd-v2/rpi-go-message/internal/model/db"
	"github.com/Edd-v2/rpi-go-message/internal/repository"
	"github.com/sirupsen/logrus"
)

func StartPrivateChat(userID, targetID string, log *logrus.Logger) (*db_model.Chat, error) {
	return repository.FindOrCreatePrivateChat(userID, targetID, log)
}

func GetUserChats(userID string, log *logrus.Logger) ([]dto.ChatSummaryResponse, error) {
	chats, err := repository.GetChatsByUserID(userID)
	if err != nil {
		log.Errorf("[service] GetUserChats DB error: %v", err)
		return nil, err
	}

	var result []dto.ChatSummaryResponse

	for _, chat := range chats {
		memberDTOs := make([]dto.UserResponse, 0)
		for _, m := range chat.Members {
			user, err := repository.FindUserByID(m.Hex())
			if err == nil {
				memberDTOs = append(memberDTOs, dto.UserResponse{
					ID:       user.ID.Hex(),
					Username: user.Username,
				})
			}
		}

		var lastMsg *dto.LastMessageDto
		if chat.LastMessage != nil {
			msg, err := repository.FindMessageByID(chat.LastMessage.Hex())
			if err == nil {
				lastMsg = &dto.LastMessageDto{
					Content:   msg.Content,
					Timestamp: msg.Timestamp.Format(time.RFC3339),
				}
			}
		}

		result = append(result, dto.ChatSummaryResponse{
			ID:          chat.ID.Hex(),
			IsGroup:     chat.IsGroup,
			Name:        chat.Name,
			Members:     memberDTOs,
			LastMessage: lastMsg,
		})
	}

	return result, nil
}
