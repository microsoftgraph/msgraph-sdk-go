package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupsItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody provides operations to call the getNotebookFromWebUrl method.
type GroupsItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The webUrl property
    webUrl *string
}
// NewGroupsItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody instantiates a new GroupsItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody and sets the default values.
func NewGroupsItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody()(*GroupsItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody) {
    m := &GroupsItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGroupsItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupsItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupsItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupsItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupsItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetWebUrl gets the webUrl property value. The webUrl property
func (m *GroupsItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *GroupsItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("webUrl", m.GetWebUrl())
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
func (m *GroupsItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetWebUrl sets the webUrl property value. The webUrl property
func (m *GroupsItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody) SetWebUrl(value *string)() {
    m.webUrl = value
}
