package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AutomaticRepliesMailTipsable 
type AutomaticRepliesMailTipsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetMessage()(*string)
    GetMessageLanguage()(LocaleInfoable)
    GetScheduledEndTime()(DateTimeTimeZoneable)
    GetScheduledStartTime()(DateTimeTimeZoneable)
    SetMessage(value *string)()
    SetMessageLanguage(value LocaleInfoable)()
    SetScheduledEndTime(value DateTimeTimeZoneable)()
    SetScheduledStartTime(value DateTimeTimeZoneable)()
}
