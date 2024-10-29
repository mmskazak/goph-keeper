package cardservices

import (
	"context"
	"errors"
	"fmt"
	"gophKeeper/internal/logger"
	"gophKeeper/internal/modules/card/carddto"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CardService struct {
	pool *pgxpool.Pool
}

func NewCardService(pool *pgxpool.Pool) *CardService {
	return &CardService{pool: pool}
}

func (cs *CardService) SaveCard(ctx context.Context, dto *carddto.SaveCardDTO) error {
	sql := `INSERT INTO cards (user_id, title, description, number, pincode, cvv, expire) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := cs.pool.Exec(ctx, sql, dto.UserID, dto.Title,
		dto.Description, dto.Number, dto.PinCode,
		dto.CVV, dto.Expire)
	if err != nil {
		logger.Log.Errorf("Error saving card: %v", err)
		return fmt.Errorf("error saving card: %w", err)
	}
	return nil
}

func (cs *CardService) GetCard(ctx context.Context, dto *carddto.GetCardDTO) (carddto.SaveCardDTO, error) {
	var card carddto.SaveCardDTO
	sql := `SELECT user_id, title, description, number, pincode, cvv, expire 
		FROM cards WHERE id = $1 AND user_id = $2`
	row := cs.pool.QueryRow(ctx, sql, dto.CardID, dto.UserID)
	err := row.Scan(&card.UserID, &card.Title, &card.Description, &card.Number, &card.PinCode, &card.CVV, &card.Expire)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Log.Errorf("Card not found: %v", err)
			return card, fmt.Errorf("card not found: %w", err)
		}
		logger.Log.Errorf("Error scaning card: %v", err)
		return card, fmt.Errorf("error scaning card: %w", err)
	}
	return card, nil
}

func (cs *CardService) UpdateCard(ctx context.Context, dto *carddto.UpdateCardDTO) error {
	sql := `UPDATE cards SET title = $1, description = $2, number = $3, pincode = $4, cvv = $5, expire = $6 
             WHERE id = $7`
	_, err := cs.pool.Exec(ctx, sql, dto.Title,
		dto.Description, dto.Number,
		dto.PinCode, dto.CVV,
		dto.Expire, dto.CardID)
	if err != nil {
		logger.Log.Errorf("Error updating card: %v", err)
		return fmt.Errorf("error updating card: %w", err)
	}
	return nil
}

func (cs *CardService) DeleteCard(ctx context.Context, dto *carddto.DeleteCardDTO) error {
	sql := `DELETE FROM cards WHERE id = $1 AND user_id = (SELECT user_id FROM cards WHERE id = $1)`
	_, err := cs.pool.Exec(ctx, sql, dto.CardID)
	if err != nil {
		logger.Log.Errorf("Error deleting card: %v", err)
		return fmt.Errorf("error deleting card: %w", err)
	}
	return nil
}

func (cs *CardService) GetAllCards(ctx context.Context, dto *carddto.GetAllCardsDTO) ([]carddto.SaveCardDTO, error) {
	var cards []carddto.SaveCardDTO
	sql := `SELECT id, title, description, number, pincode, cvv, expire FROM cards WHERE user_id = $1`
	rows, err := cs.pool.Query(ctx, sql, dto.UserID)
	if err != nil {
		logger.Log.Errorf("Error getting cards: %v", err)
		return nil, fmt.Errorf("error getting cards: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var card carddto.SaveCardDTO
		err := rows.Scan(&card.UserID,
			&card.Title, &card.Description,
			&card.Number, &card.PinCode,
			&card.CVV, &card.Expire)
		if err != nil {
			logger.Log.Errorf("Error scanning card: %v", err)
			return nil, fmt.Errorf("error scanning card: %w", err)
		}
		cards = append(cards, card)
	}

	if err = rows.Err(); err != nil {
		logger.Log.Errorf("Error iterating over cards: %v", err)
		return nil, fmt.Errorf("error iterating over cards: %w", err)
	}

	return cards, nil
}
