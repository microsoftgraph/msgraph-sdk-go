package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TopicModelingSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // To learn more, see Adjust maximum number of themes dynamically.
    dynamicallyAdjustTopicCount *bool;
    // To learn more, see Include numbers in themes.
    ignoreNumbers *bool;
    // Indicates whether themes is enabled for the case.
    isEnabled *bool;
    // To learn more, see Maximum number of themes.
    topicCount *int32;
}
// Instantiates a new topicModelingSettings and sets the default values.
func NewTopicModelingSettings()(*TopicModelingSettings) {
    m := &TopicModelingSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TopicModelingSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the dynamicallyAdjustTopicCount property value. To learn more, see Adjust maximum number of themes dynamically.
func (m *TopicModelingSettings) GetDynamicallyAdjustTopicCount()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.dynamicallyAdjustTopicCount
    }
}
// Gets the ignoreNumbers property value. To learn more, see Include numbers in themes.
func (m *TopicModelingSettings) GetIgnoreNumbers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.ignoreNumbers
    }
}
// Gets the isEnabled property value. Indicates whether themes is enabled for the case.
func (m *TopicModelingSettings) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// Gets the topicCount property value. To learn more, see Maximum number of themes.
func (m *TopicModelingSettings) GetTopicCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.topicCount
    }
}
// The deserialization information for the current model
func (m *TopicModelingSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["dynamicallyAdjustTopicCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDynamicallyAdjustTopicCount(val)
        }
        return nil
    }
    res["ignoreNumbers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIgnoreNumbers(val)
        }
        return nil
    }
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["topicCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopicCount(val)
        }
        return nil
    }
    return res
}
func (m *TopicModelingSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TopicModelingSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("dynamicallyAdjustTopicCount", m.GetDynamicallyAdjustTopicCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ignoreNumbers", m.GetIgnoreNumbers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("topicCount", m.GetTopicCount())
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
func (m *TopicModelingSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the dynamicallyAdjustTopicCount property value. To learn more, see Adjust maximum number of themes dynamically.
// Parameters:
//  - value : Value to set for the dynamicallyAdjustTopicCount property.
func (m *TopicModelingSettings) SetDynamicallyAdjustTopicCount(value *bool)() {
    m.dynamicallyAdjustTopicCount = value
}
// Sets the ignoreNumbers property value. To learn more, see Include numbers in themes.
// Parameters:
//  - value : Value to set for the ignoreNumbers property.
func (m *TopicModelingSettings) SetIgnoreNumbers(value *bool)() {
    m.ignoreNumbers = value
}
// Sets the isEnabled property value. Indicates whether themes is enabled for the case.
// Parameters:
//  - value : Value to set for the isEnabled property.
func (m *TopicModelingSettings) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// Sets the topicCount property value. To learn more, see Maximum number of themes.
// Parameters:
//  - value : Value to set for the topicCount property.
func (m *TopicModelingSettings) SetTopicCount(value *int32)() {
    m.topicCount = value
}
