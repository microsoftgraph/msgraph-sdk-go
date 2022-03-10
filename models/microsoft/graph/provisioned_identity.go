package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ProvisionedIdentity provides operations to manage the auditLogRoot singleton.
type ProvisionedIdentity struct {
    Identity
    // Details of the identity.
    details DetailsInfoable;
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
// CreateProvisionedIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisionedIdentityFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewProvisionedIdentity(), nil
}
// GetDetails gets the details property value. Details of the identity.
func (m *ProvisionedIdentity) GetDetails()(DetailsInfoable) {
    if m == nil {
        return nil
    } else {
        return m.details
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisionedIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["details"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDetailsInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetails(val.(DetailsInfoable))
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
// GetIdentityType gets the identityType property value. Type of identity that has been provisioned, such as 'user' or 'group'.
func (m *ProvisionedIdentity) GetIdentityType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.identityType
    }
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
func (m *ProvisionedIdentity) SetDetails(value DetailsInfoable)() {
    if m != nil {
        m.details = value
    }
}
// SetIdentityType sets the identityType property value. Type of identity that has been provisioned, such as 'user' or 'group'.
func (m *ProvisionedIdentity) SetIdentityType(value *string)() {
    if m != nil {
        m.identityType = value
    }
}
