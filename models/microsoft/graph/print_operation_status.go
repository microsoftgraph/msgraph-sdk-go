package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PrintOperationStatus struct {
    additionalData map[string]interface{};
    description *string;
    state *PrintOperationProcessingState;
}
func NewPrintOperationStatus()(*PrintOperationStatus) {
    m := &PrintOperationStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PrintOperationStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PrintOperationStatus) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *PrintOperationStatus) GetState()(*PrintOperationProcessingState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *PrintOperationStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintOperationProcessingState)
        if err != nil {
            return err
        }
        cast := val.(PrintOperationProcessingState)
        m.SetState(&cast)
        return nil
    }
    return res
}
func (m *PrintOperationStatus) IsNil()(bool) {
    return m == nil
}
func (m *PrintOperationStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *PrintOperationStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PrintOperationStatus) SetDescription(value *string)() {
    m.description = value
}
func (m *PrintOperationStatus) SetState(value *PrintOperationProcessingState)() {
    m.state = value
}
