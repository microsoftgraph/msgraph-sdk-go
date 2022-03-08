package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AgreementFileLocalization provides operations to manage the collection of agreement entities.
type AgreementFileLocalization struct {
    AgreementFileProperties
    // Read-only. Customized versions of the terms of use agreement in the Azure AD tenant.
    versions []AgreementFileVersionable;
}
// NewAgreementFileLocalization instantiates a new agreementFileLocalization and sets the default values.
func NewAgreementFileLocalization()(*AgreementFileLocalization) {
    m := &AgreementFileLocalization{
        AgreementFileProperties: *NewAgreementFileProperties(),
    }
    return m
}
// CreateAgreementFileLocalizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAgreementFileLocalizationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAgreementFileLocalization(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AgreementFileLocalization) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.AgreementFileProperties.GetFieldDeserializers()
    res["versions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAgreementFileVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AgreementFileVersionable, len(val))
            for i, v := range val {
                res[i] = v.(AgreementFileVersionable)
            }
            m.SetVersions(res)
        }
        return nil
    }
    return res
}
// GetVersions gets the versions property value. Read-only. Customized versions of the terms of use agreement in the Azure AD tenant.
func (m *AgreementFileLocalization) GetVersions()([]AgreementFileVersionable) {
    if m == nil {
        return nil
    } else {
        return m.versions
    }
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
    if m.GetVersions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetVersions()))
        for i, v := range m.GetVersions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("versions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetVersions sets the versions property value. Read-only. Customized versions of the terms of use agreement in the Azure AD tenant.
func (m *AgreementFileLocalization) SetVersions(value []AgreementFileVersionable)() {
    if m != nil {
        m.versions = value
    }
}
