package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PrintTask struct {
    Entity
    definition *PrintTaskDefinition;
    parentUrl *string;
    status *PrintTaskStatus;
    trigger *PrintTaskTrigger;
}
func NewPrintTask()(*PrintTask) {
    m := &PrintTask{
        Entity: *NewEntity(),
    }
    return m
}
func (m *PrintTask) GetDefinition()(*PrintTaskDefinition) {
    if m == nil {
        return nil
    } else {
        return m.definition
    }
}
func (m *PrintTask) GetParentUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentUrl
    }
}
func (m *PrintTask) GetStatus()(*PrintTaskStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *PrintTask) GetTrigger()(*PrintTaskTrigger) {
    if m == nil {
        return nil
    } else {
        return m.trigger
    }
}
func (m *PrintTask) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["definition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintTaskDefinition() })
        if err != nil {
            return err
        }
        m.SetDefinition(val.(*PrintTaskDefinition))
        return nil
    }
    res["parentUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetParentUrl(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintTaskStatus() })
        if err != nil {
            return err
        }
        m.SetStatus(val.(*PrintTaskStatus))
        return nil
    }
    res["trigger"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintTaskTrigger() })
        if err != nil {
            return err
        }
        m.SetTrigger(val.(*PrintTaskTrigger))
        return nil
    }
    return res
}
func (m *PrintTask) IsNil()(bool) {
    return m == nil
}
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
func (m *PrintTask) SetDefinition(value *PrintTaskDefinition)() {
    m.definition = value
}
func (m *PrintTask) SetParentUrl(value *string)() {
    m.parentUrl = value
}
func (m *PrintTask) SetStatus(value *PrintTaskStatus)() {
    m.status = value
}
func (m *PrintTask) SetTrigger(value *PrintTaskTrigger)() {
    m.trigger = value
}
