package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintTaskTrigger provides operations to manage the print singleton.
type PrintTaskTrigger struct {
    Entity
    // 
    definition PrintTaskDefinitionable;
    // The Universal Print event that will cause a new printTask to be triggered. Valid values are described in the following table.
    event *PrintEvent;
}
// NewPrintTaskTrigger instantiates a new printTaskTrigger and sets the default values.
func NewPrintTaskTrigger()(*PrintTaskTrigger) {
    m := &PrintTaskTrigger{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrintTaskTriggerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintTaskTriggerFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPrintTaskTrigger(), nil
}
// GetDefinition gets the definition property value. 
func (m *PrintTaskTrigger) GetDefinition()(PrintTaskDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.definition
    }
}
// GetEvent gets the event property value. The Universal Print event that will cause a new printTask to be triggered. Valid values are described in the following table.
func (m *PrintTaskTrigger) GetEvent()(*PrintEvent) {
    if m == nil {
        return nil
    } else {
        return m.event
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintTaskTrigger) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["event"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintEvent)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvent(val.(*PrintEvent))
        }
        return nil
    }
    return res
}
func (m *PrintTaskTrigger) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        cast := (*m.GetEvent()).String()
        err = writer.WriteStringValue("event", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefinition sets the definition property value. 
func (m *PrintTaskTrigger) SetDefinition(value PrintTaskDefinitionable)() {
    if m != nil {
        m.definition = value
    }
}
// SetEvent sets the event property value. The Universal Print event that will cause a new printTask to be triggered. Valid values are described in the following table.
func (m *PrintTaskTrigger) SetEvent(value *PrintEvent)() {
    if m != nil {
        m.event = value
    }
}
