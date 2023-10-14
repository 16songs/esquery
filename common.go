package esquery

// Source represents the "_source" option which is commonly accepted in ES
// queries. Currently, only the "includes" option is supported.
type Source struct {
	includes []string
	excludes []string
}

// Map returns a map representation of the Source object.
func (source Source) Map() map[string]interface{} {
	m := make(map[string]interface{})
	if len(source.includes) > 0 {
		m["includes"] = source.includes
	}
	if len(source.excludes) > 0 {
		m["excludes"] = source.excludes
	}
	return m
}

// Sort represents a list of keys to sort by.
type Sort []map[string]interface{}

// Order is the ordering for a sort key (ascending, descending).
type Order string

const (
	// OrderAsc represents sorting in ascending order.
	OrderAsc Order = "asc"

	// OrderDesc represents sorting in descending order.
	OrderDesc Order = "desc"
)

type DistanceUnit uint8

const (
	_ DistanceUnit = iota
	DistanceUnitMile
	DistanceUnitYard
	DistanceUnitFeet
	DistanceUnitInch
	DistanceUnitKilometer
	DistanceUnitMeter
	DistanceUnitCentimeter
	DistanceUnitMillimeter
	DistanceUnitNauticalMile
)

func (u DistanceUnit) String() string {
	switch u {
	case DistanceUnitMile:
		return "mi"
	case DistanceUnitYard:
		return "yd"
	case DistanceUnitFeet:
		return "ft"
	case DistanceUnitInch:
		return "in"
	case DistanceUnitKilometer:
		return "km"
	case DistanceUnitMeter:
		return "m"
	case DistanceUnitCentimeter:
		return "cm"
	case DistanceUnitMillimeter:
		return "mm"
	case DistanceUnitNauticalMile:
		return "nmi"
	}

	return ""
}

// GeoDistanceType specify how to compute the distance.
type GeoDistanceType uint8

const (
	_ GeoDistanceType = iota
	// GeoDistanceArc is default
	GeoDistanceArc
	// GeoDistancePlane is faster, but inaccurate on long distances and close to the poles
	GeoDistancePlane
)

// String returns a string representation of the RangeRelation value, as
// accepted by ElasticSearch
func (a GeoDistanceType) String() string {
	switch a {
	case GeoDistanceArc:
		return "arc"
	case GeoDistancePlane:
		return "plane"
	default:
		return ""
	}
}

type GeoPoint struct {
	Lat float64 `structs:"lat"`
	Lon float64 `structs:"lon"`
}
