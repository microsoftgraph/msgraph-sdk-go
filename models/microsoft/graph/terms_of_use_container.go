package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TermsOfUseContainer 
type TermsOfUseContainer struct {
    Entity
    // Represents the current status of a user's response to a company's customizable terms of use agreement.
    agreementAcceptances []AgreementAcceptanceable;
    // Represents a tenant's customizable terms of use agreement that's created and managed with Azure Active Directory (Azure AD).
    agreements []Agreementable;
}
// NewTermsOfUseContainer instantiates a new termsOfUseContainer and sets the default values.
func NewTermsOfUseContainer()(*TermsOfUseContainer) {
    m := &TermsOfUseContainer{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTermsOfUseContainerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTermsOfUseContainerFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTermsOfUseContainer(), nil
}
// GetAgreementAcceptances gets the agreementAcceptances property value. Represents the current status of a user's response to a company's customizable terms of use agreement.
func (m *TermsOfUseContainer) GetAgreementAcceptances()([]AgreementAcceptanceable) {
    if m == nil {
        return nil
    } else {
        return m.agreementAcceptances
    }
}
// GetAgreements gets the agreements property value. Represents a tenant's customizable terms of use agreement that's created and managed with Azure Active Directory (Azure AD).
func (m *TermsOfUseContainer) GetAgreements()([]Agreementable) {
    if m == nil {
        return nil
    } else {
        return m.agreements
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TermsOfUseContainer) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agreementAcceptances"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAgreementAcceptanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AgreementAcceptanceable, len(val))
            for i, v := range val {
                res[i] = v.(AgreementAcceptanceable)
            }
            m.SetAgreementAcceptances(res)
        }
        return nil
    }
    res["agreements"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAgreementFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Agreementable, len(val))
            for i, v := range val {
                res[i] = v.(Agreementable)
            }
            m.SetAgreements(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TermsOfUseContainer) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAgreementAcceptances() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAgreementAcceptances()))
        for i, v := range m.GetAgreementAcceptances() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("agreementAcceptances", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAgreements() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAgreements()))
        for i, v := range m.GetAgreements() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("agreements", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAgreementAcceptances sets the agreementAcceptances property value. Represents the current status of a user's response to a company's customizable terms of use agreement.
func (m *TermsOfUseContainer) SetAgreementAcceptances(value []AgreementAcceptanceable)() {
    if m != nil {
        m.agreementAcceptances = value
    }
}
// SetAgreements sets the agreements property value. Represents a tenant's customizable terms of use agreement that's created and managed with Azure Active Directory (Azure AD).
func (m *TermsOfUseContainer) SetAgreements(value []Agreementable)() {
    if m != nil {
        m.agreements = value
    }
}
