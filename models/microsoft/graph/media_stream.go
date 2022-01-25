package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MediaStream 
type MediaStream struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The direction. The possible values are inactive, sendOnly, receiveOnly, sendReceive.
    direction *MediaDirection;
    // The media stream label.
    label *string;
    // The media type. The possible value are unknown, audio, video, videoBasedScreenSharing, data.
    mediaType *Modality;
    // Indicates whether the media is muted by the server.
    serverMuted *bool;
    // The source ID.
    sourceId *string;
}
// NewMediaStream instantiates a new mediaStream and sets the default values.
func NewMediaStream()(*MediaStream) {
    m := &MediaStream{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MediaStream) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDirection gets the direction property value. The direction. The possible values are inactive, sendOnly, receiveOnly, sendReceive.
func (m *MediaStream) GetDirection()(*MediaDirection) {
    if m == nil {
        return nil
    } else {
        return m.direction
    }
}
// GetLabel gets the label property value. The media stream label.
func (m *MediaStream) GetLabel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.label
    }
}
// GetMediaType gets the mediaType property value. The media type. The possible value are unknown, audio, video, videoBasedScreenSharing, data.
func (m *MediaStream) GetMediaType()(*Modality) {
    if m == nil {
        return nil
    } else {
        return m.mediaType
    }
}
// GetServerMuted gets the serverMuted property value. Indicates whether the media is muted by the server.
func (m *MediaStream) GetServerMuted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.serverMuted
    }
}
// GetSourceId gets the sourceId property value. The source ID.
func (m *MediaStream) GetSourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MediaStream) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["direction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMediaDirection)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(MediaDirection)
            m.SetDirection(&cast)
        }
        return nil
    }
    res["label"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val)
        }
        return nil
    }
    res["mediaType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseModality)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(Modality)
            m.SetMediaType(&cast)
        }
        return nil
    }
    res["serverMuted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerMuted(val)
        }
        return nil
    }
    res["sourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceId(val)
        }
        return nil
    }
    return res
}
func (m *MediaStream) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MediaStream) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDirection sets the direction property value. The direction. The possible values are inactive, sendOnly, receiveOnly, sendReceive.
func (m *MediaStream) SetDirection(value *MediaDirection)() {
    if m != nil {
        m.direction = value
    }
}
// SetLabel sets the label property value. The media stream label.
func (m *MediaStream) SetLabel(value *string)() {
    if m != nil {
        m.label = value
    }
}
// SetMediaType sets the mediaType property value. The media type. The possible value are unknown, audio, video, videoBasedScreenSharing, data.
func (m *MediaStream) SetMediaType(value *Modality)() {
    if m != nil {
        m.mediaType = value
    }
}
// SetServerMuted sets the serverMuted property value. Indicates whether the media is muted by the server.
func (m *MediaStream) SetServerMuted(value *bool)() {
    if m != nil {
        m.serverMuted = value
    }
}
// SetSourceId sets the sourceId property value. The source ID.
func (m *MediaStream) SetSourceId(value *string)() {
    if m != nil {
        m.sourceId = value
    }
}
