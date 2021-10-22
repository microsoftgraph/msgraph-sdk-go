package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SignInStatus struct {
    additionalData map[string]interface{};
    additionalDetails *string;
    errorCode *int32;
    failureReason *string;
}
func NewSignInStatus()(*SignInStatus) {
    m := &SignInStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SignInStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SignInStatus) GetAdditionalDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.additionalDetails
    }
}
func (m *SignInStatus) GetErrorCode()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
func (m *SignInStatus) GetFailureReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.failureReason
    }
}
func (m *SignInStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["additionalDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAdditionalDetails(val)
        return nil
    }
    res["errorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetErrorCode(val)
        return nil
    }
    res["failureReason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFailureReason(val)
        return nil
    }
    return res
}
func (m *SignInStatus) IsNil()(bool) {
    return m == nil
}
func (m *SignInStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("additionalDetails", m.GetAdditionalDetails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("failureReason", m.GetFailureReason())
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
func (m *SignInStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SignInStatus) SetAdditionalDetails(value *string)() {
    m.additionalDetails = value
}
func (m *SignInStatus) SetErrorCode(value *int32)() {
    m.errorCode = value
}
func (m *SignInStatus) SetFailureReason(value *string)() {
    m.failureReason = value
}
