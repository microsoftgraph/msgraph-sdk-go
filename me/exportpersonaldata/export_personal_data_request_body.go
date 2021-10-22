package exportpersonaldata

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ExportPersonalDataRequestBody struct {
    additionalData map[string]interface{};
    storageLocation *string;
}
func NewExportPersonalDataRequestBody()(*ExportPersonalDataRequestBody) {
    m := &ExportPersonalDataRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ExportPersonalDataRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ExportPersonalDataRequestBody) GetStorageLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.storageLocation
    }
}
func (m *ExportPersonalDataRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["storageLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStorageLocation(val)
        return nil
    }
    return res
}
func (m *ExportPersonalDataRequestBody) IsNil()(bool) {
    return m == nil
}
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
func (m *ExportPersonalDataRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ExportPersonalDataRequestBody) SetStorageLocation(value *string)() {
    m.storageLocation = value
}
