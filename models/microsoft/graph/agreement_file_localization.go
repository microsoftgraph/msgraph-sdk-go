package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// agreementFileLocalization 
type AgreementFileLocalization struct {
    AgreementFileProperties
    // 
    versions []AgreementFileVersion;
}
// NewAgreementFileLocalization instantiates a new agreementFileLocalization and sets the default values.
func NewAgreementFileLocalization()(*AgreementFileLocalization) {
    m := &AgreementFileLocalization{
        AgreementFileProperties: *NewAgreementFileProperties(),
    }
    return m
}
// GetVersions gets the versions property value. 
func (m *AgreementFileLocalization) GetVersions()([]AgreementFileVersion) {
    if m == nil {
        return nil
    } else {
        return m.versions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AgreementFileLocalization) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.AgreementFileProperties.GetFieldDeserializers()
    res["versions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAgreementFileVersion() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AgreementFileVersion, len(val))
            for i, v := range val {
                res[i] = *(v.(*AgreementFileVersion))
            }
            m.SetVersions(res)
        }
        return nil
    }
    return res
}
func (m *AgreementFileLocalization) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetVersions sets the versions property value. 
func (m *AgreementFileLocalization) SetVersions(value []AgreementFileVersion)() {
    m.versions = value
}
