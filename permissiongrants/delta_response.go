package permissiongrants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type DeltaResponse struct {
    DeltaGetResponse
}
// NewDeltaResponse instantiates a new DeltaResponse and sets the default values.
func NewDeltaResponse()(*DeltaResponse) {
    m := &DeltaResponse{
        DeltaGetResponse: *NewDeltaGetResponse(),
    }
    return m
}
// CreateDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeltaResponse(), nil
}
// DeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type DeltaResponseable interface {
    DeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
