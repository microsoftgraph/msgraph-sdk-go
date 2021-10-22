package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Agreement struct {
    Entity
    acceptances []AgreementAcceptance;
    displayName *string;
    file *AgreementFile;
    files []AgreementFileLocalization;
    isPerDeviceAcceptanceRequired *bool;
    isViewingBeforeAcceptanceRequired *bool;
    termsExpiration *TermsExpiration;
    userReacceptRequiredFrequency *string;
}
func NewAgreement()(*Agreement) {
    m := &Agreement{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Agreement) GetAcceptances()([]AgreementAcceptance) {
    if m == nil {
        return nil
    } else {
        return m.acceptances
    }
}
func (m *Agreement) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *Agreement) GetFile()(*AgreementFile) {
    if m == nil {
        return nil
    } else {
        return m.file
    }
}
func (m *Agreement) GetFiles()([]AgreementFileLocalization) {
    if m == nil {
        return nil
    } else {
        return m.files
    }
}
func (m *Agreement) GetIsPerDeviceAcceptanceRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPerDeviceAcceptanceRequired
    }
}
func (m *Agreement) GetIsViewingBeforeAcceptanceRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isViewingBeforeAcceptanceRequired
    }
}
func (m *Agreement) GetTermsExpiration()(*TermsExpiration) {
    if m == nil {
        return nil
    } else {
        return m.termsExpiration
    }
}
func (m *Agreement) GetUserReacceptRequiredFrequency()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userReacceptRequiredFrequency
    }
}
func (m *Agreement) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acceptances"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAgreementAcceptance() })
        if err != nil {
            return err
        }
        res := make([]AgreementAcceptance, len(val))
        for i, v := range val {
            res[i] = *(v.(*AgreementAcceptance))
        }
        m.SetAcceptances(res)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["file"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAgreementFile() })
        if err != nil {
            return err
        }
        m.SetFile(val.(*AgreementFile))
        return nil
    }
    res["files"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAgreementFileLocalization() })
        if err != nil {
            return err
        }
        res := make([]AgreementFileLocalization, len(val))
        for i, v := range val {
            res[i] = *(v.(*AgreementFileLocalization))
        }
        m.SetFiles(res)
        return nil
    }
    res["isPerDeviceAcceptanceRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsPerDeviceAcceptanceRequired(val)
        return nil
    }
    res["isViewingBeforeAcceptanceRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsViewingBeforeAcceptanceRequired(val)
        return nil
    }
    res["termsExpiration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTermsExpiration() })
        if err != nil {
            return err
        }
        m.SetTermsExpiration(val.(*TermsExpiration))
        return nil
    }
    res["userReacceptRequiredFrequency"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserReacceptRequiredFrequency(val)
        return nil
    }
    return res
}
func (m *Agreement) IsNil()(bool) {
    return m == nil
}
func (m *Agreement) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAcceptances()))
        for i, v := range m.GetAcceptances() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("acceptances", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("file", m.GetFile())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFiles()))
        for i, v := range m.GetFiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("files", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isPerDeviceAcceptanceRequired", m.GetIsPerDeviceAcceptanceRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isViewingBeforeAcceptanceRequired", m.GetIsViewingBeforeAcceptanceRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("termsExpiration", m.GetTermsExpiration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userReacceptRequiredFrequency", m.GetUserReacceptRequiredFrequency())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Agreement) SetAcceptances(value []AgreementAcceptance)() {
    m.acceptances = value
}
func (m *Agreement) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *Agreement) SetFile(value *AgreementFile)() {
    m.file = value
}
func (m *Agreement) SetFiles(value []AgreementFileLocalization)() {
    m.files = value
}
func (m *Agreement) SetIsPerDeviceAcceptanceRequired(value *bool)() {
    m.isPerDeviceAcceptanceRequired = value
}
func (m *Agreement) SetIsViewingBeforeAcceptanceRequired(value *bool)() {
    m.isViewingBeforeAcceptanceRequired = value
}
func (m *Agreement) SetTermsExpiration(value *TermsExpiration)() {
    m.termsExpiration = value
}
func (m *Agreement) SetUserReacceptRequiredFrequency(value *string)() {
    m.userReacceptRequiredFrequency = value
}
