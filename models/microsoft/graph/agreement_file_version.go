package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AgreementFileVersion provides operations to manage the collection of agreement entities.
type AgreementFileVersion struct {
    AgreementFileProperties
}
// NewAgreementFileVersion instantiates a new agreementFileVersion and sets the default values.
func NewAgreementFileVersion()(*AgreementFileVersion) {
    m := &AgreementFileVersion{
        AgreementFileProperties: *NewAgreementFileProperties(),
    }
    return m
}
// CreateAgreementFileVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAgreementFileVersionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAgreementFileVersion(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AgreementFileVersion) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.AgreementFileProperties.GetFieldDeserializers()
    return res
}
func (m *AgreementFileVersion) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AgreementFileVersion) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.AgreementFileProperties.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
