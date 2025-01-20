package blockchain

type Block struct {
	Index     int
	Timestamp string
	VoterID   string
	Candidate string
	PrevHash  string
	Hash      string
}
