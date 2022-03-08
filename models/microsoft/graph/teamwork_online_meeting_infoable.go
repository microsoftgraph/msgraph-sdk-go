package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkOnlineMeetingInfoable 
type TeamworkOnlineMeetingInfoable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetCalendarEventId()(*string)
    GetJoinWebUrl()(*string)
    GetOrganizer()(TeamworkUserIdentityable)
    SetCalendarEventId(value *string)()
    SetJoinWebUrl(value *string)()
    SetOrganizer(value TeamworkUserIdentityable)()
}
