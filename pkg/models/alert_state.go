package models

// type AlertState struct {
// 	Id              int64            `json:"-"`
// 	OrgId           int64            `json:"-"`
// 	AlertId         int64            `json:"alertId"`
// 	State           string           `json:"state"`
// 	Created         time.Time        `json:"created"`
// 	Info            string           `json:"info"`
// 	TriggeredAlerts *simplejson.Json `json:"triggeredAlerts"`
// }
//
// func (this *UpdateAlertStateCommand) IsValidState() bool {
// 	for _, v := range alertstates.ValidStates {
// 		if this.State == v {
// 			return true
// 		}
// 	}
// 	return false
// }
//
// // Commands
//
// type UpdateAlertStateCommand struct {
// 	AlertId int64  `json:"alertId" binding:"Required"`
// 	OrgId   int64  `json:"orgId" binding:"Required"`
// 	State   string `json:"state" binding:"Required"`
// 	Info    string `json:"info"`
//
// 	Result *Alert
// }
//
// // Queries
//
// type GetAlertsStateQuery struct {
// 	OrgId   int64 `json:"orgId" binding:"Required"`
// 	AlertId int64 `json:"alertId" binding:"Required"`
//
// 	Result *[]AlertState
// }
//
// type GetLastAlertStateQuery struct {
// 	AlertId int64
// 	OrgId   int64
//
// 	Result *AlertState
// }
