package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// provisionedIdentity 
type ProvisionedIdentity struct {
    Identity
    // Details of the identity.
    details *DetailsInfo;
    // Type of identity that has been provisioned, such as 'user' or 'group'.
    identityType *string;
}
// NewProvisionedIdentity instantiates a new provisionedIdentity and sets the default values.
func NewProvisionedIdentity()(*ProvisionedIdentity) {
    m := &ProvisionedIdentity{
        Identity: *NewIdentity(),
    }
    return m
}
// GetDetails gets the details property value. Details of the identity.
func (m *ProvisionedIdentity) GetDetails()(*DetailsInfo) {
    if m == nil {
        return nil
    } else {
        return m.details
    }
}
// GetIdentityType gets the identityType property value. Type of identity that has been provisioned, such as 'user' or 'group'.
func (m *ProvisionedIdentity) GetIdentityType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.identityType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisionedIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["details"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDetailsInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetails(val.(*DetailsInfo))
        }
        return nil
    }
    res["identityType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityType(val)
        }
        return nil
    }
    return res
}
func (m *ProvisionedIdentity) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ProvisionedIdentity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("details", m.GetDetails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("identityType", m.GetIdentityType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDetails sets the details property value. Details of the identity.
func (m *ProvisionedIdentity) SetDetails(value *DetailsInfo)() {
    m.details = value
}
// SetIdentityType sets the identityType property value. Type of identity that has been provisioned, such as 'user' or 'group'.
func (m *ProvisionedIdentity) SetIdentityType(value *string)() {
    m.identityType = value
}
