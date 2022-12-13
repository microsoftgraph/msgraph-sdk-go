package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReportsGetHistoricalReportPostRequestBody provides operations to call the getHistoricalReport method.
type ReportsGetHistoricalReportPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The filter property
    filter *string
    // The groupBy property
    groupBy []string
    // The name property
    name *string
    // The orderBy property
    orderBy []string
    // The search property
    search *string
    // The select property
    select_escaped []string
    // The skip property
    skip *int32
    // The top property
    top *int32
}
// NewReportsGetHistoricalReportPostRequestBody instantiates a new ReportsGetHistoricalReportPostRequestBody and sets the default values.
func NewReportsGetHistoricalReportPostRequestBody()(*ReportsGetHistoricalReportPostRequestBody) {
    m := &ReportsGetHistoricalReportPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateReportsGetHistoricalReportPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReportsGetHistoricalReportPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReportsGetHistoricalReportPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReportsGetHistoricalReportPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReportsGetHistoricalReportPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["filter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilter(val)
        }
        return nil
    }
    res["groupBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetGroupBy(res)
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
    res["orderBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOrderBy(res)
        }
        return nil
    }
    res["search"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearch(val)
        }
        return nil
    }
    res["select"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSelect(res)
        }
        return nil
    }
    res["skip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkip(val)
        }
        return nil
    }
    res["top"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTop(val)
        }
        return nil
    }
    return res
}
// GetFilter gets the filter property value. The filter property
func (m *ReportsGetHistoricalReportPostRequestBody) GetFilter()(*string) {
    return m.filter
}
// GetGroupBy gets the groupBy property value. The groupBy property
func (m *ReportsGetHistoricalReportPostRequestBody) GetGroupBy()([]string) {
    return m.groupBy
}
// GetName gets the name property value. The name property
func (m *ReportsGetHistoricalReportPostRequestBody) GetName()(*string) {
    return m.name
}
// GetOrderBy gets the orderBy property value. The orderBy property
func (m *ReportsGetHistoricalReportPostRequestBody) GetOrderBy()([]string) {
    return m.orderBy
}
// GetSearch gets the search property value. The search property
func (m *ReportsGetHistoricalReportPostRequestBody) GetSearch()(*string) {
    return m.search
}
// GetSelect gets the select property value. The select property
func (m *ReportsGetHistoricalReportPostRequestBody) GetSelect()([]string) {
    return m.select_escaped
}
// GetSkip gets the skip property value. The skip property
func (m *ReportsGetHistoricalReportPostRequestBody) GetSkip()(*int32) {
    return m.skip
}
// GetTop gets the top property value. The top property
func (m *ReportsGetHistoricalReportPostRequestBody) GetTop()(*int32) {
    return m.top
}
// Serialize serializes information the current object
func (m *ReportsGetHistoricalReportPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("filter", m.GetFilter())
        if err != nil {
            return err
        }
    }
    if m.GetGroupBy() != nil {
        err := writer.WriteCollectionOfStringValues("groupBy", m.GetGroupBy())
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
    if m.GetOrderBy() != nil {
        err := writer.WriteCollectionOfStringValues("orderBy", m.GetOrderBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("search", m.GetSearch())
        if err != nil {
            return err
        }
    }
    if m.GetSelect() != nil {
        err := writer.WriteCollectionOfStringValues("select", m.GetSelect())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("skip", m.GetSkip())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("top", m.GetTop())
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
func (m *ReportsGetHistoricalReportPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetFilter sets the filter property value. The filter property
func (m *ReportsGetHistoricalReportPostRequestBody) SetFilter(value *string)() {
    m.filter = value
}
// SetGroupBy sets the groupBy property value. The groupBy property
func (m *ReportsGetHistoricalReportPostRequestBody) SetGroupBy(value []string)() {
    m.groupBy = value
}
// SetName sets the name property value. The name property
func (m *ReportsGetHistoricalReportPostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetOrderBy sets the orderBy property value. The orderBy property
func (m *ReportsGetHistoricalReportPostRequestBody) SetOrderBy(value []string)() {
    m.orderBy = value
}
// SetSearch sets the search property value. The search property
func (m *ReportsGetHistoricalReportPostRequestBody) SetSearch(value *string)() {
    m.search = value
}
// SetSelect sets the select property value. The select property
func (m *ReportsGetHistoricalReportPostRequestBody) SetSelect(value []string)() {
    m.select_escaped = value
}
// SetSkip sets the skip property value. The skip property
func (m *ReportsGetHistoricalReportPostRequestBody) SetSkip(value *int32)() {
    m.skip = value
}
// SetTop sets the top property value. The top property
func (m *ReportsGetHistoricalReportPostRequestBody) SetTop(value *int32)() {
    m.top = value
}
