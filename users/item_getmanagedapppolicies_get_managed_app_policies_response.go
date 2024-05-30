package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemGetmanagedapppoliciesGetManagedAppPoliciesGetResponseable instead.
type ItemGetmanagedapppoliciesGetManagedAppPoliciesResponse struct {
    ItemGetmanagedapppoliciesGetManagedAppPoliciesGetResponse
}
// NewItemGetmanagedapppoliciesGetManagedAppPoliciesResponse instantiates a new ItemGetmanagedapppoliciesGetManagedAppPoliciesResponse and sets the default values.
func NewItemGetmanagedapppoliciesGetManagedAppPoliciesResponse()(*ItemGetmanagedapppoliciesGetManagedAppPoliciesResponse) {
    m := &ItemGetmanagedapppoliciesGetManagedAppPoliciesResponse{
        ItemGetmanagedapppoliciesGetManagedAppPoliciesGetResponse: *NewItemGetmanagedapppoliciesGetManagedAppPoliciesGetResponse(),
    }
    return m
}
// CreateItemGetmanagedapppoliciesGetManagedAppPoliciesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemGetmanagedapppoliciesGetManagedAppPoliciesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGetmanagedapppoliciesGetManagedAppPoliciesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemGetmanagedapppoliciesGetManagedAppPoliciesGetResponseable instead.
type ItemGetmanagedapppoliciesGetManagedAppPoliciesResponseable interface {
    ItemGetmanagedapppoliciesGetManagedAppPoliciesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
