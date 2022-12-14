package drive

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ListContentTypesItemCopyToDefaultContentLocationPostRequestBody provides operations to call the copyToDefaultContentLocation method.
type ListContentTypesItemCopyToDefaultContentLocationPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The destinationFileName property
    destinationFileName *string
    // The sourceFile property
    sourceFile iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemReferenceable
}
// NewListContentTypesItemCopyToDefaultContentLocationPostRequestBody instantiates a new ListContentTypesItemCopyToDefaultContentLocationPostRequestBody and sets the default values.
func NewListContentTypesItemCopyToDefaultContentLocationPostRequestBody()(*ListContentTypesItemCopyToDefaultContentLocationPostRequestBody) {
    m := &ListContentTypesItemCopyToDefaultContentLocationPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateListContentTypesItemCopyToDefaultContentLocationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateListContentTypesItemCopyToDefaultContentLocationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewListContentTypesItemCopyToDefaultContentLocationPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ListContentTypesItemCopyToDefaultContentLocationPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDestinationFileName gets the destinationFileName property value. The destinationFileName property
func (m *ListContentTypesItemCopyToDefaultContentLocationPostRequestBody) GetDestinationFileName()(*string) {
    return m.destinationFileName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ListContentTypesItemCopyToDefaultContentLocationPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["destinationFileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationFileName(val)
        }
        return nil
    }
    res["sourceFile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateItemReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceFile(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemReferenceable))
        }
        return nil
    }
    return res
}
// GetSourceFile gets the sourceFile property value. The sourceFile property
func (m *ListContentTypesItemCopyToDefaultContentLocationPostRequestBody) GetSourceFile()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemReferenceable) {
    return m.sourceFile
}
// Serialize serializes information the current object
func (m *ListContentTypesItemCopyToDefaultContentLocationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("destinationFileName", m.GetDestinationFileName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("sourceFile", m.GetSourceFile())
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
func (m *ListContentTypesItemCopyToDefaultContentLocationPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDestinationFileName sets the destinationFileName property value. The destinationFileName property
func (m *ListContentTypesItemCopyToDefaultContentLocationPostRequestBody) SetDestinationFileName(value *string)() {
    m.destinationFileName = value
}
// SetSourceFile sets the sourceFile property value. The sourceFile property
func (m *ListContentTypesItemCopyToDefaultContentLocationPostRequestBody) SetSourceFile(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemReferenceable)() {
    m.sourceFile = value
}
