package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type EstimateStatisticsOperation struct {
    CaseOperation
    // The estimated count of items for the sourceCollection that matched the content query.
    indexedItemCount *int64;
    // The estimated size of items for the sourceCollection that matched the content query.
    indexedItemsSize *int64;
    // The number of mailboxes that had search hits.
    mailboxCount *int32;
    // The number of mailboxes that had search hits.
    siteCount *int32;
    // eDiscovery collection, commonly known as a search.
    sourceCollection *SourceCollection;
    // The estimated count of unindexed items for the collection.
    unindexedItemCount *int64;
    // The estimated size of unindexed items for the collection.
    unindexedItemsSize *int64;
}
// Instantiates a new estimateStatisticsOperation and sets the default values.
func NewEstimateStatisticsOperation()(*EstimateStatisticsOperation) {
    m := &EstimateStatisticsOperation{
        CaseOperation: *NewCaseOperation(),
    }
    return m
}
// Gets the indexedItemCount property value. The estimated count of items for the sourceCollection that matched the content query.
func (m *EstimateStatisticsOperation) GetIndexedItemCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.indexedItemCount
    }
}
// Gets the indexedItemsSize property value. The estimated size of items for the sourceCollection that matched the content query.
func (m *EstimateStatisticsOperation) GetIndexedItemsSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.indexedItemsSize
    }
}
// Gets the mailboxCount property value. The number of mailboxes that had search hits.
func (m *EstimateStatisticsOperation) GetMailboxCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.mailboxCount
    }
}
// Gets the siteCount property value. The number of mailboxes that had search hits.
func (m *EstimateStatisticsOperation) GetSiteCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.siteCount
    }
}
// Gets the sourceCollection property value. eDiscovery collection, commonly known as a search.
func (m *EstimateStatisticsOperation) GetSourceCollection()(*SourceCollection) {
    if m == nil {
        return nil
    } else {
        return m.sourceCollection
    }
}
// Gets the unindexedItemCount property value. The estimated count of unindexed items for the collection.
func (m *EstimateStatisticsOperation) GetUnindexedItemCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.unindexedItemCount
    }
}
// Gets the unindexedItemsSize property value. The estimated size of unindexed items for the collection.
func (m *EstimateStatisticsOperation) GetUnindexedItemsSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.unindexedItemsSize
    }
}
// The deserialization information for the current model
func (m *EstimateStatisticsOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CaseOperation.GetFieldDeserializers()
    res["indexedItemCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndexedItemCount(val)
        }
        return nil
    }
    res["indexedItemsSize"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndexedItemsSize(val)
        }
        return nil
    }
    res["mailboxCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailboxCount(val)
        }
        return nil
    }
    res["siteCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteCount(val)
        }
        return nil
    }
    res["sourceCollection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSourceCollection() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceCollection(val.(*SourceCollection))
        }
        return nil
    }
    res["unindexedItemCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnindexedItemCount(val)
        }
        return nil
    }
    res["unindexedItemsSize"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnindexedItemsSize(val)
        }
        return nil
    }
    return res
}
func (m *EstimateStatisticsOperation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *EstimateStatisticsOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CaseOperation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("indexedItemCount", m.GetIndexedItemCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("indexedItemsSize", m.GetIndexedItemsSize())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("mailboxCount", m.GetMailboxCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("siteCount", m.GetSiteCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sourceCollection", m.GetSourceCollection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("unindexedItemCount", m.GetUnindexedItemCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("unindexedItemsSize", m.GetUnindexedItemsSize())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the indexedItemCount property value. The estimated count of items for the sourceCollection that matched the content query.
// Parameters:
//  - value : Value to set for the indexedItemCount property.
func (m *EstimateStatisticsOperation) SetIndexedItemCount(value *int64)() {
    m.indexedItemCount = value
}
// Sets the indexedItemsSize property value. The estimated size of items for the sourceCollection that matched the content query.
// Parameters:
//  - value : Value to set for the indexedItemsSize property.
func (m *EstimateStatisticsOperation) SetIndexedItemsSize(value *int64)() {
    m.indexedItemsSize = value
}
// Sets the mailboxCount property value. The number of mailboxes that had search hits.
// Parameters:
//  - value : Value to set for the mailboxCount property.
func (m *EstimateStatisticsOperation) SetMailboxCount(value *int32)() {
    m.mailboxCount = value
}
// Sets the siteCount property value. The number of mailboxes that had search hits.
// Parameters:
//  - value : Value to set for the siteCount property.
func (m *EstimateStatisticsOperation) SetSiteCount(value *int32)() {
    m.siteCount = value
}
// Sets the sourceCollection property value. eDiscovery collection, commonly known as a search.
// Parameters:
//  - value : Value to set for the sourceCollection property.
func (m *EstimateStatisticsOperation) SetSourceCollection(value *SourceCollection)() {
    m.sourceCollection = value
}
// Sets the unindexedItemCount property value. The estimated count of unindexed items for the collection.
// Parameters:
//  - value : Value to set for the unindexedItemCount property.
func (m *EstimateStatisticsOperation) SetUnindexedItemCount(value *int64)() {
    m.unindexedItemCount = value
}
// Sets the unindexedItemsSize property value. The estimated size of unindexed items for the collection.
// Parameters:
//  - value : Value to set for the unindexedItemsSize property.
func (m *EstimateStatisticsOperation) SetUnindexedItemsSize(value *int64)() {
    m.unindexedItemsSize = value
}
