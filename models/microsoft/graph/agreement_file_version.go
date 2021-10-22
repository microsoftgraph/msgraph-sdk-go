package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AgreementFileVersion struct {
    AgreementFileProperties
}
func NewAgreementFileVersion()(*AgreementFileVersion) {
    m := &AgreementFileVersion{
        AgreementFileProperties: *NewAgreementFileProperties(),
    }
    return m
}
func (m *AgreementFileVersion) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.AgreementFileProperties.GetFieldDeserializers()
    return res
}
func (m *AgreementFileVersion) IsNil()(bool) {
    return m == nil
}
func (m *AgreementFileVersion) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.AgreementFileProperties.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
