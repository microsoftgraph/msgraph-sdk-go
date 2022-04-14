package reply

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ReplyRequestBody provides operations to call the reply method.
type ReplyRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Post property
    post iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Postable
}
// NewReplyRequestBody instantiates a new replyRequestBody and sets the default values.
func NewReplyRequestBody()(*ReplyRequestBody) {
    m := &ReplyRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateReplyRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReplyRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReplyRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReplyRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReplyRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["post"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePostFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPost(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Postable))
        }
        return nil
    }
    return res
}
// GetPost gets the post property value. The Post property
func (m *ReplyRequestBody) GetPost()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Postable) {
    if m == nil {
        return nil
    } else {
        return m.post
    }
}
// Serialize serializes information the current object
func (m *ReplyRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("post", m.GetPost())
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
func (m *ReplyRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPost sets the post property value. The Post property
func (m *ReplyRequestBody) SetPost(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Postable)() {
    if m != nil {
        m.post = value
    }
}
