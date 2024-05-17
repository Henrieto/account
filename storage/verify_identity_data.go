package storage

import (
	"context"

	"github.com/henrieto/account/models/database/db"
)

type VerifyIdentityData struct {
	Querier db.Querier
}

func NewVerifyIdentityData(querier db.Querier) *VerifyIdentityData {
	return &VerifyIdentityData{Querier: querier}
}

// (context , identity verification data) (db identity verification data , error)
func (strg *VerifyIdentityData) Create(ctx context.Context, verify_identity_data *db.VerifyIdentityData) (*db.VerifyIdentityData, error) {
	params := db.CreateVerifyIdentityDataParams{
		RandomString:        verify_identity_data.RandomString,
		IdentificationType:  verify_identity_data.IdentificationType,
		IdentificationValue: verify_identity_data.IdentificationValue,
		Otp:                 verify_identity_data.Otp,
		Expiry:              verify_identity_data.Expiry,
		OperationType:       verify_identity_data.OperationType,
	}
	new_verify_identity_data, err := strg.Querier.CreateVerifyIdentityData(ctx, params)
	if err != nil {
		return nil, err
	}
	return &new_verify_identity_data, nil
}

// (context) (number of verify identity datas , error )
func (strg *VerifyIdentityData) Count(ctx context.Context) (int64, error) {
	count, err := strg.Querier.CountVerifyIdentityDatas(ctx)
	if err != nil {
		return 0, err
	}
	return count, err
}

// (context , random string) (db identification data , error)
func (strg *VerifyIdentityData) Get(ctx context.Context, random_string string) (*db.VerifyIdentityData, error) {
	verify_identity_data, err := strg.Querier.GetVerifyIdentityData(ctx, random_string)
	if err != nil {
		return nil, err
	}
	return &verify_identity_data, nil
}

// (context , verification data id) (error)
func (strg *VerifyIdentityData) Delete(ctx context.Context, verify_identity_data_id int32) error {
	err := strg.Querier.DeleteVerifyIdentiyData(ctx, verify_identity_data_id)
	if err != nil {
		return err
	}
	return nil
}

// (context , order by , offset , limit)
func (strg *VerifyIdentityData) List(ctx context.Context, order_by string, offset uint, limit uint) ([]db.VerifyIdentityData, error) {
	params := db.PaginateVerifyIdentityDataParams{
		Column1: order_by,
		Offset:  int32(offset),
		Limit:   int32(limit),
	}
	verify_identity_datas, err := strg.Querier.PaginateVerifyIdentityData(ctx, params)
	if err != nil {
		return nil, err
	}
	return verify_identity_datas, nil
}
