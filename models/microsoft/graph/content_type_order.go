package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ContentTypeOrder struct {
    additionalData map[string]interface{};
    default_escpaped *bool;
    position *int32;
}
func NewContentTypeOrder()(*ContentTypeOrder) {
    m := &ContentTypeOrder{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ContentTypeOrder) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ContentTypeOrder) GetDefault_escpaped()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.default_escpaped
    }
}
func (m *ContentTypeOrder) GetPosition()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.position
    }
}
func (m *ContentTypeOrder) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["default_escpaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDefault_escpaped(val)
        return nil
    }
    res["position"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPosition(val)
        return nil
    }
    return res
}
func (m *ContentTypeOrder) IsNil()(bool) {
    return m == nil
}
func (m *ContentTypeOrder) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("default_escpaped", m.GetDefault_escpaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("position", m.GetPosition())
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
func (m *ContentTypeOrder) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ContentTypeOrder) SetDefault_escpaped(value *bool)() {
    m.default_escpaped = value
}
func (m *ContentTypeOrder) SetPosition(value *int32)() {
    m.position = value
}
