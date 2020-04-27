package integration

import "testing"

type dsClickhouseSuite struct {
	suite
}

func (d *dsClickhouseSuite) Init(t *testing.T) {
	c := &container{
		Image: "yandex/clickhouse-server",
	}
	d.addContainer(c)
}

func (d *dsClickhouseSuite) Run(t *testing.T) {

}
