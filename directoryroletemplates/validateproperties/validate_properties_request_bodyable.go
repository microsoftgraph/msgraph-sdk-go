package validateproperties

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ValidatePropertiesRequestBodyable 
type ValidatePropertiesRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetDisplayName()(*string)
    GetEntityType()(*string)
    GetMailNickname()(*string)
    GetOnBehalfOfUserId()(*string)
    SetDisplayName(value *string)()
    SetEntityType(value *string)()
    SetMailNickname(value *string)()
    SetOnBehalfOfUserId(value *string)()
}
