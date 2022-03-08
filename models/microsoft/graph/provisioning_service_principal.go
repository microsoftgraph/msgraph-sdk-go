package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ProvisioningServicePrincipal provides operations to manage the auditLogRoot singleton.
type ProvisioningServicePrincipal struct {
    Identity
}
// NewProvisioningServicePrincipal instantiates a new provisioningServicePrincipal and sets the default values.
func NewProvisioningServicePrincipal()(*ProvisioningServicePrincipal) {
    m := &ProvisioningServicePrincipal{
        Identity: *NewIdentity(),
    }
    return m
}
// CreateProvisioningServicePrincipalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningServicePrincipalFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewProvisioningServicePrincipal(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningServicePrincipal) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    return res
}
func (m *ProvisioningServicePrincipal) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ProvisioningServicePrincipal) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
