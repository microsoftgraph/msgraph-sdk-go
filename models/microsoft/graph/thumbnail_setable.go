package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ThumbnailSetable 
type ThumbnailSetable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetLarge()(Thumbnailable)
    GetMedium()(Thumbnailable)
    GetSmall()(Thumbnailable)
    GetSource()(Thumbnailable)
    SetLarge(value Thumbnailable)()
    SetMedium(value Thumbnailable)()
    SetSmall(value Thumbnailable)()
    SetSource(value Thumbnailable)()
}
