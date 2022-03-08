package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RelatedContactable 
type RelatedContactable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetAccessConsent()(*bool)
    GetDisplayName()(*string)
    GetEmailAddress()(*string)
    GetMobilePhone()(*string)
    GetRelationship()(*ContactRelationship)
    SetAccessConsent(value *bool)()
    SetDisplayName(value *string)()
    SetEmailAddress(value *string)()
    SetMobilePhone(value *string)()
    SetRelationship(value *ContactRelationship)()
}
