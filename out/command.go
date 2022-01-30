package out

type Command struct {
	Component  string        `json:"component"`
	Capability string        `json:"capability"`
	Command    string        `json:"command"`
	Arguments  []interface{} `json:"arguments,omitempty"`
}

type SamsungCommand struct {
	Commands []Command `json:"commands"`
}
