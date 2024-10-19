package pwd_services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"gophKeeper/internal/modules/pwd/pwd_services/dto/request"
	"gophKeeper/internal/modules/pwd/pwd_services/dto/response"
	"gophKeeper/internal/modules/pwd/pwd_services/value_obj"
	"gophKeeper/pkg/crypto"
)

type PwdService struct {
	pool      *pgxpool.Pool
	cryptoKey [32]byte
}

func NewPwdService(pool *pgxpool.Pool, enKey [32]byte) *PwdService {
	return &PwdService{
		pool:      pool,
		cryptoKey: enKey,
	}
}

func (pwd *PwdService) SavePassword(ctx context.Context, dto request.SavePwdDTO) error {
	sql := `INSERT INTO passwords (user_id, title, description, credentials) VALUES ($1, $2, $3, $4)`

	marshaledCredentials, err := json.Marshal(dto.Credentials)
	if err != nil {
		return fmt.Errorf("error marshalling credentials: %w", err)
	}

	// Шифруем данные
	encryptedCredentials, err := crypto.Encrypt(pwd.cryptoKey, marshaledCredentials)
	if err != nil {
		return fmt.Errorf("error while encrypting credentials: %w", err)
	}

	_, err = pwd.pool.Exec(ctx, sql, dto.UserID, dto.Title, dto.Description, encryptedCredentials)
	if err != nil {
		return fmt.Errorf("error save password from pwd service: %w", err)
	}

	return nil
}

func (pwd *PwdService) DeletePassword(ctx context.Context, dto request.DeletePwdDTO) error {
	sql := `DELETE FROM passwords WHERE user_id = $1 AND id = $2;`
	_, err := pwd.pool.Exec(ctx, sql, dto.UserID, dto.PwdID)
	if err != nil {
		return fmt.Errorf("error delete password from pwd service: %w", err)
	}
	return nil
}

func (pwd *PwdService) GetPassword(ctx context.Context, dto request.GetPwdDTO) (response.CredentialsDTO, error) {
	sql := `SELECT credentials FROM passwords WHERE user_id = $1 AND id = $2;`
	row := pwd.pool.QueryRow(ctx, sql, dto.UserID, dto.PwdID)

	var credentialsData []byte
	err := row.Scan(&credentialsData)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// Обработка случая, когда запись не найдена
			return response.CredentialsDTO{}, fmt.Errorf("password not found for user: %w", err)
		}
		// Обработка других ошибок
		return response.CredentialsDTO{}, fmt.Errorf("error scanning password from pwd service: %w", err)
	}

	//Расшифровываем данные
	decryptedCredentials, err := crypto.Decrypt(pwd.cryptoKey, string(credentialsData))
	if err != nil {
		return response.CredentialsDTO{}, err
	}

	var credentials response.CredentialsDTO
	if err := json.Unmarshal(decryptedCredentials, &credentials); err != nil {
		return response.CredentialsDTO{}, fmt.Errorf("error unmarshalling credentials: %w", err)
	}

	return credentials, nil
}

func (pwd *PwdService) GetAllPasswords(ctx context.Context, dto request.AllPwdDTO) ([]response.PwdDTO, error) {
	sql := `SELECT id,resource, login, password FROM passwords WHERE user_id = $1`
	rows, err := pwd.pool.Query(ctx, sql, dto.UserID)
	if err != nil {
		return []response.PwdDTO{}, fmt.Errorf("error query get all passwords: %w", err)
	}

	var listPasswords []response.PwdDTO

	for rows.Next() {
		var id string
		var title string
		var description string
		var credentialsData []byte
		err := rows.Scan(&id, &title, &description, &credentialsData)
		if err != nil {
			return []response.PwdDTO{}, fmt.Errorf("error scan get all passwords from pwd service: %w", err)
		}

		//Расшифровываем данные
		decryptedCredentials, err := crypto.Decrypt(pwd.cryptoKey, string(credentialsData))
		if err != nil {
			return []response.PwdDTO{}, err
		}

		var credentials value_obj.Credentials
		if err := json.Unmarshal(decryptedCredentials, &credentials); err != nil {
			return []response.PwdDTO{}, fmt.Errorf("error unmarshalling credentials: %w", err)
		}

		listPasswords = append(listPasswords, response.PwdDTO{
			ID:          id,
			Title:       title,
			Description: description,
			Credentials: credentials,
		})
	}

	return listPasswords, nil
}

func (pwd *PwdService) UpdatePassword(ctx context.Context, dto request.UpdatePwdDTO) error {
	marshaledCredentials, err := json.Marshal(dto.Credentials)
	if err != nil {
		return fmt.Errorf("error marshalling credentials: %w", err)
	}

	// Шифруем данные
	encryptedCredentials, err := crypto.Encrypt(pwd.cryptoKey, marshaledCredentials)

	sql := `UPDATE passwords SET title = $2, descriotion = $3, credentials = $4 WHERE id = $5 AND user_id = $6`
	_, err = pwd.pool.Exec(ctx, sql, dto.Title, dto.Description, encryptedCredentials, dto.ID, dto.UserID)
	if err != nil {
		return fmt.Errorf("error updating password: %w", err)
	}
	return nil
}
