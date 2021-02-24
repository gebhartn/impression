package handlers

import (
	"io"
	"net/http"
	"path/filepath"

	"github.com/gebhartn/impression/files"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

// Files is a handler for reading and writing files
type Files struct {
	log   hclog.Logger
	store files.Storage
}

// NewFiles creates a new file handler
func NewFiles(s files.Storage, l hclog.Logger) *Files {
	return &Files{store: s, log: l}
}

func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r io.ReadCloser) {
	f.log.Info("Save the file", "id", id, "path", path)

	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r)
	if err != nil {
		f.log.Error("Unable to save file", "error", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
	}
}

// UploadREST implements the http.Handler interface
func (f *Files) UploadREST(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.log.Info("Handle POST", "id", id, "filename", fn)

	f.saveFile(id, fn, rw, r.Body)
}
