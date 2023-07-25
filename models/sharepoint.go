package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Sharepoint 
type Sharepoint struct {
    Entity
}
// NewSharepoint instantiates a new sharepoint and sets the default values.
func NewSharepoint()(*Sharepoint) {
    m := &Sharepoint{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSharepointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharepointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharepoint(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Sharepoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharepointSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(SharepointSettingsable))
        }
        return nil
    }
    return res
}
// GetSettings gets the settings property value. The settings property
func (m *Sharepoint) GetSettings()(SharepointSettingsable) {
    val, err := m.GetBackingStore().Get("settings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SharepointSettingsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Sharepoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSettings sets the settings property value. The settings property
func (m *Sharepoint) SetSettings(value SharepointSettingsable)() {
    err := m.GetBackingStore().Set("settings", value)
    if err != nil {
        panic(err)
    }
}
// Sharepointable 
type Sharepointable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSettings()(SharepointSettingsable)
    SetSettings(value SharepointSettingsable)()
}
