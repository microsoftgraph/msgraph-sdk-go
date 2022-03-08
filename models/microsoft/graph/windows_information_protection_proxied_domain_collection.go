package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsInformationProtectionProxiedDomainCollection provides operations to manage the deviceAppManagement singleton.
type WindowsInformationProtectionProxiedDomainCollection struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Display name
    displayName *string;
    // Collection of proxied domains
    proxiedDomains []ProxiedDomainable;
}
// NewWindowsInformationProtectionProxiedDomainCollection instantiates a new windowsInformationProtectionProxiedDomainCollection and sets the default values.
func NewWindowsInformationProtectionProxiedDomainCollection()(*WindowsInformationProtectionProxiedDomainCollection) {
    m := &WindowsInformationProtectionProxiedDomainCollection{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsInformationProtectionProxiedDomainCollectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsInformationProtectionProxiedDomainCollectionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWindowsInformationProtectionProxiedDomainCollection(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsInformationProtectionProxiedDomainCollection) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. Display name
func (m *WindowsInformationProtectionProxiedDomainCollection) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsInformationProtectionProxiedDomainCollection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["proxiedDomains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProxiedDomainFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProxiedDomainable, len(val))
            for i, v := range val {
                res[i] = v.(ProxiedDomainable)
            }
            m.SetProxiedDomains(res)
        }
        return nil
    }
    return res
}
// GetProxiedDomains gets the proxiedDomains property value. Collection of proxied domains
func (m *WindowsInformationProtectionProxiedDomainCollection) GetProxiedDomains()([]ProxiedDomainable) {
    if m == nil {
        return nil
    } else {
        return m.proxiedDomains
    }
}
func (m *WindowsInformationProtectionProxiedDomainCollection) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WindowsInformationProtectionProxiedDomainCollection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetProxiedDomains() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProxiedDomains()))
        for i, v := range m.GetProxiedDomains() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("proxiedDomains", cast)
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
func (m *WindowsInformationProtectionProxiedDomainCollection) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. Display name
func (m *WindowsInformationProtectionProxiedDomainCollection) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetProxiedDomains sets the proxiedDomains property value. Collection of proxied domains
func (m *WindowsInformationProtectionProxiedDomainCollection) SetProxiedDomains(value []ProxiedDomainable)() {
    if m != nil {
        m.proxiedDomains = value
    }
}
