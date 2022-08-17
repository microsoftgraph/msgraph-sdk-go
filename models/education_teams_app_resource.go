package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationTeamsAppResource 
type EducationTeamsAppResource struct {
    EducationResource
    // The appIconWebUrl property
    appIconWebUrl *string
    // The appId property
    appId *string
    // The teamsEmbeddedContentUrl property
    teamsEmbeddedContentUrl *string
    // The webUrl property
    webUrl *string
}
// NewEducationTeamsAppResource instantiates a new EducationTeamsAppResource and sets the default values.
func NewEducationTeamsAppResource()(*EducationTeamsAppResource) {
    m := &EducationTeamsAppResource{
        EducationResource: *NewEducationResource(),
    }
    odataTypeValue := "#microsoft.graph.educationTeamsAppResource";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEducationTeamsAppResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationTeamsAppResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationTeamsAppResource(), nil
}
// GetAppIconWebUrl gets the appIconWebUrl property value. The appIconWebUrl property
func (m *EducationTeamsAppResource) GetAppIconWebUrl()(*string) {
    return m.appIconWebUrl
}
// GetAppId gets the appId property value. The appId property
func (m *EducationTeamsAppResource) GetAppId()(*string) {
    return m.appId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationTeamsAppResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationResource.GetFieldDeserializers()
    res["appIconWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppIconWebUrl(val)
        }
        return nil
    }
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
    res["teamsEmbeddedContentUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsEmbeddedContentUrl(val)
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetTeamsEmbeddedContentUrl gets the teamsEmbeddedContentUrl property value. The teamsEmbeddedContentUrl property
func (m *EducationTeamsAppResource) GetTeamsEmbeddedContentUrl()(*string) {
    return m.teamsEmbeddedContentUrl
}
// GetWebUrl gets the webUrl property value. The webUrl property
func (m *EducationTeamsAppResource) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *EducationTeamsAppResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationResource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appIconWebUrl", m.GetAppIconWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("teamsEmbeddedContentUrl", m.GetTeamsEmbeddedContentUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppIconWebUrl sets the appIconWebUrl property value. The appIconWebUrl property
func (m *EducationTeamsAppResource) SetAppIconWebUrl(value *string)() {
    m.appIconWebUrl = value
}
// SetAppId sets the appId property value. The appId property
func (m *EducationTeamsAppResource) SetAppId(value *string)() {
    m.appId = value
}
// SetTeamsEmbeddedContentUrl sets the teamsEmbeddedContentUrl property value. The teamsEmbeddedContentUrl property
func (m *EducationTeamsAppResource) SetTeamsEmbeddedContentUrl(value *string)() {
    m.teamsEmbeddedContentUrl = value
}
// SetWebUrl sets the webUrl property value. The webUrl property
func (m *EducationTeamsAppResource) SetWebUrl(value *string)() {
    m.webUrl = value
}
