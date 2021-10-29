package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TermsOfUseContainer struct {
    Entity
    // 
    agreementAcceptances []AgreementAcceptance;
    // 
    agreements []Agreement;
}
// Instantiates a new termsOfUseContainer and sets the default values.
func NewTermsOfUseContainer()(*TermsOfUseContainer) {
    m := &TermsOfUseContainer{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the agreementAcceptances property value. 
func (m *TermsOfUseContainer) GetAgreementAcceptances()([]AgreementAcceptance) {
    if m == nil {
        return nil
    } else {
        return m.agreementAcceptances
    }
}
// Gets the agreements property value. 
func (m *TermsOfUseContainer) GetAgreements()([]Agreement) {
    if m == nil {
        return nil
    } else {
        return m.agreements
    }
}
// The deserialization information for the current model
func (m *TermsOfUseContainer) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agreementAcceptances"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAgreementAcceptance() })
        if err != nil {
            return err
        }
        res := make([]AgreementAcceptance, len(val))
        for i, v := range val {
            res[i] = *(v.(*AgreementAcceptance))
        }
        m.SetAgreementAcceptances(res)
        return nil
    }
    res["agreements"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAgreement() })
        if err != nil {
            return err
        }
        res := make([]Agreement, len(val))
        for i, v := range val {
            res[i] = *(v.(*Agreement))
        }
        m.SetAgreements(res)
        return nil
    }
    return res
}
func (m *TermsOfUseContainer) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TermsOfUseContainer) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAgreementAcceptances()))
        for i, v := range m.GetAgreementAcceptances() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("agreementAcceptances", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAgreements()))
        for i, v := range m.GetAgreements() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("agreements", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the agreementAcceptances property value. 
// Parameters:
//  - value : Value to set for the agreementAcceptances property.
func (m *TermsOfUseContainer) SetAgreementAcceptances(value []AgreementAcceptance)() {
    m.agreementAcceptances = value
}
// Sets the agreements property value. 
// Parameters:
//  - value : Value to set for the agreements property.
func (m *TermsOfUseContainer) SetAgreements(value []Agreement)() {
    m.agreements = value
}
