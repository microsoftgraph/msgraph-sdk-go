package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TodoTaskListable 
type TodoTaskListable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetDisplayName()(*string)
    GetExtensions()([]Extensionable)
    GetIsOwner()(*bool)
    GetIsShared()(*bool)
    GetTasks()([]TodoTaskable)
    GetWellknownListName()(*WellknownListName)
    SetDisplayName(value *string)()
    SetExtensions(value []Extensionable)()
    SetIsOwner(value *bool)()
    SetIsShared(value *bool)()
    SetTasks(value []TodoTaskable)()
    SetWellknownListName(value *WellknownListName)()
}
