package db

import (
	"context"

	"github.com/deepakjacob/digester/domain"
)

type RegistrationDB interface {
	RegisterFile(context.Context, *domain.Registration) (domain.FileIDType, error)
}

type RegistrationDBImpl struct {
	*PgConn
}

func (r *RegistrationDBImpl) RegisterFile(ctx context.Context, o *domain.Registration) (domain.FileIDType, error) {
	_, err := r.db.Exec(ctx,
		"insert into FILE_REGISTER "+
			"(file_name, file_date, tower_id, location_id, postal_code, area_code)"+
			" values "+
			"($1, $2, $3, $4, $5, $6)",
		o.FileName, o.FileDate, o.TowerID, o.LocationID, o.PostalCode, o.AreaCode,
	)
	if err != nil {
		return "", err
	}
	return "1234", nil
}
