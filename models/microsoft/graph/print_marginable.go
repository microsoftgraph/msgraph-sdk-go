package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintMarginable 
type PrintMarginable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetBottom()(*int32)
    GetLeft()(*int32)
    GetRight()(*int32)
    GetTop()(*int32)
    SetBottom(value *int32)()
    SetLeft(value *int32)()
    SetRight(value *int32)()
    SetTop(value *int32)()
}
