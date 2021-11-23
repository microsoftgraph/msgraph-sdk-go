package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// persistentBrowserSessionControl 
type PersistentBrowserSessionControl struct {
    ConditionalAccessSessionControl
    // Possible values are: always, never.
    mode *PersistentBrowserSessionMode;
}
// NewPersistentBrowserSessionControl instantiates a new persistentBrowserSessionControl and sets the default values.
func NewPersistentBrowserSessionControl()(*PersistentBrowserSessionControl) {
    m := &PersistentBrowserSessionControl{
        ConditionalAccessSessionControl: *NewConditionalAccessSessionControl(),
    }
    return m
}
// GetMode gets the mode property value. Possible values are: always, never.
func (m *PersistentBrowserSessionControl) GetMode()(*PersistentBrowserSessionMode) {
    if m == nil {
        return nil
    } else {
        return m.mode
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PersistentBrowserSessionControl) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ConditionalAccessSessionControl.GetFieldDeserializers()
    res["mode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePersistentBrowserSessionMode)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(PersistentBrowserSessionMode)
            m.SetMode(&cast)
        }
        return nil
    }
    return res
}
func (m *PersistentBrowserSessionControl) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetMode sets the mode property value. Possible values are: always, never.
func (m *PersistentBrowserSessionControl) SetMode(value *PersistentBrowserSessionMode)() {
    m.mode = value
}
