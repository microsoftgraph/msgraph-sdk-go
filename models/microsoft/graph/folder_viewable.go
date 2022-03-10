package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// FolderViewable 
type FolderViewable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetSortBy()(*string)
    GetSortOrder()(*string)
    GetViewType()(*string)
    SetSortBy(value *string)()
    SetSortOrder(value *string)()
    SetViewType(value *string)()
}
