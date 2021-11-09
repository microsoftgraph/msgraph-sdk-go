package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SharingInvitation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The email address provided for the recipient of the sharing invitation. Read-only.
    email *string;
    // Provides information about who sent the invitation that created this permission, if that information is available. Read-only.
    invitedBy *IdentitySet;
    // 
    redeemedBy *string;
    // If true the recipient of the invitation needs to sign in in order to access the shared item. Read-only.
    signInRequired *bool;
}
// Instantiates a new sharingInvitation and sets the default values.
func NewSharingInvitation()(*SharingInvitation) {
    m := &SharingInvitation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SharingInvitation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the email property value. The email address provided for the recipient of the sharing invitation. Read-only.
func (m *SharingInvitation) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// Gets the invitedBy property value. Provides information about who sent the invitation that created this permission, if that information is available. Read-only.
func (m *SharingInvitation) GetInvitedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.invitedBy
    }
}
// Gets the redeemedBy property value. 
func (m *SharingInvitation) GetRedeemedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.redeemedBy
    }
}
// Gets the signInRequired property value. If true the recipient of the invitation needs to sign in in order to access the shared item. Read-only.
func (m *SharingInvitation) GetSignInRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.signInRequired
    }
}
// The deserialization information for the current model
func (m *SharingInvitation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["invitedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvitedBy(val.(*IdentitySet))
        }
        return nil
    }
    res["redeemedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRedeemedBy(val)
        }
        return nil
    }
    res["signInRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInRequired(val)
        }
        return nil
    }
    return res
}
func (m *SharingInvitation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SharingInvitation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("invitedBy", m.GetInvitedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("redeemedBy", m.GetRedeemedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("signInRequired", m.GetSignInRequired())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *SharingInvitation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the email property value. The email address provided for the recipient of the sharing invitation. Read-only.
// Parameters:
//  - value : Value to set for the email property.
func (m *SharingInvitation) SetEmail(value *string)() {
    m.email = value
}
// Sets the invitedBy property value. Provides information about who sent the invitation that created this permission, if that information is available. Read-only.
// Parameters:
//  - value : Value to set for the invitedBy property.
func (m *SharingInvitation) SetInvitedBy(value *IdentitySet)() {
    m.invitedBy = value
}
// Sets the redeemedBy property value. 
// Parameters:
//  - value : Value to set for the redeemedBy property.
func (m *SharingInvitation) SetRedeemedBy(value *string)() {
    m.redeemedBy = value
}
// Sets the signInRequired property value. If true the recipient of the invitation needs to sign in in order to access the shared item. Read-only.
// Parameters:
//  - value : Value to set for the signInRequired property.
func (m *SharingInvitation) SetSignInRequired(value *bool)() {
    m.signInRequired = value
}
