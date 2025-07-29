package integration

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

var (
    apiToken = os.Getenv("API_TOKEN")
)

func ComparePower(hero, other string) (string, error) {
	heroPower, err := getPower(hero)

    if err != nil {
        return "", fmt.Errorf("failed to get %s power: %w", hero, err)
    }

	otherPower, err := getPower(other)

	if err != nil {
		return "", fmt.Errorf("failed to get %s power: %w", other, err)
	}
	
	switch {
	case heroPower > otherPower:
		return hero, nil
	case heroPower < otherPower:
		return other, nil
	default:
		return "Both superheroes are equally powerful", nil
	}
}

func getPower(hero string) (int, error) {
	resp, err := getHeroData(hero)

	if err != nil {
		return 0, fmt.Errorf("error getting data: %v", err)
	}

	power, err := strconv.Atoi(resp.Results[0].PowerStats.Power)
    if err != nil {
        return 0, fmt.Errorf("invalid power value for %s: %v", hero, err)
    }

	return power, nil
}

func getHeroData(hero string) (APIResponse, error) {
	resp, err := http.Get("https://superheroapi.com/api/" + apiToken + "/search/" + hero)

	if err != nil {
		return APIResponse{}, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return APIResponse{}, fmt.Errorf("failed to read response body: %w", err)
	}
	
	var data APIResponse
    if err := json.Unmarshal(body, &data); err != nil {
        return APIResponse{}, fmt.Errorf("JSON decode error: %w", err)
    }

    return data, nil
}