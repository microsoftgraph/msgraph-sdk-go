package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// CasesEdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody provides operations to call the addToReviewSet method.
type CasesEdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The additionalDataOptions property
    additionalDataOptions *idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.AdditionalDataOptions
    // The search property
    search idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoverySearchable
}
// NewCasesEdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody instantiates a new CasesEdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody and sets the default values.
func NewCasesEdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody()(*CasesEdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody) {
    m := &CasesEdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCasesEdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCasesEdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCasesEdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CasesEdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAdditionalDataOptions gets the additionalDataOptions property value. The additionalDataOptions property
func (m *CasesEdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody) GetAdditionalDataOptions()(*idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.AdditionalDataOptions) {
    return m.additionalDataOptions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CasesEdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *CasesEdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody) GetSearch()(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoverySearchable) {
    return m.search
}
// Serialize serializes information the current object
func (m *CasesEdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CasesEdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAdditionalDataOptions sets the additionalDataOptions property value. The additionalDataOptions property
func (m *CasesEdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody) SetAdditionalDataOptions(value *idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.AdditionalDataOptions)() {
    m.additionalDataOptions = value
}
// SetSearch sets the search property value. The search property
func (m *CasesEdiscoveryCasesItemReviewSetsItemAddToReviewSetPostRequestBody) SetSearch(value idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoverySearchable)() {
    m.search = value
}
