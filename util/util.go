package util

import (
	"os"
)

type service struct {
	URL  string
	Host string
	Port string
}

type envData struct {
	DB struct {
		Host string
		User string
		Pass string
	}
	Services struct {
		User     service
		Frontend service
		Blog     service
	}
}

// GetENV returns all the environment variables for services to be able
// to intercommunicate and know how to send each other data without hard coding
// being required as it is setup via environment variables.
func GetENV() envData {
	res := envData{}
	res.DB.Host = os.Getenv("DB_HOST")
	res.DB.User = os.Getenv("DB_USER")
	res.DB.Pass = os.Getenv("DB_PASS")

	res.Services.User = service{
		Host: os.Getenv("SERVICE_USER_HOST"),
		Port: os.Getenv("SERVICE_USER_PORT"),
		URL:  os.Getenv("SERVICE_USER_HOST") + ":" + os.Getenv("SERVICE_USER_PORT"),
	}

	res.Services.Blog = service{
		Host: os.Getenv("SERVICE_BLOG_HOST"),
		Port: os.Getenv("SERVICE_BLOG_PORT"),
		URL:  os.Getenv("SERVICE_BLOG_HOST") + ":" + os.Getenv("SERVICE_BLOG_PORT"),
	}

	res.Services.Frontend = service{
		Host: os.Getenv("SERVICE_FRONTEND_HOST"),
		Port: os.Getenv("SERVICE_FRONTEND_PORT"),
		URL:  os.Getenv("SERVICE_FRONTEND_HOST") + ":" + os.Getenv("SERVICE_FRONTEND_PORT"),
	}
	return res
}
