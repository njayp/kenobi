package exercises

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

// Exercise: File I/O
// Task: Work with file reading, writing, and manipulation

// WriteFile: Write content to a file
func WriteFile(filename, content string) error {
	// TODO: Write content to file using os.WriteFile
	return nil
}

// ReadFile: Read entire file content
func ReadFile(filename string) (string, error) {
	// TODO: Read file content using os.ReadFile and return as string
	return "", nil
}

// AppendToFile: Append content to existing file
func AppendToFile(filename, content string) error {
	// TODO: Open file in append mode and write content
	return nil
}

// ReadLines: Read file line by line
func ReadLines(filename string) ([]string, error) {
	// TODO: Read file and return slice of lines
	// Hint: Read file, split by newlines, handle different line endings
	return nil, nil
}

// WriteLines: Write slice of strings as lines to file
func WriteLines(filename string, lines []string) error {
	// TODO: Write each string as a line to file
	return nil
}

// CopyFile: Copy content from source to destination
func CopyFile(src, dst string) error {
	// TODO: Copy file content from src to dst
	return nil
}

// FileExists: Check if file exists
func FileExists(filename string) bool {
	// TODO: Check if file exists using os.Stat
	return false
}

// GetFileSize: Get file size in bytes
func GetFileSize(filename string) (int64, error) {
	// TODO: Get file size using os.Stat
	return 0, nil
}

// CreateTempFile: Create temporary file with content
func CreateTempFile(content string) (string, error) {
	// TODO: Create temporary file and write content
	// Return the filename
	return "", nil
}

// ReadWithBuffer: Read file using buffered I/O
func ReadWithBuffer(filename string) (string, error) {
	// TODO: Read file using buffered reader
	return "", nil
}

// Test helper function to create test files
func createTestFile(t *testing.T, filename, content string) {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
}

// Test helper function to clean up test files
func removeTestFile(t *testing.T, filename string) {
	err := os.Remove(filename)
	if err != nil {
		t.Logf("Warning: Failed to remove test file %s: %v", filename, err)
	}
}

func TestWriteFile(t *testing.T) {
	filename := "test_write.txt"
	content := "Hello, World!"
	
	defer removeTestFile(t, filename)
	
	err := WriteFile(filename, content)
	if err != nil {
		t.Errorf("WriteFile returned error: %v", err)
		return
	}
	
	// Verify file was written correctly
	data, err := os.ReadFile(filename)
	if err != nil {
		t.Errorf("Failed to read test file: %v", err)
		return
	}
	
	if string(data) != content {
		t.Errorf("File content = %q, want %q", string(data), content)
	}
}

func TestReadFile(t *testing.T) {
	filename := "test_read.txt"
	content := "Test content for reading"
	
	createTestFile(t, filename, content)
	defer removeTestFile(t, filename)
	
	result, err := ReadFile(filename)
	if err != nil {
		t.Errorf("ReadFile returned error: %v", err)
		return
	}
	
	if result != content {
		t.Errorf("ReadFile = %q, want %q", result, content)
	}
}

func TestAppendToFile(t *testing.T) {
	filename := "test_append.txt"
	initialContent := "Initial content"
	appendContent := "\nAppended content"
	
	createTestFile(t, filename, initialContent)
	defer removeTestFile(t, filename)
	
	err := AppendToFile(filename, appendContent)
	if err != nil {
		t.Errorf("AppendToFile returned error: %v", err)
		return
	}
	
	// Verify content was appended
	data, err := os.ReadFile(filename)
	if err != nil {
		t.Errorf("Failed to read test file: %v", err)
		return
	}
	
	expected := initialContent + appendContent
	if string(data) != expected {
		t.Errorf("File content = %q, want %q", string(data), expected)
	}
}

func TestReadLines(t *testing.T) {
	filename := "test_lines.txt"
	content := "Line 1\nLine 2\nLine 3"
	
	createTestFile(t, filename, content)
	defer removeTestFile(t, filename)
	
	lines, err := ReadLines(filename)
	if err != nil {
		t.Errorf("ReadLines returned error: %v", err)
		return
	}
	
	expected := []string{"Line 1", "Line 2", "Line 3"}
	if len(lines) != len(expected) {
		t.Errorf("Got %d lines, want %d", len(lines), len(expected))
		return
	}
	
	for i, expectedLine := range expected {
		if lines[i] != expectedLine {
			t.Errorf("Line %d: got %q, want %q", i, lines[i], expectedLine)
		}
	}
}

