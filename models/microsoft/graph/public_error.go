package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PublicError 
type PublicError struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Represents the error code.
    code *string;
    // Details of the error.
    details []PublicErrorDetail;
    // Details of the inner error.
    innerError *PublicInnerError;
    // A non-localized message for the developer.
    message *string;
    // The target of the error.
    target *string;
}
// NewPublicError instantiates a new publicError and sets the default values.
func NewPublicError()(*PublicError) {
    m := &PublicError{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PublicError) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCode gets the code property value. Represents the error code.
func (m *PublicError) GetCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
// GetDetails gets the details property value. Details of the error.
func (m *PublicError) GetDetails()([]PublicErrorDetail) {
    if m == nil {
        return nil
    } else {
        return m.details
    }
}
// GetInnerError gets the innerError property value. Details of the inner error.
func (m *PublicError) GetInnerError()(*PublicInnerError) {
    if m == nil {
        return nil
    } else {
        return m.innerError
    }
}
// GetMessage gets the message property value. A non-localized message for the developer.
func (m *PublicError) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
// GetTarget gets the target property value. The target of the error.
func (m *PublicError) GetTarget()(*string) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PublicError) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["details"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPublicErrorDetail() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PublicErrorDetail, len(val))
            for i, v := range val {
                res[i] = *(v.(*PublicErrorDetail))
            }
            m.SetDetails(res)
        }
        return nil
    }
    res["innerError"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPublicInnerError() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInnerError(val.(*PublicInnerError))
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
func (m *PublicError) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PublicError) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDetails()))
        for i, v := range m.GetDetails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("details", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("innerError", m.GetInnerError())
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
func (m *PublicError) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCode sets the code property value. Represents the error code.
func (m *PublicError) SetCode(value *string)() {
    if m != nil {
        m.code = value
    }
}
// SetDetails sets the details property value. Details of the error.
func (m *PublicError) SetDetails(value []PublicErrorDetail)() {
    if m != nil {
        m.details = value
    }
}
// SetInnerError sets the innerError property value. Details of the inner error.
func (m *PublicError) SetInnerError(value *PublicInnerError)() {
    if m != nil {
        m.innerError = value
    }
}
// SetMessage sets the message property value. A non-localized message for the developer.
func (m *PublicError) SetMessage(value *string)() {
    if m != nil {
        m.message = value
    }
}
// SetTarget sets the target property value. The target of the error.
func (m *PublicError) SetTarget(value *string)() {
    if m != nil {
        m.target = value
    }
}
