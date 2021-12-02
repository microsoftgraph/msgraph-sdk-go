package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsInformationProtectionIPRangeCollection 
type WindowsInformationProtectionIPRangeCollection struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Display name
    displayName *string;
    // Collection of ip ranges
    ranges []IpRange;
}
// NewWindowsInformationProtectionIPRangeCollection instantiates a new windowsInformationProtectionIPRangeCollection and sets the default values.
func NewWindowsInformationProtectionIPRangeCollection()(*WindowsInformationProtectionIPRangeCollection) {
    m := &WindowsInformationProtectionIPRangeCollection{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsInformationProtectionIPRangeCollection) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. Display name
func (m *WindowsInformationProtectionIPRangeCollection) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetRanges gets the ranges property value. Collection of ip ranges
func (m *WindowsInformationProtectionIPRangeCollection) GetRanges()([]IpRange) {
    if m == nil {
        return nil
    } else {
        return m.ranges
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsInformationProtectionIPRangeCollection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["ranges"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIpRange() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IpRange, len(val))
            for i, v := range val {
                res[i] = *(v.(*IpRange))
            }
            m.SetRanges(res)
        }
        return nil
    }
    return res
}
func (m *WindowsInformationProtectionIPRangeCollection) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsInformationProtectionIPRangeCollection) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. Display name
func (m *WindowsInformationProtectionIPRangeCollection) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetRanges sets the ranges property value. Collection of ip ranges
func (m *WindowsInformationProtectionIPRangeCollection) SetRanges(value []IpRange)() {
    if m != nil {
        m.ranges = value
    }
}
