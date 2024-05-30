package directoryobjects

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use GetbyidsGetByIdsPostResponseable instead.
type GetbyidsGetByIdsResponse struct {
    GetbyidsGetByIdsPostResponse
}
// NewGetbyidsGetByIdsResponse instantiates a new GetbyidsGetByIdsResponse and sets the default values.
func NewGetbyidsGetByIdsResponse()(*GetbyidsGetByIdsResponse) {
    m := &GetbyidsGetByIdsResponse{
        GetbyidsGetByIdsPostResponse: *NewGetbyidsGetByIdsPostResponse(),
    }
    return m
}
// CreateGetbyidsGetByIdsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetbyidsGetByIdsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetbyidsGetByIdsResponse(), nil
}
// Deprecated: This class is obsolete. Use GetbyidsGetByIdsPostResponseable instead.
type GetbyidsGetByIdsResponseable interface {
    GetbyidsGetByIdsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
