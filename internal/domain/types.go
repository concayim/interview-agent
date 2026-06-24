package domain

import "time"

type Resume struct {
	ID          string    `json:"id"`
	FileName    string    `json:"fileName"`
	ContentType string    `json:"contentType"`
	Text        string    `json:"-"`
	Preview     string    `json:"preview"`
	Keywords    []string  `json:"keywords"`
	Characters  int       `json:"characters"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Question struct {
	ID             string   `json:"id"`
	Language       string   `json:"language"`
	Difficulty     string   `json:"difficulty"`
	Prompt         string   `json:"prompt"`
	StandardAnswer string   `json:"standardAnswer,omitempty"`
	KeyPoints      []string `json:"keyPoints,omitempty"`
	Tags           []string `json:"tags"`
}

type PublicQuestion struct {
	ID            string   `json:"id"`
	PromptID      string   `json:"promptId"`
	Stage         string   `json:"stage"`
	LeadIn        string   `json:"leadIn,omitempty"`
	Language      string   `json:"language"`
	Difficulty    string   `json:"difficulty"`
	Prompt        string   `json:"prompt"`
	Tags          []string `json:"tags"`
	FollowUp      bool     `json:"followUp"`
	Round         int      `json:"round"`
	FollowUpTotal int      `json:"followUpTotal"`
}

func (q Question) Public() PublicQuestion {
	return PublicQuestion{ID: q.ID, PromptID: q.ID, Stage: "technical", Language: q.Language, Difficulty: q.Difficulty, Prompt: q.Prompt, Tags: q.Tags}
}

type Evaluation struct {
	Score        int      `json:"score"`
	Summary      string   `json:"summary"`
	Strengths    []string `json:"strengths"`
	Improvements []string `json:"improvements"`
	Source       string   `json:"source"`
}

type AnswerRecord struct {
	Question       Question         `json:"question"`
	Answer         string           `json:"answer"`
	ElapsedSeconds int              `json:"elapsedSeconds"`
	Evaluation     Evaluation       `json:"evaluation"`
	FollowUps      []FollowUpRecord `json:"followUps"`
	AverageScore   int              `json:"averageScore"`
}

type FollowUpRecord struct {
	Round          int        `json:"round"`
	Prompt         string     `json:"prompt"`
	Answer         string     `json:"answer"`
	ElapsedSeconds int        `json:"elapsedSeconds"`
	Evaluation     Evaluation `json:"evaluation"`
}

type IntroductionRecord struct {
	Prompt         string `json:"prompt"`
	Answer         string `json:"answer"`
	ElapsedSeconds int    `json:"elapsedSeconds"`
}

type Session struct {
	ID                 string              `json:"id"`
	CandidateName      string              `json:"candidateName"`
	ResumeID           string              `json:"resumeId,omitempty"`
	Language           string              `json:"language"`
	Difficulty         string              `json:"difficulty"`
	Industry           string              `json:"industry"`
	DomainSkillID      string              `json:"domainSkillId"`
	DomainSkillName    string              `json:"domainSkillName"`
	InterviewerSkillID string              `json:"interviewerSkillId"`
	InterviewerName    string              `json:"interviewerName"`
	InterviewerOpening string              `json:"interviewerOpening"`
	InterviewerPrompt  string              `json:"-"`
	InterviewerIntro   string              `json:"-"`
	EvaluationFocus    []string            `json:"-"`
	FeedbackTone       string              `json:"-"`
	IncludeFoundation  bool                `json:"includeFoundation"`
	Status             string              `json:"status"`
	Phase              string              `json:"phase"`
	Questions          []Question          `json:"-"`
	ResumeKeywords     []string            `json:"-"`
	IntroductionWords  []string            `json:"-"`
	Introduction       *IntroductionRecord `json:"introduction,omitempty"`
	Current            int                 `json:"current"`
	FollowUpRound      int                 `json:"followUpRound"`
	FollowUpTotal      int                 `json:"followUpTotal"`
	CurrentPrompt      string              `json:"-"`
	Answers            []AnswerRecord      `json:"answers"`
	StartedAt          time.Time           `json:"startedAt"`
	CompletedAt        *time.Time          `json:"completedAt,omitempty"`
}

type Report struct {
	SessionID       string              `json:"sessionId"`
	CandidateName   string              `json:"candidateName"`
	Language        string              `json:"language"`
	Difficulty      string              `json:"difficulty"`
	DomainSkillName string              `json:"domainSkillName"`
	InterviewerName string              `json:"interviewerName"`
	Score           int                 `json:"score"`
	Answered        int                 `json:"answered"`
	Duration        int                 `json:"durationSeconds"`
	Highlights      []string            `json:"highlights"`
	FocusAreas      []string            `json:"focusAreas"`
	Introduction    *IntroductionRecord `json:"introduction,omitempty"`
	Answers         []AnswerRecord      `json:"answers"`
	StartedAt       time.Time           `json:"startedAt"`
	CompletedAt     time.Time           `json:"completedAt"`
}
