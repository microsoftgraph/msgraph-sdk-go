package groupsettingtemplates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCheckMemberGroupsResponse 
// Deprecated: This class is obsolete. Use checkMemberGroupsPostResponse instead.
type ItemCheckMemberGroupsResponse struct {
    ItemCheckMemberGroupsPostResponse
}
// NewItemCheckMemberGroupsResponse instantiates a new ItemCheckMemberGroupsResponse and sets the default values.
func NewItemCheckMemberGroupsResponse()(*ItemCheckMemberGroupsResponse) {
    m := &ItemCheckMemberGroupsResponse{
        ItemCheckMemberGroupsPostResponse: *NewItemCheckMemberGroupsPostResponse(),
    }
    return m
}
// CreateItemCheckMemberGroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCheckMemberGroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCheckMemberGroupsResponse(), nil
}
// ItemCheckMemberGroupsResponseable 
// Deprecated: This class is obsolete. Use checkMemberGroupsPostResponse instead.
type ItemCheckMemberGroupsResponseable interface {
    ItemCheckMemberGroupsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
