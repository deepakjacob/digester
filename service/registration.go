package service

import (
	"context"

	"github.com/deepakjacob/digester/db"
	"github.com/deepakjacob/digester/domain"
)

type RegistrationService interface {
	//RegisterFile registers an incoming file in the system
	RegisterFile(context.Context, *domain.Registration) (*domain.RegistrationStatus, error)
}

type RegistrationServiceImpl struct {
	db.RegistrationDB
}

func (r *RegistrationServiceImpl) RegisterFile(ctx context.Context, o *domain.Registration) (*domain.RegistrationStatus, error) {
	fileID, err := r.RegistrationDB.RegisterFile(ctx, o)
	if err != nil {
		return &domain.RegistrationStatus{
			FileID:     "",
			StatusCD:   "FILE_REGISTRATION_DB_ERR",
			StatusDesc: "File registration failed in DB",
		}, err
	}
	return &domain.RegistrationStatus{
		FileID:     fileID,
		StatusCD:   "FILE_REGISTRATION_DB_SUCCESS",
		StatusDesc: "File registered successfully in DB",
	}, nil
}
