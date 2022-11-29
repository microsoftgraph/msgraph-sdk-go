package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsResponse provides operations to call the getMemberGroups method.
type UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []string
}
// NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsResponse instantiates a new UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsResponse and sets the default values.
func NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsResponse()(*UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsResponse) {
    m := &UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsResponse) GetValue()([]string) {
    return m.value
}
// Serialize serializes information the current object
func (m *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        err = writer.WriteCollectionOfStringValues("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsResponse) SetValue(value []string)() {
    m.value = value
}
