package appcontext

import "github.com/google/uuid"

func generateID() string {
	id, _ := uuid.NewUUID()
	return id.String()
}
