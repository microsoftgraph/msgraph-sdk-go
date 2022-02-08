package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintOperationStatus 
type PrintOperationStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A human-readable description of the printOperation's current processing state. Read-only.
    description *string;
    // The printOperation's current processing state. Valid values are described in the following table. Read-only.
    state *PrintOperationProcessingState;
}
// NewPrintOperationStatus instantiates a new printOperationStatus and sets the default values.
func NewPrintOperationStatus()(*PrintOperationStatus) {
    m := &PrintOperationStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrintOperationStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDescription gets the description property value. A human-readable description of the printOperation's current processing state. Read-only.
func (m *PrintOperationStatus) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetState gets the state property value. The printOperation's current processing state. Valid values are described in the following table. Read-only.
func (m *PrintOperationStatus) GetState()(*PrintOperationProcessingState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintOperationStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintOperationProcessingState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*PrintOperationProcessingState))
        }
        return nil
    }
    return res
}
func (m *PrintOperationStatus) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PrintOperationStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *PrintOperationStatus) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDescription sets the description property value. A human-readable description of the printOperation's current processing state. Read-only.
func (m *PrintOperationStatus) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetState sets the state property value. The printOperation's current processing state. Valid values are described in the following table. Read-only.
func (m *PrintOperationStatus) SetState(value *PrintOperationProcessingState)() {
    if m != nil {
        m.state = value
    }
}
