package quote

import "testeFreteRapido/internal/domain/interfaces"

type Input struct {
	RecipientZip int32
	Volumes      []interfaces.Volume
}
