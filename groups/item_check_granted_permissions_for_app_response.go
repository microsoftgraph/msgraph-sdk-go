package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCheckGrantedPermissionsForAppResponse 
// Deprecated: This class is obsolete. Use checkGrantedPermissionsForAppPostResponse instead.
type ItemCheckGrantedPermissionsForAppResponse struct {
    ItemCheckGrantedPermissionsForAppPostResponse
}
// NewItemCheckGrantedPermissionsForAppResponse instantiates a new ItemCheckGrantedPermissionsForAppResponse and sets the default values.
func NewItemCheckGrantedPermissionsForAppResponse()(*ItemCheckGrantedPermissionsForAppResponse) {
    m := &ItemCheckGrantedPermissionsForAppResponse{
        ItemCheckGrantedPermissionsForAppPostResponse: *NewItemCheckGrantedPermissionsForAppPostResponse(),
    }
    return m
}
// CreateItemCheckGrantedPermissionsForAppResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCheckGrantedPermissionsForAppResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCheckGrantedPermissionsForAppResponse(), nil
}
// ItemCheckGrantedPermissionsForAppResponseable 
// Deprecated: This class is obsolete. Use checkGrantedPermissionsForAppPostResponse instead.
type ItemCheckGrantedPermissionsForAppResponseable interface {
    ItemCheckGrantedPermissionsForAppPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
