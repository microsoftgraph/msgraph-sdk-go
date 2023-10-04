package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSitesItemPermissionsItemGrantResponse 
// Deprecated: This class is obsolete. Use grantPostResponse instead.
type ItemSitesItemPermissionsItemGrantResponse struct {
    ItemSitesItemPermissionsItemGrantPostResponse
}
// NewItemSitesItemPermissionsItemGrantResponse instantiates a new ItemSitesItemPermissionsItemGrantResponse and sets the default values.
func NewItemSitesItemPermissionsItemGrantResponse()(*ItemSitesItemPermissionsItemGrantResponse) {
    m := &ItemSitesItemPermissionsItemGrantResponse{
        ItemSitesItemPermissionsItemGrantPostResponse: *NewItemSitesItemPermissionsItemGrantPostResponse(),
    }
    return m
}
// CreateItemSitesItemPermissionsItemGrantResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSitesItemPermissionsItemGrantResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemPermissionsItemGrantResponse(), nil
}
// ItemSitesItemPermissionsItemGrantResponseable 
// Deprecated: This class is obsolete. Use grantPostResponse instead.
type ItemSitesItemPermissionsItemGrantResponseable interface {
    ItemSitesItemPermissionsItemGrantPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
