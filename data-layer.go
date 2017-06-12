package main

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/namsral/flag"

	_ "github.com/lib/pq"

	"github.com/iochti/thing-group-service/models"
)

// DataLayerInterface is an interface abstracting methods to CRUD ThingGroup
type DataLayerInterface interface {
	GetGroupByID(groupID string, group *models.ThingGroup) error
	CreateGroup(group *models.ThingGroup) error
	UpdateGroup(group *models.ThingGroup) error
	DeleteGroup(groupID string) error
}

const THING_GROUP_COLLECTION = "thing_group"

// MgoDL implements DataLayerInterface
type MgoDL struct {
	DBName  string
	Session *mgo.Session
}

var (
	mainSession *mgo.Session
	mainDB      *mgo.Database
	DBName      string
)

// Init inits the DB
func (m *MgoDL) Init() error {
	mHost := flag.String("mhost", "localhost", "MongoDB database host")
	mPort := flag.String("mport", "27017", "MongoDB's port")
	mName := flag.String("mname", "crm", "MongoDB's name")
	flag.Parse()
	mainSession, err := mgo.Dial(fmt.Sprintf("mongodb://%s:%s", *mHost, *mPort))
	if err != nil {
		panic(err)
	}
	mainDB = mainSession.DB(*mName)

	m.Session = mainSession.Copy()

	s := m.Session.Copy()
	defer s.Close()
	return nil
}

// GetGroupByID gets a group by its ID
func (m *MgoDL) GetGroupByID(groupID string, group *models.ThingGroup) error {
	sess := m.Session.Copy()
	if err := sess.DB(DBName).C(THING_GROUP_COLLECTION).FindId(bson.ObjectIdHex(groupID)).One(&group); err != nil {
		return err
	}
	return nil
}

// CreateGroup creates a group and updates its ID
func (m *MgoDL) CreateGroup(group *models.ThingGroup) error {
	sess := m.Session.Copy()
	defer sess.Close()
	if err := sess.DB(DBName).C(THING_GROUP_COLLECTION).Insert(group); err != nil {
		return err
	}
	return nil
}

// UpdateGroup takes a group as parameter and uses its informations to update the DB
func (m *MgoDL) UpdateGroup(group *models.ThingGroup) error {
	sess := m.Session.Copy()
	defer sess.Clone()
	if err := sess.DB(DBName).C(THING_GROUP_COLLECTION).UpdateId(group.ID, group); err != nil {
		return err
	}
	return nil
}

// DeleteGroup deletes a group identified by its ID
func (m *MgoDL) DeleteGroup(groupID string) error {
	sess := m.Session.Copy()
	defer sess.Clone()
	if err := sess.DB(DBName).C(THING_GROUP_COLLECTION).RemoveId(bson.ObjectIdHex(groupID)); err != nil {
		return err
	}
	return nil
}
