package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SharePointIdentity provides operations to manage the drive singleton.
type SharePointIdentity struct {
    Identity
    // The sign in name of the SharePoint identity.
    loginName *string;
}
// NewSharePointIdentity instantiates a new sharePointIdentity and sets the default values.
func NewSharePointIdentity()(*SharePointIdentity) {
    m := &SharePointIdentity{
        Identity: *NewIdentity(),
    }
    return m
}
// CreateSharePointIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharePointIdentityFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSharePointIdentity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharePointIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["loginName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoginName(val)
        }
        return nil
    }
    return res
}
// GetLoginName gets the loginName property value. The sign in name of the SharePoint identity.
func (m *SharePointIdentity) GetLoginName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.loginName
    }
}
func (m *SharePointIdentity) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SharePointIdentity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("loginName", m.GetLoginName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLoginName sets the loginName property value. The sign in name of the SharePoint identity.
func (m *SharePointIdentity) SetLoginName(value *string)() {
    if m != nil {
        m.loginName = value
    }
}
