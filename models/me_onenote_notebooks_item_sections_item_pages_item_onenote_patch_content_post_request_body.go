package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeOnenoteNotebooksItemSectionsItemPagesItemOnenotePatchContentPostRequestBody provides operations to call the onenotePatchContent method.
type MeOnenoteNotebooksItemSectionsItemPagesItemOnenotePatchContentPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The commands property
    commands []OnenotePatchContentCommandable
}
// NewMeOnenoteNotebooksItemSectionsItemPagesItemOnenotePatchContentPostRequestBody instantiates a new MeOnenoteNotebooksItemSectionsItemPagesItemOnenotePatchContentPostRequestBody and sets the default values.
func NewMeOnenoteNotebooksItemSectionsItemPagesItemOnenotePatchContentPostRequestBody()(*MeOnenoteNotebooksItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) {
    m := &MeOnenoteNotebooksItemSectionsItemPagesItemOnenotePatchContentPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeOnenoteNotebooksItemSectionsItemPagesItemOnenotePatchContentPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeOnenoteNotebooksItemSectionsItemPagesItemOnenotePatchContentPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeOnenoteNotebooksItemSectionsItemPagesItemOnenotePatchContentPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeOnenoteNotebooksItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCommands gets the commands property value. The commands property
func (m *MeOnenoteNotebooksItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) GetCommands()([]OnenotePatchContentCommandable) {
    return m.commands
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeOnenoteNotebooksItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["commands"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateOnenotePatchContentCommandFromDiscriminatorValue , m.SetCommands)
    return res
}
// Serialize serializes information the current object
func (m *MeOnenoteNotebooksItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCommands() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCommands())
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
func (m *MeOnenoteNotebooksItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCommands sets the commands property value. The commands property
func (m *MeOnenoteNotebooksItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) SetCommands(value []OnenotePatchContentCommandable)() {
    m.commands = value
}
