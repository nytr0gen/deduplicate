package main

import (
    "bufio"
    "flag"
    "fmt"
    "net/url"
    "os"
    "strings"
    "sort"
)

func main() {
    var urls []string

    var hideImages bool
    flag.BoolVar(&hideImages, "hide-images", false, "hide images")

    var shouldSort bool
    flag.BoolVar(&shouldSort, "sort", true, "sort output")

    flag.Parse()

    var file *os.File
    if flag.NArg() > 0 {
        // open file
        var err error
        file, err = os.Open(flag.Arg(0))
        if err != nil {
            panic(err)
        }
    } else {
        // fetch for all urls from stdin
        file = os.Stdin
    }

    sc := bufio.NewScanner(file)
    for sc.Scan() {
        urls = append(urls, sc.Text())
    }

    if err := sc.Err(); err != nil {
        panic(err)
    }

    found := make(map[string]string)
    for _, uri := range urls {
        uri = strings.TrimSpace(uri)
        if uri == "" {
            continue
        }

        u, _ := url.Parse(uri)
        if u == nil {
            continue
        }

        // considering only images without params
        if hideImages && len(u.RawQuery) == 0 {
            pos := strings.LastIndexByte(u.Path, '.')
            if pos != -1 {
                ext := strings.ToLower(u.Path[pos+1:])
                if extIsImage(ext) {
                    continue
                }
            }
        }

        // ignore scheme, port, query values, auth info and hash
        key := u.Host + u.Path
        if len(u.RawQuery) != 0 {
            key += "?"
            for k, _ := range u.Query() {
                key += k + "&"
            }
        }

        if val, ok := found[key]; ok {
            // prefer https urls
            if u.Scheme == "https" && strings.HasPrefix(val, "http:") {
                found[key] = uri
            }
        } else {
            found[key] = uri
        }
    }

    if !shouldSort {
        for _, uri := range found {
            fmt.Println(uri)
        }
    } else {
        keys := make([]string, len(found))
        i := 0
        for k, _ := range(found) {
            keys[i] = k
            i++
        }

        sort.Strings(keys)
        for _, k := range keys {
            fmt.Println(found[k])
        }
    }
}

func extIsImage(ext string) bool {
    return (ext == "png" ||
        ext == "gif" ||
        ext == "ico" ||
        ext == "jpg" ||
        ext == "jpeg")
}
