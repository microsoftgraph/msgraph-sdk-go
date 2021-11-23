package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// claimsMappingPolicy 
type ClaimsMappingPolicy struct {
    StsPolicy
}
// NewClaimsMappingPolicy instantiates a new claimsMappingPolicy and sets the default values.
func NewClaimsMappingPolicy()(*ClaimsMappingPolicy) {
    m := &ClaimsMappingPolicy{
        StsPolicy: *NewStsPolicy(),
    }
    return m
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClaimsMappingPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.StsPolicy.GetFieldDeserializers()
    return res
}
func (m *ClaimsMappingPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ClaimsMappingPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.StsPolicy.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
