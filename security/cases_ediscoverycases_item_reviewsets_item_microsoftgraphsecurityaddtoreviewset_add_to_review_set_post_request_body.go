package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

type CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetAddToReviewSetPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetAddToReviewSetPostRequestBody instantiates a new CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetAddToReviewSetPostRequestBody and sets the default values.
func NewCasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetAddToReviewSetPostRequestBody()(*CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetAddToReviewSetPostRequestBody) {
    m := &CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetAddToReviewSetPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetAddToReviewSetPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetAddToReviewSetPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetAddToReviewSetPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetAddToReviewSetPostRequestBody) GetAdditionalData()(map[string]any) {
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
// GetAdditionalDataOptions gets the additionalDataOptions property value. The additionalDataOptions property
// returns a *AdditionalDataOptions when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetAddToReviewSetPostRequestBody) GetAdditionalDataOptions()(*idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.AdditionalDataOptions) {
    val, err := m.GetBackingStore().Get("additionalDataOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.AdditionalDataOptions)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetAddToReviewSetPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetAddToReviewSetPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["additionalDataOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ParseAdditionalDataOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalDataOptions(val.(*idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.AdditionalDataOptions))
        }
        return nil
    }
    res["search"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoverySearchFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearch(val.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoverySearchable))
        }
        return nil
    }
    return res
}
// GetSearch gets the search property value. The search property
// returns a EdiscoverySearchable when successful
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetAddToReviewSetPostRequestBody) GetSearch()(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoverySearchable) {
    val, err := m.GetBackingStore().Get("search")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoverySearchable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetAddToReviewSetPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAdditionalDataOptions() != nil {
        cast := (*m.GetAdditionalDataOptions()).String()
        err := writer.WriteStringValue("additionalDataOptions", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("search", m.GetSearch())
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
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetAddToReviewSetPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalDataOptions sets the additionalDataOptions property value. The additionalDataOptions property
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetAddToReviewSetPostRequestBody) SetAdditionalDataOptions(value *idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.AdditionalDataOptions)() {
    err := m.GetBackingStore().Set("additionalDataOptions", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetAddToReviewSetPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetSearch sets the search property value. The search property
func (m *CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetAddToReviewSetPostRequestBody) SetSearch(value idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoverySearchable)() {
    err := m.GetBackingStore().Set("search", value)
    if err != nil {
        panic(err)
    }
}
type CasesEdiscoverycasesItemReviewsetsItemMicrosoftgraphsecurityaddtoreviewsetAddToReviewSetPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalDataOptions()(*idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.AdditionalDataOptions)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetSearch()(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoverySearchable)
    SetAdditionalDataOptions(value *idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.AdditionalDataOptions)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetSearch(value idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoverySearchable)()
}
