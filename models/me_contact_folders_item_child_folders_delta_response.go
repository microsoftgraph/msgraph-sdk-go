package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeContactFoldersItemChildFoldersDeltaResponse provides operations to call the delta method.
type MeContactFoldersItemChildFoldersDeltaResponse struct {
    BaseDeltaFunctionResponse
    // The value property
    value []ContactFolderable
}
// NewMeContactFoldersItemChildFoldersDeltaResponse instantiates a new MeContactFoldersItemChildFoldersDeltaResponse and sets the default values.
func NewMeContactFoldersItemChildFoldersDeltaResponse()(*MeContactFoldersItemChildFoldersDeltaResponse) {
    m := &MeContactFoldersItemChildFoldersDeltaResponse{
        BaseDeltaFunctionResponse: *NewBaseDeltaFunctionResponse(),
    }
    return m
}
// CreateMeContactFoldersItemChildFoldersDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeContactFoldersItemChildFoldersDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeContactFoldersItemChildFoldersDeltaResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeContactFoldersItemChildFoldersDeltaResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseDeltaFunctionResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateContactFolderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ContactFolderable, len(val))
            for i, v := range val {
                res[i] = v.(ContactFolderable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *MeContactFoldersItemChildFoldersDeltaResponse) GetValue()([]ContactFolderable) {
    return m.value
}
// Serialize serializes information the current object
func (m *MeContactFoldersItemChildFoldersDeltaResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseDeltaFunctionResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *MeContactFoldersItemChildFoldersDeltaResponse) SetValue(value []ContactFolderable)() {
    m.value = value
}
