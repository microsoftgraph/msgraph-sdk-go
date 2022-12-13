package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemCalendarsItemEventsItemInstancesItemForwardPostRequestBody provides operations to call the forward method.
type ItemCalendarsItemEventsItemInstancesItemForwardPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Comment property
    comment *string
    // The ToRecipients property
    toRecipients []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Recipientable
}
// NewItemCalendarsItemEventsItemInstancesItemForwardPostRequestBody instantiates a new ItemCalendarsItemEventsItemInstancesItemForwardPostRequestBody and sets the default values.
func NewItemCalendarsItemEventsItemInstancesItemForwardPostRequestBody()(*ItemCalendarsItemEventsItemInstancesItemForwardPostRequestBody) {
    m := &ItemCalendarsItemEventsItemInstancesItemForwardPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemCalendarsItemEventsItemInstancesItemForwardPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarsItemEventsItemInstancesItemForwardPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarsItemEventsItemInstancesItemForwardPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCalendarsItemEventsItemInstancesItemForwardPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetComment gets the comment property value. The Comment property
func (m *ItemCalendarsItemEventsItemInstancesItemForwardPostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCalendarsItemEventsItemInstancesItemForwardPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["toRecipients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Recipientable, len(val))
            for i, v := range val {
                res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Recipientable)
            }
            m.SetToRecipients(res)
        }
        return nil
    }
    return res
}
// GetToRecipients gets the toRecipients property value. The ToRecipients property
func (m *ItemCalendarsItemEventsItemInstancesItemForwardPostRequestBody) GetToRecipients()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Recipientable) {
    return m.toRecipients
}
// Serialize serializes information the current object
func (m *ItemCalendarsItemEventsItemInstancesItemForwardPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    if m.GetToRecipients() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetToRecipients()))
        for i, v := range m.GetToRecipients() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("toRecipients", cast)
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
func (m *ItemCalendarsItemEventsItemInstancesItemForwardPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetComment sets the comment property value. The Comment property
func (m *ItemCalendarsItemEventsItemInstancesItemForwardPostRequestBody) SetComment(value *string)() {
    m.comment = value
}
// SetToRecipients sets the toRecipients property value. The ToRecipients property
func (m *ItemCalendarsItemEventsItemInstancesItemForwardPostRequestBody) SetToRecipients(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Recipientable)() {
    m.toRecipients = value
}
