package geo

import (
	"encoding/json"
	"errors"
)

// MarshalJSON enables lines to be encoded as JSON using the encoding/json package.
func (l *Line) MarshalJSON() ([]byte, error) {
	return json.Marshal([2]Point{l.a, l.b})
}

// UnmarshalJSON enables lines to be decoded as JSON using the encoding/json package.
func (l *Line) UnmarshalJSON(data []byte) error {
	var points [][2]float64

	err := json.Unmarshal(data, &points)
	if err != nil {
		return err
	}

	if len(points) > 2 {
		return errors.New("geo: too many points to unmarshal into line")
	}

	if len(points) < 2 {
		return errors.New("geo: not enough points to unmarshal into line")
	}

	l.a = Point{points[0][0], points[0][1]}
	l.b = Point{points[1][0], points[1][1]}

	return nil
}

// // MarshalJSON enables paths to be encoded as JSON using the encoding/json package.
// func (p Path) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(p)
// }

// UnmarshalJSON enables paths to be decoded as JSON using the encoding/json package.
// func (p Path) UnmarshalJSON(data []byte) error {
// 	err := json.Unmarshal(data, &p)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// MarshalJSON enables bounds to be encoded as JSON using the encoding/json package.
func (b Bound) MarshalJSON() ([]byte, error) {
	return json.Marshal([2]Point{b.sw, b.ne})
}

// UnmarshalJSON enables bounds to be decoded as JSON using the encoding/json package.
func (b *Bound) UnmarshalJSON(data []byte) error {
	var points []Point

	err := json.Unmarshal(data, &points)
	if err != nil {
		return err
	}

	if len(points) > 2 {
		return errors.New("geo: too many points to unmarshal into bound")
	}

	if len(points) < 2 {
		return errors.New("geo: not enough points to unmarshal into bound")
	}

	b.sw = points[0]
	b.ne = points[0]
	b.Extend(points[1])

	return nil
}
