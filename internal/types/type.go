package types

import "encoding/xml"

type Status string

const (
	NotStarted 		Status = "not_started"
	Running 		Status = "running"
	Completed 		Status = "completed"
	Failed 			Status = "failed"
)

type IntentCheckingResponse struct {
	Intent 			string `xml:"intent"`
	IsSafe 			bool `xml:"is_safe"`
	NeedsReview 	bool `xml:"needs_review"`
	Reason 			string `xml:"reason"`
	SanitizedPrompt string `xml:"sanitized_prompt"`
}

type Step struct {
	Id 			string `xml:"id,attr"` // file, command
	FilePath 	string `xml:"FilePath"`
	Content 	string `xml:"Content"`
	Command		string `xml:"Command"`
	Status  	string `xml:"-"`
}

type Project struct {
	XMLName 	xml.Name `xml:"PollenArtifact"`
	Title   	string   `xml:"ProjectTitle"`
	Steps   	[]Step   `xml:"Step"`
}

type StreamMessage struct {
	Type    	string // "thought" | "xml"
	Content 	string
}