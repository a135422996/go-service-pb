package gorpc

type AssetCheckConfig struct {
	Assets []struct {
		Path string `json:"path"`
		Key  string `json:"key"`
		Hash string `json:"hash"`
	} `json:"assets"`
	Config struct {
		CreatedAt string `json:"createdAt"`
		Id        string `json:"id"`
	} `json:"config"`
}
