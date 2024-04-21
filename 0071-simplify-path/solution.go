package _071_simplify_path

import "strings"

func SimplifyPath(path string) string {
	arr := strings.Split(path, "/")
	var stack []string
	for _, element := range arr {
		if element == "." || element == "" {
			continue
		}
		if element == ".." {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, element)
		}
	}

	var builder strings.Builder
	for _, element := range stack {
		builder.WriteString("/")
		builder.WriteString(element)
	}
	if builder.Len() == 0 {
		return "/"
	} else {
		return builder.String()
	}
}
