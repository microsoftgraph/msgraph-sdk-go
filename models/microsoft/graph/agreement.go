package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Agreement provides operations to manage the collection of agreement entities.
type Agreement struct {
    Entity
    // Read-only. Information about acceptances of this agreement.
    acceptances []AgreementAcceptanceable;
    // Display name of the agreement. The display name is used for internal tracking of the agreement but is not shown to end users who view the agreement. Supports $filter (eq).
    displayName *string;
    // Default PDF linked to this agreement.
    file AgreementFileable;
    // PDFs linked to this agreement. This property is in the process of being deprecated. Use the  file property instead. Supports $expand.
    files []AgreementFileLocalizationable;
    // Indicates whether end users are required to accept this agreement on every device that they access it from. The end user is required to register their device in Azure AD, if they haven't already done so. Supports $filter (eq).
    isPerDeviceAcceptanceRequired *bool;
    // Indicates whether the user has to expand the agreement before accepting. Supports $filter (eq).
    isViewingBeforeAcceptanceRequired *bool;
    // Expiration schedule and frequency of agreement for all users. Supports $filter (eq).
    termsExpiration TermsExpirationable;
    // The duration after which the user must re-accept the terms of use. The value is represented in ISO 8601 format for durations. Supports $filter (eq).
    userReacceptRequiredFrequency *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
}
// NewAgreement instantiates a new agreement and sets the default values.
func NewAgreement()(*Agreement) {
    m := &Agreement{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAgreementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAgreementFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAgreement(), nil
}
// GetAcceptances gets the acceptances property value. Read-only. Information about acceptances of this agreement.
func (m *Agreement) GetAcceptances()([]AgreementAcceptanceable) {
    if m == nil {
        return nil
    } else {
        return m.acceptances
    }
}
// GetDisplayName gets the displayName property value. Display name of the agreement. The display name is used for internal tracking of the agreement but is not shown to end users who view the agreement. Supports $filter (eq).
func (m *Agreement) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Agreement) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acceptances"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAgreementAcceptanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AgreementAcceptanceable, len(val))
            for i, v := range val {
                res[i] = v.(AgreementAcceptanceable)
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
        val, err := n.GetObjectValue(CreateAgreementFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFile(val.(AgreementFileable))
        }
        return nil
    }
    res["files"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAgreementFileLocalizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AgreementFileLocalizationable, len(val))
            for i, v := range val {
                res[i] = v.(AgreementFileLocalizationable)
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
        val, err := n.GetObjectValue(CreateTermsExpirationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTermsExpiration(val.(TermsExpirationable))
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
// GetFile gets the file property value. Default PDF linked to this agreement.
func (m *Agreement) GetFile()(AgreementFileable) {
    if m == nil {
        return nil
    } else {
        return m.file
    }
}
// GetFiles gets the files property value. PDFs linked to this agreement. This property is in the process of being deprecated. Use the  file property instead. Supports $expand.
func (m *Agreement) GetFiles()([]AgreementFileLocalizationable) {
    if m == nil {
        return nil
    } else {
        return m.files
    }
}
// GetIsPerDeviceAcceptanceRequired gets the isPerDeviceAcceptanceRequired property value. Indicates whether end users are required to accept this agreement on every device that they access it from. The end user is required to register their device in Azure AD, if they haven't already done so. Supports $filter (eq).
func (m *Agreement) GetIsPerDeviceAcceptanceRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPerDeviceAcceptanceRequired
    }
}
// GetIsViewingBeforeAcceptanceRequired gets the isViewingBeforeAcceptanceRequired property value. Indicates whether the user has to expand the agreement before accepting. Supports $filter (eq).
func (m *Agreement) GetIsViewingBeforeAcceptanceRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isViewingBeforeAcceptanceRequired
    }
}
// GetTermsExpiration gets the termsExpiration property value. Expiration schedule and frequency of agreement for all users. Supports $filter (eq).
func (m *Agreement) GetTermsExpiration()(TermsExpirationable) {
    if m == nil {
        return nil
    } else {
        return m.termsExpiration
    }
}
// GetUserReacceptRequiredFrequency gets the userReacceptRequiredFrequency property value. The duration after which the user must re-accept the terms of use. The value is represented in ISO 8601 format for durations. Supports $filter (eq).
func (m *Agreement) GetUserReacceptRequiredFrequency()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.userReacceptRequiredFrequency
    }
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *Agreement) SetAcceptances(value []AgreementAcceptanceable)() {
    if m != nil {
        m.acceptances = value
    }
}
// SetDisplayName sets the displayName property value. Display name of the agreement. The display name is used for internal tracking of the agreement but is not shown to end users who view the agreement. Supports $filter (eq).
func (m *Agreement) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetFile sets the file property value. Default PDF linked to this agreement.
func (m *Agreement) SetFile(value AgreementFileable)() {
    if m != nil {
        m.file = value
    }
}
// SetFiles sets the files property value. PDFs linked to this agreement. This property is in the process of being deprecated. Use the  file property instead. Supports $expand.
func (m *Agreement) SetFiles(value []AgreementFileLocalizationable)() {
    if m != nil {
        m.files = value
    }
}
// SetIsPerDeviceAcceptanceRequired sets the isPerDeviceAcceptanceRequired property value. Indicates whether end users are required to accept this agreement on every device that they access it from. The end user is required to register their device in Azure AD, if they haven't already done so. Supports $filter (eq).
func (m *Agreement) SetIsPerDeviceAcceptanceRequired(value *bool)() {
    if m != nil {
        m.isPerDeviceAcceptanceRequired = value
    }
}
// SetIsViewingBeforeAcceptanceRequired sets the isViewingBeforeAcceptanceRequired property value. Indicates whether the user has to expand the agreement before accepting. Supports $filter (eq).
func (m *Agreement) SetIsViewingBeforeAcceptanceRequired(value *bool)() {
    if m != nil {
        m.isViewingBeforeAcceptanceRequired = value
    }
}
// SetTermsExpiration sets the termsExpiration property value. Expiration schedule and frequency of agreement for all users. Supports $filter (eq).
func (m *Agreement) SetTermsExpiration(value TermsExpirationable)() {
    if m != nil {
        m.termsExpiration = value
    }
}
// SetUserReacceptRequiredFrequency sets the userReacceptRequiredFrequency property value. The duration after which the user must re-accept the terms of use. The value is represented in ISO 8601 format for durations. Supports $filter (eq).
func (m *Agreement) SetUserReacceptRequiredFrequency(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.userReacceptRequiredFrequency = value
    }
}
