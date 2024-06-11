package utils_test

import (
	"reflect"
	"testing"

	"github.com/daver-dev/quizzer/models"
	"github.com/daver-dev/quizzer/utils"
)

var testQuestions = []models.Question{
	{
		Question:       "What is the capital of France?",
		Options:        []string{"Paris", "London", "Berlin", "Rome"},
		Correct_Answer: "Paris",
		Distractors:    []string{"London", "Berlin", "Rome"},
	},
	{
		Question:       "Who wrote 'Romeo and Juliet'?",
		Options:        []string{"William Shakespeare", "Jane Austen", "Charles Dickens", "Leo Tolstoy"},
		Correct_Answer: "William Shakespeare",
		Distractors:    []string{"Jane Austen", "Charles Dickens", "Leo Tolstoy"},
	},
}

func TestLoadQuestions(t *testing.T) {
	questions := utils.LoadQuestions("test_questions.json")

	// Iterate over each property because you cant compare slice properties
	for i, question := range questions {
		expected := testQuestions[i]
		if question.Question != expected.Question ||
			!equalSlices(question.Options, expected.Options) ||
			question.Correct_Answer != expected.Correct_Answer ||
			!equalSlices(question.Distractors, expected.Distractors) {
			t.Errorf("Question %d does not match. Got %+v, expected %+v", i, question, expected)
		}
	}
}

func TestFindQuestion(t *testing.T) {
	tests := []struct {
		name         string
		searchString string
		expected     *models.Question
	}{
		{
			name:         "Existing Question",
			searchString: "France",
			expected: &models.Question{
				Question:       "What is the capital of France?",
				Options:        []string{"Paris", "London", "Berlin", "Rome"},
				Correct_Answer: "Paris",
				Distractors:    []string{"London", "Berlin", "Rome"},
			},
		},
		{
			name:         "Non-existing Question",
			searchString: "Italy",
			expected:     nil,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			foundQuestion := utils.FindQuestion(tt.searchString, testQuestions)
			if !reflect.DeepEqual(foundQuestion, tt.expected) {
				t.Errorf("Test case %s failed: expected %v, got %v", tt.name, tt.expected, foundQuestion)
			}
		})
	}
}

func TestIsAnswerCorrect(t *testing.T) {

	tests := []struct {
		name     string
		answer   int
		expected bool
	}{
		{
			name:     "Correct",
			answer:   1,
			expected: true,
		},
		{
			name:     "Incorrect",
			answer:   2,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isCorrect := utils.IsAnswerCorrect(testQuestions[0], tt.answer)
			if !reflect.DeepEqual(isCorrect, tt.expected) {
				t.Errorf("Test case %s failed: expected %v, got %v", tt.name, tt.expected, isCorrect)
			}
		})
	}
}

// Helper function to compare two slices
func equalSlices(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
