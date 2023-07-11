package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdgeSearchEngineCustom allows IT admins to set a default search engine for MDM-Controlled devices. Users can override this and change their default search engine provided the AllowSearchEngineCustomization policy is not set.
type EdgeSearchEngineCustom struct {
    EdgeSearchEngineBase
}
// NewEdgeSearchEngineCustom instantiates a new edgeSearchEngineCustom and sets the default values.
func NewEdgeSearchEngineCustom()(*EdgeSearchEngineCustom) {
    m := &EdgeSearchEngineCustom{
        EdgeSearchEngineBase: *NewEdgeSearchEngineBase(),
    }
    odataTypeValue := "#microsoft.graph.edgeSearchEngineCustom"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEdgeSearchEngineCustomFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdgeSearchEngineCustomFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdgeSearchEngineCustom(), nil
}
// GetEdgeSearchEngineOpenSearchXmlUrl gets the edgeSearchEngineOpenSearchXmlUrl property value. Points to a https link containing the OpenSearch xml file that contains, at minimum, the short name and the URL to the search Engine.
func (m *EdgeSearchEngineCustom) GetEdgeSearchEngineOpenSearchXmlUrl()(*string) {
    val, err := m.GetBackingStore().Get("edgeSearchEngineOpenSearchXmlUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdgeSearchEngineCustom) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EdgeSearchEngineBase.GetFieldDeserializers()
    res["edgeSearchEngineOpenSearchXmlUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeSearchEngineOpenSearchXmlUrl(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EdgeSearchEngineCustom) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EdgeSearchEngineCustom) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EdgeSearchEngineBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("edgeSearchEngineOpenSearchXmlUrl", m.GetEdgeSearchEngineOpenSearchXmlUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEdgeSearchEngineOpenSearchXmlUrl sets the edgeSearchEngineOpenSearchXmlUrl property value. Points to a https link containing the OpenSearch xml file that contains, at minimum, the short name and the URL to the search Engine.
func (m *EdgeSearchEngineCustom) SetEdgeSearchEngineOpenSearchXmlUrl(value *string)() {
    err := m.GetBackingStore().Set("edgeSearchEngineOpenSearchXmlUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EdgeSearchEngineCustom) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// EdgeSearchEngineCustomable 
type EdgeSearchEngineCustomable interface {
    EdgeSearchEngineBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEdgeSearchEngineOpenSearchXmlUrl()(*string)
    GetOdataType()(*string)
    SetEdgeSearchEngineOpenSearchXmlUrl(value *string)()
    SetOdataType(value *string)()
}
