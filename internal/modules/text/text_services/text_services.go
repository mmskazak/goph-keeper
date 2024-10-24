package text_services

import (
	"context"
	"errors"
	"fmt"
	"gophKeeper/internal/modules/text/text_dto"
	"gophKeeper/pkg/crypto"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TextService struct {
	pool      *pgxpool.Pool
	cryptoKey [32]byte
}

func NewTextService(pool *pgxpool.Pool, enKey [32]byte) *TextService {
	return &TextService{
		pool:      pool,
		cryptoKey: enKey,
	}
}

func (svc *TextService) SaveText(ctx context.Context, dto text_dto.SaveTextDTO) error {
	sql := `INSERT INTO texts (user_id, title, description, text_content) VALUES ($1, $2, $3, $4)`

	// Шифруем текст
	encryptedText, err := crypto.Encrypt(svc.cryptoKey, []byte(dto.TextContent))
	if err != nil {
		return fmt.Errorf("error while encrypting text: %w", err)
	}

	_, err = svc.pool.Exec(ctx, sql, dto.UserID, dto.Title, dto.Description, encryptedText)
	if err != nil {
		return fmt.Errorf("error while inserting text to DB: %w", err)
	}
	return nil
}

func (svc *TextService) GetText(ctx context.Context, dto text_dto.GetTextDTO) (text_dto.TextDTO, error) {
	sql := `SELECT title, description, text_content FROM texts WHERE id = $1 AND user_id = $2`

	var title, description string
	var encryptedText []byte

	err := svc.pool.QueryRow(ctx, sql, dto.TextID, dto.UserID).Scan(&title, &description, &encryptedText)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return text_dto.TextDTO{}, fmt.Errorf("record not found: %w", err)
		}
		return text_dto.TextDTO{}, fmt.Errorf("error while querying text from DB: %w", err)
	}

	// Расшифровка текста
	decryptedText, err := crypto.Decrypt(svc.cryptoKey, string(encryptedText))
	if err != nil {
		return text_dto.TextDTO{}, fmt.Errorf("error while decrypting text: %w", err)
	}

	return text_dto.TextDTO{
		Title:       title,
		Description: description,
		TextContent: string(decryptedText),
	}, nil
}

func (svc *TextService) DeleteText(ctx context.Context, dto text_dto.DeleteTextDTO) error {
	sql := `DELETE FROM texts WHERE id = $1 AND user_id = $2`

	commandTag, err := svc.pool.Exec(ctx, sql, dto.TextID, dto.UserID)
	if err != nil {
		return fmt.Errorf("error while deleting text from DB: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no rows affected, text not found")
	}

	return nil
}

func (svc *TextService) GetAllTexts(ctx context.Context, dto text_dto.AllTextDTO) ([]text_dto.TextDTO, error) {
	sql := `SELECT id, title, description, text_content FROM texts WHERE user_id = $1`

	rows, err := svc.pool.Query(ctx, sql, dto.UserID)
	if err != nil {
		return nil, fmt.Errorf("error while querying texts from DB: %w", err)
	}
	defer rows.Close()

	var texts []text_dto.TextDTO

	for rows.Next() {
		var text text_dto.TextDTO
		var encryptedText []byte

		err := rows.Scan(&text.ID, &text.Title, &text.Description, &encryptedText)
		if err != nil {
			return nil, fmt.Errorf("error while scanning text: %w", err)
		}

		// Расшифровка текста
		decryptedText, err := crypto.Decrypt(svc.cryptoKey, string(encryptedText))
		if err != nil {
			return nil, fmt.Errorf("error while decrypting text: %w", err)
		}
		text.TextContent = string(decryptedText)

		texts = append(texts, text)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error while iterating rows: %w", rows.Err())
	}

	return texts, nil
}

func (svc *TextService) UpdateText(ctx context.Context, dto text_dto.UpdateTextDTO) error {
	sql := `UPDATE texts SET title = $1, description = $2, text_content = $3 WHERE id = $4 AND user_id = $5`

	// Шифруем обновленный текст
	encryptedText, err := crypto.Encrypt(svc.cryptoKey, []byte(dto.TextContent))
	if err != nil {
		return fmt.Errorf("error while encrypting text: %w", err)
	}

	commandTag, err := svc.pool.Exec(ctx, sql, dto.Title, dto.Description, encryptedText, dto.ID, dto.UserID)
	if err != nil {
		return fmt.Errorf("error while updating text in DB: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("nothing to update")
	}

	return nil
}
