/* -------------------------------------------------------------------------- */
/* Copyright 2002-2020, OpenNebula Project, OpenNebula Systems                */
/*                                                                            */
/* Licensed under the Apache License, Version 2.0 (the "License"); you may    */
/* not use this file except in compliance with the License. You may obtain    */
/* a copy of the License at                                                   */
/*                                                                            */
/* http://www.apache.org/licenses/LICENSE-2.0                                 */
/*                                                                            */
/* Unless required by applicable law or agreed to in writing, software        */
/* distributed under the License is distributed on an "AS IS" BASIS,          */
/* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.   */
/* See the License for the specific language governing permissions and        */
/* limitations under the License.                                             */
/*--------------------------------------------------------------------------- */

package vm

import (
	"encoding/xml"

	dyn "github.com/OpenNebula/one/src/oca/go/src/goca/dynamic"

	"github.com/OpenNebula/one/src/oca/go/src/goca/schemas/shared"
)

// Pool represents an OpenNebula Virtual Machine pool
type Pool struct {
	XMLName xml.Name `xml:"VM_POOL"`
	VMs     []VM     `xml:"VM"`
}

// VM represents an OpenNebula Virtual Machine
type VM struct {
	XMLName         xml.Name            `xml:"VM"`
	ID              int                 `xml:"ID,omitempty"`
	UID             int                 `xml:"UID,omitempty"`
	GID             int                 `xml:"GID,omitempty"`
	UName           string              `xml:"UNAME,omitempty"`
	GName           string              `xml:"GNAME,omitempty"`
	Name            string              `xml:"NAME,omitempty"`
	Permissions     *shared.Permissions `xml:"PERMISSIONS,omitempty"`
	LastPoll        int                 `xml:"LAST_POLL,omitempty"`
	StateRaw        int                 `xml:"STATE,omitempty"`
	LCMStateRaw     int                 `xml:"LCM_STATE,omitempty"`
	PrevStateRaw    int                 `xml:"PREV_STATE,omitempty"`
	PrevLCMStateRaw int                 `xml:"PREV_LCM_STATE,omitempty"`
	ReschedValue    int                 `xml:"RESCHED,omitempty"`
	STime           int                 `xml:"STIME,omitempty"`
	ETime           int                 `xml:"ETIME,omitempty"`
	DeployID        string              `xml:"DEPLOY_ID,omitempty"`
	MonitoringInfos Monitoring          `xml:"MONITORING,omitempty"`
	Template        Template            `xml:"TEMPLATE,omitempty"`
	UserTemplate    UserTemplate        `xml:"USER_TEMPLATE,omitempty"`
	HistoryRecords  []HistoryRecord     `xml:"HISTORY_RECORDS>HISTORY,omitempty"`
	Snapshots       shared.DiskSnapshot `xml:"SNAPSHOTS,omitempty"`

	// Not filled with NewUserPool call
	LockInfos *shared.Lock `xml:"LOCK"`
}

// Monitoring is a dynamic VM part containing metrics
type Monitoring struct {
	dyn.Template
}

// History records
type HistoryRecord struct {
	OID       int    `xml:"OID"`
	SEQ       int    `xml:"SEQ"`
	Hostname  string `xml:"HOSTNAME"`
	HID       int    `xml:"HID"`
	CID       int    `xml:"CID"`
	STime     int    `xml:"STIME"`
	ETime     int    `xml:"ETIME"`
	VMMad     string `xml:"VM_MAD"`
	TMMad     string `xml:"TM_MAD"`
	DSID      int    `xml:"DS_ID"`
	PSTime    int    `xml:"PSTIME"`
	PETime    int    `xml:"PETIME"`
	RSTime    int    `xml:"RSTIME"`
	RETime    int    `xml:"RETIME"`
	ESTime    int    `xml:"ESTIME"`
	EETime    int    `xml:"EETIME"`
	Action    int    `xml:"ACTION"`
	UID       int    `xml:"UID"`
	GID       int    `xml:"GID"`
	RequestID string `xml:"REQUEST_ID"`
}

// Showback Records
type ShowbackRecord struct {
	XMLName  xml.Name `xml:"SHOWBACK_RECORDS"`
	SHOWBACK []struct {
		VMID       string `xml:"VMID"`
		VMNAME     string `xml:"VMNAME"`
		UID        string `xml:"UID"`
		GID        string `xml:"GID"`
		UNAME      string `xml:"UNAME"`
		GNAME      string `xml:"GNAME"`
		YEAR       string `xml:"YEAR"`
		MONTH      string `xml:"MONTH"`
		CPUCOST    string `xml:"CPU_COST"`
		MEMORYCOST string `xml:"MEMORY_COST"`
		DISKCOST   string `xml:"DISK_COST"`
		TOTALCOST  string `xml:"TOTAL_COST"`
		HOURS      string `xml:"HOURS"`
	} `xml:"SHOWBACK"`
}