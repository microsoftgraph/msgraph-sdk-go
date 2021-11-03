package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ExtensionProperty struct {
    DirectoryObject
    // Display name of the application object on which this extension property is defined. Read-only.
    appDisplayName *string;
    // Specifies the data type of the value the extension property can hold. Following values are supported. Not nullable. Binary - 256 bytes maximumBooleanDateTime - Must be specified in ISO 8601 format. Will be stored in UTC.Integer - 32-bit value.LargeInteger - 64-bit value.String - 256 characters maximum
    dataType *string;
    // Indicates if this extension property was sycned from onpremises directory using Azure AD Connect. Read-only.
    isSyncedFromOnPremises *bool;
    // Name of the extension property. Not nullable.
    name *string;
    // Following values are supported. Not nullable. UserGroupOrganizationDeviceApplication
    targetObjects []string;
}
// Instantiates a new extensionProperty and sets the default values.
func NewExtensionProperty()(*ExtensionProperty) {
    m := &ExtensionProperty{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// Gets the appDisplayName property value. Display name of the application object on which this extension property is defined. Read-only.
func (m *ExtensionProperty) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
// Gets the dataType property value. Specifies the data type of the value the extension property can hold. Following values are supported. Not nullable. Binary - 256 bytes maximumBooleanDateTime - Must be specified in ISO 8601 format. Will be stored in UTC.Integer - 32-bit value.LargeInteger - 64-bit value.String - 256 characters maximum
func (m *ExtensionProperty) GetDataType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dataType
    }
}
// Gets the isSyncedFromOnPremises property value. Indicates if this extension property was sycned from onpremises directory using Azure AD Connect. Read-only.
func (m *ExtensionProperty) GetIsSyncedFromOnPremises()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSyncedFromOnPremises
    }
}
// Gets the name property value. Name of the extension property. Not nullable.
func (m *ExtensionProperty) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the targetObjects property value. Following values are supported. Not nullable. UserGroupOrganizationDeviceApplication
func (m *ExtensionProperty) GetTargetObjects()([]string) {
    if m == nil {
        return nil
    } else {
        return m.targetObjects
    }
}
// The deserialization information for the current model
func (m *ExtensionProperty) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["appDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppDisplayName(val)
        return nil
    }
    res["dataType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDataType(val)
        return nil
    }
    res["isSyncedFromOnPremises"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsSyncedFromOnPremises(val)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["targetObjects"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetTargetObjects(res)
        return nil
    }
    return res
}
func (m *ExtensionProperty) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ExtensionProperty) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("dataType", m.GetDataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSyncedFromOnPremises", m.GetIsSyncedFromOnPremises())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("targetObjects", m.GetTargetObjects())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the appDisplayName property value. Display name of the application object on which this extension property is defined. Read-only.
// Parameters:
//  - value : Value to set for the appDisplayName property.
func (m *ExtensionProperty) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// Sets the dataType property value. Specifies the data type of the value the extension property can hold. Following values are supported. Not nullable. Binary - 256 bytes maximumBooleanDateTime - Must be specified in ISO 8601 format. Will be stored in UTC.Integer - 32-bit value.LargeInteger - 64-bit value.String - 256 characters maximum
// Parameters:
//  - value : Value to set for the dataType property.
func (m *ExtensionProperty) SetDataType(value *string)() {
    m.dataType = value
}
// Sets the isSyncedFromOnPremises property value. Indicates if this extension property was sycned from onpremises directory using Azure AD Connect. Read-only.
// Parameters:
//  - value : Value to set for the isSyncedFromOnPremises property.
func (m *ExtensionProperty) SetIsSyncedFromOnPremises(value *bool)() {
    m.isSyncedFromOnPremises = value
}
// Sets the name property value. Name of the extension property. Not nullable.
// Parameters:
//  - value : Value to set for the name property.
func (m *ExtensionProperty) SetName(value *string)() {
    m.name = value
}
// Sets the targetObjects property value. Following values are supported. Not nullable. UserGroupOrganizationDeviceApplication
// Parameters:
//  - value : Value to set for the targetObjects property.
func (m *ExtensionProperty) SetTargetObjects(value []string)() {
    m.targetObjects = value
}
