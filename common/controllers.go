package common

import (
	ml "git.vdo.space/berk/model"
	"git.vdo.space/foy/model"
	. "git.vdo.space/foy/proto"
	"github.com/jinzhu/gorm"
	fRegistry "github.com/llcan1120/fast-registry"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
	"os"
	"time"
)

type Controller struct {
	RDB    *gorm.DB
	MDB    *mgo.Database
	User   *ml.User
	Broker broker.Broker
	//Logger *log.MeanLogger
}

func NewMeanController(MDBSession *mgo.Session, MDB *mgo.Database) *Controller {
	mc := &Controller{
		//RDB: RDB,
		MDB: MDB,
	}
	return mc
}

type SingleEntity struct {
	Validation model.NullBool `json:"validation,omitempty"`
	QNToken    string         `json:"QN-Token,omitempty"`
	Expires    *time.Time     `json:"QN-Expires,omitempty"`
}

const (
	CONSUL_ADDRESSES_KEY = "CONSUL_ADDRS"
)

const (
	DREAM_SERVICE_KEY  = "tech-ngs-dream"
	ATHENA_SERVICE_KEY = "tech-ngs-athena"
)

var Cl DreamServicesClient
var Cathena AthenaServicesClient

type TimeWrapper struct {
	client.Client
}

func (l *TimeWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	start := time.Now()

	defer func() {
		end := time.Now()
		end.Sub(start)
	}()

	return l.Client.Call(ctx, req, rsp)
}

func Wrapper(c client.Client) client.Client {
	return &TimeWrapper{c}
}

func InitDreamClient() {
	reg := fRegistry.NewRegistry(
		fRegistry.Addrs(os.Getenv(CONSUL_ADDRESSES_KEY)),
	)
	c := client.NewClient(
		client.Registry(reg),
		client.Wrap(Wrapper),
	)
	Cl = NewDreamServicesClient(DREAM_SERVICE_KEY, c)
}
func InitAthenaClient() {
	reg := fRegistry.NewRegistry(
		fRegistry.Addrs(os.Getenv(CONSUL_ADDRESSES_KEY)),
	)
	c := client.NewClient(
		client.Registry(reg),
		client.Wrap(Wrapper),
	)
	Cathena = NewAthenaServicesClient(ATHENA_SERVICE_KEY, c)
}
