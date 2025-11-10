package types

import _ "encoding/xml"

type Status struct{
	State 	string
	Errinfo error
}

const (
	NotStarted 	 = "not_started"
	Running 	 = "running"
	Completed 	 = "completed"
	Failed 		 = "failed"
)

type IntentCheckingResponse struct {
	Intent 			string `xml:"intent"`
	IsSafe 			bool `xml:"is_safe"`
	NeedsReview 	bool `xml:"needs_review"`
	Reason 			string `xml:"reason"`
	IsDone			bool	`xml:"-"`
}

type Step struct {
	Type 		string `xml:"type,attr"` // file, command
	Path 		string `xml:"Path"`
	Content 	string `xml:"Content"`
	Command		string `xml:"Command"`
}

type Project struct {
	Title   	string   `xml:"Projectname"`
}

type CodeMessage struct {
	Type    	string // "thought" | "xml"
	Content 	string
}

type Task struct {
	Type 	string
	Details	any
	Status 	Status
}