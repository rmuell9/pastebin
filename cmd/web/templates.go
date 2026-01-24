package main

import "snippetbox.senor-crab.com/internal/models"

type templateData struct  {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}
