package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// FreeBusyError 
type FreeBusyError struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Describes the error.
    message *string;
    // The response code from querying for the availability of the user, distribution list, or resource.
    responseCode *string;
}
// NewFreeBusyError instantiates a new freeBusyError and sets the default values.
func NewFreeBusyError()(*FreeBusyError) {
    m := &FreeBusyError{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateFreeBusyErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFreeBusyErrorFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewFreeBusyError(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FreeBusyError) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FreeBusyError) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["responseCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponseCode(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. Describes the error.
func (m *FreeBusyError) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
// GetResponseCode gets the responseCode property value. The response code from querying for the availability of the user, distribution list, or resource.
func (m *FreeBusyError) GetResponseCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.responseCode
    }
}
// Serialize serializes information the current object
func (m *FreeBusyError) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("responseCode", m.GetResponseCode())
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
func (m *FreeBusyError) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMessage sets the message property value. Describes the error.
func (m *FreeBusyError) SetMessage(value *string)() {
    if m != nil {
        m.message = value
    }
}
// SetResponseCode sets the responseCode property value. The response code from querying for the availability of the user, distribution list, or resource.
func (m *FreeBusyError) SetResponseCode(value *string)() {
    if m != nil {
        m.responseCode = value
    }
}
