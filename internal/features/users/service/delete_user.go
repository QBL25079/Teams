package user_service

import (
	"context"
	"fmt"
)

func (r *UserService) DeleteUser(ctx context.Context, id int) error {
	if err := r.userRepository.DeleteUser(ctx, id); err != nil {
		return fmt.Errorf("delete user: %w", err)
	}
	return nil
}
