package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookSortFieldable 
type WorkbookSortFieldable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetAscending()(*bool)
    GetColor()(*string)
    GetDataOption()(*string)
    GetIcon()(WorkbookIconable)
    GetKey()(*int32)
    GetSortOn()(*string)
    SetAscending(value *bool)()
    SetColor(value *string)()
    SetDataOption(value *string)()
    SetIcon(value WorkbookIconable)()
    SetKey(value *int32)()
    SetSortOn(value *string)()
}
