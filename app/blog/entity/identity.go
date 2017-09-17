package entity

import (
	"github.com/duhruh/tackle/domain"
	"github.com/satori/go.uuid"
)

func NextIdentity() domain.Identity {
	return domain.NewIdentity(uuid.NewV4().String())
}
