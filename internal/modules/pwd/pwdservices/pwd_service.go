package pwdservices

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gophKeeper/internal/logger"
	"gophKeeper/internal/modules/pwd/pwddto"
	"gophKeeper/internal/modules/pwd/valueobj"
	"gophKeeper/internal/storage"
	"gophKeeper/pkg/crypto"

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
	logger.Log.Infoln("start save password")
	sql := `INSERT INTO passwords (user_id, title, description, credentials) VALUES ($1, $2, $3, $4)`

	marshaledCredentials, err := json.Marshal(dto.Credentials)
	if err != nil {
		logger.Log.Debugln("marshal password failed")
		return fmt.Errorf("error marshalling credentials: %w", err)

	}
	_, err = pwd.pool.Exec(ctx, sql, dto.UserID, dto.Title, dto.Description, marshaledCredentials)
	if err != nil {
		logger.Log.Debugln("error save password from pwd service: %v", err)
		return fmt.Errorf("error save password from pwd service: %w", err)
	}

	return nil
}

func (pwd *PwdService) DeletePassword(ctx context.Context, dto *pwddto.DeletePwdDTO) error {
	sql := `DELETE FROM passwords WHERE user_id = $1 AND id = $2;`
	_, err := pwd.pool.Exec(ctx, sql, dto.UserID, dto.PwdID)
	if err != nil {
		return fmt.Errorf("error delete password from pwd service: %w", err)
	}
	return nil
}

func (pwd *PwdService) GetPassword(ctx context.Context, dto *pwddto.GetPwdDTO) (pwddto.CredentialsDTO, error) {
	sql := `SELECT credentials FROM passwords WHERE user_id = $1 AND id = $2;`
	row := pwd.pool.QueryRow(ctx, sql, dto.UserID, dto.PwdID)
	fmt.Println(row)
	fmt.Println("QueryRow")
	var credentialsData []byte
	err := row.Scan(&credentialsData)
	fmt.Println("Scan")
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// Обработка случая, когда запись не найдена
			return pwddto.CredentialsDTO{}, fmt.Errorf("password not found for user: %w", err)
		}
		// Обработка других ошибок
		return pwddto.CredentialsDTO{}, fmt.Errorf("error scanning password from pwd service: %w", err)
	}

	var credentials pwddto.CredentialsDTO
	if err := json.Unmarshal(credentialsData, &credentials); err != nil {
		return pwddto.CredentialsDTO{}, fmt.Errorf("error unmarshalling credentials: %w", err)
	}

	return credentials, nil
}

func (pwd *PwdService) GetAllPasswords(ctx context.Context, dto *pwddto.AllPwdDTO) ([]pwddto.PwdDTO, error) {
	sql := `SELECT id,resource, login, password FROM passwords WHERE user_id = $1`
	rows, err := pwd.pool.Query(ctx, sql, dto.UserID)
	if err != nil {
		return []pwddto.PwdDTO{}, fmt.Errorf("error query get all passwords: %w", err)
	}

	var listPasswords []pwddto.PwdDTO

	for rows.Next() {
		var id string
		var title string
		var description string
		var credentialsData []byte
		err := rows.Scan(&id, &title, &description, &credentialsData)
		if err != nil {
			return []pwddto.PwdDTO{}, fmt.Errorf("error scan get all passwords from pwd service: %w", err)
		}

		// Расшифровываем данные
		decryptedCredentials, err := crypto.Decrypt(pwd.cryptoKey, string(credentialsData))
		if err != nil {
			return []pwddto.PwdDTO{}, fmt.Errorf("error decrypt for GetAllPasswords %w", err)
		}

		var credentials valueobj.Credentials
		if err := json.Unmarshal(decryptedCredentials, &credentials); err != nil {
			return []pwddto.PwdDTO{}, fmt.Errorf("error unmarshalling credentials: %w", err)
		}

		listPasswords = append(listPasswords, pwddto.PwdDTO{
			ID:          id,
			Title:       title,
			Description: description,
			Credentials: credentials,
		})
	}

	return listPasswords, nil
}

func (pwd *PwdService) UpdatePassword(ctx context.Context, dto *pwddto.UpdatePwdDTO) error {
	marshaledCredentials, err := json.Marshal(dto.Credentials)
	if err != nil {
		return fmt.Errorf("error marshalling credentials: %w", err)
	}

	// Шифруем данные
	encryptedCredentials, err := crypto.Encrypt(pwd.cryptoKey, marshaledCredentials)
	if err != nil {
		return fmt.Errorf("error while encrypting credentials: %w", err)
	}

	sql := `UPDATE passwords SET title = $2, descriotion = $3, credentials = $4 WHERE id = $5 AND user_id = $6`
	result, err := pwd.pool.Exec(ctx, sql, dto.Title, dto.Description, encryptedCredentials, dto.ID, dto.UserID)
	if err != nil {
		return fmt.Errorf("error updating password: %w", err)
	}
	if result.RowsAffected() == 0 {
		return errors.New("updated record not found") // Возвращаем ошибку, если ни одна строка не была обновлена
	}
	return nil
}
