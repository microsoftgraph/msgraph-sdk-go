package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SharingInvitation struct {
    additionalData map[string]interface{};
    email *string;
    invitedBy *IdentitySet;
    redeemedBy *string;
    signInRequired *bool;
}
func NewSharingInvitation()(*SharingInvitation) {
    m := &SharingInvitation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SharingInvitation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SharingInvitation) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
func (m *SharingInvitation) GetInvitedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.invitedBy
    }
}
func (m *SharingInvitation) GetRedeemedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.redeemedBy
    }
}
func (m *SharingInvitation) GetSignInRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.signInRequired
    }
}
func (m *SharingInvitation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmail(val)
        return nil
    }
    res["invitedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetInvitedBy(val.(*IdentitySet))
        return nil
    }
    res["redeemedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRedeemedBy(val)
        return nil
    }
    res["signInRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSignInRequired(val)
        return nil
    }
    return res
}
func (m *SharingInvitation) IsNil()(bool) {
    return m == nil
}
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
func (m *SharingInvitation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SharingInvitation) SetEmail(value *string)() {
    m.email = value
}
func (m *SharingInvitation) SetInvitedBy(value *IdentitySet)() {
    m.invitedBy = value
}
func (m *SharingInvitation) SetRedeemedBy(value *string)() {
    m.redeemedBy = value
}
func (m *SharingInvitation) SetSignInRequired(value *bool)() {
    m.signInRequired = value
}
