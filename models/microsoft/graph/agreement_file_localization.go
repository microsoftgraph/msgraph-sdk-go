package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AgreementFileLocalization struct {
    AgreementFileProperties
    // 
    versions []AgreementFileVersion;
}
// Instantiates a new agreementFileLocalization and sets the default values.
func NewAgreementFileLocalization()(*AgreementFileLocalization) {
    m := &AgreementFileLocalization{
        AgreementFileProperties: *NewAgreementFileProperties(),
    }
    return m
}
// Gets the versions property value. 
func (m *AgreementFileLocalization) GetVersions()([]AgreementFileVersion) {
    if m == nil {
        return nil
    } else {
        return m.versions
    }
}
// The deserialization information for the current model
func (m *AgreementFileLocalization) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.AgreementFileProperties.GetFieldDeserializers()
    res["versions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAgreementFileVersion() })
        if err != nil {
            return err
        }
        res := make([]AgreementFileVersion, len(val))
        for i, v := range val {
            res[i] = *(v.(*AgreementFileVersion))
        }
        m.SetVersions(res)
        return nil
    }
    return res
}
func (m *AgreementFileLocalization) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AgreementFileLocalization) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.AgreementFileProperties.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetVersions()))
        for i, v := range m.GetVersions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("versions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the versions property value. 
// Parameters:
//  - value : Value to set for the versions property.
func (m *AgreementFileLocalization) SetVersions(value []AgreementFileVersion)() {
    m.versions = value
}
