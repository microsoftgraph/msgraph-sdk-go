package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AgreementFileProperties 
type AgreementFileProperties struct {
    Entity
    // The date time representing when the file was created.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Localized display name of the policy file of an agreement. The localized display name is shown to end users who view the agreement.
    displayName *string;
    // Data that represents the terms of use PDF document. Read-only.
    fileData *AgreementFileData;
    // Name of the agreement file (for example, TOU.pdf). Read-only.
    fileName *string;
    // If none of the languages matches the client preference, indicates whether this is the default agreement file . If none of the files are marked as default, the first one is treated as the default. Read-only.
    isDefault *bool;
    // Indicates whether the agreement file is a major version update. Major version updates invalidate the agreement's acceptances on the corresponding language.
    isMajorVersion *bool;
    // The language of the agreement file in the format 'languagecode2-country/regioncode2'. 'languagecode2' is a lowercase two-letter code derived from ISO 639-1, while 'country/regioncode2' is derived from ISO 3166 and usually consists of two uppercase letters, or a BCP-47 language tag. For example, U.S. English is en-US. Read-only.
    language *string;
}
// NewAgreementFileProperties instantiates a new agreementFileProperties and sets the default values.
func NewAgreementFileProperties()(*AgreementFileProperties) {
    m := &AgreementFileProperties{
        Entity: *NewEntity(),
    }
    return m
}
// GetCreatedDateTime gets the createdDateTime property value. The date time representing when the file was created.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AgreementFileProperties) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDisplayName gets the displayName property value. Localized display name of the policy file of an agreement. The localized display name is shown to end users who view the agreement.
func (m *AgreementFileProperties) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFileData gets the fileData property value. Data that represents the terms of use PDF document. Read-only.
func (m *AgreementFileProperties) GetFileData()(*AgreementFileData) {
    if m == nil {
        return nil
    } else {
        return m.fileData
    }
}
// GetFileName gets the fileName property value. Name of the agreement file (for example, TOU.pdf). Read-only.
func (m *AgreementFileProperties) GetFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileName
    }
}
// GetIsDefault gets the isDefault property value. If none of the languages matches the client preference, indicates whether this is the default agreement file . If none of the files are marked as default, the first one is treated as the default. Read-only.
func (m *AgreementFileProperties) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// GetIsMajorVersion gets the isMajorVersion property value. Indicates whether the agreement file is a major version update. Major version updates invalidate the agreement's acceptances on the corresponding language.
func (m *AgreementFileProperties) GetIsMajorVersion()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMajorVersion
    }
}
// GetLanguage gets the language property value. The language of the agreement file in the format 'languagecode2-country/regioncode2'. 'languagecode2' is a lowercase two-letter code derived from ISO 639-1, while 'country/regioncode2' is derived from ISO 3166 and usually consists of two uppercase letters, or a BCP-47 language tag. For example, U.S. English is en-US. Read-only.
func (m *AgreementFileProperties) GetLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.language
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AgreementFileProperties) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
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
    res["fileData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAgreementFileData() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileData(val.(*AgreementFileData))
        }
        return nil
    }
    res["fileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
        }
        return nil
    }
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["isMajorVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMajorVersion(val)
        }
        return nil
    }
    res["language"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val)
        }
        return nil
    }
    return res
}
func (m *AgreementFileProperties) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AgreementFileProperties) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
        err = writer.WriteObjectValue("fileData", m.GetFileData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMajorVersion", m.GetIsMajorVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("language", m.GetLanguage())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The date time representing when the file was created.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AgreementFileProperties) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDisplayName sets the displayName property value. Localized display name of the policy file of an agreement. The localized display name is shown to end users who view the agreement.
func (m *AgreementFileProperties) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetFileData sets the fileData property value. Data that represents the terms of use PDF document. Read-only.
func (m *AgreementFileProperties) SetFileData(value *AgreementFileData)() {
    if m != nil {
        m.fileData = value
    }
}
// SetFileName sets the fileName property value. Name of the agreement file (for example, TOU.pdf). Read-only.
func (m *AgreementFileProperties) SetFileName(value *string)() {
    if m != nil {
        m.fileName = value
    }
}
// SetIsDefault sets the isDefault property value. If none of the languages matches the client preference, indicates whether this is the default agreement file . If none of the files are marked as default, the first one is treated as the default. Read-only.
func (m *AgreementFileProperties) SetIsDefault(value *bool)() {
    if m != nil {
        m.isDefault = value
    }
}
// SetIsMajorVersion sets the isMajorVersion property value. Indicates whether the agreement file is a major version update. Major version updates invalidate the agreement's acceptances on the corresponding language.
func (m *AgreementFileProperties) SetIsMajorVersion(value *bool)() {
    if m != nil {
        m.isMajorVersion = value
    }
}
// SetLanguage sets the language property value. The language of the agreement file in the format 'languagecode2-country/regioncode2'. 'languagecode2' is a lowercase two-letter code derived from ISO 639-1, while 'country/regioncode2' is derived from ISO 3166 and usually consists of two uppercase letters, or a BCP-47 language tag. For example, U.S. English is en-US. Read-only.
func (m *AgreementFileProperties) SetLanguage(value *string)() {
    if m != nil {
        m.language = value
    }
}
