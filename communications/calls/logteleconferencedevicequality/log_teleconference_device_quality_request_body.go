package logteleconferencedevicequality

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// LogTeleconferenceDeviceQualityRequestBody provides operations to call the logTeleconferenceDeviceQuality method.
type LogTeleconferenceDeviceQualityRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    quality i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TeleconferenceDeviceQualityable;
}
// NewLogTeleconferenceDeviceQualityRequestBody instantiates a new logTeleconferenceDeviceQualityRequestBody and sets the default values.
func NewLogTeleconferenceDeviceQualityRequestBody()(*LogTeleconferenceDeviceQualityRequestBody) {
    m := &LogTeleconferenceDeviceQualityRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateLogTeleconferenceDeviceQualityRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLogTeleconferenceDeviceQualityRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewLogTeleconferenceDeviceQualityRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LogTeleconferenceDeviceQualityRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LogTeleconferenceDeviceQualityRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["quality"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateTeleconferenceDeviceQualityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuality(val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TeleconferenceDeviceQualityable))
        }
        return nil
    }
    return res
}
// GetQuality gets the quality property value. 
func (m *LogTeleconferenceDeviceQualityRequestBody) GetQuality()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TeleconferenceDeviceQualityable) {
    if m == nil {
        return nil
    } else {
        return m.quality
    }
}
// Serialize serializes information the current object
func (m *LogTeleconferenceDeviceQualityRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("quality", m.GetQuality())
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
func (m *LogTeleconferenceDeviceQualityRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetQuality sets the quality property value. 
func (m *LogTeleconferenceDeviceQualityRequestBody) SetQuality(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TeleconferenceDeviceQualityable)() {
    if m != nil {
        m.quality = value
    }
}
