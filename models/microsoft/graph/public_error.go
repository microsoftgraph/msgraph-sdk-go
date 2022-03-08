package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PublicError provides operations to manage the collection of externalConnection entities.
type PublicError struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Represents the error code.
    code *string;
    // Details of the error.
    details []PublicErrorDetailable;
    // Details of the inner error.
    innerError PublicInnerErrorable;
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
// CreatePublicErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePublicErrorFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPublicError(), nil
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
func (m *PublicError) GetDetails()([]PublicErrorDetailable) {
    if m == nil {
        return nil
    } else {
        return m.details
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
        val, err := n.GetCollectionOfObjectValues(CreatePublicErrorDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PublicErrorDetailable, len(val))
            for i, v := range val {
                res[i] = v.(PublicErrorDetailable)
            }
            m.SetDetails(res)
        }
        return nil
    }
    res["innerError"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePublicInnerErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInnerError(val.(PublicInnerErrorable))
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
// GetInnerError gets the innerError property value. Details of the inner error.
func (m *PublicError) GetInnerError()(PublicInnerErrorable) {
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
    if m.GetDetails() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDetails()))
        for i, v := range m.GetDetails() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *PublicError) SetDetails(value []PublicErrorDetailable)() {
    if m != nil {
        m.details = value
    }
}
// SetInnerError sets the innerError property value. Details of the inner error.
func (m *PublicError) SetInnerError(value PublicInnerErrorable)() {
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
