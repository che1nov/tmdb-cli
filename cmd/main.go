package main

import (
	"flag"
	"fmt"
	"os"
	services "tmdb-cli/internal/service"

	"tmdb-cli/internal/cache"
	"tmdb-cli/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	if cfg.TMDBApiKey == "" {
		fmt.Println("TMDB_API_KEY is not set")
		os.Exit(1)
	}

	movieType := flag.String("type", "", "Type of movies: playing, popular, top, upcoming")
	flag.Parse()
	if *movieType == "" {
		fmt.Println("Usage: tmdb-cli --type <playing|popular|top|upcoming>")
		os.Exit(1)
	}

	cacheClient, err := cache.NewRedisCache(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)
	if err != nil {
		fmt.Printf("Failed to initialize Redis cache: %v\n", err)
		os.Exit(1)
	}

	movies, err := services.FetchMovies(*movieType, cfg, cacheClient)
	if err != nil {
		fmt.Printf("Error fetching movies: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Movies (%s):\n", *movieType)
	for _, movie := range movies {
		fmt.Printf("- %s (%s) - Rating: %.1f\n", movie.Title, movie.ReleaseDate, movie.VoteAverage)
	}
}
