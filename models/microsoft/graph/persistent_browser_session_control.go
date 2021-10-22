package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PersistentBrowserSessionControl struct {
    ConditionalAccessSessionControl
    mode *PersistentBrowserSessionMode;
}
func NewPersistentBrowserSessionControl()(*PersistentBrowserSessionControl) {
    m := &PersistentBrowserSessionControl{
        ConditionalAccessSessionControl: *NewConditionalAccessSessionControl(),
    }
    return m
}
func (m *PersistentBrowserSessionControl) GetMode()(*PersistentBrowserSessionMode) {
    if m == nil {
        return nil
    } else {
        return m.mode
    }
}
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
func (m *PersistentBrowserSessionControl) SetMode(value *PersistentBrowserSessionMode)() {
    m.mode = value
}
