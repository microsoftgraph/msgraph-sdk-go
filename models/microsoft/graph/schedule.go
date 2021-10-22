package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Schedule struct {
    Entity
    enabled *bool;
    offerShiftRequests []OfferShiftRequest;
    offerShiftRequestsEnabled *bool;
    openShiftChangeRequests []OpenShiftChangeRequest;
    openShifts []OpenShift;
    openShiftsEnabled *bool;
    provisionStatus *OperationStatus;
    provisionStatusCode *string;
    schedulingGroups []SchedulingGroup;
    shifts []Shift;
    swapShiftsChangeRequests []SwapShiftsChangeRequest;
    swapShiftsRequestsEnabled *bool;
    timeClockEnabled *bool;
    timeOffReasons []TimeOffReason;
    timeOffRequests []TimeOffRequest;
    timeOffRequestsEnabled *bool;
    timesOff []TimeOff;
    timeZone *string;
    workforceIntegrationIds []string;
}
func NewSchedule()(*Schedule) {
    m := &Schedule{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Schedule) GetEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enabled
    }
}
func (m *Schedule) GetOfferShiftRequests()([]OfferShiftRequest) {
    if m == nil {
        return nil
    } else {
        return m.offerShiftRequests
    }
}
func (m *Schedule) GetOfferShiftRequestsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.offerShiftRequestsEnabled
    }
}
func (m *Schedule) GetOpenShiftChangeRequests()([]OpenShiftChangeRequest) {
    if m == nil {
        return nil
    } else {
        return m.openShiftChangeRequests
    }
}
func (m *Schedule) GetOpenShifts()([]OpenShift) {
    if m == nil {
        return nil
    } else {
        return m.openShifts
    }
}
func (m *Schedule) GetOpenShiftsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.openShiftsEnabled
    }
}
func (m *Schedule) GetProvisionStatus()(*OperationStatus) {
    if m == nil {
        return nil
    } else {
        return m.provisionStatus
    }
}
func (m *Schedule) GetProvisionStatusCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.provisionStatusCode
    }
}
func (m *Schedule) GetSchedulingGroups()([]SchedulingGroup) {
    if m == nil {
        return nil
    } else {
        return m.schedulingGroups
    }
}
func (m *Schedule) GetShifts()([]Shift) {
    if m == nil {
        return nil
    } else {
        return m.shifts
    }
}
func (m *Schedule) GetSwapShiftsChangeRequests()([]SwapShiftsChangeRequest) {
    if m == nil {
        return nil
    } else {
        return m.swapShiftsChangeRequests
    }
}
func (m *Schedule) GetSwapShiftsRequestsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.swapShiftsRequestsEnabled
    }
}
func (m *Schedule) GetTimeClockEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.timeClockEnabled
    }
}
func (m *Schedule) GetTimeOffReasons()([]TimeOffReason) {
    if m == nil {
        return nil
    } else {
        return m.timeOffReasons
    }
}
func (m *Schedule) GetTimeOffRequests()([]TimeOffRequest) {
    if m == nil {
        return nil
    } else {
        return m.timeOffRequests
    }
}
func (m *Schedule) GetTimeOffRequestsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.timeOffRequestsEnabled
    }
}
func (m *Schedule) GetTimesOff()([]TimeOff) {
    if m == nil {
        return nil
    } else {
        return m.timesOff
    }
}
func (m *Schedule) GetTimeZone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeZone
    }
}
func (m *Schedule) GetWorkforceIntegrationIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.workforceIntegrationIds
    }
}
func (m *Schedule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["enabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnabled(val)
        return nil
    }
    res["offerShiftRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOfferShiftRequest() })
        if err != nil {
            return err
        }
        res := make([]OfferShiftRequest, len(val))
        for i, v := range val {
            res[i] = *(v.(*OfferShiftRequest))
        }
        m.SetOfferShiftRequests(res)
        return nil
    }
    res["offerShiftRequestsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetOfferShiftRequestsEnabled(val)
        return nil
    }
    res["openShiftChangeRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOpenShiftChangeRequest() })
        if err != nil {
            return err
        }
        res := make([]OpenShiftChangeRequest, len(val))
        for i, v := range val {
            res[i] = *(v.(*OpenShiftChangeRequest))
        }
        m.SetOpenShiftChangeRequests(res)
        return nil
    }
    res["openShifts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOpenShift() })
        if err != nil {
            return err
        }
        res := make([]OpenShift, len(val))
        for i, v := range val {
            res[i] = *(v.(*OpenShift))
        }
        m.SetOpenShifts(res)
        return nil
    }
    res["openShiftsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetOpenShiftsEnabled(val)
        return nil
    }
    res["provisionStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperationStatus)
        if err != nil {
            return err
        }
        cast := val.(OperationStatus)
        m.SetProvisionStatus(&cast)
        return nil
    }
    res["provisionStatusCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProvisionStatusCode(val)
        return nil
    }
    res["schedulingGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSchedulingGroup() })
        if err != nil {
            return err
        }
        res := make([]SchedulingGroup, len(val))
        for i, v := range val {
            res[i] = *(v.(*SchedulingGroup))
        }
        m.SetSchedulingGroups(res)
        return nil
    }
    res["shifts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewShift() })
        if err != nil {
            return err
        }
        res := make([]Shift, len(val))
        for i, v := range val {
            res[i] = *(v.(*Shift))
        }
        m.SetShifts(res)
        return nil
    }
    res["swapShiftsChangeRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSwapShiftsChangeRequest() })
        if err != nil {
            return err
        }
        res := make([]SwapShiftsChangeRequest, len(val))
        for i, v := range val {
            res[i] = *(v.(*SwapShiftsChangeRequest))
        }
        m.SetSwapShiftsChangeRequests(res)
        return nil
    }
    res["swapShiftsRequestsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSwapShiftsRequestsEnabled(val)
        return nil
    }
    res["timeClockEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetTimeClockEnabled(val)
        return nil
    }
    res["timeOffReasons"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeOffReason() })
        if err != nil {
            return err
        }
        res := make([]TimeOffReason, len(val))
        for i, v := range val {
            res[i] = *(v.(*TimeOffReason))
        }
        m.SetTimeOffReasons(res)
        return nil
    }
    res["timeOffRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeOffRequest() })
        if err != nil {
            return err
        }
        res := make([]TimeOffRequest, len(val))
        for i, v := range val {
            res[i] = *(v.(*TimeOffRequest))
        }
        m.SetTimeOffRequests(res)
        return nil
    }
    res["timeOffRequestsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetTimeOffRequestsEnabled(val)
        return nil
    }
    res["timesOff"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeOff() })
        if err != nil {
            return err
        }
        res := make([]TimeOff, len(val))
        for i, v := range val {
            res[i] = *(v.(*TimeOff))
        }
        m.SetTimesOff(res)
        return nil
    }
    res["timeZone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTimeZone(val)
        return nil
    }
    res["workforceIntegrationIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetWorkforceIntegrationIds(res)
        return nil
    }
    return res
}
func (m *Schedule) IsNil()(bool) {
    return m == nil
}
func (m *Schedule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOfferShiftRequests()))
        for i, v := range m.GetOfferShiftRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("offerShiftRequests", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("offerShiftRequestsEnabled", m.GetOfferShiftRequestsEnabled())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOpenShiftChangeRequests()))
        for i, v := range m.GetOpenShiftChangeRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("openShiftChangeRequests", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOpenShifts()))
        for i, v := range m.GetOpenShifts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("openShifts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("openShiftsEnabled", m.GetOpenShiftsEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetProvisionStatus() != nil {
        cast := m.GetProvisionStatus().String()
        err = writer.WriteStringValue("provisionStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("provisionStatusCode", m.GetProvisionStatusCode())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSchedulingGroups()))
        for i, v := range m.GetSchedulingGroups() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("schedulingGroups", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetShifts()))
        for i, v := range m.GetShifts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("shifts", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSwapShiftsChangeRequests()))
        for i, v := range m.GetSwapShiftsChangeRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("swapShiftsChangeRequests", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("swapShiftsRequestsEnabled", m.GetSwapShiftsRequestsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("timeClockEnabled", m.GetTimeClockEnabled())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTimeOffReasons()))
        for i, v := range m.GetTimeOffReasons() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("timeOffReasons", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTimeOffRequests()))
        for i, v := range m.GetTimeOffRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("timeOffRequests", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("timeOffRequestsEnabled", m.GetTimeOffRequestsEnabled())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTimesOff()))
        for i, v := range m.GetTimesOff() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("timesOff", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("timeZone", m.GetTimeZone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("workforceIntegrationIds", m.GetWorkforceIntegrationIds())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Schedule) SetEnabled(value *bool)() {
    m.enabled = value
}
func (m *Schedule) SetOfferShiftRequests(value []OfferShiftRequest)() {
    m.offerShiftRequests = value
}
func (m *Schedule) SetOfferShiftRequestsEnabled(value *bool)() {
    m.offerShiftRequestsEnabled = value
}
func (m *Schedule) SetOpenShiftChangeRequests(value []OpenShiftChangeRequest)() {
    m.openShiftChangeRequests = value
}
func (m *Schedule) SetOpenShifts(value []OpenShift)() {
    m.openShifts = value
}
func (m *Schedule) SetOpenShiftsEnabled(value *bool)() {
    m.openShiftsEnabled = value
}
func (m *Schedule) SetProvisionStatus(value *OperationStatus)() {
    m.provisionStatus = value
}
func (m *Schedule) SetProvisionStatusCode(value *string)() {
    m.provisionStatusCode = value
}
func (m *Schedule) SetSchedulingGroups(value []SchedulingGroup)() {
    m.schedulingGroups = value
}
func (m *Schedule) SetShifts(value []Shift)() {
    m.shifts = value
}
func (m *Schedule) SetSwapShiftsChangeRequests(value []SwapShiftsChangeRequest)() {
    m.swapShiftsChangeRequests = value
}
func (m *Schedule) SetSwapShiftsRequestsEnabled(value *bool)() {
    m.swapShiftsRequestsEnabled = value
}
func (m *Schedule) SetTimeClockEnabled(value *bool)() {
    m.timeClockEnabled = value
}
func (m *Schedule) SetTimeOffReasons(value []TimeOffReason)() {
    m.timeOffReasons = value
}
func (m *Schedule) SetTimeOffRequests(value []TimeOffRequest)() {
    m.timeOffRequests = value
}
func (m *Schedule) SetTimeOffRequestsEnabled(value *bool)() {
    m.timeOffRequestsEnabled = value
}
func (m *Schedule) SetTimesOff(value []TimeOff)() {
    m.timesOff = value
}
func (m *Schedule) SetTimeZone(value *string)() {
    m.timeZone = value
}
func (m *Schedule) SetWorkforceIntegrationIds(value []string)() {
    m.workforceIntegrationIds = value
}
