package voice_adapter

import (
	"os"
	"testing"
	"time"
)

func TestNewAdapterConfigFromEnv_Defaults(t *testing.T) {
	os.Unsetenv("SPEECH_SERVICE_URL")
	os.Unsetenv("SPEECH_API_KEY")
	os.Unsetenv("SPEECH_TIMEOUT")

	cfg := NewAdapterConfigFromEnv()

	if cfg.SpeechTimeout != 50*time.Second {
		t.Errorf("expected default timeout 50s, got %v", cfg.SpeechTimeout)
	}
}

func TestNewAdapterConfigFromEnv_Custom(t *testing.T) {
	t.Setenv("SPEECH_SERVICE_URL", "http://speech:8780")
	t.Setenv("SPEECH_API_KEY", "my-key")
	t.Setenv("SPEECH_TIMEOUT", "30")

	cfg := NewAdapterConfigFromEnv()

	if cfg.SpeechServiceURL != "http://speech:8780" {
		t.Errorf("unexpected URL: %s", cfg.SpeechServiceURL)
	}
	if cfg.SpeechAPIKey != "my-key" {
		t.Errorf("unexpected key: %s", cfg.SpeechAPIKey)
	}
	if cfg.SpeechTimeout != 30*time.Second {
		t.Errorf("expected 30s, got %v", cfg.SpeechTimeout)
	}
}

func TestNewAdapterConfigFromEnv_InvalidValues(t *testing.T) {
	t.Setenv("SPEECH_TIMEOUT", "invalid")

	cfg := NewAdapterConfigFromEnv()

	if cfg.SpeechTimeout != 50*time.Second {
		t.Errorf("expected default timeout 50s for invalid value, got %v", cfg.SpeechTimeout)
	}
}
