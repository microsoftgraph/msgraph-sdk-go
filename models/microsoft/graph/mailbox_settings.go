package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MailboxSettings 
type MailboxSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Folder ID of an archive folder for the user.
    archiveFolder *string;
    // Configuration settings to automatically notify the sender of an incoming email with a message from the signed-in user.
    automaticRepliesSetting *AutomaticRepliesSetting;
    // The date format for the user's mailbox.
    dateFormat *string;
    // If the user has a calendar delegate, this specifies whether the delegate, mailbox owner, or both receive meeting messages and meeting responses. Possible values are: sendToDelegateAndInformationToPrincipal, sendToDelegateAndPrincipal, sendToDelegateOnly.
    delegateMeetingMessageDeliveryOptions *DelegateMeetingMessageDeliveryOptions;
    // The locale information for the user, including the preferred language and country/region.
    language *LocaleInfo;
    // The time format for the user's mailbox.
    timeFormat *string;
    // The default time zone for the user's mailbox.
    timeZone *string;
    // The days of the week and hours in a specific time zone that the user works.
    workingHours *WorkingHours;
}
// NewMailboxSettings instantiates a new mailboxSettings and sets the default values.
func NewMailboxSettings()(*MailboxSettings) {
    m := &MailboxSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MailboxSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetArchiveFolder gets the archiveFolder property value. Folder ID of an archive folder for the user.
func (m *MailboxSettings) GetArchiveFolder()(*string) {
    if m == nil {
        return nil
    } else {
        return m.archiveFolder
    }
}
// GetAutomaticRepliesSetting gets the automaticRepliesSetting property value. Configuration settings to automatically notify the sender of an incoming email with a message from the signed-in user.
func (m *MailboxSettings) GetAutomaticRepliesSetting()(*AutomaticRepliesSetting) {
    if m == nil {
        return nil
    } else {
        return m.automaticRepliesSetting
    }
}
// GetDateFormat gets the dateFormat property value. The date format for the user's mailbox.
func (m *MailboxSettings) GetDateFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dateFormat
    }
}
// GetDelegateMeetingMessageDeliveryOptions gets the delegateMeetingMessageDeliveryOptions property value. If the user has a calendar delegate, this specifies whether the delegate, mailbox owner, or both receive meeting messages and meeting responses. Possible values are: sendToDelegateAndInformationToPrincipal, sendToDelegateAndPrincipal, sendToDelegateOnly.
func (m *MailboxSettings) GetDelegateMeetingMessageDeliveryOptions()(*DelegateMeetingMessageDeliveryOptions) {
    if m == nil {
        return nil
    } else {
        return m.delegateMeetingMessageDeliveryOptions
    }
}
// GetLanguage gets the language property value. The locale information for the user, including the preferred language and country/region.
func (m *MailboxSettings) GetLanguage()(*LocaleInfo) {
    if m == nil {
        return nil
    } else {
        return m.language
    }
}
// GetTimeFormat gets the timeFormat property value. The time format for the user's mailbox.
func (m *MailboxSettings) GetTimeFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeFormat
    }
}
// GetTimeZone gets the timeZone property value. The default time zone for the user's mailbox.
func (m *MailboxSettings) GetTimeZone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeZone
    }
}
// GetWorkingHours gets the workingHours property value. The days of the week and hours in a specific time zone that the user works.
func (m *MailboxSettings) GetWorkingHours()(*WorkingHours) {
    if m == nil {
        return nil
    } else {
        return m.workingHours
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MailboxSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["archiveFolder"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArchiveFolder(val)
        }
        return nil
    }
    res["automaticRepliesSetting"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAutomaticRepliesSetting() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutomaticRepliesSetting(val.(*AutomaticRepliesSetting))
        }
        return nil
    }
    res["dateFormat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateFormat(val)
        }
        return nil
    }
    res["delegateMeetingMessageDeliveryOptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDelegateMeetingMessageDeliveryOptions)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DelegateMeetingMessageDeliveryOptions)
            m.SetDelegateMeetingMessageDeliveryOptions(&cast)
        }
        return nil
    }
    res["language"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocaleInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val.(*LocaleInfo))
        }
        return nil
    }
    res["timeFormat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeFormat(val)
        }
        return nil
    }
    res["timeZone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZone(val)
        }
        return nil
    }
    res["workingHours"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkingHours() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkingHours(val.(*WorkingHours))
        }
        return nil
    }
    return res
}
func (m *MailboxSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MailboxSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetArchiveFolder sets the archiveFolder property value. Folder ID of an archive folder for the user.
func (m *MailboxSettings) SetArchiveFolder(value *string)() {
    m.archiveFolder = value
}
// SetAutomaticRepliesSetting sets the automaticRepliesSetting property value. Configuration settings to automatically notify the sender of an incoming email with a message from the signed-in user.
func (m *MailboxSettings) SetAutomaticRepliesSetting(value *AutomaticRepliesSetting)() {
    m.automaticRepliesSetting = value
}
// SetDateFormat sets the dateFormat property value. The date format for the user's mailbox.
func (m *MailboxSettings) SetDateFormat(value *string)() {
    m.dateFormat = value
}
// SetDelegateMeetingMessageDeliveryOptions sets the delegateMeetingMessageDeliveryOptions property value. If the user has a calendar delegate, this specifies whether the delegate, mailbox owner, or both receive meeting messages and meeting responses. Possible values are: sendToDelegateAndInformationToPrincipal, sendToDelegateAndPrincipal, sendToDelegateOnly.
func (m *MailboxSettings) SetDelegateMeetingMessageDeliveryOptions(value *DelegateMeetingMessageDeliveryOptions)() {
    m.delegateMeetingMessageDeliveryOptions = value
}
// SetLanguage sets the language property value. The locale information for the user, including the preferred language and country/region.
func (m *MailboxSettings) SetLanguage(value *LocaleInfo)() {
    m.language = value
}
// SetTimeFormat sets the timeFormat property value. The time format for the user's mailbox.
func (m *MailboxSettings) SetTimeFormat(value *string)() {
    m.timeFormat = value
}
// SetTimeZone sets the timeZone property value. The default time zone for the user's mailbox.
func (m *MailboxSettings) SetTimeZone(value *string)() {
    m.timeZone = value
}
// SetWorkingHours sets the workingHours property value. The days of the week and hours in a specific time zone that the user works.
func (m *MailboxSettings) SetWorkingHours(value *WorkingHours)() {
    m.workingHours = value
}
