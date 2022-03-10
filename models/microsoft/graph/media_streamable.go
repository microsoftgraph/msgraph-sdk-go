package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MediaStreamable 
type MediaStreamable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDirection()(*MediaDirection)
    GetLabel()(*string)
    GetMediaType()(*Modality)
    GetServerMuted()(*bool)
    GetSourceId()(*string)
    SetDirection(value *MediaDirection)()
    SetLabel(value *string)()
    SetMediaType(value *Modality)()
    SetServerMuted(value *bool)()
    SetSourceId(value *string)()
}
