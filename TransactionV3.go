package main

import "github.com/kjk/notionapi"

type Pointer struct {
	Table   string `json:"table"`
	ID      string `json:"id"`
	SpaceID string `json:"spaceId"`
}

type Args struct {
	Type           string   `json:"type,omitempty"`
	ListProperties []string `json:"list_properties,omitempty"`
	Aggregations   []string `json:"aggregations,omitempty"`
	BlockLocked    bool     `json:"block_locked,omitempty"`
	BlockLockedBy  string   `json:"block_locked_by,omitempty"`
	LastEditedTime int      `json:"last_edited_time,omitempty"`
}

type Operation struct {
	Pointer Pointer  `json:"pointer"`
	Path    []string `json:"path"`
	Command string   `json:"command"`
	Args    Args     `json:"args"`
}

type Debug struct {
	UserAction string `json:"userAction"`
}

type Transaction struct {
	TransactionID string      `json:"id"`
	SpaceID       string      `json:"spaceId"`
	Debug         Debug       `json:"debug"`
	Operations    []Operation `json:"operations"`
}

type UpdateRequest struct {
	RequestID    string        `json:"requestId"`
	Transactions []Transaction `json:"transactions"`
}

type UpdateResponse struct {
	RawJSON map[string]interface{} `json:"-"`
}

func TableViewTypeUpdateRequest(spaceID, viewID, desiredType string) UpdateRequest {
	pointer := Pointer{
		Table:   "collection_view",
		ID:      viewID,
		SpaceID: spaceID,
	}
	op1 := Operation{
		Pointer: pointer,
		Path:    []string{},
		Command: "update",
		Args:    Args{Type: desiredType},
	}
	op2 := Operation{
		Pointer: pointer,
		Path:    []string{"format"},
		Command: "update",
		Args:    Args{ListProperties: []string{}},
	}
	op3 := Operation{
		Pointer: pointer,
		Path:    []string{"query2"},
		Command: "update",
		Args:    Args{Aggregations: []string{}},
	}
	transaction := Transaction{
		TransactionID: "b0ecba00-1717-4513-9c11-04852c0d1c9a",
		SpaceID:       spaceID,
		Debug:         Debug{UserAction: "CollectionSettingsViewLayoutMenu.renderViewTypeButton"},
		Operations:    []Operation{op1, op2, op3},
	}
	req := UpdateRequest{
		RequestID:    "66312006-8bb3-4afa-9b32-83d29a8c0cc8",
		Transactions: []Transaction{transaction},
	}
	return req
}

func LockPageUpdateRequest(spaceID, pageID string) UpdateRequest {
	pointer := Pointer{
		Table:   "block",
		ID:      notionapi.ToDashID(pageID),
		SpaceID: spaceID,
	}
	op1 := Operation{
		Pointer: pointer,
		Path:    []string{"format"},
		Command: "update",
		Args: Args{
			BlockLocked:   true,
			BlockLockedBy: "9d1601de-8e0d-48f1-a8d7-b2a9b3a356b5",
		},
	}
	transaction := Transaction{
		TransactionID: "f6d982f2-2bfd-409b-af96-21de7c1fa28e",
		SpaceID:       spaceID,
		Debug:         Debug{UserAction: "pageBlockActions.setPageLock"},
		Operations:    []Operation{op1},
	}
	req := UpdateRequest{
		RequestID:    "aa223778-08b4-44f6-8fd1-dfe204cb2215",
		Transactions: []Transaction{transaction},
	}
	return req
}
