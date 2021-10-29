package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AlertDetection struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    detectionType *string;
    // 
    method *string;
    // 
    name *string;
}
// Instantiates a new alertDetection and sets the default values.
func NewAlertDetection()(*AlertDetection) {
    m := &AlertDetection{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlertDetection) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the detectionType property value. 
func (m *AlertDetection) GetDetectionType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.detectionType
    }
}
// Gets the method property value. 
func (m *AlertDetection) GetMethod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.method
    }
}
// Gets the name property value. 
func (m *AlertDetection) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *AlertDetection) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the detectionType property value. 
// Parameters:
//  - value : Value to set for the detectionType property.
func (m *AlertDetection) SetDetectionType(value *string)() {
    m.detectionType = value
}
// Sets the method property value. 
// Parameters:
//  - value : Value to set for the method property.
func (m *AlertDetection) SetMethod(value *string)() {
    m.method = value
}
// Sets the name property value. 
// Parameters:
//  - value : Value to set for the name property.
func (m *AlertDetection) SetName(value *string)() {
    m.name = value
}
