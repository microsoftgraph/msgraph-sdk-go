package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

type CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBody instantiates a new CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBody and sets the default values.
func NewCasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBody()(*CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBody) {
    m := &CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBody) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExportOptions gets the exportOptions property value. The exportOptions property
// returns a *ExportOptions when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBody) GetExportOptions()(*idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ExportOptions) {
    val, err := m.GetBackingStore().Get("exportOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ExportOptions)
    }
    return nil
}
// GetExportStructure gets the exportStructure property value. The exportStructure property
// returns a *ExportFileStructure when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBody) GetExportStructure()(*idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ExportFileStructure) {
    val, err := m.GetBackingStore().Get("exportStructure")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ExportFileStructure)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["exportOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ParseExportOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportOptions(val.(*idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ExportOptions))
        }
        return nil
    }
    res["exportStructure"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ParseExportFileStructure)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportStructure(val.(*idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ExportFileStructure))
        }
        return nil
    }
    res["outputName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutputName(val)
        }
        return nil
    }
    return res
}
// GetOutputName gets the outputName property value. The outputName property
// returns a *string when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBody) GetOutputName()(*string) {
    val, err := m.GetBackingStore().Get("outputName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetExportOptions() != nil {
        cast := (*m.GetExportOptions()).String()
        err := writer.WriteStringValue("exportOptions", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetExportStructure() != nil {
        cast := (*m.GetExportStructure()).String()
        err := writer.WriteStringValue("exportStructure", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("outputName", m.GetOutputName())
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
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDescription sets the description property value. The description property
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBody) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetExportOptions sets the exportOptions property value. The exportOptions property
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBody) SetExportOptions(value *idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ExportOptions)() {
    err := m.GetBackingStore().Set("exportOptions", value)
    if err != nil {
        panic(err)
    }
}
// SetExportStructure sets the exportStructure property value. The exportStructure property
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBody) SetExportStructure(value *idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ExportFileStructure)() {
    err := m.GetBackingStore().Set("exportStructure", value)
    if err != nil {
        panic(err)
    }
}
// SetOutputName sets the outputName property value. The outputName property
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBody) SetOutputName(value *string)() {
    err := m.GetBackingStore().Set("outputName", value)
    if err != nil {
        panic(err)
    }
}
type CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityexportExportPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDescription()(*string)
    GetExportOptions()(*idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ExportOptions)
    GetExportStructure()(*idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ExportFileStructure)
    GetOutputName()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDescription(value *string)()
    SetExportOptions(value *idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ExportOptions)()
    SetExportStructure(value *idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ExportFileStructure)()
    SetOutputName(value *string)()
}
