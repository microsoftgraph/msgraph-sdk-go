package education

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemAssignmentsItemCategoriesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type UsersItemAssignmentsItemCategoriesDeltaResponse struct {
    UsersItemAssignmentsItemCategoriesDeltaGetResponse
}
// NewUsersItemAssignmentsItemCategoriesDeltaResponse instantiates a new UsersItemAssignmentsItemCategoriesDeltaResponse and sets the default values.
func NewUsersItemAssignmentsItemCategoriesDeltaResponse()(*UsersItemAssignmentsItemCategoriesDeltaResponse) {
    m := &UsersItemAssignmentsItemCategoriesDeltaResponse{
        UsersItemAssignmentsItemCategoriesDeltaGetResponse: *NewUsersItemAssignmentsItemCategoriesDeltaGetResponse(),
    }
    return m
}
// CreateUsersItemAssignmentsItemCategoriesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemAssignmentsItemCategoriesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemAssignmentsItemCategoriesDeltaResponse(), nil
}
// UsersItemAssignmentsItemCategoriesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type UsersItemAssignmentsItemCategoriesDeltaResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UsersItemAssignmentsItemCategoriesDeltaGetResponseable
}
