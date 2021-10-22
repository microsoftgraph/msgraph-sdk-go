package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MediaStream struct {
    additionalData map[string]interface{};
    direction *MediaDirection;
    label *string;
    mediaType *Modality;
    serverMuted *bool;
    sourceId *string;
}
func NewMediaStream()(*MediaStream) {
    m := &MediaStream{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MediaStream) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MediaStream) GetDirection()(*MediaDirection) {
    if m == nil {
        return nil
    } else {
        return m.direction
    }
}
func (m *MediaStream) GetLabel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.label
    }
}
func (m *MediaStream) GetMediaType()(*Modality) {
    if m == nil {
        return nil
    } else {
        return m.mediaType
    }
}
func (m *MediaStream) GetServerMuted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.serverMuted
    }
}
func (m *MediaStream) GetSourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceId
    }
}
func (m *MediaStream) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["direction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMediaDirection)
        if err != nil {
            return err
        }
        cast := val.(MediaDirection)
        m.SetDirection(&cast)
        return nil
    }
    res["label"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLabel(val)
        return nil
    }
    res["mediaType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseModality)
        if err != nil {
            return err
        }
        cast := val.(Modality)
        m.SetMediaType(&cast)
        return nil
    }
    res["serverMuted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetServerMuted(val)
        return nil
    }
    res["sourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSourceId(val)
        return nil
    }
    return res
}
func (m *MediaStream) IsNil()(bool) {
    return m == nil
}
func (m *MediaStream) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDirection() != nil {
        cast := m.GetDirection().String()
        err := writer.WriteStringValue("direction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("label", m.GetLabel())
        if err != nil {
            return err
        }
    }
    if m.GetMediaType() != nil {
        cast := m.GetMediaType().String()
        err := writer.WriteStringValue("mediaType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("serverMuted", m.GetServerMuted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceId", m.GetSourceId())
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
func (m *MediaStream) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MediaStream) SetDirection(value *MediaDirection)() {
    m.direction = value
}
func (m *MediaStream) SetLabel(value *string)() {
    m.label = value
}
func (m *MediaStream) SetMediaType(value *Modality)() {
    m.mediaType = value
}
func (m *MediaStream) SetServerMuted(value *bool)() {
    m.serverMuted = value
}
func (m *MediaStream) SetSourceId(value *string)() {
    m.sourceId = value
}
