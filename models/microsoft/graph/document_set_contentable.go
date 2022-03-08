package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DocumentSetContentable 
type DocumentSetContentable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetContentType()(ContentTypeInfoable)
    GetFileName()(*string)
    GetFolderName()(*string)
    SetContentType(value ContentTypeInfoable)()
    SetFileName(value *string)()
    SetFolderName(value *string)()
}
