package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/externalconnectors"
)

// Propertyable 
type Propertyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetAliases()([]string)
    GetIsQueryable()(*bool)
    GetIsRefinable()(*bool)
    GetIsRetrievable()(*bool)
    GetIsSearchable()(*bool)
    GetLabels()([]i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.Label)
    GetName()(*string)
    GetType()(*i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.PropertyType)
    SetAliases(value []string)()
    SetIsQueryable(value *bool)()
    SetIsRefinable(value *bool)()
    SetIsRetrievable(value *bool)()
    SetIsSearchable(value *bool)()
    SetLabels(value []i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.Label)()
    SetName(value *string)()
    SetType(value *i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.PropertyType)()
}
