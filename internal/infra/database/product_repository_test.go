package database

import (
	"context"
	"database/sql"
	"testing"

	"github.com/raulsilva-tech/ProductAPI/internal/entity"
	"github.com/stretchr/testify/suite"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type ProductRepositoryTestSuite struct {
	Db *sql.DB
	suite.Suite
}

func (suite *ProductRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	_, err = db.Exec(`create table product_types(
		id varchar(40) not null ,
		name varchar(100) NOT NULL,
		description varchar(200) NOT NULL,
		created_at timestamp,
		primary key (id)
	);`)
	suite.NoError(err)
	_, err = db.Exec(`create table products(
		id varchar(40) NOT NULL, 
		name varchar(100) NOT NULL,
		description varchar(200) NOT NULL,
		product_type_id varchar(40),
		created_at timestamp,
		primary key (id),
		foreign key (product_type_id) references product_types(id)
	);`)
	suite.NoError(err)
	suite.Db = db
}

func (suite *ProductRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(ProductRepositoryTestSuite))
}

func (suite *ProductRepositoryTestSuite) TestSave() {

	pt, err := entity.NewProductType("type 1", "type 1 desc")
	suite.NoError(err)

	p, err := entity.NewProduct("test 1", "teste 1 desc", *pt)
	suite.NoError(err)

	repo := NewProductRepository(suite.Db)
	record, err := repo.Save(context.Background(), p)
	suite.NoError(err)
	suite.NotNil(record)
	suite.Equal(p.ID, record.ID)
	suite.Equal(p.Name, record.Name)

}

func (suite *ProductRepositoryTestSuite) TestDelete() {

	pt, err := entity.NewProductType("type 1", "type 1 desc")
	suite.NoError(err)

	p, err := entity.NewProduct("test 1", "teste 1 desc", *pt)
	suite.NoError(err)

	repo := NewProductRepository(suite.Db)
	record, err := repo.Save(context.Background(), p)
	suite.NoError(err)
	suite.NotNil(record)

	err = repo.Delete(context.Background(), record.ID.String())
	suite.NoError(err)

}
func (suite *ProductRepositoryTestSuite) TestUpdate() {

	pt, err := entity.NewProductType("type 1", "type 1 desc")
	suite.NoError(err)

	p, err := entity.NewProduct("test 1", "teste 1 desc", *pt)
	suite.NoError(err)

	repo := NewProductRepository(suite.Db)
	record, err := repo.Save(context.Background(), p)
	suite.NoError(err)
	suite.NotNil(record)

	p.Description = "New Description"
	err = repo.Update(context.Background(), p)
	suite.NoError(err)

	found, err := repo.GetById(context.Background(), p.ID.String())
	suite.NoError(err)
	suite.NotNil(found)
	suite.Equal(found.Description, "New Description")
}

func (suite *ProductRepositoryTestSuite) TestGetById() {

	pt, err := entity.NewProductType("type 1", "type 1 desc")
	suite.NoError(err)
	p, err := entity.NewProduct("test 1", "teste 1 desc", *pt)
	suite.NoError(err)

	repo := NewProductRepository(suite.Db)
	record, err := repo.Save(context.Background(), p)
	suite.NoError(err)
	suite.NotNil(record)

	found, err := repo.GetById(context.Background(), p.ID.String())
	suite.NoError(err)
	suite.NotNil(found)
}

func (suite *ProductRepositoryTestSuite) TestList() {

	pt, err := entity.NewProductType("type 1", "type 1 desc")
	suite.NoError(err)
	p, err := entity.NewProduct("test 1", "teste 1 desc", *pt)
	suite.NoError(err)

	repo := NewProductRepository(suite.Db)
	record, err := repo.Save(context.Background(), p)
	suite.NoError(err)
	suite.NotNil(record)

	foundList, err := repo.List(context.Background())
	suite.NoError(err)
	suite.NotNil(foundList)
	suite.Greater(len(foundList), 0)
}
