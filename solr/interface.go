package solr

// ISolr interface
type ISolr interface {
	Add(data map[string]interface{}) error
	Update(data map[string]interface{}) error
}
