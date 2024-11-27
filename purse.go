// Package purse provides utility functions for string manipulation and slice handling.
package purse

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// MakeLines splits a string into lines.
func MakeLines(s string) []string {
	return strings.Split(s, "\n")
}

// JoinLines joins an array of strings into a single string with newlines.
func JoinLines(lines []string) string {
	return strings.Join(lines, "\n")
}

// ReplaceLastSubStr replaces the last occurrence of a substring in a string.
func ReplaceLastSubStr(s, old, new string) string {
	pos := strings.LastIndex(s, old)
	if pos == -1 {
		return s
	}
	return s[:pos] + new + s[pos+len(old):]
}

// GetFirstLine returns the first line of a string.
func GetFirstLine(s string) string {
	lines := MakeLines(s)
	if len(lines) == 0 {
		return s
	}
	return lines[0]
}

// GetLastLine returns the last line of a string.
func GetLastLine(s string) string {
	lines := MakeLines(s)
	if len(lines) == 0 {
		return s
	}
	return lines[len(lines)-1]
}

// RemoveAllSubStr removes all specified substrings from a string.
func RemoveAllSubStr(s string, subs ...string) string {
	for _, sub := range subs {
		s = strings.ReplaceAll(s, sub, "")
	}
	return s
}

// CountLeadingSpaces counts the number of leading spaces in a string.
func CountLeadingSpaces(line string) int {
	count := 0
	for _, char := range line {
		if char != ' ' {
			break
		}
		count++
	}
	return count
}

// PrefixLines adds a prefix to each line of a string.
func PrefixLines(str, prefix string) string {
	lines := strings.Split(str, "\n")
	for i, line := range lines {
		lines[i] = prefix + line
	}
	return strings.Join(lines, "\n")
}

// FlattenLines removes leading spaces and tabs from each line of a slice.
func FlattenLines(lines []string) []string {
	for i, line := range lines {
		lines[i] = strings.TrimLeft(line, " \t")
	}
	return lines
}

// Flatten removes leading spaces and tabs from all lines of a string.
func Flatten(str string) string {
	lines := MakeLines(str)
	flat := FlattenLines(lines)
	return strings.Join(flat, "")
}

// TrimLeadingSpaces removes leading spaces from all lines of a string.
func TrimLeadingSpaces(str string) string {
	lines := strings.Split(str, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimLeft(line, " ")
	}
	return strings.Join(lines, "\n")
}

// SliceContains checks if a slice contains a specific item.
func SliceContains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// BackTick returns a backtick character.
func BackTick() string {
	return "`"
}

// ReplaceFirstLine replaces the first line of a string with a new line.
func ReplaceFirstLine(input, newLine string) string {
	lines := strings.Split(input, "\n")
	if len(lines) > 0 {
		lines[0] = newLine
	}
	return strings.Join(lines, "\n")
}

// ReplaceLastLine replaces the last line of a string with a new line.
func ReplaceLastLine(input, newLine string) string {
	lines := strings.Split(input, "\n")
	if len(lines) > 0 {
		lines[len(lines)-1] = newLine
	}
	return strings.Join(lines, "\n")
}

// Squeeze removes all spaces from a string.
func Squeeze(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

// ScanBetweenSubStrs extracts substrings between specified delimiters.
func ScanBetweenSubStrs(s, start, end string) []string {
	var out []string
	inSearch := false
	searchStr := ""
	i := 0
	for i < len(s) {
		if !inSearch && i+len(start) <= len(s) && s[i:i+len(start)] == start {
			inSearch = true
			searchStr = start
			i += len(start)
			continue
		}
		if inSearch {
			if i+len(end) <= len(s) && s[i:i+len(end)] == end {
				searchStr += end
				out = append(out, searchStr)
				searchStr = ""
				inSearch = false
				i += len(end)
				continue
			}
			searchStr += string(s[i])
		}
		i++
	}
	return out
}

// RemoveFirstLine removes the first line from a string.
func RemoveFirstLine(input string) string {
	index := strings.Index(input, "\n")
	if index == -1 {
		return ""
	}
	return input[index+1:]
}

// RemoveTrailingEmptyLines removes empty lines from the end of a string.
func RemoveTrailingEmptyLines(input string) string {
	lines := strings.Split(input, "\n")
	for len(lines) > 0 && strings.TrimSpace(lines[len(lines)-1]) == "" {
		lines = lines[:len(lines)-1]
	}
	return strings.Join(lines, "\n")
}

// RemoveEmptyLines removes all empty lines from a string.
func RemoveEmptyLines(input string) string {
	lines := strings.Split(input, "\n")
	var cleanedLines []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			cleanedLines = append(cleanedLines, line)
		}
	}
	return strings.Join(cleanedLines, "\n")
}

