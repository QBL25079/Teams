package domain

import (
	"fmt"
	"time"

	"github.com/QBL25079/teams/internal/core/errors"
)

type Team struct {
	ID        int
	Name      string
	ParentID  *int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TeamPatch struct {
	Name     *string
	ParentID *int
}

func (t *Team) ApplyPatch(patch TeamPatch) error {
	tmp := *t

	if patch.Name != nil {
		if len([]rune(*patch.Name)) < 1 || len([]rune(*patch.Name)) > 255 {
			return fmt.Errorf("invalid name length")
		}
		tmp.Name = *patch.Name
	}

	if patch.ParentID != nil {
		tmp.ParentID = patch.ParentID

		if tmp.ParentID != nil && *tmp.ParentID == tmp.ID {
			return fmt.Errorf("team cannot be parent of itself")
		}
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate updated team: %w", err)
	}

	*t = tmp
	return nil
}

func NewTeam(ID int, Name string, ParentID *int, CreatedAt, UpdateAt time.Time) Team {
	return Team{ID: ID, Name: Name, ParentID: ParentID, CreatedAt: CreatedAt, UpdatedAt: UpdateAt}
}

func NewTeamPatch(Name *string, ParentID *int) TeamPatch {
	return TeamPatch{Name: Name, ParentID: ParentID}
}

func NewTeamUninitialized(name string, parentID *int) Team {
	now := time.Now()

	return NewTeam(
		UninitializedID,
		name,
		parentID,
		now,
		now,
	)
}

func (t *Team) Validate() error {
	nameLen := len([]rune(t.Name))
	if nameLen < 1 || nameLen > 15 {
		return fmt.Errorf("invalid name length %d: %w", nameLen, errors.ErrInvalidArgument)
	}

	if t.ParentID != nil {
		if *t.ParentID <= 0 {
			return fmt.Errorf("invalid parent_id %d: %w", *t.ParentID, errors.ErrInvalidArgument)
		}
	}

	if t.UpdatedAt.Before(t.CreatedAt) {
		return fmt.Errorf("updated_at cannot be before created_at: %w", errors.ErrInvalidArgument)
	}

	return nil
}
