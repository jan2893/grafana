package cmsclient

import (
	"context"
	"math/rand"
)

// NewInMemoryClient returns an implementation of Client that returns canned responses
func NewInMemoryClient() Client {
	return &memoryClientImpl{}
}

type memoryClientImpl struct{}

func (c *memoryClientImpl) ValidateKey(ctx context.Context, input ValidateKeyInput) error {
	// return ErrMigrationNotDeleted
	return nil
}

func (c *memoryClientImpl) MigrateData(
	ctx context.Context,
	cm MigrateDataInput,
	request MigrateDataRequestDTO,
) (*MigrateDataResponseDTO, error) {
	//return nil, ErrMigrationNotDeleted

	result := MigrateDataResponseDTO{
		Items: make([]MigrateDataResponseItemDTO, len(request.Items)),
	}

	for i, v := range request.Items {
		result.Items[i] = MigrateDataResponseItemDTO{
			Type:   v.Type,
			RefID:  v.RefID,
			Status: "OK",
		}
	}

	// simulate flakiness on one random item
	i := rand.Intn(len(result.Items))
	failedItem := result.Items[i]
	failedItem.Status = "ERROR"
	failedItem.Error = "simulated random error"
	result.Items[i] = failedItem

	return &result, nil
}
