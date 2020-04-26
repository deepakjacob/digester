package db

import (
	"context"
	"fmt"

	"github.com/deepakjacob/digester/domain"
)

type RegistrationDB interface {
	RegisterFile(context.Context, *domain.Registration) (domain.FileIDType, error)
}

type RegistrationDBImpl struct {
	*PgConn
}

func (r *RegistrationDBImpl) RegisterFile(ctx context.Context, o *domain.Registration) (domain.FileIDType, error) {
	rows, _ := r.db.Query(ctx,
		"select file_id, file_name from file_register")
	for rows.Next() {
		var id int32
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return "", err
		}
		fmt.Printf("%d. %s\n", id, name)
	}
	return "", rows.Err()
}
