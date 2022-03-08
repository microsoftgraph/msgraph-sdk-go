package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Todo provides operations to manage the drive singleton.
type Todo struct {
    Entity
    // The task lists in the users mailbox.
    lists []TodoTaskListable;
}
// NewTodo instantiates a new todo and sets the default values.
func NewTodo()(*Todo) {
    m := &Todo{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTodoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTodoFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTodo(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Todo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lists"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTodoTaskListFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TodoTaskListable, len(val))
            for i, v := range val {
                res[i] = v.(TodoTaskListable)
            }
            m.SetLists(res)
        }
        return nil
    }
    return res
}
// GetLists gets the lists property value. The task lists in the users mailbox.
func (m *Todo) GetLists()([]TodoTaskListable) {
    if m == nil {
        return nil
    } else {
        return m.lists
    }
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("lists", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLists sets the lists property value. The task lists in the users mailbox.
func (m *Todo) SetLists(value []TodoTaskListable)() {
    if m != nil {
        m.lists = value
    }
}
