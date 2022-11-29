package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemDrivesItemRootInvitePostRequestBody provides operations to call the invite method.
type UsersItemDrivesItemRootInvitePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The expirationDateTime property
    expirationDateTime *string
    // The message property
    message *string
    // The password property
    password *string
    // The recipients property
    recipients []DriveRecipientable
    // The requireSignIn property
    requireSignIn *bool
    // The retainInheritedPermissions property
    retainInheritedPermissions *bool
    // The roles property
    roles []string
    // The sendInvitation property
    sendInvitation *bool
}
// NewUsersItemDrivesItemRootInvitePostRequestBody instantiates a new UsersItemDrivesItemRootInvitePostRequestBody and sets the default values.
func NewUsersItemDrivesItemRootInvitePostRequestBody()(*UsersItemDrivesItemRootInvitePostRequestBody) {
    m := &UsersItemDrivesItemRootInvitePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemDrivesItemRootInvitePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemDrivesItemRootInvitePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemDrivesItemRootInvitePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemDrivesItemRootInvitePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetExpirationDateTime gets the expirationDateTime property value. The expirationDateTime property
func (m *UsersItemDrivesItemRootInvitePostRequestBody) GetExpirationDateTime()(*string) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemDrivesItemRootInvitePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expirationDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExpirationDateTime)
    res["message"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMessage)
    res["password"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPassword)
    res["recipients"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDriveRecipientFromDiscriminatorValue , m.SetRecipients)
    res["requireSignIn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRequireSignIn)
    res["retainInheritedPermissions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRetainInheritedPermissions)
    res["roles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoles)
    res["sendInvitation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetSendInvitation)
    return res
}
// GetMessage gets the message property value. The message property
func (m *UsersItemDrivesItemRootInvitePostRequestBody) GetMessage()(*string) {
    return m.message
}
// GetPassword gets the password property value. The password property
func (m *UsersItemDrivesItemRootInvitePostRequestBody) GetPassword()(*string) {
    return m.password
}
// GetRecipients gets the recipients property value. The recipients property
func (m *UsersItemDrivesItemRootInvitePostRequestBody) GetRecipients()([]DriveRecipientable) {
    return m.recipients
}
// GetRequireSignIn gets the requireSignIn property value. The requireSignIn property
func (m *UsersItemDrivesItemRootInvitePostRequestBody) GetRequireSignIn()(*bool) {
    return m.requireSignIn
}
// GetRetainInheritedPermissions gets the retainInheritedPermissions property value. The retainInheritedPermissions property
func (m *UsersItemDrivesItemRootInvitePostRequestBody) GetRetainInheritedPermissions()(*bool) {
    return m.retainInheritedPermissions
}
// GetRoles gets the roles property value. The roles property
func (m *UsersItemDrivesItemRootInvitePostRequestBody) GetRoles()([]string) {
    return m.roles
}
// GetSendInvitation gets the sendInvitation property value. The sendInvitation property
func (m *UsersItemDrivesItemRootInvitePostRequestBody) GetSendInvitation()(*bool) {
    return m.sendInvitation
}
// Serialize serializes information the current object
func (m *UsersItemDrivesItemRootInvitePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    if m.GetRecipients() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRecipients())
        err := writer.WriteCollectionOfObjectValues("recipients", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("requireSignIn", m.GetRequireSignIn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("retainInheritedPermissions", m.GetRetainInheritedPermissions())
        if err != nil {
            return err
        }
    }
    if m.GetRoles() != nil {
        err := writer.WriteCollectionOfStringValues("roles", m.GetRoles())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sendInvitation", m.GetSendInvitation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemDrivesItemRootInvitePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetExpirationDateTime sets the expirationDateTime property value. The expirationDateTime property
func (m *UsersItemDrivesItemRootInvitePostRequestBody) SetExpirationDateTime(value *string)() {
    m.expirationDateTime = value
}
// SetMessage sets the message property value. The message property
func (m *UsersItemDrivesItemRootInvitePostRequestBody) SetMessage(value *string)() {
    m.message = value
}
// SetPassword sets the password property value. The password property
func (m *UsersItemDrivesItemRootInvitePostRequestBody) SetPassword(value *string)() {
    m.password = value
}
// SetRecipients sets the recipients property value. The recipients property
func (m *UsersItemDrivesItemRootInvitePostRequestBody) SetRecipients(value []DriveRecipientable)() {
    m.recipients = value
}
// SetRequireSignIn sets the requireSignIn property value. The requireSignIn property
func (m *UsersItemDrivesItemRootInvitePostRequestBody) SetRequireSignIn(value *bool)() {
    m.requireSignIn = value
}
// SetRetainInheritedPermissions sets the retainInheritedPermissions property value. The retainInheritedPermissions property
func (m *UsersItemDrivesItemRootInvitePostRequestBody) SetRetainInheritedPermissions(value *bool)() {
    m.retainInheritedPermissions = value
}
// SetRoles sets the roles property value. The roles property
func (m *UsersItemDrivesItemRootInvitePostRequestBody) SetRoles(value []string)() {
    m.roles = value
}
// SetSendInvitation sets the sendInvitation property value. The sendInvitation property
func (m *UsersItemDrivesItemRootInvitePostRequestBody) SetSendInvitation(value *bool)() {
    m.sendInvitation = value
}
