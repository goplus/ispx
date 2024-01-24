package github

import (
	"bytes"
	"context"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
)

var (
	Verbose bool
)

func NewClient(token string) *github.Client {
	if token != "" {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		return github.NewClient(tc)
	} else {
		return github.NewClient(nil)
	}
}

type FileSystem struct {
	client *github.Client
	ctx    context.Context
	root   string
	cur    time.Time
}

func IsSupport(root string) (string, bool) {
	if strings.HasPrefix(root, githubPrefix) {
		return root, true
	} else if strings.HasPrefix(root, "github.com/") {
		return "https://" + root, true
	}
	return "", false
}

const (
	githubPrefix = "https://github.com/"
)

// https://github.com/goplus/FlappyCalf
// https://api.github.com/repos/goplus/FlappyCalf/contents
func NewFileSystem(client *github.Client, root string) (*FileSystem, error) {
	return &FileSystem{client: client, ctx: context.Background(), root: root, cur: time.Now()}, nil
}

func parserUrl(dirname string) (owner, repo, path string) {
	if strings.HasPrefix(dirname, "https://github.com/") {
		dirname = dirname[19:]
	}
	list := strings.Split(dirname, "/")
	n := len(list)
	if n == 2 {
		owner = list[0]
		repo = list[1]
	} else if n >= 3 {
		owner = list[0]
		repo = list[1]
		path = strings.Join(list[2:], "/")
	}
	return
}

// type FileSystem interface {
// 	ReadDir(dirname string) ([]fs.DirEntry, error)
// 	ReadFile(filename string) ([]byte, error)
// 	Join(elem ...string) string
// }

func (f *FileSystem) ReadDir(dirname string) ([]fs.DirEntry, error) {
	owner, repo, path := parserUrl(dirname)
	_, dirs, _, err := f.client.Repositories.GetContents(f.ctx, owner, repo, path, nil)
	if Verbose {
		log.Println("ReadDir", dirname, err)
	}
	if err != nil {
		return nil, err
	}
	var list []fs.DirEntry
	for _, v := range dirs {
		list = append(list, &DirEntry{&FileInfo{c: v, t: f.cur}})
	}
	return list, nil
}

func (f *FileSystem) ReadFile(dirname string) ([]byte, error) {
	owner, repo, path := parserUrl(dirname)
	fc, _, _, err := f.client.Repositories.GetContents(f.ctx, owner, repo, path, nil)
	if Verbose {
		log.Println("ReadFile", dirname, err)
	}
	if err != nil {
		return nil, err
	}
	data, err := fc.GetContent()
	return []byte(data), nil
}

func (f *FileSystem) Join(elem ...string) string {
	return strings.Join(elem, "/")
}

func (f *FileSystem) Base(filename string) string {
	return path.Base(filename)
}

func (f *FileSystem) Abs(path string) (string, error) {
	return path, nil
}

func GetHttpFile(path string) ([]byte, error) {
	res, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return body, err
}

/*
type FileInfo interface {
	Name() string       // base name of the file
	Size() int64        // length in bytes for regular files; system-dependent for others
	Mode() FileMode     // file mode bits
	ModTime() time.Time // modification time
	IsDir() bool        // abbreviation for Mode().IsDir()
	Sys() interface{}   // underlying data source (can return nil)
}
*/

type FileInfo struct {
	c *github.RepositoryContent
	t time.Time
}

func (f *FileInfo) Name() string {
	return f.c.GetName()
}
func (f *FileInfo) Size() int64 {
	return int64(f.c.GetSize())
}
func (f *FileInfo) Mode() fs.FileMode {
	if f.IsDir() {
		return 0755
	}
	return 0644
}
func (f *FileInfo) ModTime() time.Time {
	return f.t
}
func (f *FileInfo) IsDir() bool {
	return f.c.GetType() == "dir"
}
func (f *FileInfo) Sys() interface{} {
	return nil
}

// type DirEntry interface {
// 	Name() string
// 	IsDir() bool
// 	Type() FileMode
// 	Info() (FileInfo, error)
// }

type DirEntry struct {
	info *FileInfo
}

func (p *DirEntry) Name() string {
	return p.info.Name()
}

func (p *DirEntry) IsDir() bool {
	return p.info.IsDir()
}

func (p *DirEntry) Type() fs.FileMode {
	return p.info.Mode()
}

func (p *DirEntry) Info() (fs.FileInfo, error) {
	return p.info, nil
}

var _ fs.DirEntry = &DirEntry{}

/*
type Dir interface {
	Open(file string) (io.ReadCloser, error)
	Close() error
}
*/

type Dir struct {
	ctx    context.Context
	client *github.Client
	assert string
	cache  map[string]string
}

func NewDir(client *github.Client, assert string) (*Dir, error) {
	return &Dir{
		ctx:    context.Background(),
		client: client,
		assert: assert,
		cache:  make(map[string]string),
	}, nil
}

func (d *Dir) Open(file string) (io.ReadCloser, error) {
	if data, ok := d.cache[file]; ok {
		return ioutil.NopCloser(bytes.NewReader([]byte(data))), nil
	}
	dirname := d.assert + "/" + file
	owner, repo, path := parserUrl(dirname)
	fc, _, _, err := d.client.Repositories.GetContents(d.ctx, owner, repo, path, nil)
	if Verbose {
		log.Println("Open", dirname, err)
	}
	if err != nil {
		return nil, err
	}
	data, err := fc.GetContent()
	if err != nil {
		return nil, err
	}
	d.cache[file] = data
	return ioutil.NopCloser(bytes.NewReader([]byte(data))), nil
}

func (d *Dir) Close() error {
	return nil
}
