package datatype

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Location struct {
	X, Y float64
}

func (loc Location) GormDataType() string {
	return "POINT"
}

func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_PointFromText(?)",
		Vars: []interface{}{fmt.Sprintf("Point(%f %f)", loc.X, loc.Y)},
	}
}

func (loc *Location) Scan(v interface{}) error {
	val, ok := v.([]byte)
	if ok {
		longitude, latitude, err := returnGeoPoint(val)
		if err != nil {
			return err
		}
		loc.X = longitude
		loc.Y = latitude
	}
	return nil
}

func returnGeoPoint(point []byte) (float64, float64, error) {
	hexString := hex.EncodeToString(point)
	lonHex := hexString[18:34]
	longitude, err := hexToFloat64(lonHex)
	if err != nil {
		return 0, 0, err
	}

	latHex := hexString[34:]
	latitude, err := hexToFloat64(latHex)
	if err != nil {
		return 0, 0, err
	}
	return longitude, latitude, nil
}

func hexToFloat64(hexString string) (float64, error) {
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return 0, err
	}
	bits := binary.LittleEndian.Uint64(bytes)
	f := math.Float64frombits(bits)
	return f, nil
}
