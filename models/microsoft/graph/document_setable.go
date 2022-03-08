package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DocumentSetable 
type DocumentSetable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetAllowedContentTypes()([]ContentTypeInfoable)
    GetDefaultContents()([]DocumentSetContentable)
    GetPropagateWelcomePageChanges()(*bool)
    GetSharedColumns()([]ColumnDefinitionable)
    GetShouldPrefixNameToFile()(*bool)
    GetWelcomePageColumns()([]ColumnDefinitionable)
    GetWelcomePageUrl()(*string)
    SetAllowedContentTypes(value []ContentTypeInfoable)()
    SetDefaultContents(value []DocumentSetContentable)()
    SetPropagateWelcomePageChanges(value *bool)()
    SetSharedColumns(value []ColumnDefinitionable)()
    SetShouldPrefixNameToFile(value *bool)()
    SetWelcomePageColumns(value []ColumnDefinitionable)()
    SetWelcomePageUrl(value *string)()
}
