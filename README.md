# Vinyl

Exchange CSV Records with Structs

### Examples

#### Writer

```
type Album struct {
    Name string
    Count uint
    Timestamp time.Time
}

func WriteAlbums() {
    buf := bytes.NewBuffer(nil)
    w := csv.NewWriter(buf)
    headers := vinyl.Labels(Album{}, vinyl.SnakeFormat)
    if err := w.Write(headers); err != nil {
        log.Fatal(err)
    }
    for i := 0; i < 10; i++ {
        album := Album {
            Name: fmt.Sprint("Album", i),
            Count: 2 << i,
            Timestamp: time.Now(),
        }
        record := vinyl.New(album)
        row := record.Values()
        if err := w.Write(row); err != nil {
            log.Fatal(err)
        }
    }
}

```

#### Reader

```
top_record_artists.csv
###
name, platinum, multi_platinum, diamond
"Elvis Presley", 67, 27, 1
"The Beatles", 42, 26, 6 
"George Strait", 33, 13, 0
"Garth Brooks", 30, 16, 7
"Barbra Steisand", 30, 12, 0

type Artist {
    Name            string
    Platinum        string
    MultiPlatinum   string
    Diamond         string
}
func ReadArtists() {
    f, err :=  f, err := os.Open("top_record_artists.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    exhange, err := vinyl.Exchange(Artist{})
    if err != nil {
        log.Fatal(err)
    }
    type awardCount struct {
        Artist string
        Count int
    }
    r := csv.NewReader(bufio.NewReader(f))
    for {
        rec, err := r.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }
        a := exchange.From(rec).(Artist)
        fmt.Println(a)
    }
}
```

### Benchmarks

```
goos: darwin
goarch: amd64
pkg: github.com/infatuation/vinyl
BenchmarkCSVEncoding-8   	  500000	      2694 ns/op
BenchmarkLabels-8        	 2000000	       616 ns/op
BenchmarkValues-8        	 1000000	      1770 ns/op
BenchmarkDict-8          	 1000000	      2122 ns/op
BenchmarkNewRecord-8     	30000000	        58.7 ns/op
PASS
ok  	github.com/infatuation/vinyl	9.019s

```
