package odataerrors

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InnerError 
type InnerError struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Client request Id as sent by the client application.
    clientRequestId *string;
    // Date when the error occured.
    date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Request Id as tracked internally by the service
    requestId *string;
}
// NewInnerError instantiates a new InnerError and sets the default values.
func NewInnerError()(*InnerError) {
    m := &InnerError{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateInnerErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInnerErrorFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewInnerError(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InnerError) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetClientRequestId gets the client-request-id property value. Client request Id as sent by the client application.
func (m *InnerError) GetClientRequestId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientRequestId
    }
}
// GetDate gets the date property value. Date when the error occured.
func (m *InnerError) GetDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.date
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InnerError) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["client-request-id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientRequestId(val)
        }
        return nil
    }
    res["date"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDate(val)
        }
        return nil
    }
    res["request-id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestId(val)
        }
        return nil
    }
    return res
}
// GetRequestId gets the request-id property value. Request Id as tracked internally by the service
func (m *InnerError) GetRequestId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestId
    }
}
// Serialize serializes information the current object
func (m *InnerError) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("client-request-id", m.GetClientRequestId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("date", m.GetDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("request-id", m.GetRequestId())
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
func (m *InnerError) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetClientRequestId sets the client-request-id property value. Client request Id as sent by the client application.
func (m *InnerError) SetClientRequestId(value *string)() {
    if m != nil {
        m.clientRequestId = value
    }
}
// SetDate sets the date property value. Date when the error occured.
func (m *InnerError) SetDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.date = value
    }
}
// SetRequestId sets the request-id property value. Request Id as tracked internally by the service
func (m *InnerError) SetRequestId(value *string)() {
    if m != nil {
        m.requestId = value
    }
}
