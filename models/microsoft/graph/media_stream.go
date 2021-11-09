package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MediaStream struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The direction. The possible values are inactive, sendOnly, receiveOnly, sendReceive.
    direction *MediaDirection;
    // The media stream label.
    label *string;
    // The media type. The possible value are unknown, audio, video, videoBasedScreenSharing, data.
    mediaType *Modality;
    // If the media is muted by the server.
    serverMuted *bool;
    // The source ID.
    sourceId *string;
}
// Instantiates a new mediaStream and sets the default values.
func NewMediaStream()(*MediaStream) {
    m := &MediaStream{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MediaStream) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the direction property value. The direction. The possible values are inactive, sendOnly, receiveOnly, sendReceive.
func (m *MediaStream) GetDirection()(*MediaDirection) {
    if m == nil {
        return nil
    } else {
        return m.direction
    }
}
// Gets the label property value. The media stream label.
func (m *MediaStream) GetLabel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.label
    }
}
// Gets the mediaType property value. The media type. The possible value are unknown, audio, video, videoBasedScreenSharing, data.
func (m *MediaStream) GetMediaType()(*Modality) {
    if m == nil {
        return nil
    } else {
        return m.mediaType
    }
}
// Gets the serverMuted property value. If the media is muted by the server.
func (m *MediaStream) GetServerMuted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.serverMuted
    }
}
// Gets the sourceId property value. The source ID.
func (m *MediaStream) GetSourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceId
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *MediaStream) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the direction property value. The direction. The possible values are inactive, sendOnly, receiveOnly, sendReceive.
// Parameters:
//  - value : Value to set for the direction property.
func (m *MediaStream) SetDirection(value *MediaDirection)() {
    m.direction = value
}
// Sets the label property value. The media stream label.
// Parameters:
//  - value : Value to set for the label property.
func (m *MediaStream) SetLabel(value *string)() {
    m.label = value
}
// Sets the mediaType property value. The media type. The possible value are unknown, audio, video, videoBasedScreenSharing, data.
// Parameters:
//  - value : Value to set for the mediaType property.
func (m *MediaStream) SetMediaType(value *Modality)() {
    m.mediaType = value
}
// Sets the serverMuted property value. If the media is muted by the server.
// Parameters:
//  - value : Value to set for the serverMuted property.
func (m *MediaStream) SetServerMuted(value *bool)() {
    m.serverMuted = value
}
// Sets the sourceId property value. The source ID.
// Parameters:
//  - value : Value to set for the sourceId property.
func (m *MediaStream) SetSourceId(value *string)() {
    m.sourceId = value
}
