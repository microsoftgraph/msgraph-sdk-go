package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// CallsItemUpdateRecordingStatusPostRequestBody provides operations to call the updateRecordingStatus method.
type CallsItemUpdateRecordingStatusPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The clientContext property
    clientContext *string
    // The status property
    status *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RecordingStatus
}
// NewCallsItemUpdateRecordingStatusPostRequestBody instantiates a new CallsItemUpdateRecordingStatusPostRequestBody and sets the default values.
func NewCallsItemUpdateRecordingStatusPostRequestBody()(*CallsItemUpdateRecordingStatusPostRequestBody) {
    m := &CallsItemUpdateRecordingStatusPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCallsItemUpdateRecordingStatusPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallsItemUpdateRecordingStatusPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallsItemUpdateRecordingStatusPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallsItemUpdateRecordingStatusPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetClientContext gets the clientContext property value. The clientContext property
func (m *CallsItemUpdateRecordingStatusPostRequestBody) GetClientContext()(*string) {
    return m.clientContext
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallsItemUpdateRecordingStatusPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clientContext"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientContext(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseRecordingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RecordingStatus))
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. The status property
func (m *CallsItemUpdateRecordingStatusPostRequestBody) GetStatus()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RecordingStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *CallsItemUpdateRecordingStatusPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("clientContext", m.GetClientContext())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *CallsItemUpdateRecordingStatusPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetClientContext sets the clientContext property value. The clientContext property
func (m *CallsItemUpdateRecordingStatusPostRequestBody) SetClientContext(value *string)() {
    m.clientContext = value
}
// SetStatus sets the status property value. The status property
func (m *CallsItemUpdateRecordingStatusPostRequestBody) SetStatus(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RecordingStatus)() {
    m.status = value
}
