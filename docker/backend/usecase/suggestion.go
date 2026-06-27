package usecase

import (
	"github.com/hackThacker/OWASP-AttackForge/backend/domain"
)

type suggestionUsecase struct{}

func NewSuggestionUsecase() domain.SuggestionUsecase {
	return &suggestionUsecase{}
}

var suggestionsList = []*domain.Suggestion{
	{
		Text:     "Configure crAPI to learn about Broken Object Level Authorization (BOLA).",
		Category: "API SECURITY",
		Icon:     "fas fa-shield-alt",
	},
	{
		Text:     "Test Juice Shop for DOM-based Cross-Site Scripting (XSS) and SQL Injection.",
		Category: "WEB APPLICATION",
		Icon:     "fas fa-bug",
	},
	{
		Text:     "Exploit Tomcat Manager Console with brute-force to achieve Remote Code Execution.",
		Category: "INFRASTRUCTURE",
		Icon:     "fas fa-server",
	},
	{
		Text:     "Investigate BrokenCrystals AI and Vector DB vulnerabilities using Ollama.",
		Category: "AI SECURITY",
		Icon:     "fas fa-brain",
	},
	{
		Text:     "Examine ZeroHealth to understand medical application data leaks.",
		Category: "HEALTHCARE SEC",
		Icon:     "fas fa-heartbeat",
	},
	{
		Text:     "Solve WrongSecrets challenges to find hardcoded credentials.",
		Category: "CLOUD SECURITY",
		Icon:     "fas fa-key",
	},
	{
		Text:     "Use Security Shepherd challenges to practice mobile and web API security.",
		Category: "TRAINING",
		Icon:     "fas fa-graduation-cap",
	},
	{
		Text:     "Analyze RESTaurant Swagger API playground for injection flaws.",
		Category: "API SECURITY",
		Icon:     "fas fa-utensils",
	},
	{
		Text:     "Practice session hijacking and command injection using DVWA.",
		Category: "WEB APPLICATION",
		Icon:     "fas fa-unlock-alt",
	},
	{
		Text:     "Audit the Nginx router logs for suspicious penetration testing traffic.",
		Category: "BEST PRACTICE",
		Icon:     "fas fa-eye",
	},
}

func (u *suggestionUsecase) GetSuggestions() ([]*domain.Suggestion, error) {
	return suggestionsList, nil
}
