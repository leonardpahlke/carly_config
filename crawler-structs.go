package carly_pkg

const EnvSpiderName = "NAME"
const EnvArticleBucket = "ARTICLE_BUCKET"
const EnvArticleBucketAnalytics = "ARTICLE_BUCKET_ANALYTICS"
const EnvSpiderRoleArn = "SPIDER_ROLE_ARN"
const EnvFilePrefix = "FILE_PREFIX"
const EnvLogLevel = "LOG_LEVEL"

// Downloader
type CarlyEngineLmbDownloaderEvent struct {
	ArticleReference string `json:"article_reference"`
	ArticleUrl       string `json:"article_url"`
	Newspaper        string `json:"newspaper"`
}
type CarlyEngineLmbDownloaderResponse struct {
	ArticleDom       string `json:"article_dom"`
	S3ArticleDomLink string `json:"s_3_article_dom_link"`
	ArticleReference string `json:"article_reference"`
	ArticleUrl       string `json:"article_url"`
	Newspaper        string `json:"newspaper"`
}

// Parser
type CarlyEngineLmbParserEvent struct {
	ArticleDom       string `json:"article_dom"`
	ArticleReference string `json:"article_reference"`
	S3ArticleDomLink string `json:"s_3_article_dom_link"`
	Newspaper        string `json:"newspaper"`
}
type CarlyEngineLmbParserResponse struct {
	ArticleReference string `json:"article_reference"`
	S3ArticleDomLink string `json:"s_3_article_dom_link"`
	Newspaper        string `json:"newspaper"`

	ArticleText string `json:"article_text"`
}

// ML
type CarlyEngineLmbMLEvent struct {
	ArticleReference string `json:"article_reference"`
	Newspaper        string `json:"newspaper"`
	ArticleText      string `json:"article_text"`
}
type CarlyEngineLmbMLResponse struct {
	ArticleReference string `json:"article_reference"`
	Newspaper        string `json:"newspaper"`
	S3ArticleFileUrl string `json:"s3_article_file_url"`
}

// Translator
type CarlyEngineLmbTranslatorEvent struct {
	ArticleReference string `json:"article_reference"`
	Newspaper        string `json:"newspaper"`
	ArticleText      string `json:"article_text"`
}
type CarlyEngineLmbTranslatorResponse struct {
	ArticleReference string `json:"article_reference"`
	Newspaper        string `json:"newspaper"`
	ArticleText      string `json:"article_text"`
}
