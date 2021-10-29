package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PrintTaskDefinition struct {
    Entity
    // 
    createdBy *AppIdentity;
    // The name of the printTaskDefinition.
    displayName *string;
    // A list of tasks that have been created based on this definition. The list includes currently running tasks and recently completed tasks. Read-only.
    tasks []PrintTask;
}
// Instantiates a new printTaskDefinition and sets the default values.
func NewPrintTaskDefinition()(*PrintTaskDefinition) {
    m := &PrintTaskDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the createdBy property value. 
func (m *PrintTaskDefinition) GetCreatedBy()(*AppIdentity) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// Gets the displayName property value. The name of the printTaskDefinition.
func (m *PrintTaskDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the tasks property value. A list of tasks that have been created based on this definition. The list includes currently running tasks and recently completed tasks. Read-only.
func (m *PrintTaskDefinition) GetTasks()([]PrintTask) {
    if m == nil {
        return nil
    } else {
        return m.tasks
    }
}
// The deserialization information for the current model
func (m *PrintTaskDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppIdentity() })
        if err != nil {
            return err
        }
        m.SetCreatedBy(val.(*AppIdentity))
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["tasks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintTask() })
        if err != nil {
            return err
        }
        res := make([]PrintTask, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintTask))
        }
        m.SetTasks(res)
        return nil
    }
    return res
}
func (m *PrintTaskDefinition) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PrintTaskDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
    return nil
}
// Sets the createdBy property value. 
// Parameters:
//  - value : Value to set for the createdBy property.
func (m *PrintTaskDefinition) SetCreatedBy(value *AppIdentity)() {
    m.createdBy = value
}
// Sets the displayName property value. The name of the printTaskDefinition.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *PrintTaskDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the tasks property value. A list of tasks that have been created based on this definition. The list includes currently running tasks and recently completed tasks. Read-only.
// Parameters:
//  - value : Value to set for the tasks property.
func (m *PrintTaskDefinition) SetTasks(value []PrintTask)() {
    m.tasks = value
}
