package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AgreementFile 
type AgreementFile struct {
    AgreementFileProperties
    // The localized version of the terms of use agreement files attached to the agreement.
    localizations []AgreementFileLocalizationable;
}
// NewAgreementFile instantiates a new agreementFile and sets the default values.
func NewAgreementFile()(*AgreementFile) {
    m := &AgreementFile{
        AgreementFileProperties: *NewAgreementFileProperties(),
    }
    return m
}
// CreateAgreementFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAgreementFileFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAgreementFile(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AgreementFile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.AgreementFileProperties.GetFieldDeserializers()
    res["localizations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAgreementFileLocalizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AgreementFileLocalizationable, len(val))
            for i, v := range val {
                res[i] = v.(AgreementFileLocalizationable)
            }
            m.SetLocalizations(res)
        }
        return nil
    }
    return res
}
// GetLocalizations gets the localizations property value. The localized version of the terms of use agreement files attached to the agreement.
func (m *AgreementFile) GetLocalizations()([]AgreementFileLocalizationable) {
    if m == nil {
        return nil
    } else {
        return m.localizations
    }
}
// Serialize serializes information the current object
func (m *AgreementFile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.AgreementFileProperties.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetLocalizations() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLocalizations()))
        for i, v := range m.GetLocalizations() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("localizations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLocalizations sets the localizations property value. The localized version of the terms of use agreement files attached to the agreement.
func (m *AgreementFile) SetLocalizations(value []AgreementFileLocalizationable)() {
    if m != nil {
        m.localizations = value
    }
}
