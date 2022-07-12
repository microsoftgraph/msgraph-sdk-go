package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MediaStream 
type MediaStream struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The direction property
    direction *MediaDirection
    // The media stream label.
    label *string
    // The mediaType property
    mediaType *Modality
    // Indicates whether the media is muted by the server.
    serverMuted *bool
    // The source ID.
    sourceId *string
}
// NewMediaStream instantiates a new mediaStream and sets the default values.
func NewMediaStream()(*MediaStream) {
    m := &MediaStream{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMediaStreamFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMediaStreamFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMediaStream(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MediaStream) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDirection gets the direction property value. The direction property
func (m *MediaStream) GetDirection()(*MediaDirection) {
    if m == nil {
        return nil
    } else {
        return m.direction
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MediaStream) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["direction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMediaDirection)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirection(val.(*MediaDirection))
        }
        return nil
    }
    res["label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val)
        }
        return nil
    }
    res["mediaType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseModality)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaType(val.(*Modality))
        }
        return nil
    }
    res["serverMuted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerMuted(val)
        }
        return nil
    }
    res["sourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetLabel gets the label property value. The media stream label.
func (m *MediaStream) GetLabel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.label
    }
}
// GetMediaType gets the mediaType property value. The mediaType property
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
// Serialize serializes information the current object
func (m *MediaStream) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDirection() != nil {
        cast := (*m.GetDirection()).String()
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
        cast := (*m.GetMediaType()).String()
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
// SetDirection sets the direction property value. The direction property
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
// SetMediaType sets the mediaType property value. The mediaType property
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
