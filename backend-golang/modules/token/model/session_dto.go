package tokenmodel

type SaveSessionDTO struct {
	DocID       *string `json:"doc_id"`
	ContentHash string  `json:"content_hash"`
	Proof       string  `json:"proof"`
	SessionId   int64   `json:"session_id"`
	RiskScore   int64   `json:"risk_score"`
}
