package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WindowsInformationProtectionIPRangeCollection struct {
    additionalData map[string]interface{};
    displayName *string;
    ranges []IpRange;
}
func NewWindowsInformationProtectionIPRangeCollection()(*WindowsInformationProtectionIPRangeCollection) {
    m := &WindowsInformationProtectionIPRangeCollection{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *WindowsInformationProtectionIPRangeCollection) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *WindowsInformationProtectionIPRangeCollection) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *WindowsInformationProtectionIPRangeCollection) GetRanges()([]IpRange) {
    if m == nil {
        return nil
    } else {
        return m.ranges
    }
}
func (m *WindowsInformationProtectionIPRangeCollection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["ranges"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIpRange() })
        if err != nil {
            return err
        }
        res := make([]IpRange, len(val))
        for i, v := range val {
            res[i] = *(v.(*IpRange))
        }
        m.SetRanges(res)
        return nil
    }
    return res
}
func (m *WindowsInformationProtectionIPRangeCollection) IsNil()(bool) {
    return m == nil
}
func (m *WindowsInformationProtectionIPRangeCollection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRanges()))
        for i, v := range m.GetRanges() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("ranges", cast)
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
func (m *WindowsInformationProtectionIPRangeCollection) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *WindowsInformationProtectionIPRangeCollection) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *WindowsInformationProtectionIPRangeCollection) SetRanges(value []IpRange)() {
    m.ranges = value
}
