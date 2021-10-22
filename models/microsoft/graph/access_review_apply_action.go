package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AccessReviewApplyAction struct {
    additionalData map[string]interface{};
}
func NewAccessReviewApplyAction()(*AccessReviewApplyAction) {
    m := &AccessReviewApplyAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AccessReviewApplyAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AccessReviewApplyAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    return res
}
func (m *AccessReviewApplyAction) IsNil()(bool) {
    return m == nil
}
func (m *AccessReviewApplyAction) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *AccessReviewApplyAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
