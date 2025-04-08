package openai

// common.go defines common types used throughout the OpenAI API.

// Usage Represents the total token usage per request to OpenAI.
type Usage struct {
	PromptTokens            int                      `json:"prompt_tokens"`
	CompletionTokens        int                      `json:"completion_tokens"`
	TotalTokens             int                      `json:"total_tokens"`
	PromptTokensDetails     *PromptTokensDetails     `json:"prompt_tokens_details"`
	CompletionTokensDetails *CompletionTokensDetails `json:"completion_tokens_details"`

	// CachedCreationTokens is the number of tokens written to the cache when creating a new entry.
	// Corresponds to "cache_creation_input_tokens" in Anthropic.
	CacheCreationInputTokens int `json:"cache_creation_input_tokens"`
	// CachedReadInputTokens is the number of tokens retrieved from the cache for this request.
	// Corresponds to "cached_tokens" in OpenAI and "cache_read_input_tokens" in Anthropic.
	CacheReadInputTokens int `json:"cache_read_input_tokens"`
}

// CompletionTokensDetails Breakdown of tokens used in a completion.
type CompletionTokensDetails struct {
	AudioTokens     int `json:"audio_tokens"`
	ReasoningTokens int `json:"reasoning_tokens"`
}

// PromptTokensDetails Breakdown of tokens used in the prompt.
type PromptTokensDetails struct {
	AudioTokens  int `json:"audio_tokens"`
	CachedTokens int `json:"cached_tokens"`
}
