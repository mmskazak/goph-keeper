package pwdservices

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"goph-keeper/internal/logger"
	"goph-keeper/internal/modules/pwd/pwddto"
	"goph-keeper/internal/modules/pwd/valueobj"
	"goph-keeper/internal/storage"
	"goph-keeper/pkg/crypto"

	"github.com/jackc/pgx/v5"
)

type PwdService struct {
	pool      storage.Database
	cryptoKey [32]byte
}

func NewPwdService(pool storage.Database, enKey [32]byte) *PwdService {
	return &PwdService{
		pool:      pool,
		cryptoKey: enKey,
	}
}

func (pwd *PwdService) SavePassword(ctx context.Context, dto *pwddto.SavePwdDTO) error {
	sql := `INSERT INTO passwords (user_id, title, credentials) VALUES ($1, $2, $3)`

	// Шифруем
	encryptedPassword, err := crypto.Encrypt(pwd.cryptoKey, []byte(dto.Credentials.Password))
	if err != nil {
		return fmt.Errorf("error while encrypting text: %w", err)
	}
	dto.Credentials.Password = encryptedPassword

	marshaledCredentials, err := json.Marshal(dto.Credentials)
	if err != nil {
		logger.Log.Errorf("error marshalling credentials: %v", err)
		return fmt.Errorf("error marshalling credentials: %w", err)
	}
	_, err = pwd.pool.Exec(ctx, sql, dto.UserID, dto.Title, marshaledCredentials)
	if err != nil {
		logger.Log.Errorf("error save password from pwd service: %v", err)
		return fmt.Errorf("error save password from pwd service: %w", err)
	}

	return nil
}

func (pwd *PwdService) DeletePassword(ctx context.Context, dto *pwddto.DeletePwdDTO) error {
	sql := `DELETE FROM passwords WHERE user_id = $1 AND id = $2 RETURNING id;`
	var deletedID int
	err := pwd.pool.QueryRow(ctx, sql, dto.UserID, dto.PwdID).Scan(&deletedID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Log.Errorf("no password found to delete for user %d", dto.UserID)
			return fmt.Errorf("no password found to delete for user %d", dto.UserID)
		}
		return fmt.Errorf("error deleting password from pwd service: %w", err)
	}

	return nil
}

func (pwd *PwdService) GetPassword(ctx context.Context, dto *pwddto.GetPwdDTO) (pwddto.ResponsePwdDTO, error) {
	sql := `SELECT id, title, credentials FROM passwords WHERE user_id = $1 AND id = $2;`
	row := pwd.pool.QueryRow(ctx, sql, dto.UserID, dto.PwdID)

	var id, title, description string
	var credentialsData []byte
	err := row.Scan(&id, &title, &description, &credentialsData)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// Обработка случая, когда запись не найдена
			logger.Log.Errorf("password not found for user %s: %v", dto.UserID, err)
			return pwddto.ResponsePwdDTO{}, fmt.Errorf("password not found for user: %w", err)
		}
		// Обработка других ошибок
		logger.Log.Errorf("error scanning password from pwd service: %v", err)
		return pwddto.ResponsePwdDTO{}, fmt.Errorf("error scanning password from pwd service: %w", err)
	}

	var responsePwdDTO pwddto.ResponsePwdDTO
	if err := json.Unmarshal(credentialsData, &responsePwdDTO); err != nil {
		logger.Log.Errorf("error unmarshalling credentials: %v", err)
		return pwddto.ResponsePwdDTO{}, fmt.Errorf("error unmarshalling credentials: %w", err)
	}

	// Заполнение всех необходимых данных для ответа
	responsePwdDTO.PwdID = id
	responsePwdDTO.Title = title

	return responsePwdDTO, nil
}

func (pwd *PwdService) GetAllPasswords(ctx context.Context, dto *pwddto.AllPwdDTO) ([]pwddto.ResponsePwdDTO, error) {
	sql := `SELECT id,title, description, credentials FROM passwords WHERE user_id = $1`
	rows, err := pwd.pool.Query(ctx, sql, dto.UserID)
	if err != nil {
		logger.Log.Errorf("error query get all passwords: %v", err)
		return []pwddto.ResponsePwdDTO{}, fmt.Errorf("error query get all passwords: %w", err)
	}

	var listPasswords []pwddto.ResponsePwdDTO

	for rows.Next() {
		var id string
		var title string
		var description string
		var credentialsData []byte
		err = rows.Scan(&id, &title, &description, &credentialsData)
		if err != nil {
			logger.Log.Errorf("error scan get all passwords from pwd service: %v", err)
			return []pwddto.ResponsePwdDTO{}, fmt.Errorf("error scan get all passwords from pwd service: %w", err)
		}

		var credentials valueobj.Credentials
		if err = json.Unmarshal(credentialsData, &credentials); err != nil {
			logger.Log.Errorf("error unmarshalling credentials: %v", err)
			return []pwddto.ResponsePwdDTO{}, fmt.Errorf("error unmarshalling credentials: %w", err)
		}

		// Расшифровка текста
		decryptedPassword, err := crypto.Decrypt(pwd.cryptoKey, credentials.Password)
		if err != nil {
			logger.Log.Errorf("error while decrypting text: %v", err)
			return []pwddto.ResponsePwdDTO{}, fmt.Errorf("error while decrypting text: %w", err)
		}
		credentials.Password = string(decryptedPassword)

		listPasswords = append(listPasswords, pwddto.ResponsePwdDTO{
			PwdID:       id,
			Title:       title,
			Credentials: credentials,
		})
	}

	return listPasswords, nil
}

func (pwd *PwdService) UpdatePassword(ctx context.Context, dto *pwddto.UpdatePwdDTO) error {
	marshaledCredentials, err := json.Marshal(dto.Credentials)
	if err != nil {
		logger.Log.Errorf("error marshalling credentials: %v", err)
		return fmt.Errorf("error marshalling credentials: %w", err)
	}

	// Шифруем
	encryptedPassword, err := crypto.Encrypt(pwd.cryptoKey, []byte(dto.Credentials.Password))
	if err != nil {
		logger.Log.Errorf("error while encrypting text: %v", err)
		return fmt.Errorf("error while encrypting text: %w", err)
	}
	dto.Credentials.Password = encryptedPassword

	sql := `UPDATE passwords SET title = $1, description = $2, credentials = $3 WHERE id = $4 AND user_id = $5`
	result, err := pwd.pool.Exec(ctx, sql, dto.Title, marshaledCredentials, dto.PwdID, dto.UserID)
	if err != nil {
		logger.Log.Errorf("error updating password: %v", err)
		return fmt.Errorf("error updating password: %w", err)
	}
	if result.RowsAffected() == 0 {
		logger.Log.Errorln("updated record not found")
		return errors.New("updated record not found") // Возвращаем ошибку, если ни одна строка не была обновлена
	}
	return nil
}
