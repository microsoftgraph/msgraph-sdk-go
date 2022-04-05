package logteleconferencedevicequality

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// LogTeleconferenceDeviceQualityRequestBody provides operations to call the logTeleconferenceDeviceQuality method.
type LogTeleconferenceDeviceQualityRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The quality property
    quality iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeleconferenceDeviceQualityable;
}
// NewLogTeleconferenceDeviceQualityRequestBody instantiates a new logTeleconferenceDeviceQualityRequestBody and sets the default values.
func NewLogTeleconferenceDeviceQualityRequestBody()(*LogTeleconferenceDeviceQualityRequestBody) {
    m := &LogTeleconferenceDeviceQualityRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateLogTeleconferenceDeviceQualityRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLogTeleconferenceDeviceQualityRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
func (m *LogTeleconferenceDeviceQualityRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["quality"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeleconferenceDeviceQualityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuality(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeleconferenceDeviceQualityable))
        }
        return nil
    }
    return res
}
// GetQuality gets the quality property value. The quality property
func (m *LogTeleconferenceDeviceQualityRequestBody) GetQuality()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeleconferenceDeviceQualityable) {
    if m == nil {
        return nil
    } else {
        return m.quality
    }
}
// Serialize serializes information the current object
func (m *LogTeleconferenceDeviceQualityRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetQuality sets the quality property value. The quality property
func (m *LogTeleconferenceDeviceQualityRequestBody) SetQuality(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeleconferenceDeviceQualityable)() {
    if m != nil {
        m.quality = value
    }
}
