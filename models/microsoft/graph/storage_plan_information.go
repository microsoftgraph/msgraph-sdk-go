package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// StoragePlanInformation 
type StoragePlanInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether there are higher storage quota plans available. Read-only.
    upgradeAvailable *bool;
}
// NewStoragePlanInformation instantiates a new storagePlanInformation and sets the default values.
func NewStoragePlanInformation()(*StoragePlanInformation) {
    m := &StoragePlanInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StoragePlanInformation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetUpgradeAvailable gets the upgradeAvailable property value. Indicates whether there are higher storage quota plans available. Read-only.
func (m *StoragePlanInformation) GetUpgradeAvailable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.upgradeAvailable
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StoragePlanInformation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["upgradeAvailable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeAvailable(val)
        }
        return nil
    }
    return res
}
func (m *StoragePlanInformation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *StoragePlanInformation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("upgradeAvailable", m.GetUpgradeAvailable())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StoragePlanInformation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetUpgradeAvailable sets the upgradeAvailable property value. Indicates whether there are higher storage quota plans available. Read-only.
func (m *StoragePlanInformation) SetUpgradeAvailable(value *bool)() {
    m.upgradeAvailable = value
}
