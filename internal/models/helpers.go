package models

import (
	"context"
	"errors"

	"github.com/Chadi-Mangle/templ-hmr-setup/package/utils"
	"github.com/jackc/pgx/v5"
)

func AddUserInBungalow(ctx context.Context, db *pgx.Conn, queries *Queries, userID int32, bungalowID int32) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	qtx := queries.WithTx(tx)

	b, err := qtx.GetBungalowByID(ctx, bungalowID)
	if err != nil {
		return err
	}

	currentReservations, err := qtx.GetBungalowNbReservations(ctx,
		utils.ConvertInt32ToNullInt32(bungalowID))
	if err != nil {
		return err
	}

	if currentReservations < int64(b.Capacity) {
		_, err := qtx.SetUserReservations(ctx, SetUserReservationsParams{
			BungalowID: utils.ConvertInt32ToNullInt32(bungalowID),
			ID:         userID,
		})

		if err != nil {
			return err
		}

		return tx.Commit(ctx)
	}
	return errors.New("bungalow is already at full capacity")
}
