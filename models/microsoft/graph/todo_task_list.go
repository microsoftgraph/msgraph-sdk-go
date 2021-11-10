package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TodoTaskList struct {
    Entity
    // The name of the task list.
    displayName *string;
    // The collection of open extensions defined for the task list. Nullable.
    extensions []Extension;
    // True if the user is owner of the given task list.
    isOwner *bool;
    // True if the task list is shared with other users
    isShared *bool;
    // The tasks in this task list. Read-only. Nullable.
    tasks []TodoTask;
    // Property indicating the list name if the given list is a well-known list. Possible values are: none, defaultList, flaggedEmails, unknownFutureValue.
    wellknownListName *WellknownListName;
}
// Instantiates a new todoTaskList and sets the default values.
func NewTodoTaskList()(*TodoTaskList) {
    m := &TodoTaskList{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayName property value. The name of the task list.
func (m *TodoTaskList) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the extensions property value. The collection of open extensions defined for the task list. Nullable.
func (m *TodoTaskList) GetExtensions()([]Extension) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// Gets the isOwner property value. True if the user is owner of the given task list.
func (m *TodoTaskList) GetIsOwner()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOwner
    }
}
// Gets the isShared property value. True if the task list is shared with other users
func (m *TodoTaskList) GetIsShared()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isShared
    }
}
// Gets the tasks property value. The tasks in this task list. Read-only. Nullable.
func (m *TodoTaskList) GetTasks()([]TodoTask) {
    if m == nil {
        return nil
    } else {
        return m.tasks
    }
}
// Gets the wellknownListName property value. Property indicating the list name if the given list is a well-known list. Possible values are: none, defaultList, flaggedEmails, unknownFutureValue.
func (m *TodoTaskList) GetWellknownListName()(*WellknownListName) {
    if m == nil {
        return nil
    } else {
        return m.wellknownListName
    }
}
// The deserialization information for the current model
func (m *TodoTaskList) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExtension() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extension, len(val))
            for i, v := range val {
                res[i] = *(v.(*Extension))
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["isOwner"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOwner(val)
        }
        return nil
    }
    res["isShared"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsShared(val)
        }
        return nil
    }
    res["tasks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTodoTask() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TodoTask, len(val))
            for i, v := range val {
                res[i] = *(v.(*TodoTask))
            }
            m.SetTasks(res)
        }
        return nil
    }
    res["wellknownListName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWellknownListName)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(WellknownListName)
            m.SetWellknownListName(&cast)
        }
        return nil
    }
    return res
}
func (m *TodoTaskList) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TodoTaskList) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isOwner", m.GetIsOwner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isShared", m.GetIsShared())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTasks()))
        for i, v := range m.GetTasks() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tasks", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWellknownListName() != nil {
        cast := m.GetWellknownListName().String()
        err = writer.WriteStringValue("wellknownListName", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the displayName property value. The name of the task list.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *TodoTaskList) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the extensions property value. The collection of open extensions defined for the task list. Nullable.
// Parameters:
//  - value : Value to set for the extensions property.
func (m *TodoTaskList) SetExtensions(value []Extension)() {
    m.extensions = value
}
// Sets the isOwner property value. True if the user is owner of the given task list.
// Parameters:
//  - value : Value to set for the isOwner property.
func (m *TodoTaskList) SetIsOwner(value *bool)() {
    m.isOwner = value
}
// Sets the isShared property value. True if the task list is shared with other users
// Parameters:
//  - value : Value to set for the isShared property.
func (m *TodoTaskList) SetIsShared(value *bool)() {
    m.isShared = value
}
// Sets the tasks property value. The tasks in this task list. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the tasks property.
func (m *TodoTaskList) SetTasks(value []TodoTask)() {
    m.tasks = value
}
// Sets the wellknownListName property value. Property indicating the list name if the given list is a well-known list. Possible values are: none, defaultList, flaggedEmails, unknownFutureValue.
// Parameters:
//  - value : Value to set for the wellknownListName property.
func (m *TodoTaskList) SetWellknownListName(value *WellknownListName)() {
    m.wellknownListName = value
}
