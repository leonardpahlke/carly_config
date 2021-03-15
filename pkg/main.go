package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func GetResourceName(name string) string {
	return fmt.Sprintf("%s-%s-%s", ProjectName, DeploymentEnv, name)
}

func GetTags(resourceName string) pulumi.StringMap {
	return pulumi.StringMap{
		"STAGE":      pulumi.String(DeploymentEnv),
		"RESOURCE":   pulumi.String(resourceName),
		"CREATED_BY": pulumi.String("Pulumi"),
		"PROJECT":    pulumi.String(ProjectName),
	}
}

func ListContains(element string, list []string) bool {
	for _, elem := range list {
		if element == elem {
			return true
		}
	}
	return false
}

func GetBucketFileName(bucketName string, newspaper string, filename string, fileEnding string) string {
	return fmt.Sprintf("%s/%s/%s.%s", bucketName, newspaper, filename, fileEnding)
}

func CheckEnvNotEmpty(key string) (string, bool) {
	envString, isEmpty := CheckEnv(key, "")
	if isEmpty {
		LogWarning("CheckEnvNotEmpty", fmt.Sprintf("environment variable %s is empty", key))
	}
	return envString, isEmpty
}

func CheckEnv(key string, expectedVal string) (string, bool) {
	env := os.Getenv(key)
	return env, env == expectedVal
}


func CreateKeyValuePairs(m map[string]string) string {
	b := new(bytes.Buffer)
	valEscapeString := ""
	_, _ = fmt.Fprint(b, "{\n  ")
	for key, value := range m {
		if strings.ContainsAny(value, "{}") {
			valEscapeString = ""
		} else {
			valEscapeString = "\""
		}
		_, _ = fmt.Fprintf(b, "  \"%s\": %s%s%s\n  ", key, valEscapeString, value, valEscapeString)
	}
	_, _ = fmt.Fprint(b, "}")
	return b.String()
}

func MarshalStruct(structIn interface{}) []byte {
	b, err := json.Marshal(structIn)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return b
}

func GetBucketKeyForAnalyticsBucket(newspaper string, articleReference string, fileName string, fileEnding string) string {
	return fmt.Sprintf("%s/%s/%s.%s",
		newspaper,
		articleReference,
		fileName,
		fileEnding)
}

func GetBucketUriForKey(key string) string {
	return fmt.Sprintf("s3://%s", key)
}

// clean string removes
func TrimStringAry(inputStrings []string) []string {
	var cleanedString []string
	for _, str := range inputStrings {
		str = strings.TrimSpace(str)
		if str != "" {
			cleanedString = append(cleanedString, str)
		}
	}
	LogInfo("pkg.TrimStringAry", fmt.Sprintf("cleanedString.. %v", cleanedString))
	return cleanedString
}
