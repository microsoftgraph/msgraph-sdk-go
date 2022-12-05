package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody provides operations to call the onenotePatchContent method.
type MeOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The commands property
    commands []OnenotePatchContentCommandable
}
// NewMeOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody instantiates a new MeOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody and sets the default values.
func NewMeOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody()(*MeOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) {
    m := &MeOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCommands gets the commands property value. The commands property
func (m *MeOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) GetCommands()([]OnenotePatchContentCommandable) {
    return m.commands
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["commands"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOnenotePatchContentCommandFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnenotePatchContentCommandable, len(val))
            for i, v := range val {
                res[i] = v.(OnenotePatchContentCommandable)
            }
            m.SetCommands(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *MeOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCommands() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCommands()))
        for i, v := range m.GetCommands() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("commands", cast)
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
func (m *MeOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCommands sets the commands property value. The commands property
func (m *MeOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) SetCommands(value []OnenotePatchContentCommandable)() {
    m.commands = value
}
