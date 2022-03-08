package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookNamedItemable 
type WorkbookNamedItemable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetComment()(*string)
    GetName()(*string)
    GetScope()(*string)
    GetType()(*string)
    GetValue()(Jsonable)
    GetVisible()(*bool)
    GetWorksheet()(WorkbookWorksheetable)
    SetComment(value *string)()
    SetName(value *string)()
    SetScope(value *string)()
    SetType(value *string)()
    SetValue(value Jsonable)()
    SetVisible(value *bool)()
    SetWorksheet(value WorkbookWorksheetable)()
}
