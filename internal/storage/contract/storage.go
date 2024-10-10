package contract

import "context"

type IStorage interface {
	Registration(ctx context.Context, login, password string) error
	Login(ctx context.Context, login, password string) error
	Logout(ctx context.Context) error
	Close()
}
