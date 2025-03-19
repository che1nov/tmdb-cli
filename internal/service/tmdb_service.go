package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"tmdb-cli/internal/cache"
	"tmdb-cli/internal/config"
	"tmdb-cli/internal/models"

	"log/slog"
)

type TMDBResponse struct {
	Page         int             `json:"page"`
	Results      []MovieResponse `json:"results"`
	TotalPages   int             `json:"total_pages"`
	TotalResults int             `json:"total_results"`
}

type MovieResponse struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Overview    string  `json:"overview"`
	ReleaseDate string  `json:"release_date"`
	VoteAverage float32 `json:"vote_average"`
}

// FetchMovies получает фильмы по типу с использованием Redis-кэша.
// Данные сохраняются в кэш, но не сохраняются в базу данных.
func FetchMovies(movieType string, cfg *config.Config, cacheClient *cache.RedisCache) ([]models.Movie, error) {
	cacheKey := "tmdb:" + movieType

	// Попытка получить фильмы из кэша
	cachedMovies, err := cacheClient.GetMovies(cacheKey)
	if err != nil {
		config.Logger.Error("Error getting movies from cache", slog.String("error", err.Error()))
	} else if len(cachedMovies) > 0 {
		config.Logger.Info("Movies retrieved from cache", slog.String("movieType", movieType))
		return cachedMovies, nil
	}

	var endpoint string
	switch movieType {
	case "playing":
		endpoint = "now_playing"
	case "popular":
		endpoint = "popular"
	case "top":
		endpoint = "top_rated"
	case "upcoming":
		endpoint = "upcoming"
	default:
		return nil, fmt.Errorf("invalid movie type: %s", movieType)
	}

	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%s?api_key=%s", endpoint, cfg.TMDBApiKey)
	config.Logger.Info("Fetching movies", slog.String("endpoint", endpoint), slog.String("url", url))

	resp, err := http.Get(url)
	if err != nil {
		config.Logger.Error("Failed to fetch movies", slog.String("error", err.Error()))
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("unexpected status code: %d", resp.StatusCode)
		config.Logger.Error("TMDB API Error", slog.String("status", errMsg))
		return nil, fmt.Errorf(errMsg)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		config.Logger.Error("Failed to read response body", slog.String("error", err.Error()))
		return nil, err
	}

	var tmdbResp TMDBResponse
	if err = json.Unmarshal(body, &tmdbResp); err != nil {
		config.Logger.Error("Failed to parse JSON", slog.String("error", err.Error()))
		return nil, err
	}

	var movies []models.Movie
	for _, mr := range tmdbResp.Results {
		movie := models.Movie{
			TMDBID:      mr.ID,
			Title:       mr.Title,
			Overview:    mr.Overview,
			ReleaseDate: mr.ReleaseDate,
			VoteAverage: mr.VoteAverage,
		}
		movies = append(movies, movie)
	}

	// Сохранение данных в кэш
	if err = cacheClient.SetMovies(cacheKey, movies, 10*time.Minute); err != nil {
		config.Logger.Error("Failed to set movies to cache", slog.String("error", err.Error()))
	}

	return movies, nil
}
