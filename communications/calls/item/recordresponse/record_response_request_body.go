package recordresponse

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// RecordResponseRequestBody 
type RecordResponseRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    bargeInAllowed *bool;
    // 
    clientContext *string;
    // 
    initialSilenceTimeoutInSeconds *int32;
    // 
    maxRecordDurationInSeconds *int32;
    // 
    maxSilenceTimeoutInSeconds *int32;
    // 
    playBeep *bool;
    // 
    prompts []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Prompt;
    // 
    stopTones []string;
}
// NewRecordResponseRequestBody instantiates a new recordResponseRequestBody and sets the default values.
func NewRecordResponseRequestBody()(*RecordResponseRequestBody) {
    m := &RecordResponseRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecordResponseRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBargeInAllowed gets the bargeInAllowed property value. 
func (m *RecordResponseRequestBody) GetBargeInAllowed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.bargeInAllowed
    }
}
// GetClientContext gets the clientContext property value. 
func (m *RecordResponseRequestBody) GetClientContext()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientContext
    }
}
// GetInitialSilenceTimeoutInSeconds gets the initialSilenceTimeoutInSeconds property value. 
func (m *RecordResponseRequestBody) GetInitialSilenceTimeoutInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.initialSilenceTimeoutInSeconds
    }
}
// GetMaxRecordDurationInSeconds gets the maxRecordDurationInSeconds property value. 
func (m *RecordResponseRequestBody) GetMaxRecordDurationInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxRecordDurationInSeconds
    }
}
// GetMaxSilenceTimeoutInSeconds gets the maxSilenceTimeoutInSeconds property value. 
func (m *RecordResponseRequestBody) GetMaxSilenceTimeoutInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxSilenceTimeoutInSeconds
    }
}
// GetPlayBeep gets the playBeep property value. 
func (m *RecordResponseRequestBody) GetPlayBeep()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.playBeep
    }
}
// GetPrompts gets the prompts property value. 
func (m *RecordResponseRequestBody) GetPrompts()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Prompt) {
    if m == nil {
        return nil
    } else {
        return m.prompts
    }
}
// GetStopTones gets the stopTones property value. 
func (m *RecordResponseRequestBody) GetStopTones()([]string) {
    if m == nil {
        return nil
    } else {
        return m.stopTones
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RecordResponseRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["bargeInAllowed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBargeInAllowed(val)
        }
        return nil
    }
    res["clientContext"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientContext(val)
        }
        return nil
    }
    res["initialSilenceTimeoutInSeconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitialSilenceTimeoutInSeconds(val)
        }
        return nil
    }
    res["maxRecordDurationInSeconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxRecordDurationInSeconds(val)
        }
        return nil
    }
    res["maxSilenceTimeoutInSeconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxSilenceTimeoutInSeconds(val)
        }
        return nil
    }
    res["playBeep"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlayBeep(val)
        }
        return nil
    }
    res["prompts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPrompt() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Prompt, len(val))
            for i, v := range val {
                res[i] = *(v.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Prompt))
            }
            m.SetPrompts(res)
        }
        return nil
    }
    res["stopTones"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetStopTones(res)
        }
        return nil
    }
    return res
}
func (m *RecordResponseRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RecordResponseRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("bargeInAllowed", m.GetBargeInAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("clientContext", m.GetClientContext())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("initialSilenceTimeoutInSeconds", m.GetInitialSilenceTimeoutInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxRecordDurationInSeconds", m.GetMaxRecordDurationInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxSilenceTimeoutInSeconds", m.GetMaxSilenceTimeoutInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("playBeep", m.GetPlayBeep())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPrompts()))
        for i, v := range m.GetPrompts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("prompts", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("stopTones", m.GetStopTones())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecordResponseRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBargeInAllowed sets the bargeInAllowed property value. 
func (m *RecordResponseRequestBody) SetBargeInAllowed(value *bool)() {
    if m != nil {
        m.bargeInAllowed = value
    }
}
// SetClientContext sets the clientContext property value. 
func (m *RecordResponseRequestBody) SetClientContext(value *string)() {
    if m != nil {
        m.clientContext = value
    }
}
// SetInitialSilenceTimeoutInSeconds sets the initialSilenceTimeoutInSeconds property value. 
func (m *RecordResponseRequestBody) SetInitialSilenceTimeoutInSeconds(value *int32)() {
    if m != nil {
        m.initialSilenceTimeoutInSeconds = value
    }
}
// SetMaxRecordDurationInSeconds sets the maxRecordDurationInSeconds property value. 
func (m *RecordResponseRequestBody) SetMaxRecordDurationInSeconds(value *int32)() {
    if m != nil {
        m.maxRecordDurationInSeconds = value
    }
}
// SetMaxSilenceTimeoutInSeconds sets the maxSilenceTimeoutInSeconds property value. 
func (m *RecordResponseRequestBody) SetMaxSilenceTimeoutInSeconds(value *int32)() {
    if m != nil {
        m.maxSilenceTimeoutInSeconds = value
    }
}
// SetPlayBeep sets the playBeep property value. 
func (m *RecordResponseRequestBody) SetPlayBeep(value *bool)() {
    if m != nil {
        m.playBeep = value
    }
}
// SetPrompts sets the prompts property value. 
func (m *RecordResponseRequestBody) SetPrompts(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Prompt)() {
    if m != nil {
        m.prompts = value
    }
}
// SetStopTones sets the stopTones property value. 
func (m *RecordResponseRequestBody) SetStopTones(value []string)() {
    if m != nil {
        m.stopTones = value
    }
}
