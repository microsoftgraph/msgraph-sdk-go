package odataerrors

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ErrorDetails provides operations to manage the admin singleton.
type ErrorDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    code *string;
    // 
    message *string;
    // 
    target *string;
}
// NewErrorDetails instantiates a new ErrorDetails and sets the default values.
func NewErrorDetails()(*ErrorDetails) {
    m := &ErrorDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateErrorDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateErrorDetailsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewErrorDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ErrorDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCode gets the code property value. 
func (m *ErrorDetails) GetCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ErrorDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["code"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["message"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. 
func (m *ErrorDetails) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
// GetTarget gets the target property value. 
func (m *ErrorDetails) GetTarget()(*string) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
func (m *ErrorDetails) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ErrorDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("target", m.GetTarget())
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
func (m *ErrorDetails) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCode sets the code property value. 
func (m *ErrorDetails) SetCode(value *string)() {
    if m != nil {
        m.code = value
    }
}
// SetMessage sets the message property value. 
func (m *ErrorDetails) SetMessage(value *string)() {
    if m != nil {
        m.message = value
    }
}
// SetTarget sets the target property value. 
func (m *ErrorDetails) SetTarget(value *string)() {
    if m != nil {
        m.target = value
    }
}
