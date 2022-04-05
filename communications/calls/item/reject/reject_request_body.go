package reject

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// RejectRequestBody provides operations to call the reject method.
type RejectRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The callbackUri property
    callbackUri *string;
    // The reason property
    reason *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RejectReason;
}
// NewRejectRequestBody instantiates a new rejectRequestBody and sets the default values.
func NewRejectRequestBody()(*RejectRequestBody) {
    m := &RejectRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRejectRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRejectRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRejectRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RejectRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCallbackUri gets the callbackUri property value. The callbackUri property
func (m *RejectRequestBody) GetCallbackUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.callbackUri
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RejectRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["callbackUri"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallbackUri(val)
        }
        return nil
    }
    res["reason"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseRejectReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RejectReason))
        }
        return nil
    }
    return res
}
// GetReason gets the reason property value. The reason property
func (m *RejectRequestBody) GetReason()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RejectReason) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
// Serialize serializes information the current object
func (m *RejectRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("callbackUri", m.GetCallbackUri())
        if err != nil {
            return err
        }
    }
    if m.GetReason() != nil {
        cast := (*m.GetReason()).String()
        err := writer.WriteStringValue("reason", &cast)
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
func (m *RejectRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCallbackUri sets the callbackUri property value. The callbackUri property
func (m *RejectRequestBody) SetCallbackUri(value *string)() {
    if m != nil {
        m.callbackUri = value
    }
}
// SetReason sets the reason property value. The reason property
func (m *RejectRequestBody) SetReason(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RejectReason)() {
    if m != nil {
        m.reason = value
    }
}
