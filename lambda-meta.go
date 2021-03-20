package carly_pkg

const GO_RUNTIME = "go1.x"
const DEFAULT_TIMEOUT = 3

type LambdaMeta struct {
	Name            string
	Runtime         string
	Timeout         int
	RequiredEnvKeys []string
}

// Lambda ml
var CARLY_ENGINE_LMB_ML = LambdaMeta{
	Name:            "carly_lmb_ml",
	Runtime:         GO_RUNTIME,
	Timeout:         DEFAULT_TIMEOUT,
	RequiredEnvKeys: []string{""},
}

// Lambda translator
var CARLY_ENGINE_LMB_TRANSLATOR = LambdaMeta{
	Name:            "carly_lmb_translator",
	Runtime:         GO_RUNTIME,
	Timeout:         DEFAULT_TIMEOUT,
	RequiredEnvKeys: []string{""},
}

// Lambda parser
var CARLY_ENGINE_LMB_PARSER = LambdaMeta{
	Name:            "carly_lmb_parser",
	Runtime:         GO_RUNTIME,
	Timeout:         DEFAULT_TIMEOUT,
	RequiredEnvKeys: []string{""},
}

// Lambda downloader
var CARLY_ENGINE_LMB_DOWNLOADER = LambdaMeta{
	Name:            "carly_lmb_downloader",
	Runtime:         GO_RUNTIME,
	Timeout:         DEFAULT_TIMEOUT,
	RequiredEnvKeys: []string{""},
}
