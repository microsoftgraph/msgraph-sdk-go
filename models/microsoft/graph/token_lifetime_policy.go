package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TokenLifetimePolicy 
type TokenLifetimePolicy struct {
    StsPolicy
}
// NewTokenLifetimePolicy instantiates a new tokenLifetimePolicy and sets the default values.
func NewTokenLifetimePolicy()(*TokenLifetimePolicy) {
    m := &TokenLifetimePolicy{
        StsPolicy: *NewStsPolicy(),
    }
    return m
}
// CreateTokenLifetimePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTokenLifetimePolicyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTokenLifetimePolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TokenLifetimePolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.StsPolicy.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *TokenLifetimePolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.StsPolicy.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
