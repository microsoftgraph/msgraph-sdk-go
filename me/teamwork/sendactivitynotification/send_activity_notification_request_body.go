package sendactivitynotification

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// SendActivityNotificationRequestBody provides operations to call the sendActivityNotification method.
type SendActivityNotificationRequestBody struct {
    // 
    activityType *string;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    chainId *int64;
    // 
    previewText i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ItemBodyable;
    // 
    templateParameters []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.KeyValuePairable;
    // 
    topic i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TeamworkActivityTopicable;
}
// NewSendActivityNotificationRequestBody instantiates a new sendActivityNotificationRequestBody and sets the default values.
func NewSendActivityNotificationRequestBody()(*SendActivityNotificationRequestBody) {
    m := &SendActivityNotificationRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSendActivityNotificationRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSendActivityNotificationRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSendActivityNotificationRequestBody(), nil
}
// GetActivityType gets the activityType property value. 
func (m *SendActivityNotificationRequestBody) GetActivityType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activityType
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SendActivityNotificationRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetChainId gets the chainId property value. 
func (m *SendActivityNotificationRequestBody) GetChainId()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.chainId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SendActivityNotificationRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["activityType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityType(val)
        }
        return nil
    }
    res["chainId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChainId(val)
        }
        return nil
    }
    res["previewText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewText(val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ItemBodyable))
        }
        return nil
    }
    res["templateParameters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.KeyValuePairable)
            }
            m.SetTemplateParameters(res)
        }
        return nil
    }
    res["topic"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateTeamworkActivityTopicFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopic(val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TeamworkActivityTopicable))
        }
        return nil
    }
    return res
}
// GetPreviewText gets the previewText property value. 
func (m *SendActivityNotificationRequestBody) GetPreviewText()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ItemBodyable) {
    if m == nil {
        return nil
    } else {
        return m.previewText
    }
}
// GetTemplateParameters gets the templateParameters property value. 
func (m *SendActivityNotificationRequestBody) GetTemplateParameters()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.KeyValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.templateParameters
    }
}
// GetTopic gets the topic property value. 
func (m *SendActivityNotificationRequestBody) GetTopic()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TeamworkActivityTopicable) {
    if m == nil {
        return nil
    } else {
        return m.topic
    }
}
// Serialize serializes information the current object
func (m *SendActivityNotificationRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("activityType", m.GetActivityType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("chainId", m.GetChainId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("previewText", m.GetPreviewText())
        if err != nil {
            return err
        }
    }
    if m.GetTemplateParameters() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTemplateParameters()))
        for i, v := range m.GetTemplateParameters() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("templateParameters", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("topic", m.GetTopic())
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
// SetActivityType sets the activityType property value. 
func (m *SendActivityNotificationRequestBody) SetActivityType(value *string)() {
    if m != nil {
        m.activityType = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SendActivityNotificationRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetChainId sets the chainId property value. 
func (m *SendActivityNotificationRequestBody) SetChainId(value *int64)() {
    if m != nil {
        m.chainId = value
    }
}
// SetPreviewText sets the previewText property value. 
func (m *SendActivityNotificationRequestBody) SetPreviewText(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ItemBodyable)() {
    if m != nil {
        m.previewText = value
    }
}
// SetTemplateParameters sets the templateParameters property value. 
func (m *SendActivityNotificationRequestBody) SetTemplateParameters(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.KeyValuePairable)() {
    if m != nil {
        m.templateParameters = value
    }
}
// SetTopic sets the topic property value. 
func (m *SendActivityNotificationRequestBody) SetTopic(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TeamworkActivityTopicable)() {
    if m != nil {
        m.topic = value
    }
}
