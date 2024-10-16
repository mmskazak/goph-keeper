package text_services

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"gophKeeper/internal/modules/pwd/pwd_dto"
)

type TextService struct {
	pool *pgxpool.Pool
}

func NewPwdService(pool *pgxpool.Pool) *TextService {
	return &TextService{pool: pool}
}

func (pwd *TextService) SavePassword(ctx context.Context, dto pwd_dto.SavePwdDTO) error {
	sql := `INSERT INTO passwords (user_id, resource, login, password) VALUES ($1, $2, $3, $4)`
	_, err := pwd.pool.Exec(ctx, sql,
		dto.UserID,
		dto.Resource,
		dto.Login,
		dto.Password)
	if err != nil {
		return fmt.Errorf("error save password from pwd service: %w", err)
	}

	return nil
}

func (pwd *TextService) DeletePassword(ctx context.Context, dto pwd_dto.DeletePwdDTO) error {
	sql := `DELETE FROM passwords WHERE user_id = $1 AND resource = $2;`
	_, err := pwd.pool.Exec(ctx, sql, dto.UserID, dto.Resource)
	if err != nil {
		return fmt.Errorf("error save password from pwd service: %w", err)
	}
	return nil
}

func (pwd *TextService) GetPassword(ctx context.Context, dto pwd_dto.GetPwdDTO) (string, error) {
	sql := `SELECT password FROM passwords WHERE user_id = $1 AND login = $2;`
	row := pwd.pool.QueryRow(ctx, sql, dto.UserID, dto.Login)
	var password string
	err := row.Scan(&password)
	if err != nil {
		return "", fmt.Errorf("error scan password from pwd sercvice")
	}
	return password, nil
}

func (pwd *TextService) GetAllPasswords(ctx context.Context, dto pwd_dto.AllPwdDTO) ([]InfoByText, error) {
	sql := `SELECT resource, login, password FROM passwords WHERE user_id = $1`
	rows, err := pwd.pool.Query(ctx, sql, dto.UserID)
	if err != nil {
		return []InfoByText{}, fmt.Errorf("error query get all passwords: %w", err)
	}

	var listPasswords []InfoByText

	for rows.Next() {
		var resource string
		var login string
		var password string
		err := rows.Scan(&resource, &login, &password)
		if err != nil {
			return []InfoByText{}, fmt.Errorf("error scan get all passwords from pwd service: %w", err)
		}

		listPasswords = append(listPasswords, InfoByText{
			Resource: resource,
			Login:    login,
			Password: password,
		})
	}

	return listPasswords, nil
}
