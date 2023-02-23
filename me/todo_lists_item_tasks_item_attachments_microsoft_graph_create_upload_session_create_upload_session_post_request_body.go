package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody 
type TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody instantiates a new TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody and sets the default values.
func NewTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody()(*TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody) {
    m := &TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAttachmentInfo gets the attachmentInfo property value. The attachmentInfo property
func (m *TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody) GetAttachmentInfo()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentInfoable) {
    val, err := m.GetBackingStore().Get("attachmentInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentInfoable)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attachmentInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAttachmentInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttachmentInfo(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentInfoable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("attachmentInfo", m.GetAttachmentInfo())
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
func (m *TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAttachmentInfo sets the attachmentInfo property value. The attachmentInfo property
func (m *TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody) SetAttachmentInfo(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentInfoable)() {
    err := m.GetBackingStore().Set("attachmentInfo", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBodyable 
type TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttachmentInfo()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentInfoable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    SetAttachmentInfo(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentInfoable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
}
