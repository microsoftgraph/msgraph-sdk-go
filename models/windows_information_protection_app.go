package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsInformationProtectionApp app for Windows information protection
type WindowsInformationProtectionApp struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // If true, app is denied protection or exemption.
    denied *bool
    // The app's description.
    description *string
    // App display name.
    displayName *string
    // The product name.
    productName *string
    // The publisher name
    publisherName *string
    // The type property
    type_escaped *string
}
// NewWindowsInformationProtectionApp instantiates a new windowsInformationProtectionApp and sets the default values.
func NewWindowsInformationProtectionApp()(*WindowsInformationProtectionApp) {
    m := &WindowsInformationProtectionApp{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odatatypeValue := "#microsoft.graph.windowsInformationProtectionApp";
    m.SetType(&odatatypeValue);
    return m
}
// CreateWindowsInformationProtectionAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsInformationProtectionAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.windowsInformationProtectionDesktopApp":
                        return NewWindowsInformationProtectionDesktopApp(), nil
                    case "#microsoft.graph.windowsInformationProtectionStoreApp":
                        return NewWindowsInformationProtectionStoreApp(), nil
                }
            }
        }
    }
    return NewWindowsInformationProtectionApp(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsInformationProtectionApp) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDenied gets the denied property value. If true, app is denied protection or exemption.
func (m *WindowsInformationProtectionApp) GetDenied()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.denied
    }
}
// GetDescription gets the description property value. The app's description.
func (m *WindowsInformationProtectionApp) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. App display name.
func (m *WindowsInformationProtectionApp) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsInformationProtectionApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["denied"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDenied(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["productName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductName(val)
        }
        return nil
    }
    res["publisherName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisherName(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetProductName gets the productName property value. The product name.
func (m *WindowsInformationProtectionApp) GetProductName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.productName
    }
}
// GetPublisherName gets the publisherName property value. The publisher name
func (m *WindowsInformationProtectionApp) GetPublisherName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisherName
    }
}
// GetType gets the @odata.type property value. The type property
func (m *WindowsInformationProtectionApp) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *WindowsInformationProtectionApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("denied", m.GetDenied())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("productName", m.GetProductName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("publisherName", m.GetPublisherName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetType())
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
func (m *WindowsInformationProtectionApp) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDenied sets the denied property value. If true, app is denied protection or exemption.
func (m *WindowsInformationProtectionApp) SetDenied(value *bool)() {
    if m != nil {
        m.denied = value
    }
}
// SetDescription sets the description property value. The app's description.
func (m *WindowsInformationProtectionApp) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. App display name.
func (m *WindowsInformationProtectionApp) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetProductName sets the productName property value. The product name.
func (m *WindowsInformationProtectionApp) SetProductName(value *string)() {
    if m != nil {
        m.productName = value
    }
}
// SetPublisherName sets the publisherName property value. The publisher name
func (m *WindowsInformationProtectionApp) SetPublisherName(value *string)() {
    if m != nil {
        m.publisherName = value
    }
}
// SetType sets the @odata.type property value. The type property
func (m *WindowsInformationProtectionApp) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
