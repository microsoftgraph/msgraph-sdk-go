package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AgreementFileProperties struct {
    Entity
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    displayName *string;
    fileData *AgreementFileData;
    fileName *string;
    isDefault *bool;
    isMajorVersion *bool;
    language *string;
}
func NewAgreementFileProperties()(*AgreementFileProperties) {
    m := &AgreementFileProperties{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AgreementFileProperties) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *AgreementFileProperties) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *AgreementFileProperties) GetFileData()(*AgreementFileData) {
    if m == nil {
        return nil
    } else {
        return m.fileData
    }
}
func (m *AgreementFileProperties) GetFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileName
    }
}
func (m *AgreementFileProperties) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
func (m *AgreementFileProperties) GetIsMajorVersion()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMajorVersion
    }
}
func (m *AgreementFileProperties) GetLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.language
    }
}
func (m *AgreementFileProperties) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
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
    res["fileData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAgreementFileData() })
        if err != nil {
            return err
        }
        m.SetFileData(val.(*AgreementFileData))
        return nil
    }
    res["fileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFileName(val)
        return nil
    }
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDefault(val)
        return nil
    }
    res["isMajorVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsMajorVersion(val)
        return nil
    }
    res["language"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLanguage(val)
        return nil
    }
    return res
}
func (m *AgreementFileProperties) IsNil()(bool) {
    return m == nil
}
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
func (m *AgreementFileProperties) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *AgreementFileProperties) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *AgreementFileProperties) SetFileData(value *AgreementFileData)() {
    m.fileData = value
}
func (m *AgreementFileProperties) SetFileName(value *string)() {
    m.fileName = value
}
func (m *AgreementFileProperties) SetIsDefault(value *bool)() {
    m.isDefault = value
}
func (m *AgreementFileProperties) SetIsMajorVersion(value *bool)() {
    m.isMajorVersion = value
}
func (m *AgreementFileProperties) SetLanguage(value *string)() {
    m.language = value
}
