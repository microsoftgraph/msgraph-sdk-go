package answer

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// AnswerRequestBody 
type AnswerRequestBody struct {
    // 
    acceptedModalities []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Modality;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    callbackUri *string;
    // 
    callOptions *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.IncomingCallOptions;
    // 
    mediaConfig *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MediaConfig;
    // 
    participantCapacity *int32;
}
// NewAnswerRequestBody instantiates a new answerRequestBody and sets the default values.
func NewAnswerRequestBody()(*AnswerRequestBody) {
    m := &AnswerRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAcceptedModalities gets the acceptedModalities property value. 
func (m *AnswerRequestBody) GetAcceptedModalities()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Modality) {
    if m == nil {
        return nil
    } else {
        return m.acceptedModalities
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AnswerRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCallbackUri gets the callbackUri property value. 
func (m *AnswerRequestBody) GetCallbackUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.callbackUri
    }
}
// GetCallOptions gets the callOptions property value. 
func (m *AnswerRequestBody) GetCallOptions()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.IncomingCallOptions) {
    if m == nil {
        return nil
    } else {
        return m.callOptions
    }
}
// GetMediaConfig gets the mediaConfig property value. 
func (m *AnswerRequestBody) GetMediaConfig()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MediaConfig) {
    if m == nil {
        return nil
    } else {
        return m.mediaConfig
    }
}
// GetParticipantCapacity gets the participantCapacity property value. 
func (m *AnswerRequestBody) GetParticipantCapacity()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.participantCapacity
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AnswerRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["acceptedModalities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ParseModality)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Modality, len(val))
            for i, v := range val {
                res[i] = *(v.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Modality))
            }
            m.SetAcceptedModalities(res)
        }
        return nil
    }
    res["callbackUri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallbackUri(val)
        }
        return nil
    }
    res["callOptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewIncomingCallOptions() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallOptions(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.IncomingCallOptions))
        }
        return nil
    }
    res["mediaConfig"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewMediaConfig() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaConfig(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MediaConfig))
        }
        return nil
    }
    res["participantCapacity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipantCapacity(val)
        }
        return nil
    }
    return res
}
func (m *AnswerRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AnswerRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAcceptedModalities() != nil {
        err := writer.WriteCollectionOfStringValues("acceptedModalities", i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SerializeModality(m.GetAcceptedModalities()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("callbackUri", m.GetCallbackUri())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("callOptions", m.GetCallOptions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("mediaConfig", m.GetMediaConfig())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("participantCapacity", m.GetParticipantCapacity())
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
// SetAcceptedModalities sets the acceptedModalities property value. 
func (m *AnswerRequestBody) SetAcceptedModalities(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Modality)() {
    if m != nil {
        m.acceptedModalities = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AnswerRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCallbackUri sets the callbackUri property value. 
func (m *AnswerRequestBody) SetCallbackUri(value *string)() {
    if m != nil {
        m.callbackUri = value
    }
}
// SetCallOptions sets the callOptions property value. 
func (m *AnswerRequestBody) SetCallOptions(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.IncomingCallOptions)() {
    if m != nil {
        m.callOptions = value
    }
}
// SetMediaConfig sets the mediaConfig property value. 
func (m *AnswerRequestBody) SetMediaConfig(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MediaConfig)() {
    if m != nil {
        m.mediaConfig = value
    }
}
// SetParticipantCapacity sets the participantCapacity property value. 
func (m *AnswerRequestBody) SetParticipantCapacity(value *int32)() {
    if m != nil {
        m.participantCapacity = value
    }
}
