package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkActivityTopic provides operations to call the sendActivityNotification method.
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
func CreateTeamworkActivityTopicFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
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
func (m *TeamworkActivityTopic) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["source"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkActivityTopicSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(*TeamworkActivityTopicSource))
        }
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *TeamworkActivityTopic) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkActivityTopic) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
