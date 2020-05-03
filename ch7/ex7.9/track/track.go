package track

import (
	"time"
)

type comparsion int

const (
	lt comparsion = iota
	eq
	gt
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func Length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

// func printTracks(tracks []*Track) {
// 	const format = "%v\t%v\t%v\t%v\t%v\t\n"
// 	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
// 	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
// 	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
// 	for _, t := range tracks {
// 		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
// 	}
// 	tw.Flush()
// }

type ColumnCmp func(a, b *Track) comparsion

type ByMostRecentlyColumns struct {
	t          []*Track
	columnCmp  []ColumnCmp
	maxColumns int
}

func (c *ByMostRecentlyColumns) Len() int {
	return len(c.t)
}

func (c *ByMostRecentlyColumns) Swap(i, j int) {
	c.t[i], c.t[j] = c.t[j], c.t[i]
}

func (c *ByMostRecentlyColumns) Less(i, j int) bool {
	for _, f := range c.columnCmp {
		cmp := f(c.t[i], c.t[j])
		switch cmp {
		case eq:
			continue
		case lt:
			return true
		case gt:
			return false
		}
	}

	return false
}

func lessString(a, b string) comparsion {
	switch {
	case a == b:
		return eq
	case a > b:
		return gt
	default:
		return lt
	}
}

func (c *ByMostRecentlyColumns) LessTitle(a, b *Track) comparsion {
	return lessString(a.Title, b.Title)
}

func (c *ByMostRecentlyColumns) LessArtist(a, b *Track) comparsion {
	return lessString(a.Artist, b.Artist)
}

func (c *ByMostRecentlyColumns) LessAlbum(a, b *Track) comparsion {
	return lessString(a.Album, b.Album)
}

func (c *ByMostRecentlyColumns) LessYear(a, b *Track) comparsion {
	switch {
	case a.Year == b.Year:
		return eq
	case a.Year > b.Year:
		return gt
	default:
		return lt
	}
}

func (c *ByMostRecentlyColumns) LessLength(a, b *Track) comparsion {
	switch {
	case a.Length == b.Length:
		return eq
	case a.Length > b.Length:
		return gt
	default:
		return lt
	}
}

func (c *ByMostRecentlyColumns) AddCmp(cmp ColumnCmp) {
	c.columnCmp = append([]ColumnCmp{cmp}, c.columnCmp...)
	if len(c.columnCmp) > c.maxColumns {
		c.columnCmp = c.columnCmp[:c.maxColumns]
	}
}

func NewByMostRecentlyColumns(t []*Track, n int) *ByMostRecentlyColumns {
	return &ByMostRecentlyColumns{
		t:          t,
		maxColumns: n,
	}
}
