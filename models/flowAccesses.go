package models

import "strings"

// FlowAccesses is used to tell if a flow is visible to a Profile
type FlowAccesses struct {
	Enabled bool   `xml:"enabled"`
	Flow    string `xml:"flow"`
}

type Flows []*FlowAccesses

// Implement Sort interface
func (f Flows) Len() int      { return len(f) }
func (f Flows) Swap(i, j int) { f[i], f[j] = f[j], f[i] }

// ByUserPermName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Flows value.
type ByFlowName struct{ Flows }

func (n ByFlowName) Less(i, j int) bool {
	return strings.Compare(n.Flows[i].Flow, n.Flows[j].Flow) < 0
}
