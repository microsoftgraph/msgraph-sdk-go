package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PrintTaskDefinition struct {
    Entity
    createdBy *AppIdentity;
    displayName *string;
    tasks []PrintTask;
}
func NewPrintTaskDefinition()(*PrintTaskDefinition) {
    m := &PrintTaskDefinition{
        Entity: *NewEntity(),
    }
    return m
}
func (m *PrintTaskDefinition) GetCreatedBy()(*AppIdentity) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
func (m *PrintTaskDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *PrintTaskDefinition) GetTasks()([]PrintTask) {
    if m == nil {
        return nil
    } else {
        return m.tasks
    }
}
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
func (m *PrintTaskDefinition) SetCreatedBy(value *AppIdentity)() {
    m.createdBy = value
}
func (m *PrintTaskDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *PrintTaskDefinition) SetTasks(value []PrintTask)() {
    m.tasks = value
}
