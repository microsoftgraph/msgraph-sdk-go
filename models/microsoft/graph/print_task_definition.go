package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintTaskDefinition provides operations to manage the print singleton.
type PrintTaskDefinition struct {
    Entity
    // 
    createdBy AppIdentityable;
    // The name of the printTaskDefinition.
    displayName *string;
    // A list of tasks that have been created based on this definition. The list includes currently running tasks and recently completed tasks. Read-only.
    tasks []PrintTaskable;
}
// NewPrintTaskDefinition instantiates a new printTaskDefinition and sets the default values.
func NewPrintTaskDefinition()(*PrintTaskDefinition) {
    m := &PrintTaskDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrintTaskDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintTaskDefinitionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPrintTaskDefinition(), nil
}
// GetCreatedBy gets the createdBy property value. 
func (m *PrintTaskDefinition) GetCreatedBy()(AppIdentityable) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetDisplayName gets the displayName property value. The name of the printTaskDefinition.
func (m *PrintTaskDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintTaskDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateAppIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(AppIdentityable))
        }
        return nil
    }
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
    res["tasks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintTaskFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintTaskable, len(val))
            for i, v := range val {
                res[i] = v.(PrintTaskable)
            }
            m.SetTasks(res)
        }
        return nil
    }
    return res
}
// GetTasks gets the tasks property value. A list of tasks that have been created based on this definition. The list includes currently running tasks and recently completed tasks. Read-only.
func (m *PrintTaskDefinition) GetTasks()([]PrintTaskable) {
    if m == nil {
        return nil
    } else {
        return m.tasks
    }
}
func (m *PrintTaskDefinition) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    return nil
}
// SetCreatedBy sets the createdBy property value. 
func (m *PrintTaskDefinition) SetCreatedBy(value AppIdentityable)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetDisplayName sets the displayName property value. The name of the printTaskDefinition.
func (m *PrintTaskDefinition) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetTasks sets the tasks property value. A list of tasks that have been created based on this definition. The list includes currently running tasks and recently completed tasks. Read-only.
func (m *PrintTaskDefinition) SetTasks(value []PrintTaskable)() {
    if m != nil {
        m.tasks = value
    }
}
