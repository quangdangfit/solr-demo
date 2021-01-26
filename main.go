package main

import (
	"encoding/json"
	"fmt"
	"time"

	gosolr "github.com/vanng822/go-solr/solr"

	"github.com/quangdangfit/solr-demo/config"
)

type Order struct {
	ID        string `json:"id"` // solr index
	OrderID   uint   `json:"order_id_i"`
	CreatedAt int64  `json:"created_at_i"`
	UpdatedAt int64  `json:"updated_at_i"`
	DeletedAt int64  `json:"deleted_at_i"`

	UserID       uint   `json:"user.id_i"`
	UserFullName string `json:"user.full_name_s"`
	UserUserName string `json:"user.username_s"`
	UserEmail    string `json:"user.email_s"`
	UserEnable   bool   `json:"user.enable_b"`

	OrderLine []OrderLine `json:"order_lines"`
}

type OrderLine struct {
	ID   string `json:"id"`
	Name string `json:"name_s"`
}

func add() {
	conf := config.GetConfig()

	//container := app.BuildContainer()
	//router.InitGinEngine(container)

	solrInterface, err := gosolr.NewSolrInterface(conf.SolrURL, fmt.Sprintf("solr/user_test1"))
	solrInterface.SetBasicAuth(conf.SolrUser, conf.SolrPwd)

	orderLine := OrderLine{
		ID:   "64298759-e35d-11e9-8801-02d4bc037",
		Name: "Order Line",
	}

	orderLine2 := OrderLine{
		ID:   "64298759-e35d-11e9-8801-02d4bc038",
		Name: "Order Line 2",
	}

	var data map[string]interface{}
	b, err := json.Marshal(&orderLine)

	json.Unmarshal(b, &data)

	order := &Order{
		ID:        "64298759-e35d-11e9-8801-02d4bc037090",
		OrderID:   123,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		DeletedAt: time.Now().Unix(),

		UserID:       1,
		UserUserName: "quang.dang",
		UserFullName: "Quang",
		UserEmail:    "quang.dang@tokoin.io",
		UserEnable:   true,

		OrderLine: []OrderLine{orderLine, orderLine2},

		//OrderLine: OrderLine{
		//	ID: "64298759-e35d-11e9-8801-02d4bc037",
		//	Name: "Order Line",
		//},
	}

	bytes, err := json.Marshal(&order)
	if err != nil {
		fmt.Println(err)
	}

	var doc gosolr.Document
	err = json.Unmarshal(bytes, &doc)
	if err != nil {
		fmt.Println(err)
	}

	var docs = make([]gosolr.Document, 0, 1)
	docs = append(docs, doc)

	res, err := solrInterface.Add(docs, 0, nil)
	if err != nil {
		fmt.Println(err)
	}

	solrInterface.Commit()

	fmt.Println(res)
}

func query() {
	conf := config.GetConfig()

	solrInterface, _ := gosolr.NewSolrInterface(conf.SolrURL, fmt.Sprintf("solr/user_test"))
	solrInterface.SetBasicAuth(conf.SolrUser, conf.SolrPwd)

	query := gosolr.NewQuery()
	query.Q("*:*")
	q := query.String()
	print(q)

	s := solrInterface.Search(query)
	r, _ := s.Result(nil)
	fmt.Println(r.Results.Docs)

	fmt.Println(r.Results.Docs)
}

func main() {
	//add()
	query()
}
