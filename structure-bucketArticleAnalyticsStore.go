package carly_pkg

/*
	This file contains the structure of the S3 Bucket - Bucket-Article-Analytics-Store
*/

// TOP STRUCTURES

// document structure stored in s3 bucket
type BucketAnalytics_TEXT struct {
	ArticleReference string `json:"article_reference"`
	ArticleText      string `json:"article_text"`
	Language         string `json:"language"`
	Newspaper        string `json:"newspaper"`
}

// document to store comprehend analytics results
type BucketAnalytics_COMPREHEND struct {
	KeyPhrases       []BucketAnalytics_COMPREHEND_key_phrases `json:"key_phrases"`
	Entities         []BucketAnalytics_COMPREHEND_entities    `json:"entities"`
	Sentiment        []BucketAnalytics_COMPREHEND_sentiment   `json:"sentiment"`
	ArticleReference string                                   `json:"article_reference"`
}

// SUB STRUCTURES

// struct that describes how the stored sentiment results are structured
type BucketAnalytics_COMPREHEND_sentiment struct {
	Sentiment      string
	SentimentScore BucketAnalytics_COMPREHEND_sentiment_scoredetails
	Sentence       string
}

type BucketAnalytics_COMPREHEND_sentiment_scoredetails struct {
	// shows percentage of confidence if this sentence has a mixed sentiment
	Mixed    float64
	// shows percentage of confidence if this sentence has a negative sentiment
	Negative float64
	// shows percentage of confidence if this sentence has a neutral sentiment
	Neutral  float64
	// shows percentage of confidence if this sentence has a positive sentiment
	Positive float64
}

type BucketAnalytics_COMPREHEND_key_phrases struct {
	// A character offset in the input text that shows where the key phrase begins
	// (the first character is at position 0). The offset returns the position of
	// each UTF-8 code point in the string. A code point is the abstract character
	// from a particular graphical representation. For example, a multi-byte UTF-8
	// character maps to a single code point.
	BeginOffset *int64 `type:"integer"`
	// A character offset in the input text where the key phrase ends. The offset
	// returns the position of each UTF-8 code point in the string. A code point
	// is the abstract character from a particular graphical representation. For
	// example, a multi-byte UTF-8 character maps to a single code point.
	EndOffset *int64 `type:"integer"`
	// The level of confidence that Amazon Comprehend has in the accuracy of the
	// detection.
	Score *float64 `type:"float"`
	// The text of a key noun phrase.
	Text *string `min:"1" type:"string"`
}

type BucketAnalytics_COMPREHEND_entities struct {
	// A character offset in the input text that shows where the entity begins (the
	// first character is at position 0). The offset returns the position of each
	// UTF-8 code point in the string. A code point is the abstract character from
	// a particular graphical representation. For example, a multi-byte UTF-8 character
	// maps to a single code point.
	BeginOffset *int64 `type:"integer"`

	// A character offset in the input text that shows where the entity ends. The
	// offset returns the position of each UTF-8 code point in the string. A code
	// point is the abstract character from a particular graphical representation.
	// For example, a multi-byte UTF-8 character maps to a single code point.
	EndOffset *int64 `type:"integer"`

	// The level of confidence that Amazon Comprehend has in the accuracy of the
	// detection.
	Score *float64 `type:"float"`

	// The text of the entity.
	Text *string `min:"1" type:"string"`

	// The entity's type.
	Type *string `type:"string" enum:"EntityType"`
}
