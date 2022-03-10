package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MailboxSettingsable 
type MailboxSettingsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetArchiveFolder()(*string)
    GetAutomaticRepliesSetting()(AutomaticRepliesSettingable)
    GetDateFormat()(*string)
    GetDelegateMeetingMessageDeliveryOptions()(*DelegateMeetingMessageDeliveryOptions)
    GetLanguage()(LocaleInfoable)
    GetTimeFormat()(*string)
    GetTimeZone()(*string)
    GetUserPurpose()(*UserPurpose)
    GetWorkingHours()(WorkingHoursable)
    SetArchiveFolder(value *string)()
    SetAutomaticRepliesSetting(value AutomaticRepliesSettingable)()
    SetDateFormat(value *string)()
    SetDelegateMeetingMessageDeliveryOptions(value *DelegateMeetingMessageDeliveryOptions)()
    SetLanguage(value LocaleInfoable)()
    SetTimeFormat(value *string)()
    SetTimeZone(value *string)()
    SetUserPurpose(value *UserPurpose)()
    SetWorkingHours(value WorkingHoursable)()
}
