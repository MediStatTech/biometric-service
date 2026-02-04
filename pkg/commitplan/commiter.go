package commitplan

import (
	"github.com/MediStatTech/commitplan"
	"github.com/MediStatTech/commitplan/drivers/postgres"
)

// NewCommitter creates the commitplan facade with registered stores
// Accepts specific dependencies to avoid import cycles
func NewCommitter(staffDB postgres.DB) *commitplan.Facade {
	staffStore := postgres.NewPostgresStore(int(StoreStaff), staffDB)

	return commitplan.NewCommitter(
		commitplan.WithStoreTyped(int(StoreStaff), staffStore),
	)
}
