package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AutomaticRepliesSettingable 
type AutomaticRepliesSettingable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetExternalAudience()(*ExternalAudienceScope)
    GetExternalReplyMessage()(*string)
    GetInternalReplyMessage()(*string)
    GetScheduledEndDateTime()(DateTimeTimeZoneable)
    GetScheduledStartDateTime()(DateTimeTimeZoneable)
    GetStatus()(*AutomaticRepliesStatus)
    SetExternalAudience(value *ExternalAudienceScope)()
    SetExternalReplyMessage(value *string)()
    SetInternalReplyMessage(value *string)()
    SetScheduledEndDateTime(value DateTimeTimeZoneable)()
    SetScheduledStartDateTime(value DateTimeTimeZoneable)()
    SetStatus(value *AutomaticRepliesStatus)()
}
