package DirectoryCollector

import (
	"regexp"
	"testing"
)

func TestDirectoryCollectorRegex(t *testing.T) {
	normalizedPath := "/home/username/Documents/dev/KoNekoD/go-deptrac/pkg/src/Core/Ast/Parser/NikicPhpParser/NikicPhpParser/NikicPhpParser.go"

	validatedPattern := "pkg/src/Core/Ast/.*"

	r, err := regexp.Compile(validatedPattern)
	if err != nil {
		t.Error(err)
	}

	match := r.FindStringSubmatch(normalizedPath)

	if len(match) == 0 {
		t.Error("Failed to match")
	}
}
