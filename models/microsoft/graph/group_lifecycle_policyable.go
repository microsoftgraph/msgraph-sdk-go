package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GroupLifecyclePolicyable 
type GroupLifecyclePolicyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetAlternateNotificationEmails()(*string)
    GetGroupLifetimeInDays()(*int32)
    GetManagedGroupTypes()(*string)
    SetAlternateNotificationEmails(value *string)()
    SetGroupLifetimeInDays(value *int32)()
    SetManagedGroupTypes(value *string)()
}
