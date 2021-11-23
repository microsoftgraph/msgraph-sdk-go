package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// licenseUnitsDetail 
type LicenseUnitsDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The number of units that are enabled for the active subscription of the service SKU.
    enabled *int32;
    // The number of units that are suspended because the subscription of the service SKU has been cancelled. The units cannot be assigned but can still be reactivated before they are deleted.
    suspended *int32;
    // The number of units that are in warning status. When the subscription of the service SKU has expired, the customer has a grace period to renew their subscription before it is cancelled (moved to a suspended state).
    warning *int32;
}
// NewLicenseUnitsDetail instantiates a new licenseUnitsDetail and sets the default values.
func NewLicenseUnitsDetail()(*LicenseUnitsDetail) {
    m := &LicenseUnitsDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LicenseUnitsDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEnabled gets the enabled property value. The number of units that are enabled for the active subscription of the service SKU.
func (m *LicenseUnitsDetail) GetEnabled()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.enabled
    }
}
// GetSuspended gets the suspended property value. The number of units that are suspended because the subscription of the service SKU has been cancelled. The units cannot be assigned but can still be reactivated before they are deleted.
func (m *LicenseUnitsDetail) GetSuspended()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.suspended
    }
}
// GetWarning gets the warning property value. The number of units that are in warning status. When the subscription of the service SKU has expired, the customer has a grace period to renew their subscription before it is cancelled (moved to a suspended state).
func (m *LicenseUnitsDetail) GetWarning()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.warning
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LicenseUnitsDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["enabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["suspended"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuspended(val)
        }
        return nil
    }
    res["warning"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWarning(val)
        }
        return nil
    }
    return res
}
func (m *LicenseUnitsDetail) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *LicenseUnitsDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("suspended", m.GetSuspended())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("warning", m.GetWarning())
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
func (m *LicenseUnitsDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEnabled sets the enabled property value. The number of units that are enabled for the active subscription of the service SKU.
func (m *LicenseUnitsDetail) SetEnabled(value *int32)() {
    m.enabled = value
}
// SetSuspended sets the suspended property value. The number of units that are suspended because the subscription of the service SKU has been cancelled. The units cannot be assigned but can still be reactivated before they are deleted.
func (m *LicenseUnitsDetail) SetSuspended(value *int32)() {
    m.suspended = value
}
// SetWarning sets the warning property value. The number of units that are in warning status. When the subscription of the service SKU has expired, the customer has a grace period to renew their subscription before it is cancelled (moved to a suspended state).
func (m *LicenseUnitsDetail) SetWarning(value *int32)() {
    m.warning = value
}
