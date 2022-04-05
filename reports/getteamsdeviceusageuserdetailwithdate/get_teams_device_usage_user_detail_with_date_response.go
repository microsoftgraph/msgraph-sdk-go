package getteamsdeviceusageuserdetailwithdate

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetTeamsDeviceUsageUserDetailWithDateResponse provides operations to call the getTeamsDeviceUsageUserDetail method.
type GetTeamsDeviceUsageUserDetailWithDateResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The value property
    value []byte;
}
// NewGetTeamsDeviceUsageUserDetailWithDateResponse instantiates a new getTeamsDeviceUsageUserDetailWithDateResponse and sets the default values.
func NewGetTeamsDeviceUsageUserDetailWithDateResponse()(*GetTeamsDeviceUsageUserDetailWithDateResponse) {
    m := &GetTeamsDeviceUsageUserDetailWithDateResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGetTeamsDeviceUsageUserDetailWithDateResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetTeamsDeviceUsageUserDetailWithDateResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetTeamsDeviceUsageUserDetailWithDateResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetTeamsDeviceUsageUserDetailWithDateResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetTeamsDeviceUsageUserDetailWithDateResponse) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["value"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetValue gets the value property value. The value property
func (m *GetTeamsDeviceUsageUserDetailWithDateResponse) GetValue()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// Serialize serializes information the current object
func (m *GetTeamsDeviceUsageUserDetailWithDateResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *GetTeamsDeviceUsageUserDetailWithDateResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetValue sets the value property value. The value property
func (m *GetTeamsDeviceUsageUserDetailWithDateResponse) SetValue(value []byte)() {
    if m != nil {
        m.value = value
    }
}
