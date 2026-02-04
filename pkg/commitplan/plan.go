package commitplan

import (
	"github.com/MediStatTech/commitplan"
	"github.com/MediStatTech/commitplan/drivers/postgres"
)

type StoreKey int

const (
	StoreStaff StoreKey = iota
)

type StaffPlan struct {
	*commitplan.Plan
	staffHandle postgres.HandlePostgres
}

func NewPlan() *StaffPlan {
	return &StaffPlan{
		Plan:        &commitplan.Plan{},
		staffHandle: postgres.HandleByKey(int(StoreStaff)),
	}
}

// AddMuts adds mutations to the staff store
func (p *StaffPlan) AddMuts(muts ...*postgres.Mutation) {
	p.staffHandle.Add(p.Plan, muts...)
}

// AddMut adds a single mutation to the staff store
func (p *StaffPlan) AddMut(mut *postgres.Mutation) {
	p.staffHandle.Add(p.Plan, mut)
}
