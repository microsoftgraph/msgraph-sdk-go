package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Todo 
type Todo struct {
    Entity
    // The task lists in the users mailbox.
    lists []TodoTaskList;
}
// NewTodo instantiates a new todo and sets the default values.
func NewTodo()(*Todo) {
    m := &Todo{
        Entity: *NewEntity(),
    }
    return m
}
// GetLists gets the lists property value. The task lists in the users mailbox.
func (m *Todo) GetLists()([]TodoTaskList) {
    if m == nil {
        return nil
    } else {
        return m.lists
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Todo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lists"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTodoTaskList() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TodoTaskList, len(val))
            for i, v := range val {
                res[i] = *(v.(*TodoTaskList))
            }
            m.SetLists(res)
        }
        return nil
    }
    return res
}
func (m *Todo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Todo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetLists() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLists()))
        for i, v := range m.GetLists() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("lists", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLists sets the lists property value. The task lists in the users mailbox.
func (m *Todo) SetLists(value []TodoTaskList)() {
    if m != nil {
        m.lists = value
    }
}
