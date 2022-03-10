package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ResourceVisualizationable 
type ResourceVisualizationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetContainerDisplayName()(*string)
    GetContainerType()(*string)
    GetContainerWebUrl()(*string)
    GetMediaType()(*string)
    GetPreviewImageUrl()(*string)
    GetPreviewText()(*string)
    GetTitle()(*string)
    GetType()(*string)
    SetContainerDisplayName(value *string)()
    SetContainerType(value *string)()
    SetContainerWebUrl(value *string)()
    SetMediaType(value *string)()
    SetPreviewImageUrl(value *string)()
    SetPreviewText(value *string)()
    SetTitle(value *string)()
    SetType(value *string)()
}
