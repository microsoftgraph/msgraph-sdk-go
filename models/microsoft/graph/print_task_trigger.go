package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PrintTaskTrigger struct {
    Entity
    // 
    definition *PrintTaskDefinition;
    // The Universal Print event that will cause a new printTask to be triggered. Valid values are described in the following table.
    event *PrintEvent;
}
// Instantiates a new printTaskTrigger and sets the default values.
func NewPrintTaskTrigger()(*PrintTaskTrigger) {
    m := &PrintTaskTrigger{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the definition property value. 
func (m *PrintTaskTrigger) GetDefinition()(*PrintTaskDefinition) {
    if m == nil {
        return nil
    } else {
        return m.definition
    }
}
// Gets the event property value. The Universal Print event that will cause a new printTask to be triggered. Valid values are described in the following table.
func (m *PrintTaskTrigger) GetEvent()(*PrintEvent) {
    if m == nil {
        return nil
    } else {
        return m.event
    }
}
// The deserialization information for the current model
func (m *PrintTaskTrigger) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["definition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintTaskDefinition() })
        if err != nil {
            return err
        }
        m.SetDefinition(val.(*PrintTaskDefinition))
        return nil
    }
    res["event"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintEvent)
        if err != nil {
            return err
        }
        cast := val.(PrintEvent)
        m.SetEvent(&cast)
        return nil
    }
    return res
}
func (m *PrintTaskTrigger) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PrintTaskTrigger) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m.GetEvent() != nil {
        cast := m.GetEvent().String()
        err = writer.WriteStringValue("event", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the definition property value. 
// Parameters:
//  - value : Value to set for the definition property.
func (m *PrintTaskTrigger) SetDefinition(value *PrintTaskDefinition)() {
    m.definition = value
}
// Sets the event property value. The Universal Print event that will cause a new printTask to be triggered. Valid values are described in the following table.
// Parameters:
//  - value : Value to set for the event property.
func (m *PrintTaskTrigger) SetEvent(value *PrintEvent)() {
    m.event = value
}
