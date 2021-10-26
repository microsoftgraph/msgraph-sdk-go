package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookOperationError struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The error code.
    code *string;
    // 
    innerError *WorkbookOperationError;
    // The error message.
    message *string;
}
// Instantiates a new workbookOperationError and sets the default values.
func NewWorkbookOperationError()(*WorkbookOperationError) {
    m := &WorkbookOperationError{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbookOperationError) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the code property value. The error code.
func (m *WorkbookOperationError) GetCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
// Gets the innerError property value. 
func (m *WorkbookOperationError) GetInnerError()(*WorkbookOperationError) {
    if m == nil {
        return nil
    } else {
        return m.innerError
    }
}
// Gets the message property value. The error message.
func (m *WorkbookOperationError) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
// The deserialization information for the current model
func (m *WorkbookOperationError) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["code"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCode(val)
        return nil
    }
    res["innerError"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookOperationError() })
        if err != nil {
            return err
        }
        m.SetInnerError(val.(*WorkbookOperationError))
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
    return res
}
func (m *WorkbookOperationError) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkbookOperationError) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
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
func (m *WorkbookOperationError) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the code property value. The error code.
// Parameters:
//  - value : Value to set for the code property.
func (m *WorkbookOperationError) SetCode(value *string)() {
    m.code = value
}
// Sets the innerError property value. 
// Parameters:
//  - value : Value to set for the innerError property.
func (m *WorkbookOperationError) SetInnerError(value *WorkbookOperationError)() {
    m.innerError = value
}
// Sets the message property value. The error message.
// Parameters:
//  - value : Value to set for the message property.
func (m *WorkbookOperationError) SetMessage(value *string)() {
    m.message = value
}
