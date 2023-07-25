package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookFunctionResult 
type WorkbookFunctionResult struct {
    Entity
}
// NewWorkbookFunctionResult instantiates a new workbookFunctionResult and sets the default values.
func NewWorkbookFunctionResult()(*WorkbookFunctionResult) {
    m := &WorkbookFunctionResult{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookFunctionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookFunctionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookFunctionResult(), nil
}
// GetError gets the error property value. The error property
func (m *WorkbookFunctionResult) GetError()(*string) {
    val, err := m.GetBackingStore().Get("error")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookFunctionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(Jsonable))
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *WorkbookFunctionResult) GetValue()(Jsonable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Jsonable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WorkbookFunctionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetError sets the error property value. The error property
func (m *WorkbookFunctionResult) SetError(value *string)() {
    err := m.GetBackingStore().Set("error", value)
    if err != nil {
        panic(err)
    }
}
// SetValue sets the value property value. The value property
func (m *WorkbookFunctionResult) SetValue(value Jsonable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// WorkbookFunctionResultable 
type WorkbookFunctionResultable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetError()(*string)
    GetValue()(Jsonable)
    SetError(value *string)()
    SetValue(value Jsonable)()
}
