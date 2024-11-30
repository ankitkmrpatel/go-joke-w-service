package business

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/ankitkmrpatel/go-joke-w-service/internal/infra"
	"github.com/ankitkmrpatel/go-joke-w-service/internal/models"
)

func PrintRandomJoke(cfg *models.Config, logger infra.Logger) {
	cfg.RLock()
	defer cfg.RUnlock()

	if len(cfg.Jokes) == 0 {
		logger.Info("No jokes available in the configuration.")
		return
	}
	joke := cfg.Jokes[rand.Intn(len(cfg.Jokes))]
	logger.Info("Random Joke: " + joke)
}

func PrintAPIJoke(logger infra.Logger) {
	apiURL := "https://v2.jokeapi.dev/joke/Programming?type=single"
	startTime := time.Now()

	resp, err := http.Get(apiURL)
	duration := time.Since(startTime).Seconds()
	infra.APIRequestDuration.Observe(duration)

	if err != nil {
		logger.Info("Failed to fetch joke from API: " + err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		logger.Info("API responded with status: " + resp.Status)
		return
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		logger.Info("Failed to parse joke from API: " + err.Error())
		return
	}

	if joke, ok := result["joke"].(string); ok {
		logger.Info("API Joke: " + joke)
		infra.JokesFromAPICounter.Inc()
	} else {
		logger.Info("No joke received from API.")
	}
}
