package waitlist

import (
	"context"
	"github.com/mixarchitecture/i18np"
)

type Repository interface {
	Join(ctx context.Context, entity *Entity) *i18np.Error
	Leave(ctx context.Context, token string) (*Entity, *i18np.Error)
}
