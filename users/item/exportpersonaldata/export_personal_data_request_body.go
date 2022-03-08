package exportpersonaldata

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ExportPersonalDataRequestBody provides operations to call the exportPersonalData method.
type ExportPersonalDataRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    storageLocation *string;
}
// NewExportPersonalDataRequestBody instantiates a new exportPersonalDataRequestBody and sets the default values.
func NewExportPersonalDataRequestBody()(*ExportPersonalDataRequestBody) {
    m := &ExportPersonalDataRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateExportPersonalDataRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExportPersonalDataRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
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
func (m *ExportPersonalDataRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["storageLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
// GetStorageLocation gets the storageLocation property value. 
func (m *ExportPersonalDataRequestBody) GetStorageLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.storageLocation
    }
}
func (m *ExportPersonalDataRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ExportPersonalDataRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetStorageLocation sets the storageLocation property value. 
func (m *ExportPersonalDataRequestBody) SetStorageLocation(value *string)() {
    if m != nil {
        m.storageLocation = value
    }
}
