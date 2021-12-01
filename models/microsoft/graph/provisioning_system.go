package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ProvisioningSystem 
type ProvisioningSystem struct {
    Identity
    // Details of the system.
    details *DetailsInfo;
}
// NewProvisioningSystem instantiates a new provisioningSystem and sets the default values.
func NewProvisioningSystem()(*ProvisioningSystem) {
    m := &ProvisioningSystem{
        Identity: *NewIdentity(),
    }
    return m
}
// GetDetails gets the details property value. Details of the system.
func (m *ProvisioningSystem) GetDetails()(*DetailsInfo) {
    if m == nil {
        return nil
    } else {
        return m.details
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningSystem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    return res
}
func (m *ProvisioningSystem) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ProvisioningSystem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    return nil
}
// SetDetails sets the details property value. Details of the system.
func (m *ProvisioningSystem) SetDetails(value *DetailsInfo)() {
    if m != nil {
        m.details = value
    }
}
