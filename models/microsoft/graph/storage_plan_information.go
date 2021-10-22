package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type StoragePlanInformation struct {
    additionalData map[string]interface{};
    upgradeAvailable *bool;
}
func NewStoragePlanInformation()(*StoragePlanInformation) {
    m := &StoragePlanInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *StoragePlanInformation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *StoragePlanInformation) GetUpgradeAvailable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.upgradeAvailable
    }
}
func (m *StoragePlanInformation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["upgradeAvailable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUpgradeAvailable(val)
        return nil
    }
    return res
}
func (m *StoragePlanInformation) IsNil()(bool) {
    return m == nil
}
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
func (m *StoragePlanInformation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *StoragePlanInformation) SetUpgradeAvailable(value *bool)() {
    m.upgradeAvailable = value
}
