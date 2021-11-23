package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AgreementFile 
type AgreementFile struct {
    AgreementFileProperties
    // 
    localizations []AgreementFileLocalization;
}
// NewAgreementFile instantiates a new agreementFile and sets the default values.
func NewAgreementFile()(*AgreementFile) {
    m := &AgreementFile{
        AgreementFileProperties: *NewAgreementFileProperties(),
    }
    return m
}
// GetLocalizations gets the localizations property value. 
func (m *AgreementFile) GetLocalizations()([]AgreementFileLocalization) {
    if m == nil {
        return nil
    } else {
        return m.localizations
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AgreementFile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.AgreementFileProperties.GetFieldDeserializers()
    res["localizations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAgreementFileLocalization() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AgreementFileLocalization, len(val))
            for i, v := range val {
                res[i] = *(v.(*AgreementFileLocalization))
            }
            m.SetLocalizations(res)
        }
        return nil
    }
    return res
}
func (m *AgreementFile) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AgreementFile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.AgreementFileProperties.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLocalizations()))
        for i, v := range m.GetLocalizations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("localizations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLocalizations sets the localizations property value. 
func (m *AgreementFile) SetLocalizations(value []AgreementFileLocalization)() {
    m.localizations = value
}
