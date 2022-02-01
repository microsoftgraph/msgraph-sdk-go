package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Agreement 
type Agreement struct {
    Entity
    // Read-only. Information about acceptances of this agreement.
    acceptances []AgreementAcceptance;
    // Display name of the agreement. The display name is used for internal tracking of the agreement but is not shown to end users who view the agreement.
    displayName *string;
    // Default PDF linked to this agreement.
    file *AgreementFile;
    // PDFs linked to this agreement. Note: This property is in the process of being deprecated. Use the  file property instead.
    files []AgreementFileLocalization;
    // This setting enables you to require end users to accept this agreement on every device that they are accessing it from. The end user will be required to register their device in Azure AD, if they haven't already done so.
    isPerDeviceAcceptanceRequired *bool;
    // Indicates whether the user has to expand the agreement before accepting.
    isViewingBeforeAcceptanceRequired *bool;
    // Expiration schedule and frequency of agreement for all users.
    termsExpiration *TermsExpiration;
    // The duration after which the user must re-accept the terms of use. The value is represented in ISO 8601 format for durations.
    userReacceptRequiredFrequency *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
}
// NewAgreement instantiates a new agreement and sets the default values.
func NewAgreement()(*Agreement) {
    m := &Agreement{
        Entity: *NewEntity(),
    }
    return m
}
// GetAcceptances gets the acceptances property value. Read-only. Information about acceptances of this agreement.
func (m *Agreement) GetAcceptances()([]AgreementAcceptance) {
    if m == nil {
        return nil
    } else {
        return m.acceptances
    }
}
// GetDisplayName gets the displayName property value. Display name of the agreement. The display name is used for internal tracking of the agreement but is not shown to end users who view the agreement.
func (m *Agreement) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFile gets the file property value. Default PDF linked to this agreement.
func (m *Agreement) GetFile()(*AgreementFile) {
    if m == nil {
        return nil
    } else {
        return m.file
    }
}
// GetFiles gets the files property value. PDFs linked to this agreement. Note: This property is in the process of being deprecated. Use the  file property instead.
func (m *Agreement) GetFiles()([]AgreementFileLocalization) {
    if m == nil {
        return nil
    } else {
        return m.files
    }
}
// GetIsPerDeviceAcceptanceRequired gets the isPerDeviceAcceptanceRequired property value. This setting enables you to require end users to accept this agreement on every device that they are accessing it from. The end user will be required to register their device in Azure AD, if they haven't already done so.
func (m *Agreement) GetIsPerDeviceAcceptanceRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPerDeviceAcceptanceRequired
    }
}
// GetIsViewingBeforeAcceptanceRequired gets the isViewingBeforeAcceptanceRequired property value. Indicates whether the user has to expand the agreement before accepting.
func (m *Agreement) GetIsViewingBeforeAcceptanceRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isViewingBeforeAcceptanceRequired
    }
}
// GetTermsExpiration gets the termsExpiration property value. Expiration schedule and frequency of agreement for all users.
func (m *Agreement) GetTermsExpiration()(*TermsExpiration) {
    if m == nil {
        return nil
    } else {
        return m.termsExpiration
    }
}
// GetUserReacceptRequiredFrequency gets the userReacceptRequiredFrequency property value. The duration after which the user must re-accept the terms of use. The value is represented in ISO 8601 format for durations.
func (m *Agreement) GetUserReacceptRequiredFrequency()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.userReacceptRequiredFrequency
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Agreement) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acceptances"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAgreementAcceptance() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AgreementAcceptance, len(val))
            for i, v := range val {
                res[i] = *(v.(*AgreementAcceptance))
            }
            m.SetAcceptances(res)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["file"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAgreementFile() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFile(val.(*AgreementFile))
        }
        return nil
    }
    res["files"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAgreementFileLocalization() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AgreementFileLocalization, len(val))
            for i, v := range val {
                res[i] = *(v.(*AgreementFileLocalization))
            }
            m.SetFiles(res)
        }
        return nil
    }
    res["isPerDeviceAcceptanceRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPerDeviceAcceptanceRequired(val)
        }
        return nil
    }
    res["isViewingBeforeAcceptanceRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsViewingBeforeAcceptanceRequired(val)
        }
        return nil
    }
    res["termsExpiration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTermsExpiration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTermsExpiration(val.(*TermsExpiration))
        }
        return nil
    }
    res["userReacceptRequiredFrequency"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserReacceptRequiredFrequency(val)
        }
        return nil
    }
    return res
}
func (m *Agreement) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Agreement) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAcceptances() != nil {
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
    if m.GetFiles() != nil {
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
        err = writer.WriteISODurationValue("userReacceptRequiredFrequency", m.GetUserReacceptRequiredFrequency())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcceptances sets the acceptances property value. Read-only. Information about acceptances of this agreement.
func (m *Agreement) SetAcceptances(value []AgreementAcceptance)() {
    if m != nil {
        m.acceptances = value
    }
}
// SetDisplayName sets the displayName property value. Display name of the agreement. The display name is used for internal tracking of the agreement but is not shown to end users who view the agreement.
func (m *Agreement) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetFile sets the file property value. Default PDF linked to this agreement.
func (m *Agreement) SetFile(value *AgreementFile)() {
    if m != nil {
        m.file = value
    }
}
// SetFiles sets the files property value. PDFs linked to this agreement. Note: This property is in the process of being deprecated. Use the  file property instead.
func (m *Agreement) SetFiles(value []AgreementFileLocalization)() {
    if m != nil {
        m.files = value
    }
}
// SetIsPerDeviceAcceptanceRequired sets the isPerDeviceAcceptanceRequired property value. This setting enables you to require end users to accept this agreement on every device that they are accessing it from. The end user will be required to register their device in Azure AD, if they haven't already done so.
func (m *Agreement) SetIsPerDeviceAcceptanceRequired(value *bool)() {
    if m != nil {
        m.isPerDeviceAcceptanceRequired = value
    }
}
// SetIsViewingBeforeAcceptanceRequired sets the isViewingBeforeAcceptanceRequired property value. Indicates whether the user has to expand the agreement before accepting.
func (m *Agreement) SetIsViewingBeforeAcceptanceRequired(value *bool)() {
    if m != nil {
        m.isViewingBeforeAcceptanceRequired = value
    }
}
// SetTermsExpiration sets the termsExpiration property value. Expiration schedule and frequency of agreement for all users.
func (m *Agreement) SetTermsExpiration(value *TermsExpiration)() {
    if m != nil {
        m.termsExpiration = value
    }
}
// SetUserReacceptRequiredFrequency sets the userReacceptRequiredFrequency property value. The duration after which the user must re-accept the terms of use. The value is represented in ISO 8601 format for durations.
func (m *Agreement) SetUserReacceptRequiredFrequency(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.userReacceptRequiredFrequency = value
    }
}
