package exportpersonaldata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExportPersonalDataRequestBody provides operations to call the exportPersonalData method.
type ExportPersonalDataRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The storageLocation property
    storageLocation *string
}
// NewExportPersonalDataRequestBody instantiates a new exportPersonalDataRequestBody and sets the default values.
func NewExportPersonalDataRequestBody()(*ExportPersonalDataRequestBody) {
    m := &ExportPersonalDataRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateExportPersonalDataRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExportPersonalDataRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExportPersonalDataRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExportPersonalDataRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExportPersonalDataRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["storageLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageLocation(val)
        }
        return nil
    }
    return res
}
// GetStorageLocation gets the storageLocation property value. The storageLocation property
func (m *ExportPersonalDataRequestBody) GetStorageLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.storageLocation
    }
}
// Serialize serializes information the current object
func (m *ExportPersonalDataRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("storageLocation", m.GetStorageLocation())
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
func (m *ExportPersonalDataRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetStorageLocation sets the storageLocation property value. The storageLocation property
func (m *ExportPersonalDataRequestBody) SetStorageLocation(value *string)() {
    if m != nil {
        m.storageLocation = value
    }
}
