package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkActivityTopic 
type TeamworkActivityTopic struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Type of source. Possible values are: entityUrl, text. For supported Microsoft Graph URLs, use entityUrl. For custom text, use text.
    source *TeamworkActivityTopicSource;
    // The topic value. If the value of the source property is entityUrl, this must be a Microsoft Graph URL. If the vaule is text, this must be a plain text value.
    value *string;
    // The link the user clicks when they select the notification. Optional when source is entityUrl; required when source is text.
    webUrl *string;
}
// NewTeamworkActivityTopic instantiates a new teamworkActivityTopic and sets the default values.
func NewTeamworkActivityTopic()(*TeamworkActivityTopic) {
    m := &TeamworkActivityTopic{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkActivityTopicFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkActivityTopicFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkActivityTopic(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkActivityTopic) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkActivityTopic) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["source"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkActivityTopicSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(*TeamworkActivityTopicSource))
        }
        return nil
    }
    res["value"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    res["webUrl"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetSource gets the source property value. Type of source. Possible values are: entityUrl, text. For supported Microsoft Graph URLs, use entityUrl. For custom text, use text.
func (m *TeamworkActivityTopic) GetSource()(*TeamworkActivityTopicSource) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
// GetValue gets the value property value. The topic value. If the value of the source property is entityUrl, this must be a Microsoft Graph URL. If the vaule is text, this must be a plain text value.
func (m *TeamworkActivityTopic) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// GetWebUrl gets the webUrl property value. The link the user clicks when they select the notification. Optional when source is entityUrl; required when source is text.
func (m *TeamworkActivityTopic) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// Serialize serializes information the current object
func (m *TeamworkActivityTopic) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSource() != nil {
        cast := (*m.GetSource()).String()
        err := writer.WriteStringValue("source", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("webUrl", m.GetWebUrl())
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
func (m *TeamworkActivityTopic) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetSource sets the source property value. Type of source. Possible values are: entityUrl, text. For supported Microsoft Graph URLs, use entityUrl. For custom text, use text.
func (m *TeamworkActivityTopic) SetSource(value *TeamworkActivityTopicSource)() {
    if m != nil {
        m.source = value
    }
}
// SetValue sets the value property value. The topic value. If the value of the source property is entityUrl, this must be a Microsoft Graph URL. If the vaule is text, this must be a plain text value.
func (m *TeamworkActivityTopic) SetValue(value *string)() {
    if m != nil {
        m.value = value
    }
}
// SetWebUrl sets the webUrl property value. The link the user clicks when they select the notification. Optional when source is entityUrl; required when source is text.
func (m *TeamworkActivityTopic) SetWebUrl(value *string)() {
    if m != nil {
        m.webUrl = value
    }
}
