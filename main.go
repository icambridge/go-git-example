package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"

	"github.com/icambridge/cartel"
)

func main() {
	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)
	fmt.Printf("1. Allocated and not free - %v\n", m.Alloc)
	fmt.Printf("2. Allocated and including free - %v\n", m.TotalAlloc)
	p := cartel.NewPool(5)
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/symfony/symfony", "11")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/cartel", "1")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/phony", "2")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/kata", "3")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/version.nim", "4")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/genkins", "5")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/dotfiles", "6")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/hoard", "7")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/gomposer", "8")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/config", "9")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/diff", "10")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/cartel", "1")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/phony", "2")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/kata", "3")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/version.nim", "4")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/genkins", "5")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/dotfiles", "6")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/hoard", "7")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/gomposer", "8")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/config", "9")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/diff", "10")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/cartel", "1")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/phony", "2")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/kata", "3")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/version.nim", "4")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/genkins", "5")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/dotfiles", "6")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/hoard", "7")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/gomposer", "8")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/config", "9")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/diff", "10")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/cartel", "1")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/phony", "2")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/kata", "3")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/version.nim", "4")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/genkins", "5")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/dotfiles", "6")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/hoard", "7")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/gomposer", "8")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/config", "9")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/diff", "10")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/cartel", "1")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/phony", "2")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/kata", "3")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/version.nim", "4")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/genkins", "5")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/dotfiles", "6")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/hoard", "7")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/gomposer", "8")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/config", "9")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/diff", "10")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/cartel", "1")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/phony", "2")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/kata", "3")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/version.nim", "4")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/genkins", "5")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/dotfiles", "6")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/hoard", "7")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/gomposer", "8")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/config", "9")})
	p.Do(ProcessTask{repository: NewGitRepository("https://github.com/icambridge/diff", "10")})
	p.End()
	m2 := runtime.MemStats{}
	runtime.ReadMemStats(&m2)
	fmt.Printf("1. Allocated and not free - %v\n", m2.Alloc)
	fmt.Printf("2. Allocated and including free - %v\n", m2.TotalAlloc)
	fmt.Printf("3. Allocated to the heap - %v\n", m2.HeapAlloc)
	fmt.Printf("4. Frees - %v\n", m2.Frees)
	runtime.GC()

}

type Commit struct {
	Hash           string
	AuthorName     string
	AuthorEmail    string
	CommitterName  string
	CommitterEmail string
	Message        string
	CommitedAt     time.Time
}

func convertCommit(commit object.Commit) Commit {

	var authorName, authorEmail, comitterName, commiterEmail string
	authorName = commit.Author.Name

	authorEmail = commit.Author.Email

	comitterName = commit.Committer.Name
	commiterEmail = commit.Committer.Email

	newCommit := Commit{
		Hash:           commit.Hash.String(),
		AuthorName:     authorName,
		AuthorEmail:    authorEmail,
		CommitterName:  comitterName,
		CommitterEmail: commiterEmail,
		Message:        commit.Message,
		CommitedAt:     commit.Committer.When,
	}

	return newCommit
}

type File interface {
	Name() string
	Content() string
}

type FileList struct {
	Files map[string]File
}

type Repository interface {
	Id() string
	FetchFiles() (*FileList, error)
	FetchFile(filename string) (File, error)
	Commits() (*object.CommitIter, error)
	Name() string
	Clone() error
}

type GitRepository struct {
	repository *git.Repository
	id         string
	name       string
}

type GitFile struct {
	fo *object.File
}

func (f GitFile) Name() string {
	return f.fo.Name
}
func (f GitFile) Content() string {

	content, err := f.fo.Contents()

	if err != nil {
		return ""
	}

	return content
}

func (gc GitRepository) Clone() error {
	err := gc.repository.Clone(&git.CloneOptions{
		URL: gc.name,
	})

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (gc GitRepository) Id() string {
	return gc.id
}

func (gc GitRepository) Name() string {
	return gc.name
}

func (gc GitRepository) Commits() (*object.CommitIter, error) {
	return gc.repository.Commits()
}

func (gc GitRepository) FetchFiles() (*FileList, error) {

	ref, err := gc.repository.Head()
	if err != nil {
		return nil, err
	}

	// ... retrieving the commit object
	commit, err := gc.repository.Commit(ref.Hash())

	if err != nil {
		return nil, err
	}
	tree, err := commit.Tree()

	if err != nil {
		return nil, err
	}
	ff := tree.Files()
	repoFiles := map[string]File{}
	for {
		fo, err := ff.Next()

		if err != nil {
			break
		}
		repoFile := GitFile{
			fo: fo,
		}
		repoFiles[fo.Name] = repoFile
	}
	return &FileList{repoFiles}, nil
}

func (gc GitRepository) FetchFile(filename string) (File, error) {

	fileList, err := gc.FetchFiles()

	if err != nil {
		return nil, err
	}

	file, ok := fileList.Files[filename]

	if ok == false {
		return nil, fmt.Errorf("Unable to find file: %s", filename)
	}

	return file, nil
}

func NewGitRepository(url string, id string) GitRepository {

	r := git.NewMemoryRepository()

	gf := GitRepository{
		repository: r,
		name:       url,
		id:         id,
	}
	return gf
}

type ProcessTask struct {
	repository Repository
}

func (pt ProcessTask) Execute() cartel.OutputValue {
	// fmt.Println(pt.repository.Name())
	pt.repository.Clone()
	// silly check
	file, err := pt.repository.FetchFile("composer.json")

	counter := 0
	if err == nil {
		counter++
		s := file.Content()
		// silly check
		if strings.Contains(s, "symfony") {
			counter++
		}
	}

	commits, err := pt.repository.Commits()
	// for {
	cmt, err := commits.Next()
	// if err == nil {
	// break
	// }
	commit := convertCommit(*cmt)
	if commit.AuthorName == "Iain Cambridge" {
		counter++
	}
	cmt = nil
	commit = Commit{}
	runtime.GC()
	// }

	m2 := runtime.MemStats{}
	runtime.ReadMemStats(&m2)
	fmt.Printf("1. Allocated and not free - %v\n", m2.Alloc)
	fmt.Printf("2. Allocated and including free - %v\n", m2.TotalAlloc)
	runtime.GC()
	return MockOutput{value: pt.repository.Name()}
}

type MockOutput struct {
	value string
}

func (mo MockOutput) Value() interface{} {
	return mo.value
}
