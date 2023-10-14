package esquery

import "github.com/fatih/structs"

type GeoSorting struct {
	pointField string
	params     geoSortParams
}

type geoSortParams struct {
	Order          Order           `structs:"order,omitempty"`
	Point          *GeoPoint       `structs:"point,omitempty"`
	DistanceType   GeoDistanceType `structs:"distance_type,string,omitempty"`
	Unit           DistanceUnit    `structs:"unit,string,omitempty"`
	IgnoreUnmapped *bool           `structs:"ignore_unmapped,omitempty"`
}

func GeoSort(pointField string) *GeoSorting {
	return &GeoSorting{pointField: pointField}
}

func (s *GeoSorting) DistanceType(distanceType GeoDistanceType) *GeoSorting {
	s.params.DistanceType = distanceType
	return s
}

func (s *GeoSorting) DistanceUnit(unit DistanceUnit) *GeoSorting {
	s.params.Unit = unit
	return s
}

func (s *GeoSorting) IgnoreUnmapped(value bool) *GeoSorting {
	s.params.IgnoreUnmapped = &value
	return s
}

func (s *GeoSorting) Order(order Order) *GeoSorting {
	s.params.Order = order
	return s
}

func (s *GeoSorting) GeoPoint(point *GeoPoint) *GeoSorting {
	s.params.Point = point
	return s
}

func (s *GeoSorting) Map() map[string]interface{} {
	m := structs.Map(s.params)
	m[s.pointField] = m["point"]
	delete(m, "point")
	return map[string]interface{}{
		"_geo_distance": m,
	}
}

type FieldSorting struct {
	field string
	order Order
}

func FieldSort(field string, order Order) *FieldSorting {
	return &FieldSorting{
		field: field,
		order: order,
	}
}

func (s *FieldSorting) Map() map[string]interface{} {
	return map[string]interface{}{
		s.field: map[string]interface{}{
			"order": s.order,
		},
	}
}
