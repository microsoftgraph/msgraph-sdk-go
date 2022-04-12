package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintTaskStatus 
type PrintTaskStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A human-readable description of the current processing state of the printTask.
    description *string;
    // The current processing state of the printTask. Valid values are described in the following table.
    state *PrintTaskProcessingState;
}
// NewPrintTaskStatus instantiates a new printTaskStatus and sets the default values.
func NewPrintTaskStatus()(*PrintTaskStatus) {
    m := &PrintTaskStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePrintTaskStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintTaskStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrintTaskStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrintTaskStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDescription gets the description property value. A human-readable description of the current processing state of the printTask.
func (m *PrintTaskStatus) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintTaskStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintTaskProcessingState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*PrintTaskProcessingState))
        }
        return nil
    }
    return res
}
// GetState gets the state property value. The current processing state of the printTask. Valid values are described in the following table.
func (m *PrintTaskStatus) GetState()(*PrintTaskProcessingState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Serialize serializes information the current object
func (m *PrintTaskStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrintTaskStatus) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDescription sets the description property value. A human-readable description of the current processing state of the printTask.
func (m *PrintTaskStatus) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetState sets the state property value. The current processing state of the printTask. Valid values are described in the following table.
func (m *PrintTaskStatus) SetState(value *PrintTaskProcessingState)() {
    if m != nil {
        m.state = value
    }
}
