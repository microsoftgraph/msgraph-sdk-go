package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsInformationProtectionResourceCollection provides operations to manage the deviceAppManagement singleton.
type WindowsInformationProtectionResourceCollection struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Display name
    displayName *string;
    // Collection of resources
    resources []string;
}
// NewWindowsInformationProtectionResourceCollection instantiates a new windowsInformationProtectionResourceCollection and sets the default values.
func NewWindowsInformationProtectionResourceCollection()(*WindowsInformationProtectionResourceCollection) {
    m := &WindowsInformationProtectionResourceCollection{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsInformationProtectionResourceCollectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsInformationProtectionResourceCollectionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWindowsInformationProtectionResourceCollection(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsInformationProtectionResourceCollection) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. Display name
func (m *WindowsInformationProtectionResourceCollection) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsInformationProtectionResourceCollection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["resources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetResources(res)
        }
        return nil
    }
    return res
}
// GetResources gets the resources property value. Collection of resources
func (m *WindowsInformationProtectionResourceCollection) GetResources()([]string) {
    if m == nil {
        return nil
    } else {
        return m.resources
    }
}
func (m *WindowsInformationProtectionResourceCollection) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WindowsInformationProtectionResourceCollection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetResources() != nil {
        err := writer.WriteCollectionOfStringValues("resources", m.GetResources())
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
func (m *WindowsInformationProtectionResourceCollection) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. Display name
func (m *WindowsInformationProtectionResourceCollection) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetResources sets the resources property value. Collection of resources
func (m *WindowsInformationProtectionResourceCollection) SetResources(value []string)() {
    if m != nil {
        m.resources = value
    }
}
