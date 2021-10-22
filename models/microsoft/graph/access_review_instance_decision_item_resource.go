package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AccessReviewInstanceDecisionItemResource struct {
    additionalData map[string]interface{};
    displayName *string;
    id *string;
    type_escpaped *string;
}
func NewAccessReviewInstanceDecisionItemResource()(*AccessReviewInstanceDecisionItemResource) {
    m := &AccessReviewInstanceDecisionItemResource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AccessReviewInstanceDecisionItemResource) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AccessReviewInstanceDecisionItemResource) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *AccessReviewInstanceDecisionItemResource) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
func (m *AccessReviewInstanceDecisionItemResource) GetType_escpaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escpaped
    }
}
func (m *AccessReviewInstanceDecisionItemResource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    res["type_escpaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetType_escpaped(val)
        return nil
    }
    return res
}
func (m *AccessReviewInstanceDecisionItemResource) IsNil()(bool) {
    return m == nil
}
func (m *AccessReviewInstanceDecisionItemResource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type_escpaped", m.GetType_escpaped())
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
func (m *AccessReviewInstanceDecisionItemResource) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AccessReviewInstanceDecisionItemResource) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *AccessReviewInstanceDecisionItemResource) SetId(value *string)() {
    m.id = value
}
func (m *AccessReviewInstanceDecisionItemResource) SetType_escpaped(value *string)() {
    m.type_escpaped = value
}
