package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PublicInnerError struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The error code.
    code *string;
    // A collection of error details.
    details []PublicErrorDetail;
    // The error message.
    message *string;
    // The target of the error.
    target *string;
}
// Instantiates a new publicInnerError and sets the default values.
func NewPublicInnerError()(*PublicInnerError) {
    m := &PublicInnerError{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PublicInnerError) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the code property value. The error code.
func (m *PublicInnerError) GetCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
// Gets the details property value. A collection of error details.
func (m *PublicInnerError) GetDetails()([]PublicErrorDetail) {
    if m == nil {
        return nil
    } else {
        return m.details
    }
}
// Gets the message property value. The error message.
func (m *PublicInnerError) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
// Gets the target property value. The target of the error.
func (m *PublicInnerError) GetTarget()(*string) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// The deserialization information for the current model
func (m *PublicInnerError) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["code"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCode(val)
        return nil
    }
    res["details"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPublicErrorDetail() })
        if err != nil {
            return err
        }
        res := make([]PublicErrorDetail, len(val))
        for i, v := range val {
            res[i] = *(v.(*PublicErrorDetail))
        }
        m.SetDetails(res)
        return nil
    }
    res["message"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMessage(val)
        return nil
    }
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTarget(val)
        return nil
    }
    return res
}
func (m *PublicInnerError) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PublicInnerError) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *PublicInnerError) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the code property value. The error code.
// Parameters:
//  - value : Value to set for the code property.
func (m *PublicInnerError) SetCode(value *string)() {
    m.code = value
}
// Sets the details property value. A collection of error details.
// Parameters:
//  - value : Value to set for the details property.
func (m *PublicInnerError) SetDetails(value []PublicErrorDetail)() {
    m.details = value
}
// Sets the message property value. The error message.
// Parameters:
//  - value : Value to set for the message property.
func (m *PublicInnerError) SetMessage(value *string)() {
    m.message = value
}
// Sets the target property value. The target of the error.
// Parameters:
//  - value : Value to set for the target property.
func (m *PublicInnerError) SetTarget(value *string)() {
    m.target = value
}
