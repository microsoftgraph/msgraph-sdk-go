package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppListItem represents an app in the list of managed applications
type AppListItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The application or bundle identifier of the application
    appId *string
    // The Store URL of the application
    appStoreUrl *string
    // The application name
    name *string
    // The publisher of the application
    publisher *string
}
// NewAppListItem instantiates a new appListItem and sets the default values.
func NewAppListItem()(*AppListItem) {
    m := &AppListItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAppListItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppListItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppListItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppListItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAppId gets the appId property value. The application or bundle identifier of the application
func (m *AppListItem) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
// GetAppStoreUrl gets the appStoreUrl property value. The Store URL of the application
func (m *AppListItem) GetAppStoreUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appStoreUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppListItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["appStoreUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppStoreUrl(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["publisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The application name
func (m *AppListItem) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetPublisher gets the publisher property value. The publisher of the application
func (m *AppListItem) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// Serialize serializes information the current object
func (m *AppListItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("appStoreUrl", m.GetAppStoreUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("publisher", m.GetPublisher())
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
func (m *AppListItem) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAppId sets the appId property value. The application or bundle identifier of the application
func (m *AppListItem) SetAppId(value *string)() {
    if m != nil {
        m.appId = value
    }
}
// SetAppStoreUrl sets the appStoreUrl property value. The Store URL of the application
func (m *AppListItem) SetAppStoreUrl(value *string)() {
    if m != nil {
        m.appStoreUrl = value
    }
}
// SetName sets the name property value. The application name
func (m *AppListItem) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetPublisher sets the publisher property value. The publisher of the application
func (m *AppListItem) SetPublisher(value *string)() {
    if m != nil {
        m.publisher = value
    }
}
