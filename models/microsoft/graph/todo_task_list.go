package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TodoTaskList provides operations to manage the drive singleton.
type TodoTaskList struct {
    Entity
    // The name of the task list.
    displayName *string;
    // The collection of open extensions defined for the task list. Nullable.
    extensions []Extensionable;
    // True if the user is owner of the given task list.
    isOwner *bool;
    // True if the task list is shared with other users
    isShared *bool;
    // The tasks in this task list. Read-only. Nullable.
    tasks []TodoTaskable;
    // Property indicating the list name if the given list is a well-known list. Possible values are: none, defaultList, flaggedEmails, unknownFutureValue.
    wellknownListName *WellknownListName;
}
// NewTodoTaskList instantiates a new todoTaskList and sets the default values.
func NewTodoTaskList()(*TodoTaskList) {
    m := &TodoTaskList{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTodoTaskListFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTodoTaskListFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTodoTaskList(), nil
}
// GetDisplayName gets the displayName property value. The name of the task list.
func (m *TodoTaskList) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the task list. Nullable.
func (m *TodoTaskList) GetExtensions()([]Extensionable) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
        val, err := n.GetCollectionOfObjectValues(CreateExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extensionable, len(val))
            for i, v := range val {
                res[i] = v.(Extensionable)
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
        val, err := n.GetCollectionOfObjectValues(CreateTodoTaskFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TodoTaskable, len(val))
            for i, v := range val {
                res[i] = v.(TodoTaskable)
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
            m.SetWellknownListName(val.(*WellknownListName))
        }
        return nil
    }
    return res
}
// GetIsOwner gets the isOwner property value. True if the user is owner of the given task list.
func (m *TodoTaskList) GetIsOwner()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOwner
    }
}
// GetIsShared gets the isShared property value. True if the task list is shared with other users
func (m *TodoTaskList) GetIsShared()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isShared
    }
}
// GetTasks gets the tasks property value. The tasks in this task list. Read-only. Nullable.
func (m *TodoTaskList) GetTasks()([]TodoTaskable) {
    if m == nil {
        return nil
    } else {
        return m.tasks
    }
}
// GetWellknownListName gets the wellknownListName property value. Property indicating the list name if the given list is a well-known list. Possible values are: none, defaultList, flaggedEmails, unknownFutureValue.
func (m *TodoTaskList) GetWellknownListName()(*WellknownListName) {
    if m == nil {
        return nil
    } else {
        return m.wellknownListName
    }
}
func (m *TodoTaskList) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetExtensions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
    if m.GetTasks() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTasks()))
        for i, v := range m.GetTasks() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tasks", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWellknownListName() != nil {
        cast := (*m.GetWellknownListName()).String()
        err = writer.WriteStringValue("wellknownListName", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The name of the task list.
func (m *TodoTaskList) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the task list. Nullable.
func (m *TodoTaskList) SetExtensions(value []Extensionable)() {
    if m != nil {
        m.extensions = value
    }
}
// SetIsOwner sets the isOwner property value. True if the user is owner of the given task list.
func (m *TodoTaskList) SetIsOwner(value *bool)() {
    if m != nil {
        m.isOwner = value
    }
}
// SetIsShared sets the isShared property value. True if the task list is shared with other users
func (m *TodoTaskList) SetIsShared(value *bool)() {
    if m != nil {
        m.isShared = value
    }
}
// SetTasks sets the tasks property value. The tasks in this task list. Read-only. Nullable.
func (m *TodoTaskList) SetTasks(value []TodoTaskable)() {
    if m != nil {
        m.tasks = value
    }
}
// SetWellknownListName sets the wellknownListName property value. Property indicating the list name if the given list is a well-known list. Possible values are: none, defaultList, flaggedEmails, unknownFutureValue.
func (m *TodoTaskList) SetWellknownListName(value *WellknownListName)() {
    if m != nil {
        m.wellknownListName = value
    }
}
