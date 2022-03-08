package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookTableColumnable 
type WorkbookTableColumnable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetFilter()(WorkbookFilterable)
    GetIndex()(*int32)
    GetName()(*string)
    GetValues()(Jsonable)
    SetFilter(value WorkbookFilterable)()
    SetIndex(value *int32)()
    SetName(value *string)()
    SetValues(value Jsonable)()
}
