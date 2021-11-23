package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// homeRealmDiscoveryPolicy 
type HomeRealmDiscoveryPolicy struct {
    StsPolicy
}
// NewHomeRealmDiscoveryPolicy instantiates a new homeRealmDiscoveryPolicy and sets the default values.
func NewHomeRealmDiscoveryPolicy()(*HomeRealmDiscoveryPolicy) {
    m := &HomeRealmDiscoveryPolicy{
        StsPolicy: *NewStsPolicy(),
    }
    return m
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HomeRealmDiscoveryPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.StsPolicy.GetFieldDeserializers()
    return res
}
func (m *HomeRealmDiscoveryPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *HomeRealmDiscoveryPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.StsPolicy.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
