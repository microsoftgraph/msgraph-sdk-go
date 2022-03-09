package externalconnectors

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Propertyable 
type Propertyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAliases()([]string)
    GetIsQueryable()(*bool)
    GetIsRefinable()(*bool)
    GetIsRetrievable()(*bool)
    GetIsSearchable()(*bool)
    GetLabels()([]Label)
    GetName()(*string)
    GetType()(*PropertyType)
    SetAliases(value []string)()
    SetIsQueryable(value *bool)()
    SetIsRefinable(value *bool)()
    SetIsRetrievable(value *bool)()
    SetIsSearchable(value *bool)()
    SetLabels(value []Label)()
    SetName(value *string)()
    SetType(value *PropertyType)()
}
