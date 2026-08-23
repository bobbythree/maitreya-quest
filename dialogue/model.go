package dialogue

type Dialogue struct {
	ID        string
	StartNode string
	Nodes     map[string]Node
}

type Node struct {
	Speaker string
	Text    string
	Choices []Choice
}

type Choice struct {
	Text     string
	NextNode string
}
