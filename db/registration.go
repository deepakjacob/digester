package db

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/deepakjacob/digester/domain"
)

type RegistrationDB interface {
	RegisterFile(context.Context, *domain.Registration) (domain.FileIDType, error)
}

type RegistrationDBImpl struct {
	*PgConn
}

func (r *RegistrationDBImpl) RegisterFile(ctx context.Context, o *domain.Registration) (domain.FileIDType, error) {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		log.Error().Err(err).Msg("error acquire db connection from pool")
		return "", err
	}
	defer conn.Release()
	_, err = conn.Exec(ctx,
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
