package deck

import (
	"github.com/kioku-project/kioku/pkg/group"
	"time"
)

type Deck struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	GroupID   uint
	Group     group.Group
}
