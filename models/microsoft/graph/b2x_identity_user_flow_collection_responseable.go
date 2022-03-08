package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// B2xIdentityUserFlowCollectionResponseable 
type B2xIdentityUserFlowCollectionResponseable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetNextLink()(*string)
    GetValue()([]B2xIdentityUserFlowable)
    SetNextLink(value *string)()
    SetValue(value []B2xIdentityUserFlowable)()
}
