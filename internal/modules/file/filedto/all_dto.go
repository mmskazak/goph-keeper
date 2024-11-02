package filedto

import (
	"fmt"
	"goph-keeper/internal/helpers"
	"goph-keeper/internal/modules/file/proto"
	"net/http"
)

type AllFilesDTO struct {
	UserID int `json:"user_id"`
}

func AllFileDTOFromHTTP(r *http.Request) (AllFilesDTO, error) {
	// Извлекаем userID из контекста
	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return AllFilesDTO{}, fmt.Errorf("error AllFileDTOFromHTTP GetUserIDFromContext AllFileDTOFromHTTP: %w", err)
	}
	return AllFilesDTO{
		UserID: userID,
	}, nil
}

func AllFileDTOFromAllFileRequestGRPC(req *proto.GetAllFilesRequest) (proto.GetAllFilesResponse, error) {
	return proto.GetAllFilesResponse{}, nil
}
