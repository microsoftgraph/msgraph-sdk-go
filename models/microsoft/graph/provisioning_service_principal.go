package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ProvisioningServicePrincipal struct {
    Identity
}
// Instantiates a new provisioningServicePrincipal and sets the default values.
func NewProvisioningServicePrincipal()(*ProvisioningServicePrincipal) {
    m := &ProvisioningServicePrincipal{
        Identity: *NewIdentity(),
    }
    return m
}
// The deserialization information for the current model
func (m *ProvisioningServicePrincipal) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    return res
}
func (m *ProvisioningServicePrincipal) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ProvisioningServicePrincipal) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
