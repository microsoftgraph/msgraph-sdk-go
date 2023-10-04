package contacts

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetByIdsResponse 
// Deprecated: This class is obsolete. Use getByIdsPostResponse instead.
type GetByIdsResponse struct {
    GetByIdsPostResponse
}
// NewGetByIdsResponse instantiates a new GetByIdsResponse and sets the default values.
func NewGetByIdsResponse()(*GetByIdsResponse) {
    m := &GetByIdsResponse{
        GetByIdsPostResponse: *NewGetByIdsPostResponse(),
    }
    return m
}
// CreateGetByIdsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetByIdsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetByIdsResponse(), nil
}
// GetByIdsResponseable 
// Deprecated: This class is obsolete. Use getByIdsPostResponse instead.
type GetByIdsResponseable interface {
    GetByIdsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
