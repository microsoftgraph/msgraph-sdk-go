package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/callrecords"
)

type FailureInfo struct {
    additionalData map[string]interface{};
    reason *string;
    stage *i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.FailureStage;
}
func NewFailureInfo()(*FailureInfo) {
    m := &FailureInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *FailureInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *FailureInfo) GetReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
func (m *FailureInfo) GetStage()(*i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.FailureStage) {
    if m == nil {
        return nil
    } else {
        return m.stage
    }
}
func (m *FailureInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["reason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReason(val)
        return nil
    }
    res["stage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.ParseFailureStage)
        if err != nil {
            return err
        }
        cast := val.(i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.FailureStage)
        m.SetStage(&cast)
        return nil
    }
    return res
}
func (m *FailureInfo) IsNil()(bool) {
    return m == nil
}
func (m *FailureInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("reason", m.GetReason())
        if err != nil {
            return err
        }
    }
    if m.GetStage() != nil {
        cast := m.GetStage().String()
        err := writer.WriteStringValue("stage", &cast)
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
func (m *FailureInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *FailureInfo) SetReason(value *string)() {
    m.reason = value
}
func (m *FailureInfo) SetStage(value *i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.FailureStage)() {
    m.stage = value
}
