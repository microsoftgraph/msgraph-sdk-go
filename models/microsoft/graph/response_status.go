package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ResponseStatus struct {
    additionalData map[string]interface{};
    response *ResponseType;
    time *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
func NewResponseStatus()(*ResponseStatus) {
    m := &ResponseStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ResponseStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ResponseStatus) GetResponse()(*ResponseType) {
    if m == nil {
        return nil
    } else {
        return m.response
    }
}
func (m *ResponseStatus) GetTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.time
    }
}
func (m *ResponseStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["response"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseResponseType)
        if err != nil {
            return err
        }
        cast := val.(ResponseType)
        m.SetResponse(&cast)
        return nil
    }
    res["time"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetTime(val)
        return nil
    }
    return res
}
func (m *ResponseStatus) IsNil()(bool) {
    return m == nil
}
func (m *ResponseStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetResponse() != nil {
        cast := m.GetResponse().String()
        err := writer.WriteStringValue("response", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("time", m.GetTime())
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
func (m *ResponseStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ResponseStatus) SetResponse(value *ResponseType)() {
    m.response = value
}
func (m *ResponseStatus) SetTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.time = value
}
