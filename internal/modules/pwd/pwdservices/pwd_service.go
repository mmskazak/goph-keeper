package pwdservices

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"goph-keeper/internal/modules/pwd/pwddto"
	"goph-keeper/internal/modules/pwd/valueobj"
	"goph-keeper/internal/storage"
	"goph-keeper/pkg/crypto"

	"go.uber.org/zap"

	"github.com/jackc/pgx/v5"
)

type PwdService struct {
	zapLogger *zap.SugaredLogger
	pool      storage.Database
	cryptoKey [32]byte
}

func NewPwdService(pool storage.Database, enKey [32]byte, zapLogger *zap.SugaredLogger) *PwdService {
	return &PwdService{
		pool:      pool,
		cryptoKey: enKey,
		zapLogger: zapLogger,
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
		pwd.zapLogger.Errorf("error marshalling credentials: %v", err)
		return fmt.Errorf("error marshalling credentials: %w", err)
	}
	_, err = pwd.pool.Exec(ctx, sql, dto.UserID, dto.Title, marshaledCredentials)
	if err != nil {
		pwd.zapLogger.Errorf("error save password from pwd service: %v", err)
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
			pwd.zapLogger.Errorf("no password found to delete for user %d", dto.UserID)
			return fmt.Errorf("no password found to delete for user %d", dto.UserID)
		}
		return fmt.Errorf("error deleting password from pwd service: %w", err)
	}

	return nil
}

func (pwd *PwdService) GetPassword(ctx context.Context, dto *pwddto.GetPwdDTO) (pwddto.ResponsePwdDTO, error) {
	sql := `SELECT id, title, credentials FROM passwords WHERE user_id = $1 AND id = $2;`
	row := pwd.pool.QueryRow(ctx, sql, dto.UserID, dto.PwdID)

	var id, title string
	var credentialsData []byte
	err := row.Scan(&id, &title, &credentialsData)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// Обработка случая, когда запись не найдена
			pwd.zapLogger.Errorf("password not found for user %v: %v", dto.UserID, err)
			return pwddto.ResponsePwdDTO{}, fmt.Errorf("password not found for user: %w", err)
		}
		// Обработка других ошибок
		pwd.zapLogger.Errorf("error scanning password from pwd service: %v", err)
		return pwddto.ResponsePwdDTO{}, fmt.Errorf("error scanning password from pwd service: %w", err)
	}

	var credentials valueobj.Credentials
	if err := json.Unmarshal(credentialsData, &credentials); err != nil {
		pwd.zapLogger.Errorf("error unmarshalling credentials: %v", err)
		return pwddto.ResponsePwdDTO{}, fmt.Errorf("error unmarshalling credentials: %w", err)
	}

	// Расшифровка текста
	decryptedPassword, err := crypto.Decrypt(pwd.cryptoKey, credentials.Password)
	if err != nil {
		pwd.zapLogger.Errorf("error while decrypting text: %v", err)
		return pwddto.ResponsePwdDTO{}, fmt.Errorf("error while decrypting text: %w", err)
	}
	credentials.Password = string(decryptedPassword)

	// Заполнение всех необходимых данных для ответа
	var responsePwdDTO pwddto.ResponsePwdDTO
	responsePwdDTO.PwdID = id
	responsePwdDTO.Title = title
	responsePwdDTO.Credentials = credentials

	return responsePwdDTO, nil
}

func (pwd *PwdService) GetAllPasswords(ctx context.Context, dto *pwddto.AllPwdDTO) ([]pwddto.ResponsePwdDTO, error) {
	sql := `SELECT id,title, credentials FROM passwords WHERE user_id = $1`
	rows, err := pwd.pool.Query(ctx, sql, dto.UserID)
	if err != nil {
		pwd.zapLogger.Errorf("error query get all passwords: %v", err)
		return []pwddto.ResponsePwdDTO{}, fmt.Errorf("error query get all passwords: %w", err)
	}

	var listPasswords []pwddto.ResponsePwdDTO

	for rows.Next() {
		var id string
		var title string
		var credentialsData []byte
		err = rows.Scan(&id, &title, &credentialsData)
		if err != nil {
			pwd.zapLogger.Errorf("error scan get all passwords from pwd service: %v", err)
			return []pwddto.ResponsePwdDTO{}, fmt.Errorf("error scan get all passwords from pwd service: %w", err)
		}

		var credentials valueobj.Credentials
		if err = json.Unmarshal(credentialsData, &credentials); err != nil {
			pwd.zapLogger.Errorf("error unmarshalling credentials: %v", err)
			return []pwddto.ResponsePwdDTO{}, fmt.Errorf("error unmarshalling credentials: %w", err)
		}

		// Расшифровка текста
		decryptedPassword, err := crypto.Decrypt(pwd.cryptoKey, credentials.Password)
		if err != nil {
			pwd.zapLogger.Errorf("error while decrypting text: %v", err)
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
	// Шифруем
	encryptedPassword, err := crypto.Encrypt(pwd.cryptoKey, []byte(dto.Credentials.Password))
	if err != nil {
		pwd.zapLogger.Errorf("error while encrypting text: %v", err)
		return fmt.Errorf("error while encrypting text: %w", err)
	}
	dto.Credentials.Password = encryptedPassword

	sql := `UPDATE passwords SET title = $1, credentials = $2 WHERE id = $3 AND user_id = $4`
	result, err := pwd.pool.Exec(ctx, sql, dto.Title, dto.Credentials, dto.PwdID, dto.UserID)
	if err != nil {
		pwd.zapLogger.Errorf("error updating password: %v", err)
		return fmt.Errorf("error updating password: %w", err)
	}
	if result.RowsAffected() == 0 {
		pwd.zapLogger.Errorln("updated record not found")
		return errors.New("updated record not found") // Возвращаем ошибку, если ни одна строка не была обновлена
	}
	return nil
}
