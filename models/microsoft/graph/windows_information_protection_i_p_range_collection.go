package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsInformationProtectionIPRangeCollection windows Information Protection IP Range Collection
type WindowsInformationProtectionIPRangeCollection struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Display name
    displayName *string;
    // Collection of ip ranges
    ranges []IpRangeable;
}
// NewWindowsInformationProtectionIPRangeCollection instantiates a new windowsInformationProtectionIPRangeCollection and sets the default values.
func NewWindowsInformationProtectionIPRangeCollection()(*WindowsInformationProtectionIPRangeCollection) {
    m := &WindowsInformationProtectionIPRangeCollection{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsInformationProtectionIPRangeCollectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsInformationProtectionIPRangeCollectionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWindowsInformationProtectionIPRangeCollection(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
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
        val, err := n.GetCollectionOfObjectValues(CreateIpRangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IpRangeable, len(val))
            for i, v := range val {
                res[i] = v.(IpRangeable)
            }
            m.SetRanges(res)
        }
        return nil
    }
    return res
}
// GetRanges gets the ranges property value. Collection of ip ranges
func (m *WindowsInformationProtectionIPRangeCollection) GetRanges()([]IpRangeable) {
    if m == nil {
        return nil
    } else {
        return m.ranges
    }
}
// Serialize serializes information the current object
func (m *WindowsInformationProtectionIPRangeCollection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetRanges() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRanges()))
        for i, v := range m.GetRanges() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
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
func (m *WindowsInformationProtectionIPRangeCollection) SetRanges(value []IpRangeable)() {
    if m != nil {
        m.ranges = value
    }
}
