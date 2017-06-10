package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/namsral/flag"

	_ "github.com/lib/pq"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	pb "github.com/iochti/thing-group-service/proto"
)

// DataLayerInterface is an interface abstracting methods to CRUD ThingGroup
type DataLayerInterface interface {
	GetGroupByID(groupID int32, group *pb.ThingGroup) error
	CreateGroup(group *pb.ThingGroup) error
	UpdateGroup(group *pb.ThingGroup) error
	DeleteGroup(groupID int32) error
}

// PostgresDL implements the DataLayerInterface for PostgresSQL mappings
type PostgresDL struct {
	DBName string
	Host   string
	Db     *sql.DB
}

const THING_GROUP_TABLE = "thing_group"

// Init the DB host,user,etc...
func (p *PostgresDL) Init() error {
	host := flag.String("pq-host", "localhost", "PostgresSQL database host")
	user := flag.String("pq-user", "postgres", "PostgresSQL user")
	dbName := flag.String("pq-db", "iochti", "PostgresSQL DBName")
	password := flag.String("pq-pwd", "", "PostgresSQL user password")
	flag.Parse()

	// Create a db connection
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", *user, *password, *host, *dbName))
	if err != nil {
		return err
	}
	p.Db = db
	return nil
}

// GetGroupByID gets a group by its ID
func (p *PostgresDL) GetGroupByID(groupID int32, group *pb.ThingGroup) error {
	var timeCreated time.Time
	var timeUpdated time.Time
	err := p.Db.QueryRow(
		"SELECT id, group_id, name, description, created_at, updated_at FROM "+THING_GROUP_TABLE+" WHERE id=$1;",
		groupID).Scan(
		&group.ID,
		&group.Group_ID,
		&group.Name,
		&group.Description,
		&timeCreated,
		&timeUpdated,
	)
	if err != nil {
		return err
	}
	group.CreatedAt = &timestamp.Timestamp{Seconds: timeCreated.Unix()}
	group.UpdatedAt = &timestamp.Timestamp{Seconds: timeUpdated.Unix()}
	return nil
}

// CreateGroup creates a group and updates its ID
func (p *PostgresDL) CreateGroup(group *pb.ThingGroup) error {
	timeNow := time.Now()
	var groupID int32
	err := p.Db.QueryRow("INSERT INTO "+THING_GROUP_TABLE+`(name, group_id, description, created_at, updated_at)
        VALUES($1, $2, $3, $4, $4) RETURNING id;
    `, group.GetName(), group.GetGroup_ID(), group.GetDescription(), timeNow).Scan(&groupID)
	if err != nil {
		return err
	}
	setTime(group, timeNow, timeNow)
	group.ID = groupID
	return nil
}

// UpdateGroup takes a group as parameter and uses its informations to update the DB
func (p *PostgresDL) UpdateGroup(group *pb.ThingGroup) error {
	updateTime := time.Now()
	var createdTime time.Time
	err := p.Db.QueryRow("UPDATE "+THING_GROUP_TABLE+` SET
        name=$1,
		group_id=$2
        description=$3,
        updated_at=$4
        WHERE id=$5 RETURNING created_at;
    `, group.GetName(), group.GetGroup_ID(), group.GetDescription(), updateTime, group.GetID()).Scan(&createdTime)
	if err != nil {
		return err
	}
	setTime(group, createdTime, updateTime)
	return nil
}

// DeleteGroup deletes a group identified by its ID
func (p *PostgresDL) DeleteGroup(groupID int32) error {
	if _, err := p.Db.Query("DELETE FROM "+THING_GROUP_TABLE+" WHERE id=$1;", groupID); err != nil {
		return err
	}
	return nil
}

func setTime(group *pb.ThingGroup, create time.Time, update time.Time) {
	group.CreatedAt = &timestamp.Timestamp{Seconds: create.Unix()}
	group.UpdatedAt = &timestamp.Timestamp{Seconds: update.Unix()}
}
