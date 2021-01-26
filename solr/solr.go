package solr

import (
	"encoding/json"
	"fmt"

	gosolr "github.com/vanng822/go-solr/solr"

	"github.com/quangdangfit/solr-demo/config"
)

// Solr struct
type Solr struct {
	solr *gosolr.SolrInterface
}

// New Solr object
func New() (ISolr, error) {
	conf := config.Conf
	solr, err := gosolr.NewSolrInterface(conf.SolrURL, fmt.Sprintf("solr/%s", conf.SolrCore))
	if err != nil {
		return nil, fmt.Errorf("connect to solr: %s", err)
	}
	solr.SetBasicAuth(conf.SolrUser, conf.SolrPwd)

	return &Solr{solr: solr}, nil
}

func (s *Solr) mapToDocument(data map[string]interface{}) (gosolr.Document, error) {
	bData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var newDoc gosolr.Document
	err = json.Unmarshal(bData, &newDoc)
	if err != nil {
		return nil, err
	}

	return newDoc, nil
}

func (s *Solr) bytesToDocument(data []byte) (gosolr.Document, error) {
	var newDoc gosolr.Document
	err := json.Unmarshal(data, &newDoc)
	if err != nil {
		return nil, err
	}

	return newDoc, nil
}

// Add document to Solr
func (s *Solr) Add(data map[string]interface{}) error {
	docs := make([]gosolr.Document, 0, 1)

	newDoc, err := s.mapToDocument(data)
	if err != nil {
		return err
	}

	docs = append(docs, newDoc)

	res, err := s.solr.Add(docs, 0, nil)
	if err != nil {
		return err
	}

	if !res.Success {
		return fmt.Errorf("solr add: %v", res.Result)
	}

	res, err = s.solr.Commit()
	if err != nil {
		return err
	}
	if !res.Success {
		return fmt.Errorf("solr commit add: %v", res.Result)
	}

	return nil
}

// Update document in Solr
func (s *Solr) Update(data map[string]interface{}) error {
	docs := make([]gosolr.Document, 0, 1)

	newDoc, err := s.mapToDocument(data)
	if err != nil {
		return err
	}

	docs = append(docs, newDoc)

	res, err := s.solr.Update(docs, nil)
	if err != nil {
		return err
	}

	if !res.Success {
		return fmt.Errorf("solr update: %v", res.Result)
	}

	res, err = s.solr.Commit()
	if err != nil {
		return err
	}
	if !res.Success {
		return fmt.Errorf("solr commit update: %v", res.Result)
	}

	return nil
}
