package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeDrivesItemItemsItemInvitePostRequestBody provides operations to call the invite method.
type MeDrivesItemItemsItemInvitePostRequestBody struct {
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
// NewMeDrivesItemItemsItemInvitePostRequestBody instantiates a new MeDrivesItemItemsItemInvitePostRequestBody and sets the default values.
func NewMeDrivesItemItemsItemInvitePostRequestBody()(*MeDrivesItemItemsItemInvitePostRequestBody) {
    m := &MeDrivesItemItemsItemInvitePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeDrivesItemItemsItemInvitePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeDrivesItemItemsItemInvitePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeDrivesItemItemsItemInvitePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeDrivesItemItemsItemInvitePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetExpirationDateTime gets the expirationDateTime property value. The expirationDateTime property
func (m *MeDrivesItemItemsItemInvitePostRequestBody) GetExpirationDateTime()(*string) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeDrivesItemItemsItemInvitePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    res["recipients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDriveRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveRecipientable, len(val))
            for i, v := range val {
                res[i] = v.(DriveRecipientable)
            }
            m.SetRecipients(res)
        }
        return nil
    }
    res["requireSignIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireSignIn(val)
        }
        return nil
    }
    res["retainInheritedPermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRetainInheritedPermissions(val)
        }
        return nil
    }
    res["roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoles(res)
        }
        return nil
    }
    res["sendInvitation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSendInvitation(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The message property
func (m *MeDrivesItemItemsItemInvitePostRequestBody) GetMessage()(*string) {
    return m.message
}
// GetPassword gets the password property value. The password property
func (m *MeDrivesItemItemsItemInvitePostRequestBody) GetPassword()(*string) {
    return m.password
}
// GetRecipients gets the recipients property value. The recipients property
func (m *MeDrivesItemItemsItemInvitePostRequestBody) GetRecipients()([]DriveRecipientable) {
    return m.recipients
}
// GetRequireSignIn gets the requireSignIn property value. The requireSignIn property
func (m *MeDrivesItemItemsItemInvitePostRequestBody) GetRequireSignIn()(*bool) {
    return m.requireSignIn
}
// GetRetainInheritedPermissions gets the retainInheritedPermissions property value. The retainInheritedPermissions property
func (m *MeDrivesItemItemsItemInvitePostRequestBody) GetRetainInheritedPermissions()(*bool) {
    return m.retainInheritedPermissions
}
// GetRoles gets the roles property value. The roles property
func (m *MeDrivesItemItemsItemInvitePostRequestBody) GetRoles()([]string) {
    return m.roles
}
// GetSendInvitation gets the sendInvitation property value. The sendInvitation property
func (m *MeDrivesItemItemsItemInvitePostRequestBody) GetSendInvitation()(*bool) {
    return m.sendInvitation
}
// Serialize serializes information the current object
func (m *MeDrivesItemItemsItemInvitePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRecipients()))
        for i, v := range m.GetRecipients() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
func (m *MeDrivesItemItemsItemInvitePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetExpirationDateTime sets the expirationDateTime property value. The expirationDateTime property
func (m *MeDrivesItemItemsItemInvitePostRequestBody) SetExpirationDateTime(value *string)() {
    m.expirationDateTime = value
}
// SetMessage sets the message property value. The message property
func (m *MeDrivesItemItemsItemInvitePostRequestBody) SetMessage(value *string)() {
    m.message = value
}
// SetPassword sets the password property value. The password property
func (m *MeDrivesItemItemsItemInvitePostRequestBody) SetPassword(value *string)() {
    m.password = value
}
// SetRecipients sets the recipients property value. The recipients property
func (m *MeDrivesItemItemsItemInvitePostRequestBody) SetRecipients(value []DriveRecipientable)() {
    m.recipients = value
}
// SetRequireSignIn sets the requireSignIn property value. The requireSignIn property
func (m *MeDrivesItemItemsItemInvitePostRequestBody) SetRequireSignIn(value *bool)() {
    m.requireSignIn = value
}
// SetRetainInheritedPermissions sets the retainInheritedPermissions property value. The retainInheritedPermissions property
func (m *MeDrivesItemItemsItemInvitePostRequestBody) SetRetainInheritedPermissions(value *bool)() {
    m.retainInheritedPermissions = value
}
// SetRoles sets the roles property value. The roles property
func (m *MeDrivesItemItemsItemInvitePostRequestBody) SetRoles(value []string)() {
    m.roles = value
}
// SetSendInvitation sets the sendInvitation property value. The sendInvitation property
func (m *MeDrivesItemItemsItemInvitePostRequestBody) SetSendInvitation(value *bool)() {
    m.sendInvitation = value
}
