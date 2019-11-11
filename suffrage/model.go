package suffrage

//Vote represents a vote
type Vote struct {
	VoterName   	string `json:"votername"`
	CandidateID 	string `json:"candidateid"`
	CandidateName 	string `json:"candidatename"`
}

//Votes is an array of Vote
type Votes []Vote

//Block ...
type Block struct {
	Index             int    `json:"index"`
	Timestamp         int64  `json:"timestamp"`
	Votes             Votes  `json:"votes"`
	Nonce             int    `json:"nonce"`
	Hash              string `json:"hash"`
	PreviousBlockHash string `json:"previousblockhash"`
}

//Blocks is an array of Block
type Blocks []Block

//Blockchain ...
type Blockchain struct {
	Chain        	Blocks   `json:"chain"`
	PendingVotes  	Votes     `json:"pending_votes"`
	NetworkNodes 	[]string `json:"network_nodes"`
}

//BlockData is used in hash calculations
type BlockData struct {
	Index 	string
	Votes  	Votes
}
