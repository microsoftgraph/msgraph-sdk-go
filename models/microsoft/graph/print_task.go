package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintTask 
type PrintTask struct {
    Entity
    // 
    definition *PrintTaskDefinition;
    // The URL for the print entity that triggered this task. For example, https://graph.microsoft.com/beta/print/printers/{printerId}/jobs/{jobId}. Read-only.
    parentUrl *string;
    // 
    status *PrintTaskStatus;
    // 
    trigger *PrintTaskTrigger;
}
// NewPrintTask instantiates a new printTask and sets the default values.
func NewPrintTask()(*PrintTask) {
    m := &PrintTask{
        Entity: *NewEntity(),
    }
    return m
}
// GetDefinition gets the definition property value. 
func (m *PrintTask) GetDefinition()(*PrintTaskDefinition) {
    if m == nil {
        return nil
    } else {
        return m.definition
    }
}
// GetParentUrl gets the parentUrl property value. The URL for the print entity that triggered this task. For example, https://graph.microsoft.com/beta/print/printers/{printerId}/jobs/{jobId}. Read-only.
func (m *PrintTask) GetParentUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentUrl
    }
}
// GetStatus gets the status property value. 
func (m *PrintTask) GetStatus()(*PrintTaskStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetTrigger gets the trigger property value. 
func (m *PrintTask) GetTrigger()(*PrintTaskTrigger) {
    if m == nil {
        return nil
    } else {
        return m.trigger
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintTask) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["definition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintTaskDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinition(val.(*PrintTaskDefinition))
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
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintTaskStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*PrintTaskStatus))
        }
        return nil
    }
    res["trigger"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintTaskTrigger() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrigger(val.(*PrintTaskTrigger))
        }
        return nil
    }
    return res
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
func (m *PrintTask) SetDefinition(value *PrintTaskDefinition)() {
    if m != nil {
        m.definition = value
    }
}
// SetParentUrl sets the parentUrl property value. The URL for the print entity that triggered this task. For example, https://graph.microsoft.com/beta/print/printers/{printerId}/jobs/{jobId}. Read-only.
func (m *PrintTask) SetParentUrl(value *string)() {
    if m != nil {
        m.parentUrl = value
    }
}
// SetStatus sets the status property value. 
func (m *PrintTask) SetStatus(value *PrintTaskStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetTrigger sets the trigger property value. 
func (m *PrintTask) SetTrigger(value *PrintTaskTrigger)() {
    if m != nil {
        m.trigger = value
    }
}
