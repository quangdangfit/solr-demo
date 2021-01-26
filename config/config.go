package config

import (
	"os"
)

// Conf global config object
var Conf *Config

func init() {
	// solr
	solrURL := os.Getenv("solr_url")
	solrUser := os.Getenv("solr_user")
	solrPwd := os.Getenv("solr_pwd")
	solrCore := os.Getenv("solr_core")

	Conf = &Config{
		// solr
		SolrURL:  solrURL,
		SolrUser: solrUser,
		SolrPwd:  solrPwd,
		SolrCore: solrCore,
	}
}

// GetConfig :
func GetConfig() *Config {
	return Conf
}

// Config : struct
type Config struct {
	// solr config
	SolrURL  string `json:"solr_url"`
	SolrUser string `json:"solr_user"`
	SolrPwd  string `json:"solr_pwd"`
	SolrCore string `json:"solr_core"`
}
