package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ImportedWindowsAutopilotDeviceIdentitiesImportPostRequestBody 
type ImportedWindowsAutopilotDeviceIdentitiesImportPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The importedWindowsAutopilotDeviceIdentities property
    importedWindowsAutopilotDeviceIdentities []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ImportedWindowsAutopilotDeviceIdentityable
}
// NewImportedWindowsAutopilotDeviceIdentitiesImportPostRequestBody instantiates a new ImportedWindowsAutopilotDeviceIdentitiesImportPostRequestBody and sets the default values.
func NewImportedWindowsAutopilotDeviceIdentitiesImportPostRequestBody()(*ImportedWindowsAutopilotDeviceIdentitiesImportPostRequestBody) {
    m := &ImportedWindowsAutopilotDeviceIdentitiesImportPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateImportedWindowsAutopilotDeviceIdentitiesImportPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateImportedWindowsAutopilotDeviceIdentitiesImportPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewImportedWindowsAutopilotDeviceIdentitiesImportPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ImportedWindowsAutopilotDeviceIdentitiesImportPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ImportedWindowsAutopilotDeviceIdentitiesImportPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["importedWindowsAutopilotDeviceIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateImportedWindowsAutopilotDeviceIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ImportedWindowsAutopilotDeviceIdentityable, len(val))
            for i, v := range val {
                res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ImportedWindowsAutopilotDeviceIdentityable)
            }
            m.SetImportedWindowsAutopilotDeviceIdentities(res)
        }
        return nil
    }
    return res
}
// GetImportedWindowsAutopilotDeviceIdentities gets the importedWindowsAutopilotDeviceIdentities property value. The importedWindowsAutopilotDeviceIdentities property
func (m *ImportedWindowsAutopilotDeviceIdentitiesImportPostRequestBody) GetImportedWindowsAutopilotDeviceIdentities()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ImportedWindowsAutopilotDeviceIdentityable) {
    return m.importedWindowsAutopilotDeviceIdentities
}
// Serialize serializes information the current object
func (m *ImportedWindowsAutopilotDeviceIdentitiesImportPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetImportedWindowsAutopilotDeviceIdentities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetImportedWindowsAutopilotDeviceIdentities()))
        for i, v := range m.GetImportedWindowsAutopilotDeviceIdentities() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("importedWindowsAutopilotDeviceIdentities", cast)
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
func (m *ImportedWindowsAutopilotDeviceIdentitiesImportPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetImportedWindowsAutopilotDeviceIdentities sets the importedWindowsAutopilotDeviceIdentities property value. The importedWindowsAutopilotDeviceIdentities property
func (m *ImportedWindowsAutopilotDeviceIdentitiesImportPostRequestBody) SetImportedWindowsAutopilotDeviceIdentities(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ImportedWindowsAutopilotDeviceIdentityable)() {
    m.importedWindowsAutopilotDeviceIdentities = value
}
