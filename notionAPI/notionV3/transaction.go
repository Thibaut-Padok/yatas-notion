package notionV3

import "github.com/kjk/notionapi"

type Pointer struct {
	Table   string `json:"table"`
	ID      string `json:"id"`
	SpaceID string `json:"spaceId"`
}

type Args struct {
	Type           string     `json:"type,omitempty"`
	ListProperties []Property `json:"list_properties,omitempty"`
	Aggregations   []string   `json:"aggregations,omitempty"`
	BlockLocked    bool       `json:"block_locked,omitempty"`
	BlockLockedBy  string     `json:"block_locked_by,omitempty"`
	LastEditedTime int        `json:"last_edited_time,omitempty"`
}

type Property struct {
	Name    string `json:"property"`
	Visible bool   `json:"visible"`
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
	transaction := Transaction{
		TransactionID: "b0ecba00-1717-4513-9c11-04852c0d1c9a",
		SpaceID:       spaceID,
		Debug:         Debug{UserAction: "CollectionSettingsViewLayoutMenu.renderViewTypeButton"},
		Operations:    []Operation{op1},
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

func ShowPropertiesUpdateRequest(spaceID, viewID string, properties []string) UpdateRequest {
	pointer := Pointer{
		Table:   "collection_view",
		ID:      notionapi.ToDashID(viewID),
		SpaceID: spaceID,
	}
	var props []Property
	for _, name := range properties {
		prop := Property{
			Name:    name,
			Visible: true,
		}
		props = append(props, prop)
	}
	op1 := Operation{
		Pointer: pointer,
		Path:    []string{"format"},
		Command: "update",
		Args: Args{
			ListProperties: props,
		},
	}
	transaction := Transaction{
		TransactionID: "439a8b93-5173-4e80-8650-94b3750ab444",
		SpaceID:       spaceID,
		Debug:         Debug{UserAction: "CollectionSettingsViewProperties.setAllPropertiesVisibility"},
		Operations:    []Operation{op1},
	}
	req := UpdateRequest{
		RequestID:    "c78d6ad4-1e7f-46bf-aa6d-fd891f8ba074",
		Transactions: []Transaction{transaction},
	}
	return req
}
