package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AlertDetection 
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
// NewAlertDetection instantiates a new alertDetection and sets the default values.
func NewAlertDetection()(*AlertDetection) {
    m := &AlertDetection{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAlertDetectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAlertDetectionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAlertDetection(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlertDetection) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDetectionType gets the detectionType property value. 
func (m *AlertDetection) GetDetectionType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.detectionType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AlertDetection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["detectionType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionType(val)
        }
        return nil
    }
    res["method"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMethod(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetMethod gets the method property value. 
func (m *AlertDetection) GetMethod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.method
    }
}
// GetName gets the name property value. 
func (m *AlertDetection) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlertDetection) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDetectionType sets the detectionType property value. 
func (m *AlertDetection) SetDetectionType(value *string)() {
    if m != nil {
        m.detectionType = value
    }
}
// SetMethod sets the method property value. 
func (m *AlertDetection) SetMethod(value *string)() {
    if m != nil {
        m.method = value
    }
}
// SetName sets the name property value. 
func (m *AlertDetection) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
