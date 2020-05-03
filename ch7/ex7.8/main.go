package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
	"bufio"
	"strconv"
	"sort"
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

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type _ColumnCmp func(a, b *Track) comparsion

type _ByMostRecentlyColumns struct {
	t          []*Track
	columnCmp  []_ColumnCmp
	maxColumns int
}

func (c *_ByMostRecentlyColumns) Len() int {
	return len(c.t)
}

func (c *_ByMostRecentlyColumns) Swap(i, j int) {
	c.t[i], c.t[j] = c.t[j], c.t[i]
}

func (c *_ByMostRecentlyColumns) Less(i, j int) bool {
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

func (c *_ByMostRecentlyColumns) lessTitle(a, b *Track) comparsion {
	return lessString(a.Title, b.Title)
}

func (c *_ByMostRecentlyColumns) lessArtist(a, b *Track) comparsion {
	return lessString(a.Artist, b.Artist)
}

func (c *_ByMostRecentlyColumns) lessAlbum(a, b *Track) comparsion {
	return lessString(a.Album, b.Album)
}

func (c *_ByMostRecentlyColumns) lessYear(a, b *Track) comparsion {
	switch {
	case a.Year == b.Year:
		return eq
	case a.Year > b.Year:
		return gt
	default:
		return lt
	}
}

func (c *_ByMostRecentlyColumns) lessLength(a, b *Track) comparsion {
	switch {
	case a.Length == b.Length:
		return eq
	case a.Length > b.Length:
		return gt
	default:
		return lt
	}
}

func (c *_ByMostRecentlyColumns) addCmp(cmp _ColumnCmp) {
	c.columnCmp = append([]_ColumnCmp{cmp}, c.columnCmp...)
	if len(c.columnCmp) > c.maxColumns {
		c.columnCmp = c.columnCmp[:c.maxColumns]
	}
}

func main() {
	var tracks = []*Track {
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}

	bmrc := &_ByMostRecentlyColumns{
		t: tracks,
		maxColumns: 5,
	}

	reader := bufio.NewReader(os.Stdin)


	for {
		printTracks(tracks)

		var idx int
		var err error 
		for {
			fmt.Print("\nSort by: Title[1], Artist[2], Album[3], Year[4], Length[5] exit(0)\n")
			text, _ := reader.ReadString('\n')
			text = text[:len(text)-1]
			idx, err = strconv.Atoi(text)
			if err != nil || idx > 5 || idx < 0 {
				fmt.Printf("unknown input %s, please select again", text)
			} else {
				break
			}
		}
		

		switch idx {
		case 0:
			return
		case 1:
			bmrc.addCmp(bmrc.lessTitle)
		case 2:
			bmrc.addCmp(bmrc.lessArtist)
		case 3:
			bmrc.addCmp(bmrc.lessAlbum)
		case 4:
			bmrc.addCmp(bmrc.lessYear)
		case 5:
			bmrc.addCmp(bmrc.lessLength)
		}
		
		sort.Sort(bmrc)
	}
}
