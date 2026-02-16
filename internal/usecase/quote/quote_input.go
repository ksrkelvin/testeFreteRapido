package quote

import "testeFreteRapido/internal/domain/entity"

type Input struct {
	RecipientZip int32
	Volumes      []entity.VolumeEntity
}
