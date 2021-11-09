package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Schedule struct {
    Entity
    // Indicates whether the schedule is enabled for the team. Required.
    enabled *bool;
    // 
    offerShiftRequests []OfferShiftRequest;
    // Indicates whether offer shift requests are enabled for the schedule.
    offerShiftRequestsEnabled *bool;
    // 
    openShiftChangeRequests []OpenShiftChangeRequest;
    // 
    openShifts []OpenShift;
    // Indicates whether open shifts are enabled for the schedule.
    openShiftsEnabled *bool;
    // The status of the schedule provisioning. The possible values are notStarted, running, completed, failed.
    provisionStatus *OperationStatus;
    // Additional information about why schedule provisioning failed.
    provisionStatusCode *string;
    // The logical grouping of users in the schedule (usually by role).
    schedulingGroups []SchedulingGroup;
    // The shifts in the schedule.
    shifts []Shift;
    // 
    swapShiftsChangeRequests []SwapShiftsChangeRequest;
    // Indicates whether swap shifts requests are enabled for the schedule.
    swapShiftsRequestsEnabled *bool;
    // Indicates whether time clock is enabled for the schedule.
    timeClockEnabled *bool;
    // The set of reasons for a time off in the schedule.
    timeOffReasons []TimeOffReason;
    // 
    timeOffRequests []TimeOffRequest;
    // Indicates whether time off requests are enabled for the schedule.
    timeOffRequestsEnabled *bool;
    // The instances of times off in the schedule.
    timesOff []TimeOff;
    // Indicates the time zone of the schedule team using tz database format. Required.
    timeZone *string;
    // 
    workforceIntegrationIds []string;
}
// Instantiates a new schedule and sets the default values.
func NewSchedule()(*Schedule) {
    m := &Schedule{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the enabled property value. Indicates whether the schedule is enabled for the team. Required.
func (m *Schedule) GetEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enabled
    }
}
// Gets the offerShiftRequests property value. 
func (m *Schedule) GetOfferShiftRequests()([]OfferShiftRequest) {
    if m == nil {
        return nil
    } else {
        return m.offerShiftRequests
    }
}
// Gets the offerShiftRequestsEnabled property value. Indicates whether offer shift requests are enabled for the schedule.
func (m *Schedule) GetOfferShiftRequestsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.offerShiftRequestsEnabled
    }
}
// Gets the openShiftChangeRequests property value. 
func (m *Schedule) GetOpenShiftChangeRequests()([]OpenShiftChangeRequest) {
    if m == nil {
        return nil
    } else {
        return m.openShiftChangeRequests
    }
}
// Gets the openShifts property value. 
func (m *Schedule) GetOpenShifts()([]OpenShift) {
    if m == nil {
        return nil
    } else {
        return m.openShifts
    }
}
// Gets the openShiftsEnabled property value. Indicates whether open shifts are enabled for the schedule.
func (m *Schedule) GetOpenShiftsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.openShiftsEnabled
    }
}
// Gets the provisionStatus property value. The status of the schedule provisioning. The possible values are notStarted, running, completed, failed.
func (m *Schedule) GetProvisionStatus()(*OperationStatus) {
    if m == nil {
        return nil
    } else {
        return m.provisionStatus
    }
}
// Gets the provisionStatusCode property value. Additional information about why schedule provisioning failed.
func (m *Schedule) GetProvisionStatusCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.provisionStatusCode
    }
}
// Gets the schedulingGroups property value. The logical grouping of users in the schedule (usually by role).
func (m *Schedule) GetSchedulingGroups()([]SchedulingGroup) {
    if m == nil {
        return nil
    } else {
        return m.schedulingGroups
    }
}
// Gets the shifts property value. The shifts in the schedule.
func (m *Schedule) GetShifts()([]Shift) {
    if m == nil {
        return nil
    } else {
        return m.shifts
    }
}
// Gets the swapShiftsChangeRequests property value. 
func (m *Schedule) GetSwapShiftsChangeRequests()([]SwapShiftsChangeRequest) {
    if m == nil {
        return nil
    } else {
        return m.swapShiftsChangeRequests
    }
}
// Gets the swapShiftsRequestsEnabled property value. Indicates whether swap shifts requests are enabled for the schedule.
func (m *Schedule) GetSwapShiftsRequestsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.swapShiftsRequestsEnabled
    }
}
// Gets the timeClockEnabled property value. Indicates whether time clock is enabled for the schedule.
func (m *Schedule) GetTimeClockEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.timeClockEnabled
    }
}
// Gets the timeOffReasons property value. The set of reasons for a time off in the schedule.
func (m *Schedule) GetTimeOffReasons()([]TimeOffReason) {
    if m == nil {
        return nil
    } else {
        return m.timeOffReasons
    }
}
// Gets the timeOffRequests property value. 
func (m *Schedule) GetTimeOffRequests()([]TimeOffRequest) {
    if m == nil {
        return nil
    } else {
        return m.timeOffRequests
    }
}
// Gets the timeOffRequestsEnabled property value. Indicates whether time off requests are enabled for the schedule.
func (m *Schedule) GetTimeOffRequestsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.timeOffRequestsEnabled
    }
}
// Gets the timesOff property value. The instances of times off in the schedule.
func (m *Schedule) GetTimesOff()([]TimeOff) {
    if m == nil {
        return nil
    } else {
        return m.timesOff
    }
}
// Gets the timeZone property value. Indicates the time zone of the schedule team using tz database format. Required.
func (m *Schedule) GetTimeZone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeZone
    }
}
// Gets the workforceIntegrationIds property value. 
func (m *Schedule) GetWorkforceIntegrationIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.workforceIntegrationIds
    }
}
// The deserialization information for the current model
func (m *Schedule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["enabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["offerShiftRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOfferShiftRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OfferShiftRequest, len(val))
            for i, v := range val {
                res[i] = *(v.(*OfferShiftRequest))
            }
            m.SetOfferShiftRequests(res)
        }
        return nil
    }
    res["offerShiftRequestsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfferShiftRequestsEnabled(val)
        }
        return nil
    }
    res["openShiftChangeRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOpenShiftChangeRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OpenShiftChangeRequest, len(val))
            for i, v := range val {
                res[i] = *(v.(*OpenShiftChangeRequest))
            }
            m.SetOpenShiftChangeRequests(res)
        }
        return nil
    }
    res["openShifts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOpenShift() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OpenShift, len(val))
            for i, v := range val {
                res[i] = *(v.(*OpenShift))
            }
            m.SetOpenShifts(res)
        }
        return nil
    }
    res["openShiftsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOpenShiftsEnabled(val)
        }
        return nil
    }
    res["provisionStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(OperationStatus)
            m.SetProvisionStatus(&cast)
        }
        return nil
    }
    res["provisionStatusCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisionStatusCode(val)
        }
        return nil
    }
    res["schedulingGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSchedulingGroup() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SchedulingGroup, len(val))
            for i, v := range val {
                res[i] = *(v.(*SchedulingGroup))
            }
            m.SetSchedulingGroups(res)
        }
        return nil
    }
    res["shifts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewShift() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Shift, len(val))
            for i, v := range val {
                res[i] = *(v.(*Shift))
            }
            m.SetShifts(res)
        }
        return nil
    }
    res["swapShiftsChangeRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSwapShiftsChangeRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SwapShiftsChangeRequest, len(val))
            for i, v := range val {
                res[i] = *(v.(*SwapShiftsChangeRequest))
            }
            m.SetSwapShiftsChangeRequests(res)
        }
        return nil
    }
    res["swapShiftsRequestsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSwapShiftsRequestsEnabled(val)
        }
        return nil
    }
    res["timeClockEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeClockEnabled(val)
        }
        return nil
    }
    res["timeOffReasons"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeOffReason() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimeOffReason, len(val))
            for i, v := range val {
                res[i] = *(v.(*TimeOffReason))
            }
            m.SetTimeOffReasons(res)
        }
        return nil
    }
    res["timeOffRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeOffRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimeOffRequest, len(val))
            for i, v := range val {
                res[i] = *(v.(*TimeOffRequest))
            }
            m.SetTimeOffRequests(res)
        }
        return nil
    }
    res["timeOffRequestsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeOffRequestsEnabled(val)
        }
        return nil
    }
    res["timesOff"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeOff() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimeOff, len(val))
            for i, v := range val {
                res[i] = *(v.(*TimeOff))
            }
            m.SetTimesOff(res)
        }
        return nil
    }
    res["timeZone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZone(val)
        }
        return nil
    }
    res["workforceIntegrationIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetWorkforceIntegrationIds(res)
        }
        return nil
    }
    return res
}
func (m *Schedule) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the enabled property value. Indicates whether the schedule is enabled for the team. Required.
// Parameters:
//  - value : Value to set for the enabled property.
func (m *Schedule) SetEnabled(value *bool)() {
    m.enabled = value
}
// Sets the offerShiftRequests property value. 
// Parameters:
//  - value : Value to set for the offerShiftRequests property.
func (m *Schedule) SetOfferShiftRequests(value []OfferShiftRequest)() {
    m.offerShiftRequests = value
}
// Sets the offerShiftRequestsEnabled property value. Indicates whether offer shift requests are enabled for the schedule.
// Parameters:
//  - value : Value to set for the offerShiftRequestsEnabled property.
func (m *Schedule) SetOfferShiftRequestsEnabled(value *bool)() {
    m.offerShiftRequestsEnabled = value
}
// Sets the openShiftChangeRequests property value. 
// Parameters:
//  - value : Value to set for the openShiftChangeRequests property.
func (m *Schedule) SetOpenShiftChangeRequests(value []OpenShiftChangeRequest)() {
    m.openShiftChangeRequests = value
}
// Sets the openShifts property value. 
// Parameters:
//  - value : Value to set for the openShifts property.
func (m *Schedule) SetOpenShifts(value []OpenShift)() {
    m.openShifts = value
}
// Sets the openShiftsEnabled property value. Indicates whether open shifts are enabled for the schedule.
// Parameters:
//  - value : Value to set for the openShiftsEnabled property.
func (m *Schedule) SetOpenShiftsEnabled(value *bool)() {
    m.openShiftsEnabled = value
}
// Sets the provisionStatus property value. The status of the schedule provisioning. The possible values are notStarted, running, completed, failed.
// Parameters:
//  - value : Value to set for the provisionStatus property.
func (m *Schedule) SetProvisionStatus(value *OperationStatus)() {
    m.provisionStatus = value
}
// Sets the provisionStatusCode property value. Additional information about why schedule provisioning failed.
// Parameters:
//  - value : Value to set for the provisionStatusCode property.
func (m *Schedule) SetProvisionStatusCode(value *string)() {
    m.provisionStatusCode = value
}
// Sets the schedulingGroups property value. The logical grouping of users in the schedule (usually by role).
// Parameters:
//  - value : Value to set for the schedulingGroups property.
func (m *Schedule) SetSchedulingGroups(value []SchedulingGroup)() {
    m.schedulingGroups = value
}
// Sets the shifts property value. The shifts in the schedule.
// Parameters:
//  - value : Value to set for the shifts property.
func (m *Schedule) SetShifts(value []Shift)() {
    m.shifts = value
}
// Sets the swapShiftsChangeRequests property value. 
// Parameters:
//  - value : Value to set for the swapShiftsChangeRequests property.
func (m *Schedule) SetSwapShiftsChangeRequests(value []SwapShiftsChangeRequest)() {
    m.swapShiftsChangeRequests = value
}
// Sets the swapShiftsRequestsEnabled property value. Indicates whether swap shifts requests are enabled for the schedule.
// Parameters:
//  - value : Value to set for the swapShiftsRequestsEnabled property.
func (m *Schedule) SetSwapShiftsRequestsEnabled(value *bool)() {
    m.swapShiftsRequestsEnabled = value
}
// Sets the timeClockEnabled property value. Indicates whether time clock is enabled for the schedule.
// Parameters:
//  - value : Value to set for the timeClockEnabled property.
func (m *Schedule) SetTimeClockEnabled(value *bool)() {
    m.timeClockEnabled = value
}
// Sets the timeOffReasons property value. The set of reasons for a time off in the schedule.
// Parameters:
//  - value : Value to set for the timeOffReasons property.
func (m *Schedule) SetTimeOffReasons(value []TimeOffReason)() {
    m.timeOffReasons = value
}
// Sets the timeOffRequests property value. 
// Parameters:
//  - value : Value to set for the timeOffRequests property.
func (m *Schedule) SetTimeOffRequests(value []TimeOffRequest)() {
    m.timeOffRequests = value
}
// Sets the timeOffRequestsEnabled property value. Indicates whether time off requests are enabled for the schedule.
// Parameters:
//  - value : Value to set for the timeOffRequestsEnabled property.
func (m *Schedule) SetTimeOffRequestsEnabled(value *bool)() {
    m.timeOffRequestsEnabled = value
}
// Sets the timesOff property value. The instances of times off in the schedule.
// Parameters:
//  - value : Value to set for the timesOff property.
func (m *Schedule) SetTimesOff(value []TimeOff)() {
    m.timesOff = value
}
// Sets the timeZone property value. Indicates the time zone of the schedule team using tz database format. Required.
// Parameters:
//  - value : Value to set for the timeZone property.
func (m *Schedule) SetTimeZone(value *string)() {
    m.timeZone = value
}
// Sets the workforceIntegrationIds property value. 
// Parameters:
//  - value : Value to set for the workforceIntegrationIds property.
func (m *Schedule) SetWorkforceIntegrationIds(value []string)() {
    m.workforceIntegrationIds = value
}
