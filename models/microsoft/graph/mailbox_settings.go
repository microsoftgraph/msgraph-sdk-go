package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MailboxSettings struct {
    additionalData map[string]interface{};
    archiveFolder *string;
    automaticRepliesSetting *AutomaticRepliesSetting;
    dateFormat *string;
    delegateMeetingMessageDeliveryOptions *DelegateMeetingMessageDeliveryOptions;
    language *LocaleInfo;
    timeFormat *string;
    timeZone *string;
    workingHours *WorkingHours;
}
func NewMailboxSettings()(*MailboxSettings) {
    m := &MailboxSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MailboxSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MailboxSettings) GetArchiveFolder()(*string) {
    if m == nil {
        return nil
    } else {
        return m.archiveFolder
    }
}
func (m *MailboxSettings) GetAutomaticRepliesSetting()(*AutomaticRepliesSetting) {
    if m == nil {
        return nil
    } else {
        return m.automaticRepliesSetting
    }
}
func (m *MailboxSettings) GetDateFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dateFormat
    }
}
func (m *MailboxSettings) GetDelegateMeetingMessageDeliveryOptions()(*DelegateMeetingMessageDeliveryOptions) {
    if m == nil {
        return nil
    } else {
        return m.delegateMeetingMessageDeliveryOptions
    }
}
func (m *MailboxSettings) GetLanguage()(*LocaleInfo) {
    if m == nil {
        return nil
    } else {
        return m.language
    }
}
func (m *MailboxSettings) GetTimeFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeFormat
    }
}
func (m *MailboxSettings) GetTimeZone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeZone
    }
}
func (m *MailboxSettings) GetWorkingHours()(*WorkingHours) {
    if m == nil {
        return nil
    } else {
        return m.workingHours
    }
}
func (m *MailboxSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["archiveFolder"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetArchiveFolder(val)
        return nil
    }
    res["automaticRepliesSetting"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAutomaticRepliesSetting() })
        if err != nil {
            return err
        }
        m.SetAutomaticRepliesSetting(val.(*AutomaticRepliesSetting))
        return nil
    }
    res["dateFormat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDateFormat(val)
        return nil
    }
    res["delegateMeetingMessageDeliveryOptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDelegateMeetingMessageDeliveryOptions)
        if err != nil {
            return err
        }
        cast := val.(DelegateMeetingMessageDeliveryOptions)
        m.SetDelegateMeetingMessageDeliveryOptions(&cast)
        return nil
    }
    res["language"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocaleInfo() })
        if err != nil {
            return err
        }
        m.SetLanguage(val.(*LocaleInfo))
        return nil
    }
    res["timeFormat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTimeFormat(val)
        return nil
    }
    res["timeZone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTimeZone(val)
        return nil
    }
    res["workingHours"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkingHours() })
        if err != nil {
            return err
        }
        m.SetWorkingHours(val.(*WorkingHours))
        return nil
    }
    return res
}
func (m *MailboxSettings) IsNil()(bool) {
    return m == nil
}
func (m *MailboxSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("archiveFolder", m.GetArchiveFolder())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("automaticRepliesSetting", m.GetAutomaticRepliesSetting())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dateFormat", m.GetDateFormat())
        if err != nil {
            return err
        }
    }
    if m.GetDelegateMeetingMessageDeliveryOptions() != nil {
        cast := m.GetDelegateMeetingMessageDeliveryOptions().String()
        err := writer.WriteStringValue("delegateMeetingMessageDeliveryOptions", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("language", m.GetLanguage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timeFormat", m.GetTimeFormat())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timeZone", m.GetTimeZone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("workingHours", m.GetWorkingHours())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *MailboxSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MailboxSettings) SetArchiveFolder(value *string)() {
    m.archiveFolder = value
}
func (m *MailboxSettings) SetAutomaticRepliesSetting(value *AutomaticRepliesSetting)() {
    m.automaticRepliesSetting = value
}
func (m *MailboxSettings) SetDateFormat(value *string)() {
    m.dateFormat = value
}
func (m *MailboxSettings) SetDelegateMeetingMessageDeliveryOptions(value *DelegateMeetingMessageDeliveryOptions)() {
    m.delegateMeetingMessageDeliveryOptions = value
}
func (m *MailboxSettings) SetLanguage(value *LocaleInfo)() {
    m.language = value
}
func (m *MailboxSettings) SetTimeFormat(value *string)() {
    m.timeFormat = value
}
func (m *MailboxSettings) SetTimeZone(value *string)() {
    m.timeZone = value
}
func (m *MailboxSettings) SetWorkingHours(value *WorkingHours)() {
    m.workingHours = value
}
