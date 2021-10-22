package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type LicenseUnitsDetail struct {
    additionalData map[string]interface{};
    enabled *int32;
    suspended *int32;
    warning *int32;
}
func NewLicenseUnitsDetail()(*LicenseUnitsDetail) {
    m := &LicenseUnitsDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *LicenseUnitsDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *LicenseUnitsDetail) GetEnabled()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.enabled
    }
}
func (m *LicenseUnitsDetail) GetSuspended()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.suspended
    }
}
func (m *LicenseUnitsDetail) GetWarning()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.warning
    }
}
func (m *LicenseUnitsDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["enabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetEnabled(val)
        return nil
    }
    res["suspended"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSuspended(val)
        return nil
    }
    res["warning"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetWarning(val)
        return nil
    }
    return res
}
func (m *LicenseUnitsDetail) IsNil()(bool) {
    return m == nil
}
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
func (m *LicenseUnitsDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *LicenseUnitsDetail) SetEnabled(value *int32)() {
    m.enabled = value
}
func (m *LicenseUnitsDetail) SetSuspended(value *int32)() {
    m.suspended = value
}
func (m *LicenseUnitsDetail) SetWarning(value *int32)() {
    m.warning = value
}
