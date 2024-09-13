package bigip

import (
	"context"
	"fmt"
)

type ILXWorkspace struct {
	Name            string      `json:"name,omitempty"`
	FullPath        string      `json:"fullPath,omitempty"`
	SelfLink        string      `json:"selfLink,omitempty"`
	NodeVersion     string      `json:"nodeVersion,omitempty"`
	StagedDirectory string      `json:"stagedDirectory,omitempty"`
	Version         string      `json:"version,omitempty"`
	Extensions      []Extension `json:"extensions,omitempty"`
	Rules           []ILXFile   `json:"rules,omitempty"`
	Generation      int         `json:"generation,omitempty"`
}

type ILXFile struct {
	Name    string `json:"name,omitempty"`
	Content string `json:"content,omitempty"`
}

type Extension struct {
	Name  string    `json:"name,omitempty"`
	Files []ILXFile `json:"files,omitempty"`
}

func (b *BigIP) GetWorkspace(ctx context.Context, path string) (*ILXWorkspace, error) {
	spc := &ILXWorkspace{}
	err, exists := b.getForEntity(spc, uriMgmt, uriTm, uriIlx, uriWorkspace, path)
	if !exists {
		return nil, fmt.Errorf("workspace does not exist: %w", err)
	}
	if err != nil {
		return nil, fmt.Errorf("error getting ILX Workspace: %w", err)
	}

	return spc, nil
}

func (b *BigIP) CreateWorkspace(ctx context.Context, path string) error {
	err := b.post(ILXWorkspace{Name: path}, uriMgmt, uriTm, uriIlx, uriWorkspace, "")
	if err != nil {
		return fmt.Errorf("error creating ILX Workspace: %w", err)
	}

	return nil
}

func (b *BigIP) DeleteWorkspace(ctx context.Context, name string) error {
	err := b.delete(uriMgmt, uriTm, uriIlx, uriWorkspace, name)
	if err != nil {
		return fmt.Errorf("error deleting ILX Workspace: %w", err)
	}
	return nil
}

type ExtensionConfig struct {
	Name          string `json:"name,omitempty"`
	Partition     string `json:"partition,omitempty"`
	WorkspaceName string `json:"workspaceName,omitempty"`
}

func (b *BigIP) CreateExtension(ctx context.Context, opts ExtensionConfig) error {
	err := b.post(ILXWorkspace{Name: opts.WorkspaceName}, uriMgmt, uriTm, uriIlx, uriWorkspace+"?options=extension,"+opts.Name)
	if err != nil {
		return fmt.Errorf("error creating ILX Extension: %w", err)
	}
	return nil
}

type ExtensionFile string

func (e ExtensionFile) Validate() error {
	if e != PackageJSON && e != IndexJS {
		return fmt.Errorf("invalid extension file")
	}
	return nil
}

const (
	PackageJSON ExtensionFile = "package.json"
	IndexJS     ExtensionFile = "index.js"
)

func (b *BigIP) WriteExtensionFile(ctx context.Context, opts ExtensionConfig, content string, filename ExtensionFile) error {
	if err := filename.Validate(); err != nil {
		return err
	}
	destination := fmt.Sprintf("%s/%s/%s/extensions/%s/%s", WORKSPACE_UPLOAD_PATH, opts.Partition, opts.WorkspaceName, opts.Name, filename)
	err := b.WriteFile(content, destination)
	if err != nil {
		return fmt.Errorf("error uploading packagejson: %w", err)
	}
	return nil
}

func (b *BigIP) ReadExtensionFile(ctx context.Context, opts ExtensionConfig, filename ExtensionFile) (*ILXFile, error) {
	if err := filename.Validate(); err != nil {
		return nil, err
	}
	destination := fmt.Sprintf("%s/%s/%s/extensions/%s/%s", WORKSPACE_UPLOAD_PATH, opts.Partition, opts.WorkspaceName, opts.Name, filename)
	files, err := b.ReadFile(destination)
	if err != nil {
		return nil, err
	}
	return files, nil
}
