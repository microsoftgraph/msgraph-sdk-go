package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintTask provides operations to manage the print singleton.
type PrintTask struct {
    Entity
    // 
    definition PrintTaskDefinitionable;
    // The URL for the print entity that triggered this task. For example, https://graph.microsoft.com/v1.0/print/printers/{printerId}/jobs/{jobId}. Read-only.
    parentUrl *string;
    // 
    status PrintTaskStatusable;
    // 
    trigger PrintTaskTriggerable;
}
// NewPrintTask instantiates a new printTask and sets the default values.
func NewPrintTask()(*PrintTask) {
    m := &PrintTask{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrintTaskFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintTaskFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPrintTask(), nil
}
// GetDefinition gets the definition property value. 
func (m *PrintTask) GetDefinition()(PrintTaskDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.definition
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintTask) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["definition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrintTaskDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinition(val.(PrintTaskDefinitionable))
        }
        return nil
    }
    res["parentUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentUrl(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrintTaskStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(PrintTaskStatusable))
        }
        return nil
    }
    res["trigger"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrintTaskTriggerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrigger(val.(PrintTaskTriggerable))
        }
        return nil
    }
    return res
}
// GetParentUrl gets the parentUrl property value. The URL for the print entity that triggered this task. For example, https://graph.microsoft.com/v1.0/print/printers/{printerId}/jobs/{jobId}. Read-only.
func (m *PrintTask) GetParentUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentUrl
    }
}
// GetStatus gets the status property value. 
func (m *PrintTask) GetStatus()(PrintTaskStatusable) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetTrigger gets the trigger property value. 
func (m *PrintTask) GetTrigger()(PrintTaskTriggerable) {
    if m == nil {
        return nil
    } else {
        return m.trigger
    }
}
func (m *PrintTask) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PrintTask) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("definition", m.GetDefinition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentUrl", m.GetParentUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("trigger", m.GetTrigger())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefinition sets the definition property value. 
func (m *PrintTask) SetDefinition(value PrintTaskDefinitionable)() {
    if m != nil {
        m.definition = value
    }
}
// SetParentUrl sets the parentUrl property value. The URL for the print entity that triggered this task. For example, https://graph.microsoft.com/v1.0/print/printers/{printerId}/jobs/{jobId}. Read-only.
func (m *PrintTask) SetParentUrl(value *string)() {
    if m != nil {
        m.parentUrl = value
    }
}
// SetStatus sets the status property value. 
func (m *PrintTask) SetStatus(value PrintTaskStatusable)() {
    if m != nil {
        m.status = value
    }
}
// SetTrigger sets the trigger property value. 
func (m *PrintTask) SetTrigger(value PrintTaskTriggerable)() {
    if m != nil {
        m.trigger = value
    }
}
