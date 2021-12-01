package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
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
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
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
// GetState gets the state property value. The current processing state of the printTask. Valid values are described in the following table.
func (m *PrintTaskStatus) GetState()(*PrintTaskProcessingState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintTaskStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
        val, err := n.GetEnumValue(ParsePrintTaskProcessingState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(PrintTaskProcessingState)
            m.SetState(&cast)
        }
        return nil
    }
    return res
}
func (m *PrintTaskStatus) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PrintTaskStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
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