func TestWriteLines(t *testing.T) {
	filename := "test_write_lines.txt"
	lines := []string{"First line", "Second line", "Third line"}
	
	defer removeTestFile(t, filename)
	
	err := WriteLines(filename, lines)
	if err != nil {
		t.Errorf("WriteLines returned error: %v", err)
		return
	}
	
	// Verify lines were written correctly
	data, err := os.ReadFile(filename)
	if err != nil {
		t.Errorf("Failed to read test file: %v", err)
		return
	}
	
	expected := strings.Join(lines, "\n")
	if string(data) != expected {
		t.Errorf("File content = %q, want %q", string(data), expected)
	}
}

func TestCopyFile(t *testing.T) {
	srcFile := "test_src.txt"
	dstFile := "test_dst.txt"
	content := "Content to copy"
	
	createTestFile(t, srcFile, content)
	defer removeTestFile(t, srcFile)
	defer removeTestFile(t, dstFile)
	
	err := CopyFile(srcFile, dstFile)
	if err != nil {
		t.Errorf("CopyFile returned error: %v", err)
		return
	}
	
	// Verify destination file has same content
	data, err := os.ReadFile(dstFile)
	if err != nil {
		t.Errorf("Failed to read destination file: %v", err)
		return
	}
	
	if string(data) != content {
		t.Errorf("Destination file content = %q, want %q", string(data), content)
	}
}

func TestFileExists(t *testing.T) {
	filename := "test_exists.txt"
	
	// Test non-existent file
	if FileExists(filename) {
		t.Error("FileExists should return false for non-existent file")
	}
	
	// Create file and test again
	createTestFile(t, filename, "test")
	defer removeTestFile(t, filename)
	
	if !FileExists(filename) {
		t.Error("FileExists should return true for existing file")
	}
}

func TestGetFileSize(t *testing.T) {
	filename := "test_size.txt"
	content := "Hello, World!" // 13 bytes
	
	createTestFile(t, filename, content)
	defer removeTestFile(t, filename)
	
	size, err := GetFileSize(filename)
	if err != nil {
		t.Errorf("GetFileSize returned error: %v", err)
		return
	}
	
	expectedSize := int64(len(content))
	if size != expectedSize {
		t.Errorf("File size = %d, want %d", size, expectedSize)
	}
}

func TestCreateTempFile(t *testing.T) {
	content := "Temporary content"
	
	filename, err := CreateTempFile(content)
	if err != nil {
		t.Errorf("CreateTempFile returned error: %v", err)
		return
	}
	
	defer os.Remove(filename) // Clean up
	
	// Verify temp file was created with correct content
	data, err := os.ReadFile(filename)
	if err != nil {
		t.Errorf("Failed to read temp file: %v", err)
		return
	}
	
	if string(data) != content {
		t.Errorf("Temp file content = %q, want %q", string(data), content)
	}
}

func TestReadWithBuffer(t *testing.T) {
	filename := "test_buffer.txt"
	content := "Buffered reading test content"
	
	createTestFile(t, filename, content)
	defer removeTestFile(t, filename)
	
	result, err := ReadWithBuffer(filename)
	if err != nil {
		t.Errorf("ReadWithBuffer returned error: %v", err)
		return
	}
	
	if result != content {
		t.Errorf("ReadWithBuffer = %q, want %q", result, content)
	}
}

// Advanced file operations
func TestProcessLargeFile(t *testing.T) {
	// TODO: Students should implement processing large files chunk by chunk
}

func ProcessLargeFile(filename string, chunkSize int, processor func([]byte) error) error {
	// TODO: Read file in chunks and process each chunk
	// This teaches streaming file processing for large files
	return nil
}

func TestProcessLargeFile(t *testing.T) {
	filename := "test_large.txt"
	content := strings.Repeat("line\n", 1000) // 5000 bytes
	
	createTestFile(t, filename, content)
	defer removeTestFile(t, filename)
	
	var processedBytes int
	processor := func(chunk []byte) error {
		processedBytes += len(chunk)
		return nil
	}
	
	err := ProcessLargeFile(filename, 1024, processor)
	if err != nil {
		t.Errorf("ProcessLargeFile returned error: %v", err)
		return
	}
	
	if processedBytes != len(content) {
		t.Errorf("Processed %d bytes, want %d", processedBytes, len(content))
	}
}
