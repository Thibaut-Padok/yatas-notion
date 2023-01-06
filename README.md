<p align="center">
<img src="docs/auditory.png" alt="yatas-logo" width="30%">
<p align="center">

# YATAS Notion export

A simple idea store easily my yatas report in a notion database

## Getting Started

### Plugin Usage

Add in .yatas.yml file your plugin configuration:

```yaml
plugins:
  - name: "notion"
    enabled: true
    type: "report"
    source: "github.com/Thibaut-Padok/yatas-notion"
    version: "latest"
    description: "Add your yatas report in a notion database"
pluginsConfiguration:
  - pluginName: notion
    pageID: $NotionPageID #The one you want the yatas database to be, it will create the database if not exist. This page must use a Notion connection YATAS-Notion
    token: $NotionConnectionToken # The secret to be able to use Notion Connection
    authToken: $token_v2 # The token from the cookies notion web page to be able to call notionapi/v3 and use yatas-notion advanced options. (optional)
```

Run ```yatas```

- The Yatas Database will be created if needed
- The Yatas report will be created in this Yatas database
- For each test you will have a dedicated page with informations related result.
- If the authToken is provided and it is valid, the page will be automatically Locked.
- Enjoy !

### Important

Such as the Yatas Database is found thanks to his name, please do not change it else a new one will be created.
Is not in the road map to use a variable because I want to keep the pluginConfiguration as simple as possible.

## Usage for development

Useful ```export YATAS_LOG_LEVEL=debug```

Use ```make install```

### Information

As I use 2 differents version of notionapi, I also use 2 differents Go library:

- notionapi/v1 : "github.com/jomei/notionapi"
- notionapi/v3 : "github.com/kjk/notionapi"

But also, this library does not implement all method I need, so I create custom clients with custom methods. Then, I have 4 differents type of client, so to be able to use one of the four client at any time in the code, I choose to manage them thanks to the structure NotionClient. The function take often this objet as argument, like that they could call the different API as they want.

<!-- ## Example
<p align="center">
<img src="docs/demo-html.png" alt="yatas-logo" width="30%">
<p align="center"> -->