package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OrganizationalBrandingLocalization struct {
    OrganizationalBrandingProperties
}
// Instantiates a new organizationalBrandingLocalization and sets the default values.
func NewOrganizationalBrandingLocalization()(*OrganizationalBrandingLocalization) {
    m := &OrganizationalBrandingLocalization{
        OrganizationalBrandingProperties: *NewOrganizationalBrandingProperties(),
    }
    return m
}
// The deserialization information for the current model
func (m *OrganizationalBrandingLocalization) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OrganizationalBrandingProperties.GetFieldDeserializers()
    return res
}
func (m *OrganizationalBrandingLocalization) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OrganizationalBrandingLocalization) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.OrganizationalBrandingProperties.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
