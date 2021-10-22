package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AlertDetection struct {
    additionalData map[string]interface{};
    detectionType *string;
    method *string;
    name *string;
}
func NewAlertDetection()(*AlertDetection) {
    m := &AlertDetection{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AlertDetection) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AlertDetection) GetDetectionType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.detectionType
    }
}
func (m *AlertDetection) GetMethod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.method
    }
}
func (m *AlertDetection) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *AlertDetection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["detectionType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDetectionType(val)
        return nil
    }
    res["method"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMethod(val)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    return res
}
func (m *AlertDetection) IsNil()(bool) {
    return m == nil
}
func (m *AlertDetection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("detectionType", m.GetDetectionType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("method", m.GetMethod())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *AlertDetection) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AlertDetection) SetDetectionType(value *string)() {
    m.detectionType = value
}
func (m *AlertDetection) SetMethod(value *string)() {
    m.method = value
}
func (m *AlertDetection) SetName(value *string)() {
    m.name = value
}
