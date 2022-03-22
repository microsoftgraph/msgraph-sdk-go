package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TokenIssuancePolicy 
type TokenIssuancePolicy struct {
    StsPolicy
}
// NewTokenIssuancePolicy instantiates a new tokenIssuancePolicy and sets the default values.
func NewTokenIssuancePolicy()(*TokenIssuancePolicy) {
    m := &TokenIssuancePolicy{
        StsPolicy: *NewStsPolicy(),
    }
    return m
}
// CreateTokenIssuancePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTokenIssuancePolicyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTokenIssuancePolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TokenIssuancePolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.StsPolicy.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *TokenIssuancePolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.StsPolicy.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
