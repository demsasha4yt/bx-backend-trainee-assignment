package avitoapi

import "testing"

// TestAvitoAPI returns testing AVITO API
func TestAvitoAPI(t *testing.T) AvitoAPI {
	conf := &Config{
		URL: "http://192.168.99.100:9000",
		Key: "rijfgjgfnidhjiuddcctntijlvngjdddigcgbcced0fa",
	}
	return New(conf)
}
