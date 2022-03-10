package getdevicemanagementintentpersettingcontributingprofiles

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GetDeviceManagementIntentPerSettingContributingProfilesResponse provides operations to call the getDeviceManagementIntentPerSettingContributingProfiles method.
type GetDeviceManagementIntentPerSettingContributingProfilesResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    value []byte;
}
// NewGetDeviceManagementIntentPerSettingContributingProfilesResponse instantiates a new getDeviceManagementIntentPerSettingContributingProfilesResponse and sets the default values.
func NewGetDeviceManagementIntentPerSettingContributingProfilesResponse()(*GetDeviceManagementIntentPerSettingContributingProfilesResponse) {
    m := &GetDeviceManagementIntentPerSettingContributingProfilesResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGetDeviceManagementIntentPerSettingContributingProfilesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetDeviceManagementIntentPerSettingContributingProfilesResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewGetDeviceManagementIntentPerSettingContributingProfilesResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetDeviceManagementIntentPerSettingContributingProfilesResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetDeviceManagementIntentPerSettingContributingProfilesResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. 
func (m *GetDeviceManagementIntentPerSettingContributingProfilesResponse) GetValue()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *GetDeviceManagementIntentPerSettingContributingProfilesResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetDeviceManagementIntentPerSettingContributingProfilesResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("value", m.GetValue())
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
func (m *GetDeviceManagementIntentPerSettingContributingProfilesResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetValue sets the value property value. 
func (m *GetDeviceManagementIntentPerSettingContributingProfilesResponse) SetValue(value []byte)() {
    if m != nil {
        m.value = value
    }
}
