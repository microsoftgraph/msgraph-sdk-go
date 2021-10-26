package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PersistentBrowserSessionControl struct {
    ConditionalAccessSessionControl
    // Possible values are: always, never.
    mode *PersistentBrowserSessionMode;
}
// Instantiates a new persistentBrowserSessionControl and sets the default values.
func NewPersistentBrowserSessionControl()(*PersistentBrowserSessionControl) {
    m := &PersistentBrowserSessionControl{
        ConditionalAccessSessionControl: *NewConditionalAccessSessionControl(),
    }
    return m
}
// Gets the mode property value. Possible values are: always, never.
func (m *PersistentBrowserSessionControl) GetMode()(*PersistentBrowserSessionMode) {
    if m == nil {
        return nil
    } else {
        return m.mode
    }
}
// The deserialization information for the current model
func (m *PersistentBrowserSessionControl) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ConditionalAccessSessionControl.GetFieldDeserializers()
    res["mode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePersistentBrowserSessionMode)
        if err != nil {
            return err
        }
        cast := val.(PersistentBrowserSessionMode)
        m.SetMode(&cast)
        return nil
    }
    return res
}
func (m *PersistentBrowserSessionControl) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PersistentBrowserSessionControl) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ConditionalAccessSessionControl.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMode() != nil {
        cast := m.GetMode().String()
        err = writer.WriteStringValue("mode", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the mode property value. Possible values are: always, never.
// Parameters:
//  - value : Value to set for the mode property.
func (m *PersistentBrowserSessionControl) SetMode(value *PersistentBrowserSessionMode)() {
    m.mode = value
}
