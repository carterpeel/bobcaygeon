package raop

import (
	"fmt"
	"testing"
)

func TestDAAPParse(t *testing.T) {
	input := []byte{109, 108, 105, 116, 0, 0, 6, 17, 109, 105, 107, 100, 0, 0, 0, 1, 2, 97, 115, 97, 108, 0, 0, 0, 13, 80, 104, 97, 110, 116, 111, 109, 32, 80, 111, 119, 101, 114, 97, 115, 97, 114, 0, 0, 0, 18, 84, 104, 101, 32, 84, 114, 97, 103, 105, 99, 97, 108, 108, 121, 32, 72, 105, 112, 97, 115, 98, 114, 0, 0, 0, 2, 1, 0, 97, 115, 99, 109, 0, 0, 0, 0, 97, 115, 99, 111, 0, 0, 0, 1, 0, 97, 115, 99, 112, 0, 0, 0, 85, 84, 104, 101, 32, 84, 114, 97, 103, 105, 99, 97, 108, 108, 121, 32, 72, 105, 112, 44, 32, 71, 111, 114, 100, 32, 68, 111, 119, 110, 105, 101, 44, 32, 82, 111, 98, 32, 66, 97, 107, 101, 114, 44, 32, 74, 111, 104, 110, 110, 121, 32, 70, 97, 121, 44, 32, 80, 97, 117, 108, 32, 76, 97, 110, 103, 108, 111, 105, 115, 32, 38, 32, 71, 111, 114, 100, 32, 83, 105, 110, 99, 108, 97, 105, 114, 109, 101, 105, 97, 0, 0, 0, 4, 90, 156, 21, 211, 97, 115, 100, 97, 0, 0, 0, 4, 90, 156, 21, 211, 109, 101, 105, 112, 0, 0, 0, 4, 131, 218, 135, 192, 97, 115, 112, 108, 0, 0, 0, 4, 131, 218, 135, 192, 97, 115, 100, 109, 0, 0, 0, 4, 90, 156, 97, 42, 97, 115, 100, 99, 0, 0, 0, 2, 0, 1, 97, 115, 100, 110, 0, 0, 0, 2, 0, 1, 97, 115, 101, 113, 0, 0, 0, 0, 97, 115, 103, 110, 0, 0, 0, 3, 80, 111, 112, 97, 115, 100, 116, 0, 0, 0, 24, 80, 117, 114, 99, 104, 97, 115, 101, 100, 32, 65, 65, 67, 32, 97, 117, 100, 105, 111, 32, 102, 105, 108, 101, 97, 115, 114, 118, 0, 0, 0, 1, 0, 97, 115, 115, 114, 0, 0, 0, 4, 0, 0, 172, 68, 97, 115, 115, 122, 0, 0, 0, 4, 0, 168, 248, 12, 97, 115, 115, 116, 0, 0, 0, 4, 0, 0, 0, 0, 97, 115, 115, 112, 0, 0, 0, 4, 0, 0, 0, 0, 97, 115, 116, 109, 0, 0, 0, 4, 0, 4, 129, 205, 97, 115, 116, 99, 0, 0, 0, 2, 0, 12, 97, 115, 116, 110, 0, 0, 0, 2, 0, 4, 97, 115, 117, 114, 0, 0, 0, 1, 0, 97, 115, 121, 114, 0, 0, 0, 2, 7, 206, 97, 115, 102, 109, 0, 0, 0, 3, 109, 52, 97, 109, 105, 105, 100, 0, 0, 0, 4, 0, 0, 193, 161, 109, 105, 110, 109, 0, 0, 0, 10, 66, 111, 98, 99, 97, 121, 103, 101, 111, 110, 109, 112, 101, 114, 0, 0, 0, 8, 54, 178, 28, 207, 201, 245, 87, 79, 97, 115, 100, 98, 0, 0, 0, 1, 0, 97, 101, 78, 86, 0, 0, 0, 4, 0, 0, 10, 60, 97, 115, 100, 107, 0, 0, 0, 1, 0, 97, 115, 98, 116, 0, 0, 0, 2, 0, 0, 97, 103, 114, 112, 0, 0, 0, 0, 97, 101, 83, 73, 0, 0, 0, 8, 0, 0, 0, 0, 58, 50, 211, 210, 97, 101, 65, 73, 0, 0, 0, 4, 0, 2, 113, 152, 97, 101, 80, 73, 0, 0, 0, 4, 58, 50, 211, 206, 97, 101, 67, 73, 0, 0, 0, 4, 1, 181, 202, 54, 97, 101, 71, 73, 0, 0, 0, 4, 0, 0, 0, 14, 97, 115, 99, 100, 0, 0, 0, 4, 109, 112, 52, 97, 97, 115, 99, 115, 0, 0, 0, 4, 0, 0, 0, 2, 97, 101, 83, 70, 0, 0, 0, 4, 0, 2, 48, 95, 97, 101, 80, 67, 0, 0, 0, 1, 0, 97, 115, 99, 116, 0, 0, 0, 0, 97, 115, 99, 110, 0, 0, 0, 0, 97, 115, 99, 114, 0, 0, 0, 1, 0, 97, 101, 72, 86, 0, 0, 0, 1, 0, 97, 101, 77, 75, 0, 0, 0, 1, 1, 97, 101, 83, 78, 0, 0, 0, 0, 97, 101, 69, 78, 0, 0, 0, 0, 97, 101, 69, 83, 0, 0, 0, 4, 0, 0, 0, 0, 97, 101, 83, 85, 0, 0, 0, 4, 0, 0, 0, 0, 97, 101, 71, 72, 0, 0, 0, 4, 0, 0, 0, 1, 97, 101, 71, 68, 0, 0, 0, 4, 0, 0, 1, 20, 97, 101, 71, 85, 0, 0, 0, 8, 0, 0, 0, 0, 0, 198, 194, 172, 97, 101, 71, 82, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 97, 101, 71, 69, 0, 0, 0, 4, 0, 0, 8, 64, 97, 115, 97, 97, 0, 0, 0, 18, 84, 104, 101, 32, 84, 114, 97, 103, 105, 99, 97, 108, 108, 121, 32, 72, 105, 112, 97, 115, 103, 112, 0, 0, 0, 1, 0, 109, 101, 120, 116, 0, 0, 0, 2, 0, 1, 97, 115, 101, 100, 0, 0, 0, 2, 0, 1, 97, 115, 100, 114, 0, 0, 0, 4, 53, 171, 1, 240, 97, 115, 100, 112, 0, 0, 0, 4, 90, 156, 92, 35, 97, 115, 104, 112, 0, 0, 0, 1, 1, 97, 115, 115, 110, 0, 0, 0, 10, 66, 111, 98, 99, 97, 121, 103, 101, 111, 110, 97, 115, 115, 97, 0, 0, 0, 14, 84, 114, 97, 103, 105, 99, 97, 108, 108, 121, 32, 72, 105, 112, 97, 115, 115, 108, 0, 0, 0, 14, 84, 114, 97, 103, 105, 99, 97, 108, 108, 121, 32, 72, 105, 112, 97, 115, 115, 117, 0, 0, 0, 13, 80, 104, 97, 110, 116, 111, 109, 32, 80, 111, 119, 101, 114, 97, 115, 115, 99, 0, 0, 0, 81, 84, 114, 97, 103, 105, 99, 97, 108, 108, 121, 32, 72, 105, 112, 44, 32, 71, 111, 114, 100, 32, 68, 111, 119, 110, 105, 101, 44, 32, 82, 111, 98, 32, 66, 97, 107, 101, 114, 44, 32, 74, 111, 104, 110, 110, 121, 32, 70, 97, 121, 44, 32, 80, 97, 117, 108, 32, 76, 97, 110, 103, 108, 111, 105, 115, 32, 38, 32, 71, 111, 114, 100, 32, 83, 105, 110, 99, 108, 97, 105, 114, 97, 115, 115, 115, 0, 0, 0, 0, 97, 115, 98, 107, 0, 0, 0, 1, 0, 97, 115, 112, 117, 0, 0, 0, 0, 97, 101, 67, 82, 0, 0, 0, 0, 97, 115, 97, 105, 0, 0, 0, 8, 208, 203, 58, 24, 226, 64, 152, 237, 97, 115, 108, 115, 0, 0, 0, 8, 0, 0, 0, 0, 0, 168, 248, 12, 97, 101, 83, 69, 0, 0, 0, 8, 0, 0, 0, 0, 1, 182, 58, 229, 97, 101, 68, 86, 0, 0, 0, 4, 0, 0, 0, 0, 97, 101, 68, 80, 0, 0, 0, 4, 0, 0, 0, 0, 97, 101, 68, 82, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 97, 101, 78, 68, 0, 0, 0, 8, 0, 0, 0, 0, 10, 81, 194, 42, 97, 101, 75, 49, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 97, 101, 75, 50, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 97, 101, 68, 76, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 97, 101, 70, 65, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 97, 101, 88, 68, 0, 0, 0, 27, 85, 110, 105, 118, 101, 114, 115, 97, 108, 58, 105, 115, 114, 99, 58, 67, 65, 77, 49, 57, 57, 55, 48, 48, 48, 55, 55, 97, 101, 77, 107, 0, 0, 0, 4, 0, 0, 0, 1, 97, 101, 77, 88, 0, 0, 0, 0, 97, 115, 112, 99, 0, 0, 0, 4, 0, 0, 0, 0, 97, 115, 114, 105, 0, 0, 0, 8, 172, 234, 51, 131, 12, 228, 253, 219, 97, 101, 67, 83, 0, 0, 0, 4, 0, 2, 195, 138, 97, 115, 107, 112, 0, 0, 0, 4, 0, 0, 0, 0, 97, 115, 97, 99, 0, 0, 0, 2, 0, 1, 97, 115, 107, 100, 0, 0, 0, 4, 131, 218, 135, 192, 109, 100, 115, 116, 0, 0, 0, 1, 1, 97, 115, 101, 115, 0, 0, 0, 1, 0, 97, 101, 67, 100, 0, 0, 0, 8, 0, 0, 191, 7, 202, 214, 154, 229, 97, 101, 67, 85, 0, 0, 0, 8, 0, 0, 0, 0, 10, 81, 194, 42, 97, 115, 114, 115, 0, 0, 0, 1, 0, 97, 115, 108, 114, 0, 0, 0, 1, 0, 97, 115, 97, 115, 0, 0, 0, 1, 32, 97, 101, 67, 70, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 2, 97, 101, 67, 75, 0, 0, 0, 1, 2, 97, 101, 71, 115, 0, 0, 0, 1, 1, 97, 101, 108, 115, 0, 0, 0, 1, 0, 97, 106, 97, 108, 0, 0, 0, 1, 0, 97, 106, 99, 65, 0, 0, 0, 1, 0, 97, 119, 114, 107, 0, 0, 0, 0, 97, 109, 118, 109, 0, 0, 0, 0, 97, 109, 118, 99, 0, 0, 0, 2, 0, 0, 97, 109, 118, 110, 0, 0, 0, 2, 0, 0, 97, 106, 117, 119, 0, 0, 0, 1, 0}
	parsed := parseDaap(input)
	if len(parsed) != 4 {
		t.Error(fmt.Sprintf("Expected: 4 entries\r\n Got: %d", len(parsed)))
	}
	val, ok := parsed["dmap.itemkind"]
	if !ok {
		t.Error(fmt.Sprintf("Expected to have dmap.itemkind key"))
	}
	if val.(uint8) != 2 {
		t.Error(fmt.Sprintf("Expected: %v\r\n Received: %v\r\n", 2, val))
	}
	val, ok = parsed["daap.songalbum"]
	if !ok {
		t.Error(fmt.Sprintf("Expected to have daap.songalbum key"))
	}
	if val.(string) != "Phantom Power" {
		t.Error(fmt.Sprintf("Expected: %v\r\n Received: %s\r\n", "Phantom Power", val))
	}
	val, ok = parsed["dmap.itemname"]
	if !ok {
		t.Error(fmt.Sprintf("Expected to have dmap.itemname key"))
	}
	if val.(string) != "Bobcaygeon" {
		t.Error(fmt.Sprintf("Expected: %v\r\n Received: %s\r\n", "Bobcaygeon", val))
	}
	val, ok = parsed["daap.songartist"]
	if !ok {
		t.Error(fmt.Sprintf("Expected to have daap.songartist key"))
	}
	if val.(string) != "The Tragically Hip" {
		t.Error(fmt.Sprintf("Expected: %v\r\n Received: %s\r\n", "The Tragically Hip", val))
	}
}
