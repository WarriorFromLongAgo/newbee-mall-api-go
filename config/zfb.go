package config

type AlipayConfig struct {
	AppID        string `mapstructure:"appid" json:"appid" yaml:"appid"`
	AliPublic    string `mapstructure:"aliPublic" json:"aliPublic" yaml:"aliPublic"`
	AliServerURL string `mapstructure:"aliServerUrl" json:"aliServerUrl" yaml:"aliServerUrl"`
	MyPrivateKey string `mapstructure:"myPrivateKey" json:"myPrivateKey" yaml:"myPrivateKey"`
	Charset      string `mapstructure:"charset" json:"charset" yaml:"charset"`
	SignType     string `mapstructure:"signType" json:"signType" yaml:"signType"`
	Profile      string `mapstructure:"profile" json:"profile" yaml:"profile"`
}
