package invite

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// InviteRequestBody provides operations to call the invite method.
type InviteRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    expirationDateTime *string;
    // 
    message *string;
    // 
    password *string;
    // 
    recipients []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveRecipientable;
    // 
    requireSignIn *bool;
    // 
    retainInheritedPermissions *bool;
    // 
    roles []string;
    // 
    sendInvitation *bool;
}
// NewInviteRequestBody instantiates a new inviteRequestBody and sets the default values.
func NewInviteRequestBody()(*InviteRequestBody) {
    m := &InviteRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateInviteRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInviteRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewInviteRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InviteRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExpirationDateTime gets the expirationDateTime property value. 
func (m *InviteRequestBody) GetExpirationDateTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InviteRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["message"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["password"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    res["recipients"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateDriveRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveRecipientable, len(val))
            for i, v := range val {
                res[i] = v.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveRecipientable)
            }
            m.SetRecipients(res)
        }
        return nil
    }
    res["requireSignIn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireSignIn(val)
        }
        return nil
    }
    res["retainInheritedPermissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRetainInheritedPermissions(val)
        }
        return nil
    }
    res["roles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    res["sendInvitation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
// GetMessage gets the message property value. 
func (m *InviteRequestBody) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
// GetPassword gets the password property value. 
func (m *InviteRequestBody) GetPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.password
    }
}
// GetRecipients gets the recipients property value. 
func (m *InviteRequestBody) GetRecipients()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveRecipientable) {
    if m == nil {
        return nil
    } else {
        return m.recipients
    }
}
// GetRequireSignIn gets the requireSignIn property value. 
func (m *InviteRequestBody) GetRequireSignIn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.requireSignIn
    }
}
// GetRetainInheritedPermissions gets the retainInheritedPermissions property value. 
func (m *InviteRequestBody) GetRetainInheritedPermissions()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.retainInheritedPermissions
    }
}
// GetRoles gets the roles property value. 
func (m *InviteRequestBody) GetRoles()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roles
    }
}
// GetSendInvitation gets the sendInvitation property value. 
func (m *InviteRequestBody) GetSendInvitation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sendInvitation
    }
}
func (m *InviteRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *InviteRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRecipients()))
        for i, v := range m.GetRecipients() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *InviteRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. 
func (m *InviteRequestBody) SetExpirationDateTime(value *string)() {
    if m != nil {
        m.expirationDateTime = value
    }
}
// SetMessage sets the message property value. 
func (m *InviteRequestBody) SetMessage(value *string)() {
    if m != nil {
        m.message = value
    }
}
// SetPassword sets the password property value. 
func (m *InviteRequestBody) SetPassword(value *string)() {
    if m != nil {
        m.password = value
    }
}
// SetRecipients sets the recipients property value. 
func (m *InviteRequestBody) SetRecipients(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveRecipientable)() {
    if m != nil {
        m.recipients = value
    }
}
// SetRequireSignIn sets the requireSignIn property value. 
func (m *InviteRequestBody) SetRequireSignIn(value *bool)() {
    if m != nil {
        m.requireSignIn = value
    }
}
// SetRetainInheritedPermissions sets the retainInheritedPermissions property value. 
func (m *InviteRequestBody) SetRetainInheritedPermissions(value *bool)() {
    if m != nil {
        m.retainInheritedPermissions = value
    }
}
// SetRoles sets the roles property value. 
func (m *InviteRequestBody) SetRoles(value []string)() {
    if m != nil {
        m.roles = value
    }
}
// SetSendInvitation sets the sendInvitation property value. 
func (m *InviteRequestBody) SetSendInvitation(value *bool)() {
    if m != nil {
        m.sendInvitation = value
    }
}
