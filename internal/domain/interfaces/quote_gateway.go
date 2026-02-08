package interfaces

import (
	"context"
	"testeFreteRapido/internal/domain/entity"
)

type QuoteGateway interface {
	Quote(
		ctx context.Context,
		recipientZip int32,
		volumes []entity.VolumeEntity,
	) (quoteEntity []entity.QuoteEntity, err error)
}
