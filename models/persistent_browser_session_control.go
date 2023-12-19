package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersistentBrowserSessionControl 
type PersistentBrowserSessionControl struct {
    ConditionalAccessSessionControl
}
// NewPersistentBrowserSessionControl instantiates a new persistentBrowserSessionControl and sets the default values.
func NewPersistentBrowserSessionControl()(*PersistentBrowserSessionControl) {
    m := &PersistentBrowserSessionControl{
        ConditionalAccessSessionControl: *NewConditionalAccessSessionControl(),
    }
    odataTypeValue := "#microsoft.graph.persistentBrowserSessionControl"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePersistentBrowserSessionControlFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePersistentBrowserSessionControlFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPersistentBrowserSessionControl(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PersistentBrowserSessionControl) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ConditionalAccessSessionControl.GetFieldDeserializers()
    res["mode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePersistentBrowserSessionControl_mode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMode(val.(*PersistentBrowserSessionControl_mode))
        }
        return nil
    }
    return res
}
// GetMode gets the mode property value. Possible values are: always, never.
func (m *PersistentBrowserSessionControl) GetMode()(*PersistentBrowserSessionControl_mode) {
    val, err := m.GetBackingStore().Get("mode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PersistentBrowserSessionControl_mode)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PersistentBrowserSessionControl) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ConditionalAccessSessionControl.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMode() != nil {
        cast := (*m.GetMode()).String()
        err = writer.WriteStringValue("mode", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMode sets the mode property value. Possible values are: always, never.
func (m *PersistentBrowserSessionControl) SetMode(value *PersistentBrowserSessionControl_mode)() {
    err := m.GetBackingStore().Set("mode", value)
    if err != nil {
        panic(err)
    }
}
// PersistentBrowserSessionControlable 
type PersistentBrowserSessionControlable interface {
    ConditionalAccessSessionControlable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMode()(*PersistentBrowserSessionControl_mode)
    SetMode(value *PersistentBrowserSessionControl_mode)()
}
