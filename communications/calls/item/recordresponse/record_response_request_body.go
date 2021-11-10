package recordresponse

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// 
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
// Instantiates a new recordResponseRequestBody and sets the default values.
func NewRecordResponseRequestBody()(*RecordResponseRequestBody) {
    m := &RecordResponseRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecordResponseRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the bargeInAllowed property value. 
func (m *RecordResponseRequestBody) GetBargeInAllowed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.bargeInAllowed
    }
}
// Gets the clientContext property value. 
func (m *RecordResponseRequestBody) GetClientContext()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientContext
    }
}
// Gets the initialSilenceTimeoutInSeconds property value. 
func (m *RecordResponseRequestBody) GetInitialSilenceTimeoutInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.initialSilenceTimeoutInSeconds
    }
}
// Gets the maxRecordDurationInSeconds property value. 
func (m *RecordResponseRequestBody) GetMaxRecordDurationInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxRecordDurationInSeconds
    }
}
// Gets the maxSilenceTimeoutInSeconds property value. 
func (m *RecordResponseRequestBody) GetMaxSilenceTimeoutInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxSilenceTimeoutInSeconds
    }
}
// Gets the playBeep property value. 
func (m *RecordResponseRequestBody) GetPlayBeep()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.playBeep
    }
}
// Gets the prompts property value. 
func (m *RecordResponseRequestBody) GetPrompts()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Prompt) {
    if m == nil {
        return nil
    } else {
        return m.prompts
    }
}
// Gets the stopTones property value. 
func (m *RecordResponseRequestBody) GetStopTones()([]string) {
    if m == nil {
        return nil
    } else {
        return m.stopTones
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *RecordResponseRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the bargeInAllowed property value. 
// Parameters:
//  - value : Value to set for the bargeInAllowed property.
func (m *RecordResponseRequestBody) SetBargeInAllowed(value *bool)() {
    m.bargeInAllowed = value
}
// Sets the clientContext property value. 
// Parameters:
//  - value : Value to set for the clientContext property.
func (m *RecordResponseRequestBody) SetClientContext(value *string)() {
    m.clientContext = value
}
// Sets the initialSilenceTimeoutInSeconds property value. 
// Parameters:
//  - value : Value to set for the initialSilenceTimeoutInSeconds property.
func (m *RecordResponseRequestBody) SetInitialSilenceTimeoutInSeconds(value *int32)() {
    m.initialSilenceTimeoutInSeconds = value
}
// Sets the maxRecordDurationInSeconds property value. 
// Parameters:
//  - value : Value to set for the maxRecordDurationInSeconds property.
func (m *RecordResponseRequestBody) SetMaxRecordDurationInSeconds(value *int32)() {
    m.maxRecordDurationInSeconds = value
}
// Sets the maxSilenceTimeoutInSeconds property value. 
// Parameters:
//  - value : Value to set for the maxSilenceTimeoutInSeconds property.
func (m *RecordResponseRequestBody) SetMaxSilenceTimeoutInSeconds(value *int32)() {
    m.maxSilenceTimeoutInSeconds = value
}
// Sets the playBeep property value. 
// Parameters:
//  - value : Value to set for the playBeep property.
func (m *RecordResponseRequestBody) SetPlayBeep(value *bool)() {
    m.playBeep = value
}
// Sets the prompts property value. 
// Parameters:
//  - value : Value to set for the prompts property.
func (m *RecordResponseRequestBody) SetPrompts(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Prompt)() {
    m.prompts = value
}
// Sets the stopTones property value. 
// Parameters:
//  - value : Value to set for the stopTones property.
func (m *RecordResponseRequestBody) SetStopTones(value []string)() {
    m.stopTones = value
}