// RemoveDuplicatesInSlice removes duplicate items from a slice.
func RemoveDuplicatesInSlice(strSlice []string) []string {
	unique := make(map[string]bool)
	var result []string
	for _, item := range strSlice {
		if _, found := unique[item]; !found {
			unique[item] = true
			result = append(result, item)
		}
	}
	return result
}

// WrapStr wraps a string with a prefix and a suffix.
func WrapStr(s, prefix, suffix string) string {
	return prefix + s + suffix
}

// MatchLeadingSpaces matches the leading spaces of one string to another.
func MatchLeadingSpaces(str1, str2 string) string {
	leadingSpaces := len(str2) - len(strings.TrimLeft(str2, " "))
	padding := strings.Repeat(" ", leadingSpaces)
	return padding + str1
}

// SnipStrAtIndex truncates a string at a given index.
func SnipStrAtIndex(s string, x int) string {
	if x > len(s) {
		x = len(s)
	}
	return s[:x]
}

// TargetSearch finds a substring between two specified strings.
func TargetSearch(input, primarySearch, secondarySearch string) (string, bool) {
	startIndex := strings.Index(input, primarySearch)
	if startIndex == -1 {
		return "", false
	}
	substring := input[startIndex:]
	endIndex := strings.Index(substring, secondarySearch)
	if endIndex == -1 {
		return "", false
	}
	finalEndIndex := startIndex + endIndex + len(secondarySearch)
	return input[startIndex:finalEndIndex], true
}

// SplitWithTargetInclusion splits a string and includes the target as separate items.
func SplitWithTargetInclusion(str, target string) []string {
	var parts []string
	start := 0
	for {
		index := strings.Index(str[start:], target)
		if index == -1 {
			parts = append(parts, str[start:])
			break
		}
		index += start
		parts = append(parts, str[start:index])
		parts = append(parts, target)
		start = index + len(target)
	}
	return parts
}

// PrefixSliceItems prefixes each item in a slice with a string.
func PrefixSliceItems(items []string, prefix string) string {
	var prefixedItems []string
	for _, item := range items {
		prefixedItems = append(prefixedItems, prefix+item)
	}
	return strings.Join(prefixedItems, "")
}

// ReverseSlice reverses the order of elements in a slice.
func ReverseSlice[T any](slice []T) []T {
	n := len(slice)
	reversed := make([]T, n)
	for i, v := range slice {
		reversed[n-1-i] = v
	}
	return reversed
}

// RandStr generates a random string of specified length.
func RandStr(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// MustEqualOneOf checks if a string matches any of the provided options.
func MustEqualOneOf(str string, options ...string) bool {
	for _, option := range options {
		if str == option {
			return true
		}
	}
	return false
}

// ReplaceFirstInstanceOf replaces the first occurrence of `old` with `new` in `s`.
func ReplaceFirstInstanceOf(s, old, new string) string {
	index := strings.Index(s, old)
	if index == -1 {
		return s // Return the original string if the substring is not found
	}
	return s[:index] + new + s[index+len(old):]
}

// ReplaceLastInstanceOf replaces the last occurrence of `old` with `new` in `s`.
func ReplaceLastInstanceOf(s, old, new string) string {
	index := strings.LastIndex(s, old)
	if index == -1 {
		return s // Return the original string if the substring is not found
	}
	return s[:index] + new + s[index+len(old):]
}

// Split a string by " " spaces and work on each chunck
func WorkOnStrChunks(input string, processFunc func(string) error) error {
	// Split the input string by spaces
	chunks := strings.Fields(input)

	// Iterate over each chunk
	for _, chunk := range chunks {
		// Apply the provided function to the chunk
		if err := processFunc(chunk); err != nil {
			return fmt.Errorf("error processing chunk %q: %w", chunk, err)
		}
	}

	return nil
}

func KebabToCamelCase(input string) string {
	parts := strings.Split(input, "-")
	if len(parts) == 0 {
		return ""
	}
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(string(parts[i][0])) + strings.ToLower(parts[i][1:])
		}
	}
	return strings.Join(parts, "")
}

func StrIsFilePath(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsNotExist(err)
}
