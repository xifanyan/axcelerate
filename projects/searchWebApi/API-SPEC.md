```
{
  "openapi": "3.0.1",
  "info": {
    "description": "The Search Web API. It offers access to records, more specifically Create, Read, Update, Delete, and List (CRUDL) in all mindservers, including documents or categories.",
    "version": "1.4",
    "title": "Search Web API",
    "termsOfService": "The Terms of Service for this API, and for the software which runs it, are defined by a contract with OpenText.",
    "contact": {
      "name": "OpenText",
      "url": "www.opentext.com",
      "email": "opentext support contact"
    },
    "license": {
      "url": "Contract with OpenText",
      "name": "Contract with OpenText"
    }
  },
  "servers": [
    {
      "url": "/searchWebApi"
    }
  ],
  "tags": [
    {
      "name": "search",
      "description": "Read-only operations on the API, typically around searching and fetching records.\n\nThe 'project id' is a mindserver identifier. The 'collection id' is either 'default' (which resembles the documents) or some child identifier within the mindserver. A child can be a selected server, a taxonomy, or even a binary storage. This definition applies to all endpoints.\n\nNote: read-only operations typically accept GET as http method. However, you can also use POST with query parameters in the body! This will be accepted as alternative to GET. To this end, specify 'Content-Type: application/x-www-form-urlencoded' and the body contains the query parameters in URL encoded form. This is convenient e.g. for long search expressions, and generally for circumventing URL length restrictions. This applies to some modifying endpoints as well."
    },
    {
      "name": "change",
      "description": "Modifications to existing records (which can be documents, categories, or even binary storage references, depending on the collection id in context)."
    },
    {
      "name": "insert / remove",
      "description": "Allows to create/remove categories or binary artifacts. There is also limited support to create or remove documents."
    },
    {
      "name": "session",
      "description": "Methods to control session handling. ATTENTION: search web api will IMPLICITLY create a new session whenever an old one timed out and credentials have been specified (or an oauth2 token). Thus, login can be done by means of any endpoint, not just session-related ones!\n\nIn order to log in, you can use HTTP basic authentication or send a valid oauth2 token in the Authentication header (which will then start with 'Bearer').\n\nSearch web api is STATEFUL: every login creates a session. The session's id is returned in the HTTP response header SWA-SESSION. All following calls must send the session id in the request header SWA-SESSION. The session consumes resources until it times out or is logged out explicitly. The state allows the API to to-use search results between successive calls.\n\nSearch web api allows sessions of type 'monitoring'. In order to create a monitoring session, send the request header SWA-SESSION-TYPE=MONITORING (alternative values are USER and DEFAULT). In this case, mindservers can be processed by the 'sleeping projects' feature even if your session is active.\n\nLogging: you can forward context identifiers by means of the request header SWA-MDC-TOKEN=tracerId where 'tracerId=SLONG.SLONG [.STRING]*' and 'SLONG' is compatible with Long.valueOf(SLONG, Character.MAX_RADIX). It also accepts the request header SWA-MDC-METHOD=STRING . Both these identifiers end up in all mindserver log files and allow to trace actions throughout the processes."
    }
  ],
  "paths": {
    "/projects": {
      "get": {
        "tags": [
          "search"
        ],
        "summary": "Discover available projects.",
        "description": "Discover available projects, i.e. all projects that are present in an installation.",
        "operationId": "getProjects",
        "responses": {
          "default": {
            "description": "Usually there will be a response, even if an error occurs.",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ProjectsResult"
                }
              }
            }
          }
        }
      }
    },
    "/projects/{projectId}": {
      "get": {
        "tags": [
          "search"
        ],
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          }
        ],
        "summary": "Discover available resources for a given project.",
        "description": "Discover available project-level resources.",
        "operationId": "getProjectResources",
        "responses": {
          "default": {
            "description": "Usually there will be a response, even if an error occurs.",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ProjectResourcesResult"
                }
              }
            }
          }
        }
      }
    },
    "/projects/{projectId}/collections": {
      "get": {
        "tags": [
          "search"
        ],
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          }
        ],
        "summary": "Discover available collections of a project.",
        "description": "Decisiv Search offers multiple domains (e.g. Documents, Matter, People, etc) belonging to a project. With this operation, the set of available domains is returned, e.g. for display as tabs in the UI. For any project that is not a non-merging meta-engine, the collection 'default' will be returned as the only collection.",
        "operationId": "getCollections",
        "responses": {
          "default": {
            "description": "Usually there will be a response, even if an error occurs.",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/CollectionsResult"
                }
              }
            }
          }
        }
      }
    },
    "/projects/{projectId}/collections/{collectionId}": {
      "get": {
        "tags": [
          "search"
        ],
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          }
        ],
        "summary": "Discover available resources for a given collection.",
        "description": "Discover available collection-level resources. Examples are 'records' which describes the actual searchable objects in the domain, 'filters' which describes the folder fields (taxonomies, smart filters), and 'fields' which describs the available fields.",
        "operationId": "getCollectionResources",
        "responses": {
          "default": {
            "description": "Usually there will be a response, even if an error occurs.",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/CollectionResourcesResult"
                }
              }
            }
          }
        }
      }
    },
    "/projects/{projectId}/collections/{collectionId}/fields": {
      "get": {
        "tags": [
          "search"
        ],
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          }
        ],
        "summary": "List the available fields in the given search domain.",
        "description": "Returns the list of available fields and their properties. Note that a folder collection (i.e. a taxonomy) typically has no defined data model such that this end point is unsupported for folder collections. Folder collections always know the field rm_display_name; hierarchical folder collections also know rm_prop_parent (which resembles the folder id of the hierarchical parent).",
        "operationId": "getFields",
        "responses": {
          "default": {
            "description": "Usually there will be a response, even if an error occurs.",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/FieldsResult"
                }
              }
            }
          }
        }
      }
    },
    "/projects/{projectId}/collections/{collectionId}/filters": {
      "get": {
        "tags": [
          "search"
        ],
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          }
        ],
        "summary": "Shows the available folder fields (taxonomies, smart filters) in the given search domain.",
        "description": "Returns the list of available folder fields (taxonomies, smart filters) and their properties. Folder fields can be used for filtering a search result by restricting to documents having a given folder in a given folder field. Also values can be accumulated across documents having a certain folder in a given folder field.",
        "operationId": "getFolderFields",
        "responses": {
          "default": {
            "description": "Usually there will be a response, even if an error occurs.",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/FieldsResult"
                }
              }
            }
          }
        }
      }
    },
    "/projects/{projectId}/collections/{collectionId}/filters/{fieldId}": {
      "get": {
        "tags": [
          "search"
        ],
        "summary": "Discover available resources for a given folder field (taxonomy, smart filter).",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          },
          {
            "$ref": "#/components/parameters/FolderFieldParameter"
          }
        ],
        "description": "Discover available resources for a folder field (taxonomy, smart filter).",
        "operationId": "getFolderFieldResources",
        "responses": {
          "default": {
            "description": "Usually there will be a response, even if an error occurs.",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/FolderFieldResourcesResult"
                }
              }
            }
          }
        }
      }
    },
    "/projects/{projectId}/collections/{collectionId}/filters/{fieldId}/values": {
      "get": {
        "tags": [
          "search"
        ],
        "summary": "Computes values for the selected folder field (taxonomy, smart filter) in the given search domain.",
        "description": "Returns a list of folders (categories) in the given folder field (taxonomy, smart filter). Each resulting folder entry has a count of associated documents. The counts are based on the set of documents found for the given query and restrictions",
        "operationId": "getFolderValues",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          },
          {
            "$ref": "#/components/parameters/FolderFieldParameter"
          },
          {
            "$ref": "#/components/parameters/QueryParameter"
          },
          {
            "$ref": "#/components/parameters/LanguageParameter"
          },
          {
            "$ref": "#/components/parameters/JoinRestriction"
          },
          {
            "in": "query",
            "name": "prefix",
            "description": "Optional prefix to constrain the folders returned. This is equivalent to restrictFoldersByQuery=rm_display_name=VALUE* where VALUE is the prefix of choice.",
            "schema": {
              "type": "string"
            }
          },
          {
            "in": "query",
            "name": "restrictFoldersByQuery",
            "description": "Optional query to constrain the folders returned. Examples are rm_folder_id=uniquevalue or rm_display_name=displayname",
            "schema": {
              "type": "string"
            }
          },
          {
            "in": "query",
            "name": "returnEmptyFolders",
            "description": "Defines if folders with count=0 should be returned. Default is false.",
            "schema": {
              "type": "boolean"
            }
          },
          {
            "in": "query",
            "name": "order",
            "description": "Provide an order criterion for the results. Needs to be one of 'count', 'relevance', 'name', 'name:asc', 'name:desc'. The latter three correspond to the display name.",
            "schema": {
              "type": "string"
            }
          },
          {
            "in": "query",
            "name": "offset",
            "description": "Define a relative position from which to start fetching results. Can be positive or zero (default).",
            "schema": {
              "type": "integer",
              "format": "int32"
            }
          },
          {
            "in": "query",
            "name": "limit",
            "description": "Define a limit on the number of results returned. Default is 20, maximally allowed value is 1000.",
            "schema": {
              "type": "integer",
              "format": "int32"
            }
          },
          {
            "$ref": "#/components/parameters/SearchCacheControl"
          }
        ],
        "responses": {
          "default": {
            "description": "Usually there will be a response, even if an error occurs.",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/FolderValuesResult"
                }
              }
            }
          }
        }
      }
    },
    "/projects/{projectId}/collections/{collectionId}/records": {
      "get": {
        "tags": [
          "search"
        ],
        "summary": "Search for records in the given search domain.",
        "description": "Returns the list of records as specified by the parameters.",
        "operationId": "getRecords",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          },
          {
            "$ref": "#/components/parameters/QueryParameter"
          },
          {
            "$ref": "#/components/parameters/LanguageParameter"
          },
          {
            "$ref": "#/components/parameters/JoinRestriction"
          },
          {
            "in": "query",
            "name": "order",
            "description": "Provide order criteria. Format - \u003Cfield1\u003E[:asc,:desc],field2\u003E[:asc,:desc],...,\u003CfieldN\u003E[:asc,:desc] Examples - 'order=title',  order=custodian:asc,title. Default sorting is as returned by the engine, which is descending by relevancy. The default order for given field sort criteria is 'ascending'.",
            "required": false,
            "schema": {
              "type": "string"
            }
          },
          {
            "$ref": "#/components/parameters/FieldListParameter"
          },
          {
            "$ref": "#/components/parameters/FolderFieldListParameter"
          },
          {
            "$ref": "#/components/parameters/FolderPropertiesRequestParameter"
          },
          {
            "in": "query",
            "name": "body",
            "description": "Indicate whether to retrieve summarized content for the record, default is false. This argument retrieves only parts of the content, namely the part which is relevant for the query or the first sentences of the document.",
            "required": false,
            "schema": {
              "type": "boolean"
            }
          },
          {
            "in": "query",
            "name": "highlight",
            "description": "An argument to 'body': Indicate whether to provide highlighting information (via XML tags of the form \u003CrecomDescriptiveWord\u003E, default is true",
            "required": false,
            "schema": {
              "type": "boolean"
            }
          },
          {
            "in": "query",
            "name": "page",
            "description": "Define which page to retrieve (counting starts at 1). Default is 1.",
            "required": false,
            "schema": {
              "type": "integer",
              "format": "int32"
            }
          },
          {
            "in": "query",
            "name": "limit",
            "description": "Define a page size / limit on the number of results returned. Default is 20. The maximally allowed value is 1000 for all standard paging use-cases. However, if the 'accept' header is set to 'application/x-ndjson', the result will be returned in streaming mode which allows to specify page=1, limit=-1 in order to return all documents (see response documentation for details).",
            "required": false,
            "schema": {
              "type": "integer",
              "format": "int32"
            }
          },
          {
            "in": "query",
            "name": "sponsoredLinks",
            "description": "Computes and returns any configured sponsored links for the provided query. Note that this requires sponsored links to be configured in the project's data model. Default is to not compute sponsored links.",
            "required": false,
            "schema": {
              "type": "boolean",
              "default": "false"
            }
          },
          {
            "in": "query",
            "name": "spellingSuggestions",
            "description": "Computes and returns any configured spelling suggestions for the provided query. Note that this requires spelling suggestions to be configured in the project's data model. Default is to not compute spelling suggestions.",
            "required": false,
            "schema": {
              "type": "boolean",
              "default": "false"
            }
          },
          {
            "$ref": "#/components/parameters/SearchCacheControl"
          }
        ],
        "responses": {
          "default": {
            "description": "The result is a SearchResult containing the specified records and some meta data. Syntax errors or other search errors are returned in the resulting object. The usual response is given in JSON format, but streamed JSON responses are also supported if the 'accept' request header is set to 'application/x-ndjson'. Requesting stream JSON has the following effects: first, it is possible to return all entries of a search result efficiently by using page=1, limit=-1. Second, the output format is NDJSON (newline-delimited json): the first returned record is the SearchResult without records (i.e. only meta data), all following json records are the returned records and the individual records are separated by newlines. Streaming mode always bypasses SearchCacheControl.",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/SearchResult"
                }
              },
              "application/x-ndjson": {
                "schema": {
                  "type": "string",
                  "format": "binary"
                }
              }
            }
          }
        }
      },
      "put": {
        "tags": [
          "change"
        ],
        "summary": "Changes properties of documents",
        "description": "Allows to change properties of documents, for example by tagging them into folders or by modifying textual content.",
        "operationId": "changeAllInSearchResult",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          },
          {
            "$ref": "#/components/parameters/QueryParameter"
          },
          {
            "$ref": "#/components/parameters/LanguageParameter"
          },
          {
            "$ref": "#/components/parameters/BlockUntilCompleteParameter"
          },
          {
            "$ref": "#/components/parameters/SearchCacheControl"
          }
        ],
        "requestBody": {
          "description": "1-n change requests describing the changes. The request body is an array of 1-n ChangeRequest objects. ",
          "content": {
            "application/json": {
              "schema": {
                "type": "array",
                "items": {
                  "$ref": "#/components/schemas/ChangeRequest"
                }
              }
            }
          },
          "required": true
        },
        "responses": {
          "default": {
            "description": "Successful operation.",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ChangeResult"
                }
              }
            }
          }
        },
        "x-codegen-request-body-name": "changes"
      }
    },
    "/projects/{projectId}/collections/{collectionId}/searchToken": {
      "get": {
        "tags": [
          "search"
        ],
        "summary": "Search for records in the given search domain and return a referencable token. Such a token can be used in search queries using SEARCH_IN_SEARCHRESULT=\u003CTOKEN\u003E.",
        "description": "Returns a referencable search token for the search.",
        "operationId": "getSearchResultToken",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          },
          {
            "$ref": "#/components/parameters/QueryParameter"
          },
          {
            "$ref": "#/components/parameters/LanguageParameter"
          },
          {
            "$ref": "#/components/parameters/JoinRestriction"
          },
          {
            "in": "query",
            "name": "order",
            "description": "Provide an optional order criteria. In most cases, search tokens need no ordering as they are simply used to define a set of documents. Format - \u003Cfield1\u003E[:asc,:desc],field2\u003E[:asc,:desc],...,\u003CfieldN\u003E[:asc,:desc] Examples - 'order=title',  order=custodian:asc,title. Default sorting is as returned by the engine, which is descending by relevancy. The default order for given field sort criteria is 'ascending'.",
            "required": false,
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "default": {
            "description": "The search token.",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/SearchResultTokenResponse"
                }
              }
            }
          }
        }
      },
      "delete": {
        "tags": [
          "search"
        ],
        "summary": "Deletes a search token",
        "description": "Deletes a search token.",
        "operationId": "deleteSearchResultToken",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          }
        ],
        "requestBody": {
          "description": "The search token",
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/SearchResultToken"
              }
            }
          },
          "required": true
        }
      },
      "put": {
        "tags": [
          "search"
        ],
        "summary": "Keeps a search token alive",
        "description": "Renews a search token's life time.",
        "operationId": "touchSearchResultToken",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          }
        ],
        "requestBody": {
          "description": "The search token",
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/SearchResultToken"
              }
            }
          },
          "required": true
        },
        "responses": {
          "default": {
            "description": "The search token.",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/SearchResultTokenResponse"
                }
              }
            }
          },
          "410": {
            "description": "The token is expired."
          }
        }
      }
    },
    "/projects/{projectId}/collections/{collectionId}/searchToken/sortOrderSnapshot": {
      "post": {
        "tags": [
          "search"
        ],
        "summary": "Creates a sort order snapshot for the specified search result token.",
        "description": "Computes a sort order snapshot of the specified search result token and stores it for the token. A 'sort order snapshot' is an atomically created sorted list of documents of the token, sorted according to its ordering. Its purpose is to allow a consistent sort order even if the sort keys are subject to concurrent modification. Keep in mind that mindserver always uses partial sort when requesting a page from a search result, i.e. any modifications to sort keys between two page fetches can and will affect the resulting sort order! The sort order snapshot allows to create a presorted consistent result.\n \n Sort order snapshots can and need to be dereferenced explicitly. The suggested workflow is as follows:\n \u003Col\u003E\n \u003Cli\u003Esearch result token A is based on an ordering which is known to be subject to concurrent\n modification,\u003C/li\u003E\n \u003Cli\u003Eyou explicitly create a sort order snapshot for token A,\u003C/li\u003E\n \u003Cli\u003Esearch result token B is based on a field-based-search expression including SEARCH_IN_SEARCHRESULT_RANGE referencing token A using a range of your choice. example:\n SEARCH_IN_SEARCHRESULT_RANGE=A;0;10 . In this context, the results returned from token A use the sort order snapshot.\u003C/li\u003E\n \u003C/ol\u003E\n Thus, a sort order snapshots allows to access the ordering by means of this field-based-search expression.\n It does _not_ allow to access the scores of A.\n\n The sorted list of documents is stored in memory, next to the token's state. It is cleared if the token is closed or the engine gets restarted.\n \n There can be at most one sort order snapshot per token; any previously existing sort order snap shot\n will be overwritten.",
        "operationId": "createSortOrderSnapshot",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          },
          {
            "in": "query",
            "name": "topN",
            "description": "an optional limit describing how many documents should become part of the sort order snapshot. A negative value (or a value which is larger than the search result) include all documents. All other values restrict the sort order snapshot to the N top-ranked documents of the last search. Consequently, usage of the sort order snapshot will return at most these N top ranked documents. Default is to use the entire search result.",
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          }
        ],
        "requestBody": {
          "description": "The search token",
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/SearchResultToken"
              }
            }
          },
          "required": true
        }
      }
    },
    "/projects/{projectId}/collections/{collectionId}/search/highlightExpression": {
      "get": {
        "tags": [
          "search"
        ],
        "summary": "Returns highlight expressions for a search request.",
        "description": "Returns regular expression(s) describing the words associated with a search request. Depending on the search request, this can become large. The return value is grouped by field. If an expression is to be applied to all fields, its field name is '*'.",
        "operationId": "getHighlightingForSearchResult",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          },
          {
            "$ref": "#/components/parameters/QueryParameter"
          },
          {
            "$ref": "#/components/parameters/LanguageParameter"
          },
          {
            "$ref": "#/components/parameters/JoinRestriction"
          },
          {
            "$ref": "#/components/parameters/SearchCacheControl"
          }
        ],
        "responses": {
          "default": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/SearchResultHighlightingResult"
                }
              }
            }
          }
        }
      }
    },
    "/projects/{projectId}/collections/{collectionId}/records/{recordId}": {
      "get": {
        "tags": [
          "search"
        ],
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          },
          {
            "$ref": "#/components/parameters/RecordIdParameter"
          }
        ],
        "summary": "Discover available resources for a given record ID (document ID).",
        "description": "Discover available resources for a record ID (document ID).",
        "operationId": "getRecordResources",
        "responses": {
          "default": {
            "description": "Usually there will be a response, even if an error occurs.",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/RecordResourcesResult"
                }
              }
            }
          }
        }
      }
    },
    "/projects/{projectId}/collections/{collectionId}/records/{recordId}/content": {
      "get": {
        "tags": [
          "search"
        ],
        "summary": "Fetches properties (or the paged document body) of a selected record in the given search domain.",
        "description": "Returns selected properties (possibly including paged body content) for the record in question.",
        "operationId": "fetch",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          },
          {
            "$ref": "#/components/parameters/RecordIdParameter"
          },
          {
            "$ref": "#/components/parameters/FieldListParameter"
          },
          {
            "$ref": "#/components/parameters/FolderFieldListParameter"
          },
          {
            "$ref": "#/components/parameters/FolderPropertiesRequestParameter"
          },
          {
            "in": "query",
            "name": "body",
            "description": "Indicate whether to retrieve content for the record, default is false. This parameter is required in order to display highlighting results. It fetches fields configured as 'Content Fields' in the data model, applies highlighting, paging, and summarization as needed, and returns the resulting XML.",
            "required": false,
            "schema": {
              "type": "boolean"
            }
          },
          {
            "in": "query",
            "name": "page",
            "description": "Define which page to retrieve (counting starts at 1). Default is 1. A page number which is too large will automatically be replaced by the highest page number.",
            "required": false,
            "schema": {
              "type": "integer",
              "format": "int32"
            }
          },
          {
            "in": "query",
            "name": "highlightSearchTermQuery",
            "description": "A query expression which should be used to highlight hits in the requested document. The query expression resembles the value of 'query' for search operations, but it is only used for search term highlighting and concept highlighting. Note that the record id in question does not necessarily need to match the query in question.",
            "required": false,
            "schema": {
              "type": "string"
            }
          },
          {
            "in": "query",
            "name": "highlightSearchTermLanguage",
            "description": "Specify a language to be used when interpreting the query. The default is to use the result of the query language detection. The argument is a two-letter ISO 639 language code like 'en'. The language is used in order to define stemming rules, for example.",
            "schema": {
              "type": "string"
            }
          },
          {
            "in": "query",
            "name": "highlightSearchTermJoinRestriction",
            "description": "This is a part of search-term based highlighting, see 'highlightSearchTermQuery'. It has the same syntax as 'joinRestriction' for searches.",
            "schema": {
              "type": "string"
            }
          },
          {
            "in": "query",
            "name": "fieldsHighlighted",
            "description": "An URI-encoded, comma-separated list of fields which will be fetched in XML form for which highlighting has been applied. The resulting field content is in XML form with \u003CrecomDescriptiveWord\u003E tags for highlighted words. The response object will contain one field for each item in this list, and the result items will be have the suffix 'Xml' to indicate that the highlighted value is in XML form (example 'titleXml'). Note that fields specified in this context do not necessarily need to be known in the data model. Note furthermore that highlighting will only match for indexed fields.",
            "schema": {
              "type": "string"
            }
          },
          {
            "in": "query",
            "name": "highlightHitNavigation",
            "description": "An argument to 'body': Allows to return specific highlighted positions within the document. Valid choices are 'first', 'previous', 'next', 'last', 'firstUser', 'nextUser'. Default is to use no hit navigation. Note that 'first', 'last', and 'firstUser' ignore the page number and jump to the first resp. last hit. The choice 'next' jumps to the first page containing a hit which has a page number bigger than the 'page' parameter. If 'page' is the last page, it continues searching at the first page.",
            "schema": {
              "type": "string"
            }
          },
          {
            "in": "query",
            "name": "highlightUserTerms",
            "description": "An argument to 'body': A comma-separated list of terms to highlight.",
            "schema": {
              "type": "string"
            }
          },
          {
            "in": "query",
            "name": "highlightFolderFieldList",
            "description": "An argument to 'body': A comma-separated list of folder field names for which folder-specific highlighting is to be applied.",
            "schema": {
              "type": "string"
            }
          },
          {
            "in": "query",
            "name": "summarize",
            "description": "An argument to 'body': indicates whether to provide a textual summary for the record, default is false in this context. The summary shows text portions around matching highlighted words in the document.",
            "schema": {
              "type": "boolean"
            }
          },
          {
            "$ref": "#/components/parameters/SearchCacheControl"
          }
        ],
        "responses": {
          "default": {
            "description": "Successful operation. Returns a record object which encapsulates a list of result objects, including the specified field values. The response object differs from the one returned by a search in the following aspects: it has no relevance and highlighted field names are returned with the string suffix 'Xml'.",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/Record"
                }
              }
            }
          }
        }
      },
      "put": {
        "tags": [
          "change"
        ],
        "summary": "Changes properties of documents",
        "description": "Allows to change properties of documents, for example by tagging them into folders or by modifying textual content.",
        "operationId": "change",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          },
          {
            "$ref": "#/components/parameters/RecordIdParameter"
          },
          {
            "$ref": "#/components/parameters/BlockUntilCompleteParameter"
          }
        ],
        "requestBody": {
          "description": "1-n change requests describing the changes. The request body is an array of 1-n ChangeRequest objects. The request accepts either application/json or multipart/form-data. The multi-part request is required in order to update streamed fields in which case the binaries need to be referenced by writing their index as value into a SET_TEXT change request.",
          "content": {
            "application/json": {
              "schema": {
                "type": "array",
                "items": {
                  "$ref": "#/components/schemas/ChangeRequest"
                }
              }
            },
            "multipart/form-data": {
              "schema": {
                "type": "object",
                "properties": {
                  "request": {
                    "type": "array",
                    "items": {
                      "type": "object",
                      "schema": {
                        "$ref": "#/components/schemas/ChangeRequest"
                      }
                    }
                  },
                  "binaries": {
                    "type": "array",
                    "items": {
                      "type": "string",
                      "format": "binary"
                    }
                  }
                }
              }
            }
          },
          "required": true
        },
        "responses": {
          "default": {
            "description": "Successful operation.",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ChangeResult"
                }
              }
            }
          }
        },
        "x-codegen-request-body-name": "changes"
      }
    },
    "/projects/{projectId}/collections/{collectionId}/records/{recordId}/inDocumentSearch": {
      "get": {
        "tags": [
          "search"
        ],
        "summary": "Returns strings/regular expression for words matching inside of a document.",
        "description": "Searches within the record and returns the matches as string and/or regular expression",
        "operationId": "searchInDocumentText",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          },
          {
            "$ref": "#/components/parameters/RecordIdParameter"
          },
          {
            "in": "query",
            "name": "highlightSearchTermQuery",
            "description": "A query expression which should be used to highlight hits in the requested document. The query expression resembles the value of 'query' for search operations, but it is only used for search term highlighting and concept highlighting. Note that the record id in question does not necessarily need to match the query in question.",
            "required": false,
            "schema": {
              "type": "string"
            }
          },
          {
            "in": "query",
            "name": "highlightSearchTermLanguage",
            "description": "Specify a language to be used when interpreting the query. The default is to use the result of the query language detection. The argument is a two-letter ISO 639 language code like 'en'. The language is used in order to define stemming rules, for example.",
            "schema": {
              "type": "string"
            }
          },
          {
            "in": "query",
            "name": "highlightSearchTermJoinRestriction",
            "description": "This is a part of search-term based highlighting, see 'highlightSearchTermQuery'. It has the same syntax as 'joinRestriction' for searches.",
            "schema": {
              "type": "string"
            }
          },
          {
            "in": "query",
            "name": "highlightUserTerms",
            "description": "An argument to 'body': A comma-separated list of terms to highlight.",
            "schema": {
              "type": "string"
            }
          },
          {
            "in": "query",
            "name": "highlightFolderFieldList",
            "description": "An argument to 'body': A comma-separated list of folder field names for which folder-specific highlighting is to be applied.",
            "schema": {
              "type": "string"
            }
          },
          {
            "in": "query",
            "name": "highlightFolderFieldsAggregation",
            "description": "A comma-separated list which is as long as highlightFolderFieldList. For each element of highlightFolderFieldList, it defines how to return values of that field: either aggregated 'by field' or 'by folder'. Example: highlightFolderFieldList='field1,field2'  and highlightFolderFieldsAggregation='by field, by folder'",
            "required": false,
            "schema": {
              "type": "string"
            }
          },
          {
            "in": "query",
            "name": "contentFieldNames",
            "description": "(Optional) Specifies a comma-separated list of XML content tag which are to be searched. Defaults to the configured content tags of the target project.",
            "required": false,
            "schema": {
              "type": "string"
            }
          },
          {
            "in": "query",
            "name": "pageTag",
            "description": "Specifies an XML tag name which defines page boundaries. Everything within this page tag is considered to be part of one page.\n\nPaging is only used in order to determine how many hits were found on each page.\n\nThe default is to page automatically (i.e. just as when the XML was text generated using the normal content endpoint).\n\nUse omitHitsPerPage in order to disable paging entirely.\n\nThe argument must be an XML tag name. Note that the page tag must be WITHIN the content tag, otherwise it will be ignored! It is valid to use pagetag name == content tag name.",
            "required": false,
            "schema": {
              "type": "string"
            }
          },
          {
            "in": "query",
            "name": "omitHitsPerPage",
            "description": "Instructs the backend to not compute hits by page, i.e. the return value will be null. This essentially disables paging.",
            "required": false,
            "schema": {
              "type": "boolean"
            }
          },
          {
            "in": "query",
            "name": "requestHitLocationsPageRelative",
            "description": "Instructs the backend to return hit locations by page.\n\nThere will be one location for each hit. The location is a relative number if the hit appears at the beginning of its page (0%) or at the end (100%).",
            "required": false,
            "schema": {
              "type": "boolean"
            }
          },
          {
            "in": "query",
            "name": "requestHitLocationsDocumentRelative",
            "description": "Instructs the backend to return hit locatons for each hit.\n\nThere will be one location for each hit. The location is a relative number if the hit appears at the beginning of the document (0%) or at the end (100%).",
            "required": false,
            "schema": {
              "type": "boolean"
            }
          },
          {
            "$ref": "#/components/parameters/SearchCacheControl"
          }
        ],
        "responses": {
          "default": {
            "description": "The response",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/HighlightedWordResult"
                }
              }
            }
          }
        }
      },
      "put": {
        "tags": [
          "change"
        ],
        "summary": "Changes properties of documents",
        "description": "Allows to change properties of documents, for example by tagging them into folders or by modifying textual content.",
        "operationId": "change",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          },
          {
            "$ref": "#/components/parameters/RecordIdParameter"
          },
          {
            "$ref": "#/components/parameters/BlockUntilCompleteParameter"
          }
        ],
        "requestBody": {
          "description": "1-n change requests describing the changes. The request body is an array of 1-n ChangeRequest objects. The request accepts either application/json or multipart/form-data. The multi-part request is required in order to update streamed fields in which case the binaries need to be referenced by writing their index as value into a SET_TEXT change request.",
          "content": {
            "application/json": {
              "schema": {
                "type": "array",
                "items": {
                  "$ref": "#/components/schemas/ChangeRequest"
                }
              }
            },
            "multipart/form-data": {
              "schema": {
                "type": "object",
                "properties": {
                  "request": {
                    "type": "array",
                    "items": {
                      "type": "object",
                      "schema": {
                        "$ref": "#/components/schemas/ChangeRequest"
                      }
                    }
                  },
                  "binaries": {
                    "type": "array",
                    "items": {
                      "type": "string",
                      "format": "binary"
                    }
                  }
                }
              }
            }
          },
          "required": true
        },
        "responses": {
          "default": {
            "description": "Successful operation.",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ChangeResult"
                }
              }
            }
          }
        },
        "x-codegen-request-body-name": "changes"
      }
    },
    "/projects/{projectId}/collections/{collectionId}/binary": {
      "get": {
        "tags": [
          "search"
        ],
        "summary": "Search for records in the given search domain and return selected binary content for a selected field.",
        "description": "Returns one binary which matches the specified search criteria. The field containing the binary data must be specified. If more than one record matches the query, the binary of the record with index 'selectedIndex' is returned. You can inspect the ordering and the number of hits by sending the same query to the records endpoint. The returned filename always starts with the value which is returned when fetching the field via the records endpoint. It ends with the file extension as available to the storage. If the storage knows some original file name, the original file name becomes an infix of sorts ___fileName just before the file extension. The mindserver storage can be accessed by using the singleMindServer identifier as projectId and the collectionId should be 'rm_storage:Image files' or 'rm_storage:Native files' . The precise value after the colon resembles the storage element type in the mindserver configuration.",
        "operationId": "getBinary",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          },
          {
            "$ref": "#/components/parameters/QueryParameter"
          },
          {
            "$ref": "#/components/parameters/LanguageParameter"
          },
          {
            "$ref": "#/components/parameters/JoinRestriction"
          },
          {
            "in": "query",
            "name": "order",
            "description": "Provide order criteria. Format - \u003Cfield1\u003E[:asc,:desc],field2\u003E[:asc,:desc],...,\u003CfieldN\u003E[:asc,:desc] Examples - 'order=title',  order=custodian:asc,title. Default sorting is as returned by the engine, which is descending by relevancy. The default order for given field sort criteria is 'ascending'.",
            "required": false,
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "field",
            "in": "query",
            "description": "A single field name storing binary data.",
            "schema": {
              "type": "string"
            }
          },
          {
            "in": "query",
            "name": "selectedIndex",
            "description": "Define the stream to select (counting starts at 1). This is useful if and only if the query returns more than one result. Default is 1.",
            "required": false,
            "schema": {
              "type": "integer",
              "format": "int32"
            }
          },
          {
            "$ref": "#/components/parameters/SearchCacheControl"
          }
        ],
        "responses": {
          "default": {
            "description": "Usually there will be a response, even if an error occurs.",
            "headers": {
              "Transfer-Encoding": {
                "schema": {
                  "type": "string"
                },
                "description": "chunked"
              }
            },
            "content": {
              "application/octet-stream": {
                "schema": {
                  "type": "string",
                  "format": "binary"
                }
              }
            }
          }
        }
      }
    },
    "/projects/{projectId}/collections/{collectionId}/binary/{recordId}/content": {
      "get": {
        "tags": [
          "search"
        ],
        "summary": "Search for records in the given search domain and return selected binary content for a selected field.",
        "description": "Returns one binary which matches the specified search criteria. The field containing the binary data must be specified. If more than one record matches the query, the binary of the record with index 'selectedIndex' is returned. You can inspect the ordering and the number of hits by sending the same query to the records endpoint. The returned filename always starts with the value which is returned when fetching the field via the records endpoint. It ends with the file extension as available to the storage. If the storage knows some original file name, the original file name becomes an infix of sorts ___fileName just before the file extension. The mindserver storage can be accessed by using the singleMindServer identifier as projectId and the collectionId should be 'rm_storage:Image files' or 'rm_storage:Native files' . The precise value after the colon resembles the storage element type in the mindserver configuration.",
        "operationId": "getBinaryByRecordId",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          },
          {
            "$ref": "#/components/parameters/RecordIdParameter"
          },
          {
            "name": "field",
            "in": "query",
            "description": "A single field name storing binary data.",
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "default": {
            "description": "Usually there will be a response, even if an error occurs.",
            "headers": {
              "Transfer-Encoding": {
                "schema": {
                  "type": "string"
                },
                "description": "chunked"
              }
            },
            "content": {
              "application/octet-stream": {
                "schema": {
                  "type": "string",
                  "format": "binary"
                }
              }
            }
          }
        }
      }
    },
    "/projects/{projectId}/collections/{collectionId}/measures": {
      "post": {
        "tags": [
          "search"
        ],
        "summary": "Computes counts or other aggregates along one or two dimensions. A query identifies documents, and the result contains the aggregated value for each returned bucket within the dimension(s). The endpoint supports one dimension (for counts, this case is essentially the same as the filters endpoint with a different return value). The endpoint also supports two dimensions in which case it will compute 'pair counts' or 'pair aggregates'. The dimensions support fields of type numeric, date, or folder fields. The most simple measure is to simply count documents falling into the buckets.",
        "description": "Returns the measure values. For zero-dimensional measures, it just computes a single value as aggregate over the entire search result. For one-dimensional measures, the result resembles a sorted table in which measure dimensions (like folders) are associated with the measure. For two-dimensional measures, the result is a matrix (a cube) which returns measure value associated with a pair of values. The most simple measure is the count of documents which fall into one of the buckets.",
        "operationId": "getMeasureCube",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          },
          {
            "$ref": "#/components/parameters/QueryParameter"
          },
          {
            "$ref": "#/components/parameters/LanguageParameter"
          },
          {
            "$ref": "#/components/parameters/JoinRestriction"
          },
          {
            "$ref": "#/components/parameters/MeasureTypeParameter"
          },
          {
            "$ref": "#/components/parameters/SearchCacheControl"
          }
        ],
        "requestBody": {
          "description": "An array of 0, 1 or 2 elements. If the array has size 0, the request is to aggregate a single number over the entire search result. If the array has size 1, the request is to compute a table of values. For measure type 'count', this returns 'normal' counts (sometimes called 'structured view'). If the array has size 2, the request is to compute pair counts / pair aggregates, i.e. the buckets are defined as a tuple and the aggregate is computed for documents in which fall into both tuple elements. The request body is an array of DimensionRequest objects. ",
          "content": {
            "application/json": {
              "schema": {
                "type": "array",
                "items": {
                  "$ref": "#/components/schemas/DimensionRequest"
                }
              }
            }
          },
          "required": true
        },
        "responses": {
          "default": {
            "description": "Usually there will be a response, even if an error occurs.",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/MeasureCube"
                }
              }
            }
          }
        },
        "x-codegen-request-body-name": "dimensions"
      }
    },
    "/projects/{projectId}/collections/{collectionId}/records/insertRemoveTransaction": {
      "post": {
        "tags": [
          "insert / remove"
        ],
        "summary": "Inserts new records. The primary application is to create folders in folder collections (taxonomies).",
        "description": "Creates new records in the specified collection.\n\n Note that this end point comes with two flavors: the first accepts application/json. It is intended to insert folders (=categories) into folder collections (=taxonomies). To this end, the projectId should be something like singleMindServer.NAME (or metaMindServer.NAME) and the collectionId should resemble the folder collection name (=taxonomy id). The second flavor accepts multipart/form-data. It is intended to populate collections containing fields requiring some sort of file streaming, typically fields with binary data or with full text content (streamed xml).\n\nThe mindserver storage can be accessed by using the singleMindServer identifier as projectId and the collectionId should be 'rm_storage:Image files' or 'rm_storage:Native files' (the precise value after the colon resembles the storage element type in the mindserver configuration). The data model for these storage-related collections can be retrieved using the fields endpoint.\n\nNote that it is also supported to create new documents. In this case, fields of type CONTENT accept a binary of type application/xml which will be streamed. XML streaming is special in several aspects: \n 1. The XML contains structured information which is inserted as-is, and the content field(s) are somewhere inside of it.\n2. The XML can also contain data for standard fields (typically in the form &lt;document&gt;&lt;meta&gt; &lt;field&gt;VALUE&lt;/field&gt;&lt;/meta&gt; .... &lt;/document&gt;). If a field is present in both XML and fields of the RecordData, the RecordData wins.\n3. It is _mandatory_ that some fields are specified as fields in the RecordData (if they are needed)! This includes rm_attachmentroot, rm_crawlid, rm_checksum, rm_modificationdate.\n4. If a native is to be inserted together with the document, its value can be specified as field rm_native. More precisely: specify a field with name rm_native, and its value resembles the index of its binary inside of the multipart/formdata.\n5. If more than one XML content stream is attached, the server will insert an artificial &lt;document&gt; ...&lt;/document&gt; around all streams and will concatenate them.\n\nAny streamed fields (i.e. either binary fields or XML content streams) are to be specified as follows: such a field must be specified into the RecordData, and the value for such a field is the integer index of the stream inside of the multipart/formdata. Example : 'fieldName': 'p', 'value': 0.",
        "operationId": "insertRemoveTransaction",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          },
          {
            "$ref": "#/components/parameters/SearchCacheControl"
          }
        ],
        "requestBody": {
          "description": "The request which defines the data for new records and/or ids of old records which are to be deleted (replaced). In case of multipart/form-data, each attached binary record must be referenced by a field data. In order to reference an attached binary record, its 0-based index must be used as value in a FieldData. ",
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/InsertRemoveRequest"
              }
            },
            "multipart/form-data": {
              "schema": {
                "type": "object",
                "properties": {
                  "request": {
                    "type": "object",
                    "schema": {
                      "$ref": "#/components/schemas/InsertRemoveRequest"
                    }
                  },
                  "binaries": {
                    "type": "array",
                    "items": {
                      "type": "string",
                      "format": "binary"
                    }
                  }
                }
              }
            }
          },
          "required": true
        },
        "responses": {
          "default": {
            "description": "Usually there will be a response, even if an error occurs.",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/InsertRemoveResult"
                }
              }
            }
          }
        },
        "x-codegen-request-body-name": "request"
      }
    },
    "/projects/{projectId}/collections/{collectionId}/records/bulkInsertRemoveTransaction": {
      "post": {
        "tags": [
          "insert / remove"
        ],
        "summary": "Starts an insert/remove transaction. The transaction must be closed by means of a POST on the associated .../{indexingBufferId}/end endpoint",
        "description": "Prepare a server-sided buffer for an upcoming insert/remove transaction. The transaction must be finished by the user.",
        "operationId": "startInsertRemoveTransaction",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          }
        ],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/StartTransactionRequest"
              }
            }
          },
          "required": true
        },
        "responses": {
          "default": {
            "description": "Usually there will be a response, even if an error occurs.",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/StartTransactionResult"
                }
              }
            }
          }
        },
        "x-codegen-request-body-name": "request"
      }
    },
    "/projects/{projectId}/collections/{collectionId}/records/bulkInsertRemoveTransaction/{indexingBufferId}/end": {
      "post": {
        "tags": [
          "insert / remove"
        ],
        "summary": "Finishes an insert remove transaction asynchronously. The endpoint returns a job id can be polled afterwards.",
        "description": "Finishes an insert remove transaction. This can take a long time and the operation is executed asynchronously. The endpoint returns the job id.",
        "operationId": "commitInsertRemoveTransaction",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          },
          {
            "name": "indexingBufferId",
            "in": "path",
            "required": true,
            "description": "A transaction ID name. Must resemble a value returned in a StartTransactionResult.",
            "schema": {
              "type": "string"
            }
          }
        ],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/FinishTransactionRequest"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "Successful operation",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/FinishTransactionResponse"
                }
              }
            }
          }
        }
      }
    },
    "/projects/{projectId}/collections/{collectionId}/records/bulkInsertRemoveTransaction/{indexingBufferId}/end/{jobId}": {
      "get": {
        "tags": [
          "insert / remove"
        ],
        "summary": "Returns the status of a previously submitted 'end transaction' request.",
        "description": "Returns the status of a previously submitted 'end transaction' request.",
        "operationId": "getFlushJobStatus",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          },
          {
            "name": "indexingBufferId",
            "in": "path",
            "required": true,
            "description": "A transaction ID name. Must resemble a value returned in a StartTransactionResult.",
            "schema": {
              "type": "string"
            }
          },
          {
            "name": "jobId",
            "in": "path",
            "required": true,
            "description": "The identifier used to submit a previous 'end transaction' job.",
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/JobStatusResponse"
                }
              }
            }
          }
        }
      }
    },
    "/projects/{projectId}/collections/{collectionId}/records/bulkInsertRemoveTransaction/{indexingBufferId}/buffer": {
      "post": {
        "tags": [
          "insert / remove"
        ],
        "summary": "A bulk variant of the .../insertRemoveTransaction: it has exactly the same arguments and behavior, except that it merely buffers results into a previously prepared bulk transaction.",
        "description": "Creates new records in the specified collection.\n\n Note that this end point comes with two flavors: the first accepts application/json. It is intended to insert folders (=categories) into folder collections (=taxonomies). To this end, the projectId should be something like singleMindServer.NAME (or metaMindServer.NAME) and the collectionId should resemble the folder collection name (=taxonomy id). The second flavor accepts multipart/form-data. It is intended to populate collections containing fields requiring some sort of file streaming, typically fields with binary data or with full text content (streamed xml).\n\nThe mindserver storage can be accessed by using the singleMindServer identifier as projectId and the collectionId should be 'rm_storage:Image files' or 'rm_storage:Native files' (the precise value after the colon resembles the storage element type in the mindserver configuration). The data model for these storage-related collections can be retrieved using the fields endpoint.\n\nNote that it is also supported to create new documents. In this case, fields of type CONTENT accept a binary of type application/xml which will be streamed. XML streaming is special in several aspects: \n 1. The XML contains structured information which is inserted as-is, and the content field(s) are somewhere inside of it.\n2. The XML can also contain data for standard fields (typically in the form &lt;document&gt;&lt;meta&gt; &lt;field&gt;VALUE&lt;/field&gt;&lt;/meta&gt; .... &lt;/document&gt;). If a field is present in both XML and fields of the RecordData, the RecordData wins.\n3. It is _mandatory_ that some fields are specified as fields in the RecordData (if they are needed)! This includes rm_attachmentroot, rm_crawlid, rm_checksum, rm_modificationdate.\n4. If a native is to be inserted together with the document, its value can be specified as field rm_native. More precisely: specify a field with name rm_native, and its value resembles the index of its binary inside of the multipart/formdata.\n5. If more than one XML content stream is attached, the server will insert an artificial &lt;document&gt; ...&lt;/document&gt; around all streams and will concatenate them.\n\nAny streamed fields (i.e. either binary fields or XML content streams) are to be specified as follows: such a field must be specified into the RecordData, and the value for such a field is the integer index of the stream inside of the multipart/formdata. Example : 'fieldName': 'p', 'value': 0.",
        "operationId": "addToInsertRemoveTransaction",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          },
          {
            "name": "indexingBufferId",
            "in": "path",
            "required": true,
            "description": "A transaction ID name. Must resemble a value returned in a StartTransactionResult.",
            "schema": {
              "type": "string"
            }
          },
          {
            "$ref": "#/components/parameters/SearchCacheControl"
          }
        ],
        "requestBody": {
          "description": "The request which defines the data for new records and/or ids of old records which are to be deleted (replaced). In case of multipart/form-data, each attached binary record must be referenced by a field data. In order to reference an attached binary record, its 0-based index must be used as value in a FieldData. ",
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/InsertRemoveRequest"
              }
            },
            "multipart/form-data": {
              "schema": {
                "type": "object",
                "properties": {
                  "request": {
                    "type": "object",
                    "schema": {
                      "$ref": "#/components/schemas/InsertRemoveRequest"
                    }
                  },
                  "binaries": {
                    "type": "array",
                    "items": {
                      "type": "string",
                      "format": "binary"
                    }
                  }
                }
              }
            }
          },
          "required": true
        },
        "responses": {
          "default": {
            "description": "Success",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/InsertRemoveResult"
                }
              }
            }
          },
          "400": {
            "description": "If the requested buffer id has not been started."
          }
        },
        "x-codegen-request-body-name": "request"
      }
    },
    "/projects/{projectId}/collections/{collectionId}/cachedSearches": {
      "get": {
        "tags": [
          "search"
        ],
        "summary": "Return a list of cached searches.",
        "description": "Returns a list of cached searches. Search web api automatically caches search results such that paging or following actions are fast without re-executing the same query. This method allows to inspect the available caches for the current session and for the selected project / collection. The result value resembles a trace id for each cached search; more precisely: the trace id which created the cached search result. Trace ids can be set using the request header SWA-MDC-TOKEN; they typically serve as log file analysis utility - in case of searches, they also identify the search initiator context.",
        "operationId": "getCachedSearches",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          }
        ],
        "responses": {
          "default": {
            "description": "Usually there will be a response, even if an error occurs.",
            "content": {
              "application/json": {
                "schema": {
                  "type": "array",
                  "items": {
                    "$ref": "#/components/schemas/CachedSearchDescription"
                  }
                }
              }
            }
          }
        }
      },
      "delete": {
        "tags": [
          "search"
        ],
        "summary": "Drops cached searches",
        "description": "Drops cached searches.",
        "operationId": "dropCachedSearches",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          },
          {
            "name": "creationTraceIds",
            "in": "query",
            "description": "A comma-separated list of creation trace ids. Only these associated caches will be cleared. Omitting this value will close all available searches for the specified collection. Such values can be retrieved using the GET endpoint.",
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "default": {
            "description": "Successful operation.",
            "content": {
              "application/json": {
                "schema": {
                  "type": "array",
                  "items": {
                    "$ref": "#/components/schemas/CachedSearchDescription"
                  }
                }
              }
            }
          }
        }
      }
    },
    "/projects/{projectId}/collections/{collectionId}/changes/queue": {
      "get": {
        "tags": [
          "change"
        ],
        "summary": "Waits until all currently queued changes are processed.",
        "description": "Change requests submitted with 'blockUntilComplete=false' are queued. Their value is immediately visible when fetching it, but searches are still based on their old state until the queue has been processed. This method allows to wait for queued change requests; more precisely: for all change requests which are queued (or in progress) at the time when this wait request is submitted.",
        "operationId": "waitForAllCurrentlyScheduledChangeRequests",
        "parameters": [
          {
            "$ref": "#/components/parameters/ProjectParameter"
          },
          {
            "$ref": "#/components/parameters/CollectionParameter"
          },
          {
            "in": "query",
            "name": "timeoutMillis",
            "description": "Defines an optional timeout in milliseconds. Default is 60000. Use a negative value to wait indefinitely.",
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          },
          {
            "in": "query",
            "name": "onlyHighPriorityChanges",
            "description": "Defines that only scheduled high priority change requests should be considered. 'High Priority' means: change requests involving single documents.",
            "schema": {
              "type": "boolean"
            }
          }
        ],
        "responses": {
          "default": {
            "description": "Successful operation.",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/WaitForPendingChangesResult"
                }
              }
            }
          }
        }
      }
    },
    "/login": {
      "post": {
        "tags": [
          "session"
        ],
        "summary": "Creates a new session.",
        "description": "Creates a new session. NOTE: all other endpoints implicitly create new sessions based on request headers as well. This method merely allows an explicit login without using side effects. The session id is returned as response header SWA-SESSION and must be provided to all following requests.",
        "operationId": "login",
        "responses": {
          "default": {
            "description": "Usually there will be a response, even if an error occurs.",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/LoginResult"
                }
              }
            }
          }
        }
      }
    },
    "/logout": {
      "delete": {
        "tags": [
          "session"
        ],
        "summary": "Delete current session.",
        "description": "Delete the current session, close open resources.",
        "operationId": "logout",
        "responses": {
          "default": {
            "description": "Usually there will be a response, even if an error occurs.",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/LogoutResult"
                }
              }
            }
          }
        }
      }
    }
  },
  "components": {
    "schemas": {
      "StatusObject": {
        "type": "object",
        "description": "The status of the response.",
        "properties": {
          "successful": {
            "type": "boolean",
            "description": "Will be 'true' if the request was fully successful."
          },
          "backendStatus": {
            "type": "string",
            "enum": [
              "ok",
              "meaningless search terms",
              "unknown error",
              "all search terms unknown",
              "child collection invalid",
              "incomplete search results due to pending synchronization",
              "unexpandable wildcard",
              "expanded query too long",
              "syntax error",
              "unsupported expression"
            ],
            "description": "The status of the backend response if there is one. This can be 'null'. The value 'meaningless search terms' indicates that the query consists only of stop words and the-like and results in HTTP code 200. The value 'all search terms unknown' is similiar, but the server does not know any of the provided terms. The special value 'incomplete search results due to pending synchronization' can only occur in early startup phases of the target server (if synchronizes searchable folder properties asynchronously); it results in HTTP code 200. The value 'child collection invalid' means that a sub-engine is currently unreachable and its results are not part of the result and also results in code 200. The value 'unexpandable wildcard' indicates that a wildcard prefix is too short (it results in HTTP code 400). The value 'expanded query too long' indicates that a wildcard expression resulted in too many terms (which is also code 400). The values here are constants and are subject to api stability requirements. The value 'syntax error' indicates a syntax error in a query. The value 'unsupported expression' indicates that the query as such is valid, but cannot be used in the current context."
          },
          "httpStatus": {
            "type": "integer",
            "description": "The http status of the response. The status depends on 'backendStatus': 'unknown error' is mapped to 'BAD REQUEST' and all others are mapped to 'OK' "
          },
          "errorMessage": {
            "type": "string",
            "description": "A short description of the error. This can be 'null'"
          }
        },
        "required": [
          "successful",
          "backendStatus",
          "errorMessage"
        ]
      },
      "CachedSearchDescription": {
        "type": "object",
        "properties": {
          "creationTraceId": {
            "type": "string",
            "description": "The trace is which was used when creating the search. It starts with the value of the request header SWA-MDC-TOKEN which was used when creating the search. If the search was created without this header, it is auto-generated."
          }
        }
      },
      "ProjectsResult": {
        "type": "object",
        "properties": {
          "status": {
            "$ref": "#/components/schemas/StatusObject"
          },
          "numberResults": {
            "type": "integer",
            "format": "int64",
            "description": "The number of results found by the search. With paging this may be more than available in the result."
          },
          "results": {
            "type": "array",
            "description": "The list of projects.",
            "items": {
              "$ref": "#/components/schemas/Project"
            }
          }
        },
        "required": [
          "status",
          "numberResults",
          "results"
        ]
      },
      "Project": {
        "type": "object",
        "properties": {
          "id": {
            "type": "string",
            "description": "The project ID, without any type prefix. The ID is the one to be used where the API requires a 'projectId'."
          }
        },
        "required": [
          "id"
        ]
      },
      "ProjectResourcesResult": {
        "type": "object",
        "properties": {
          "status": {
            "$ref": "#/components/schemas/StatusObject"
          },
          "numberResults": {
            "type": "integer",
            "format": "int64",
            "description": "The number of results found by the search. With paging this may be more than available in the result."
          },
          "results": {
            "type": "array",
            "description": "The list of project resources.",
            "items": {
              "$ref": "#/components/schemas/ProjectResource"
            }
          }
        },
        "required": [
          "status",
          "numberResults",
          "results"
        ]
      },
      "ProjectResource": {
        "type": "object",
        "properties": {
          "id": {
            "type": "string",
            "enum": [
              "collections"
            ],
            "description": "The ID of the resource."
          },
          "description": {
            "type": "string",
            "description": "A description of the resource."
          }
        },
        "required": [
          "id",
          "description"
        ]
      },
      "CollectionsResult": {
        "type": "object",
        "properties": {
          "status": {
            "$ref": "#/components/schemas/StatusObject"
          },
          "numberResults": {
            "type": "integer",
            "format": "int64",
            "description": "The number of results found by the search. With paging this may be more than available in the result."
          },
          "results": {
            "type": "array",
            "description": "The list of collections.",
            "items": {
              "$ref": "#/components/schemas/Collection"
            }
          }
        },
        "required": [
          "status",
          "numberResults",
          "results"
        ]
      },
      "Collection": {
        "type": "object",
        "properties": {
          "id": {
            "type": "string",
            "description": "The sub collection ID, as defined in the meta engine configuration in 'Meta engine sources: Server definition'. The ID resembles the configuration column 'Identifier' and is to be used where the API requires a 'collectionId'. Use 'default' for single engines and merging meta engines."
          },
          "displayName": {
            "type": "string",
            "description": "The associated display name. Can be the same as id."
          }
        },
        "required": [
          "id"
        ]
      },
      "CollectionResourcesResult": {
        "type": "object",
        "properties": {
          "status": {
            "$ref": "#/components/schemas/StatusObject"
          },
          "numberResults": {
            "type": "integer",
            "format": "int64",
            "description": "The number of results found by the search. With paging this may be more than available in the result."
          },
          "results": {
            "type": "array",
            "description": "The list of collection resources.",
            "items": {
              "$ref": "#/components/schemas/CollectionResource"
            }
          }
        },
        "required": [
          "status",
          "numberResults",
          "results"
        ]
      },
      "CollectionResource": {
        "type": "object",
        "properties": {
          "id": {
            "type": "string",
            "enum": [
              "fields",
              "filters",
              "records"
            ],
            "description": "The ID of the resource. Typical resources available will be 'records' which provides access to the actual searchable objects, 'filters' (folder fields, taxonomies, smart filters) which can be used for filtering, etc."
          },
          "description": {
            "type": "string",
            "description": "A description of the resource."
          },
          "availableQueryParameters": {
            "type": "array",
            "description": "This provides a description of available query parameters ('...&\u003Cname\u003E=\u003Cvalue\u003E&...') for the resource (e.g. query expressions, paging requests, requested fields)",
            "items": {
              "$ref": "#/components/schemas/QueryParameter"
            }
          }
        },
        "required": [
          "id",
          "description",
          "availableQueryParameters"
        ]
      },
      "QueryParameter": {
        "type": "object",
        "properties": {
          "name": {
            "type": "string",
            "description": "The parameter name as used when setting it ('...&\u003Cname\u003E=\u003Cvalue\u003E&...')."
          },
          "description": {
            "type": "string",
            "description": "The description of the query parameter."
          }
        },
        "required": [
          "name",
          "description"
        ]
      },
      "FieldsResult": {
        "type": "object",
        "properties": {
          "status": {
            "$ref": "#/components/schemas/StatusObject"
          },
          "numberResults": {
            "type": "integer",
            "format": "int64",
            "description": "The number of results found by the search. With paging this may be more than available in the result."
          },
          "results": {
            "type": "array",
            "description": "The list of fields, including folder fields.",
            "items": {
              "$ref": "#/components/schemas/FieldDescription"
            }
          }
        },
        "required": [
          "status",
          "numberResults",
          "results"
        ]
      },
      "FieldDescription": {
        "type": "object",
        "properties": {
          "id": {
            "type": "string",
            "description": "The field ID. The ID is the one to be used where the API requires a field ID (or folder field ID, depending on the context)."
          },
          "type": {
            "type": "string",
            "enum": [
              "identifier",
              "content",
              "text",
              "integer",
              "float",
              "date",
              "singleValue",
              "multiValue",
              "binary",
              "hierarchy",
              "parameterizedLong",
              "texts"
            ],
            "description": "The data type of the field: identifier (text), content (text), text, integer, float, date, single-value taxonomy, multi-value taxonomy. \n\nThe type 'binary' is special in that it accepts and returns binary streams. A main use-case is to access Binary storages configured for a singleMindServer. To this end, use projectId=singleMindServer.NAME and collectionId=rm_storage:Native files for 'Native files' or collectionId=rm_storage:Image files to access image files. The syntax is 'rm_storage:' followed by the value configured as 'Storage file type' in the single mind server configuration. \n\nThe type 'hierarchy' allows to implement tree structures. Its value is a pointer to the nodes parent id. It is intended to be used in folder collections only. \n\nThe type 'parameterizedLong' is a multi-value type: a single document can have multiple values in such a field, and each value has a 'parameter' (a string) and a long value. Such fields are supposed to contain 'few' parameters, but many different long values. It is possible to insert/change values of the form '\u003Cparameter\u003E.\u003Clong\u003E' where '\u003Cparameter\u003E' is some string chosen by the API user (example: rank.1234). Assuming such a field is called ranks, the associated values are returned by the records endpoint with fields=ranks.A, ranks.B where 'A' and 'B' are parameters. It is possible to search for documents with such values by means of field bases searches of the form 'ranks.A \u003E 42' or 'ranks.B = 1234'. It is also possible to sort by these values: specify the sort field as 'ranks.A:desc' or 'ranks.B:asc'. Fields of type 'texts' contain a list of strings."
          },
          "isSortable": {
            "type": "boolean",
            "description": "True if the field can be used for sorting."
          },
          "isFolderField": {
            "type": "boolean",
            "description": "True if the field is a folder field (taxonomy, smart filter). A document can be associated with a folder in such a folder field (or with multiple folders, depending on folder field). A folder field can be used for filtering. Values (e.g. the document count) can be accumulated for folders."
          },
          "isMultivalueFolderCollection": {
            "type": "boolean",
            "description": "True if the field is a multi-value folder collection (taxonomy). A folder collection can be used as filter. Each document can be associated with any number of folders of such a folder collection."
          },
          "isPrefixSearchable": {
            "type": "boolean",
            "description": "True if a folder search on the field can be restricted by prefix."
          },
          "displayName": {
            "type": "string",
            "description": "A field display name. The value is never null, it falls back to the field id if there is no display name. Display names can be configured in the server configuration."
          }
        },
        "required": [
          "id",
          "type",
          "isSortable",
          "isFolderField",
          "isMultivalueFolderCollection",
          "isPrefixSearchable"
        ]
      },
      "FolderFieldResourcesResult": {
        "type": "object",
        "properties": {
          "status": {
            "$ref": "#/components/schemas/StatusObject"
          },
          "numberResults": {
            "type": "integer",
            "format": "int64",
            "description": "The number of results found by the search. With paging this may be more than available in the result."
          },
          "results": {
            "type": "array",
            "description": "The list of folder field (taxonomy, smart    filter) resources.",
            "items": {
              "$ref": "#/components/schemas/FolderFieldResource"
            }
          }
        },
        "required": [
          "status",
          "numberResults",
          "results"
        ]
      },
      "FolderFieldResource": {
        "type": "object",
        "properties": {
          "id": {
            "type": "string",
            "enum": [
              "values"
            ],
            "description": "The ID of the resource."
          },
          "description": {
            "type": "string",
            "description": "A description of the resource."
          },
          "availableQueryParameters": {
            "type": "array",
            "description": "This provides a description of available query parameters ('...&\u003Cname\u003E=\u003Cvalue\u003E&...') for the resource.",
            "items": {
              "$ref": "#/components/schemas/QueryParameter"
            }
          }
        },
        "required": [
          "id",
          "description",
          "availableQueryParameters"
        ]
      },
      "FolderValuesResult": {
        "type": "object",
        "properties": {
          "status": {
            "$ref": "#/components/schemas/StatusObject"
          },
          "numberResults": {
            "type": "integer",
            "format": "int64",
            "description": "The number of results found by the search. With paging this may be more than available in the result. NOTE: this number will be 2147483647 (max int) if paging through folder results of a merging meta engine. This is due to a performance optimization when merging results. In such a case, you will need to page forward until the system has received all results (in which case the final page will contain the correct numberResults)."
          },
          "results": {
            "type": "array",
            "description": "The list of folders (categories).",
            "items": {
              "$ref": "#/components/schemas/FolderRecord"
            }
          }
        },
        "required": [
          "status",
          "numberResults",
          "results"
        ]
      },
      "FolderRecord": {
        "type": "object",
        "properties": {
          "rank": {
            "type": "integer",
            "format": "int64",
            "description": "The list position of the record in the sorted list of folders. The first folder has rank 1."
          },
          "relevance": {
            "type": "number",
            "format": "float",
            "description": "The accumulated relevancy (between 0 and 1) of collection records in the search result relating to this folder."
          },
          "id": {
            "type": "string",
            "description": "The folder ID."
          },
          "displayName": {
            "type": "string",
            "description": "The display name of the folder. Generally, for most folders, this will be the same as the ID."
          },
          "count": {
            "type": "integer",
            "format": "int64",
            "description": "The number of documents associated with the given folder (category)."
          }
        },
        "required": [
          "rank",
          "relevance",
          "id",
          "displayName",
          "count"
        ]
      },
      "FolderPropertyNames": {
        "type": "object",
        "properties": {
          "fieldName": {
            "type": "string",
            "description": "The name of a field."
          },
          "propertyNames": {
            "type": "string",
            "description": "A comma-separated list of property names. A property is nothing but a field which exists in the collection referenced by 'fieldName'. Use the endpoint /projects/{projectId}/collections/{collectionId}/fields with {collectionId} set to the value of 'fieldName' in order to inspect the list of available property names."
          }
        },
        "required": [
          "fieldName",
          "propertyNames"
        ]
      },
      "FolderPropertiesRequestParameter": {
        "type": "array",
        "items": {
          "$ref": "#/components/schemas/FolderPropertyNames"
        },
        "description": "An array of requested folder property names."
      },
      "MeasureTypeParameter": {
        "required": [
          "typeName"
        ],
        "type": "object",
        "properties": {
          "typeName": {
            "type": "string",
            "enum": [
              "count",
              "sum",
              "max",
              "min",
              "avg",
              "predictive correlation"
            ],
            "description": "Defines the type of aggregate (measure) which is to be computed. The value 'count' is the default; it merely counts the number of documents falling into one of the buckets. The choice 'sum' expects a 'fieldName' of numeric type (integer, long, double, or date) and sums values per bucket. The choices 'max', 'min', 'avg' are similar to sum and aggregate the specified value per bucket. The choice 'predictive correlation' is special: for each bucket, it computes the correlation between two properties of an arbitrary document in the \u003Cuniverse\u003E. The first property is 'document falls into the bucket'. The second property is 'document has one of \u003CselectedFieldValues\u003E in \u003CfieldName\u003E'. Example: we define the measure dimension as folders in the phrase taxonomy (i.e. buckets are phrases). In addition, the predictive correlation is set \u003CfieldName\u003E=importance and \u003CselectedFieldValues\u003E is 'High' with \u003Cuniverse\u003E='importance=*'. In this case, the system will return phrases. For each phrase, it returns the correlation as double value between -1.0 and +1.0 . A value close to -1.0 indicates that the phrase in question is anti-correlated with 'importance=High'. A value close to +1.0 indicates that the phrase in question is correlated with 'importance=High'. A value close to 0.0 indicates that the phrase in question is independent of 'importance=High'.",
            "default": "count"
          },
          "fieldName": {
            "type": "string",
            "description": "The field to consider for the measure type in question. The type 'count' has no field. All other measure types require a field name."
          },
          "universe": {
            "type": "string",
            "description": "This argument is evaluated if and only if typeName='predictive correlation'. For 'predictive correlation', the semantics is as follows. It defines the documents to consider for the correlation. Use '*' if all documents are requested. Note that correlation values are only dependent on the 'universe', not on the 'query' used by the measure endpoint. The 'query' of the measure endpoint is merely used to select or sort the buckets in question. If buckets are prescribed, the 'query' has no effect at all."
          },
          "selectedFieldValues": {
            "type": "array",
            "description": "This argument is evaluated if and only if typeName='predictive correlation'. It defines one or more folder name(s) (the value of rm_folder_id) within the specified fieldName. At least one folder name must be given. If more than one folder name is given, the system applies an 'OR semantics'.",
            "items": {
              "type": "string"
            }
          }
        }
      },
      "SearchResult": {
        "required": [
          "numberResults",
          "results",
          "status"
        ],
        "type": "object",
        "properties": {
          "status": {
            "$ref": "#/components/schemas/StatusObject"
          },
          "numberResults": {
            "type": "integer",
            "format": "int64",
            "description": "The number of results found by the search. With paging this may be more than available in the result."
          },
          "results": {
            "type": "array",
            "description": "The list of records.",
            "items": {
              "$ref": "#/components/schemas/Record"
            }
          },
          "sponsoredLinks": {
            "type": "array",
            "description": "A list of matching sponsored links (if configured for the target project).",
            "items": {
              "$ref": "#/components/schemas/SponsoredLink"
            }
          },
          "spellingSuggestions": {
            "$ref": "#/components/schemas/SpellingSuggestionResult"
          }
        }
      },
      "SearchResultHighlight": {
        "type": "object",
        "properties": {
          "field": {
            "type": "string",
            "description": "The field in which this expression is to be applied."
          },
          "regularExpression": {
            "type": "string",
            "description": "The associated regular expression"
          }
        }
      },
      "SearchResultHighlightingResult": {
        "type": "object",
        "properties": {
          "results": {
            "type": "array",
            "description": "The list of returned highlights.",
            "items": {
              "$ref": "#/components/schemas/SearchResultHighlight"
            }
          }
        }
      },
      "HighlightResultEntity": {
        "type": "object",
        "properties": {
          "termsToHighlight": {
            "type": "array",
            "description": "The list of returned highlights.",
            "items": {
              "type": "string"
            }
          },
          "regularExpression": {
            "type": "string",
            "description": "a regular expression to identify all these words"
          },
          "numberHits": {
            "type": "integer",
            "format": "int32",
            "description": ""
          },
          "numberHitsByPage": {
            "type": "array",
            "description": "The number of hits by page.",
            "items": {
              "type": "integer",
              "format": "int32"
            }
          },
          "hitLocationsRel": {
            "type": "array",
            "description": "",
            "items": {
              "type": "number",
              "format": "double"
            }
          },
          "termToHighlightIsWordBoundaryBefore": {
            "type": "array",
            "description": "",
            "items": {
              "type": "boolean"
            }
          },
          "termToHighlightIsWordBoundaryAfter": {
            "type": "array",
            "description": "",
            "items": {
              "type": "boolean"
            }
          }
        }
      },
      "StoredSearchHighlightAggregation": {
        "type": "string",
        "enum": [
          "by field",
          "by folder"
        ]
      },
      "StoredSearchHighlightResult": {
        "type": "object",
        "properties": {
          "field": {
            "type": "string",
            "description": ""
          },
          "folder": {
            "type": "string",
            "description": "The folder (can be null if hits were aggregated by field only)"
          },
          "terms": {
            "$ref": "#/components/schemas/HighlightResultEntity"
          }
        }
      },
      "HighlightedWordResult": {
        "type": "object",
        "properties": {
          "searchTerms": {
            "$ref": "#/components/schemas/HighlightResultEntity"
          },
          "conceptTerms": {
            "$ref": "#/components/schemas/HighlightResultEntity"
          },
          "trainingTerms": {
            "$ref": "#/components/schemas/HighlightResultEntity"
          },
          "userTerms": {
            "$ref": "#/components/schemas/HighlightResultEntity"
          },
          "storedSearchHighlightingByCategory": {
            "type": "array",
            "description": "",
            "items": {
              "$ref": "#/components/schemas/StoredSearchHighlightResult"
            }
          }
        }
      },
      "SearchResultTokenResponse": {
        "type": "object",
        "properties": {
          "status": {
            "$ref": "#/components/schemas/StatusObject"
          },
          "token": {
            "type": "string",
            "description": "The search result token."
          },
          "numberResults": {
            "type": "integer",
            "format": "int64",
            "description": "The number of results found by the search."
          },
          "eol": {
            "type": "string",
            "description": "The date at which this token expires and will be deleted. It has to be touched or used before this timestamp in order to keep it alive. The format is 'yyyy-MM-dd'T'HH:mm:ss.SSSXXX' in UTC timezone. Note that a search token is bound to the user session as well: it will be closed if the user session times out or if the users logs out."
          }
        }
      },
      "SearchResultToken": {
        "type": "object",
        "properties": {
          "token": {
            "type": "string",
            "description": "The search result token."
          }
        }
      },
      "MeasureCube": {
        "type": "object",
        "properties": {
          "dimensions": {
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/MeasureDimensionResult"
            }
          },
          "values": {
            "type": "array",
            "description": "The matrix of values. For one-dimension aggregates, the matrix has just one row. This member is populated if the result resembles a long integer. It will be 'null' if the value is a double. A value is long if the underlying field type is long and the measureType results in long values.",
            "items": {
              "type": "array",
              "items": {
                "type": "integer",
                "format": "int64"
              }
            }
          },
          "valuesDouble": {
            "type": "array",
            "description": "The matrix of values. For one-dimension aggregates, the matrix has just one row. This member is populated if the result resembles a floating point number. It will be 'null' if the value is a long. A value is double if the underlying field type is double or the measureType results in double values.",
            "items": {
              "type": "array",
              "items": {
                "type": "number",
                "format": "double"
              }
            }
          }
        }
      },
      "MeasureDimensionResult": {
        "type": "object",
        "properties": {
          "size": {
            "type": "integer",
            "description": "The size of this dimension. For one-dimensional counts, this resembles all available results and can be larger than the elements in 'members' (due to paging). For two-dimensional counts, this value is the same as the size of 'members'.",
            "format": "int64"
          },
          "fieldName": {
            "type": "string",
            "description": "the associated field name"
          },
          "documentsWithAnyValue": {
            "type": "integer",
            "description": "Documents with any value in this dimension. Negative if this value is unsupported.",
            "format": "int64"
          },
          "documentsWithNoValue": {
            "type": "integer",
            "description": "Documents without any value in this dimension. Negative if this value is unsupported.",
            "format": "int64"
          },
          "members": {
            "type": "array",
            "description": "The dimension members requested by the request.",
            "items": {
              "$ref": "#/components/schemas/MeasureDimensionMember"
            }
          }
        }
      },
      "MeasureDimensionMember": {
        "type": "object",
        "properties": {
          "identifier": {
            "type": "string",
            "description": "The member's identifier."
          },
          "fields": {
            "type": "array",
            "description": "A list of field-value pairs, if any requested.",
            "items": {
              "$ref": "#/components/schemas/Field"
            }
          }
        }
      },
      "Record": {
        "type": "object",
        "properties": {
          "rank": {
            "type": "integer",
            "format": "int64",
            "description": "The list position of the record in the list of found records. The first record has rank 1."
          },
          "relevance": {
            "type": "number",
            "format": "float",
            "description": "Calculated relevance (between 0 and 1) corresponding to how well content and properties match the search expression."
          },
          "id": {
            "type": "string",
            "description": "The ID of the record (document)."
          },
          "uniqueField": {
            "type": "string",
            "description": "A human-readable unique field. This field should only be used for logging or investigation, there is no guarantee about the value."
          },
          "fields": {
            "type": "array",
            "description": "A list of field-value pairs, if any requested.",
            "items": {
              "$ref": "#/components/schemas/Field"
            }
          },
          "folderSets": {
            "type": "array",
            "description": "A list of sets of folder-value pairs (also including the folder display name), if any requested.",
            "items": {
              "$ref": "#/components/schemas/FolderSet"
            }
          },
          "body": {
            "type": "string",
            "description": "Content or part of the content of the record (document). This can be 'null'."
          },
          "page": {
            "type": "integer",
            "description": "Shows the page number of the current result. Only populated if body for single documents has been retrieved (not in the result listing of a search)."
          },
          "pageCount": {
            "type": "integer",
            "description": "Shows the number of pages in the document. Only populated if body for single documents has been retrieved (not in the result listing of a search)."
          }
        },
        "required": [
          "rank",
          "relevance",
          "id",
          "uniqueField",
          "fields",
          "folderSets",
          "body",
          "page",
          "pageCount"
        ]
      },
      "Field": {
        "type": "object",
        "properties": {
          "id": {
            "type": "string",
            "description": "The field ID."
          },
          "value": {
            "type": "string",
            "description": "The value for the given document in the given field or null if there is no such value. Date values are formatted as string according to ISO_8601, i.e. in the date format 'yyyy-MM-dd'T'HH:mm:ss.SSSXXX' in UTC timezone. Folder ids are formatted as comma-separated list without any escape characters (see valueObject for a type-safe access)."
          },
          "valueObject": {
            "type": "object",
            "properties": {

            },
            "description": "This is either null or redundant: it provides the very same data as in 'value', but as typed object. This property is populated if and only if the result is not a plain string. Examples where this property is populated and contains useful values are: integer fields, date fields (which return the number of milliseconds since January 1, 1970, 00:00:00 GMT, formatted as string) or an array of folder values (if they have been requested as 'fields'). Note that 'value' will always be populated with a string-representation of the value. Use 'valueObject' in order to have type-safe access (especially if you retrieved a list of folder ids)."
          }
        },
        "required": [
          "id",
          "value"
        ]
      },
      "FolderSet": {
        "type": "object",
        "properties": {
          "id": {
            "type": "string",
            "description": "The folder field ID."
          },
          "value": {
            "type": "array",
            "description": "The value(s) for the given document in the given field.",
            "items": {
              "$ref": "#/components/schemas/Folder"
            }
          }
        },
        "required": [
          "id",
          "value"
        ]
      },
      "Folder": {
        "type": "object",
        "properties": {
          "id": {
            "type": "string",
            "description": "The folder ID."
          },
          "displayName": {
            "type": "string",
            "description": "The display name of the folder. Generally, for most folders, this will be the same as the ID."
          },
          "properties": {
            "type": "array",
            "description": "A list of field-value pairs, if any requested.",
            "items": {
              "$ref": "#/components/schemas/Field"
            }
          }
        },
        "required": [
          "id",
          "displayName"
        ]
      },
      "SponsoredLink": {
        "type": "object",
        "properties": {
          "externalLink": {
            "type": "string",
            "description": "A link into some external system (typically HTTP). Can be null."
          },
          "recordId": {
            "type": "string",
            "description": "If present, this is the record id associated with the link. It is present if and only if the link resembles a record id in the collection."
          },
          "title": {
            "type": "string",
            "description": "Link text describing the result."
          },
          "description": {
            "type": "string",
            "description": "A detailed description of the result."
          },
          "relevance": {
            "type": "number",
            "format": "float",
            "description": "The confidence for this link (between 0 and 1)."
          }
        }
      },
      "SpellingSuggestionResult": {
        "type": "object",
        "description": "A list of matching spelling suggestions (if configured for the target project). The return value 'null' means that such suggestions have not been requested.",
        "properties": {
          "totalImprovementPercent": {
            "type": "number",
            "format": "float",
            "description": "The total improvement in percent if all suggestions are applied (between 0 and 1). Note that this number is unrelated to the individual improvements of suggested words as it also takes those words into account for which no suggestion has been returned."
          },
          "suggestedWords": {
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/SpellingSuggestion"
            }
          }
        }
      },
      "SpellingSuggestion": {
        "type": "object",
        "properties": {
          "originalWord": {
            "type": "string",
            "description": "The word is it was found in the original search request."
          },
          "suggestedWord": {
            "type": "string",
            "description": "The suggested alternative word."
          },
          "improvementPercent": {
            "type": "number",
            "format": "float",
            "description": "The improvement in percent achieved by applying the suggestion (between 0 and 1), but limited only to this word."
          },
          "isOriginalFound": {
            "type": "boolean",
            "description": "Will be true if and only if the originalWord was found in some record."
          }
        }
      },
      "RecordResourcesResult": {
        "type": "object",
        "properties": {
          "status": {
            "$ref": "#/components/schemas/StatusObject"
          },
          "numberResults": {
            "type": "integer",
            "format": "int64",
            "description": "The number of results found by the search. With paging this may be more than available in the result."
          },
          "results": {
            "type": "array",
            "description": "The list of record (document) resources.",
            "items": {
              "$ref": "#/components/schemas/RecordResource"
            }
          }
        },
        "required": [
          "status",
          "numberResults",
          "results"
        ]
      },
      "RecordResource": {
        "type": "object",
        "properties": {
          "id": {
            "type": "string",
            "enum": [
              "content"
            ],
            "description": "The ID of the resource."
          },
          "description": {
            "type": "string",
            "description": "A description of the resource."
          },
          "availableQueryParameters": {
            "type": "array",
            "description": "This provides a description of available query parameters ('...&\u003Cname\u003E=\u003Cvalue\u003E&...') for the resource.",
            "items": {
              "$ref": "#/components/schemas/QueryParameter"
            }
          }
        },
        "required": [
          "id",
          "description",
          "availableQueryParameters"
        ]
      },
      "LogoutResult": {
        "required": [
          "status"
        ],
        "type": "object",
        "properties": {
          "status": {
            "$ref": "#/components/schemas/StatusObject"
          }
        }
      },
      "LoginResult": {
        "required": [
          "status"
        ],
        "type": "object",
        "properties": {
          "status": {
            "$ref": "#/components/schemas/StatusObject"
          }
        }
      },
      "ChangeResult": {
        "required": [
          "status"
        ],
        "type": "object",
        "properties": {
          "status": {
            "$ref": "#/components/schemas/StatusObject"
          }
        }
      },
      "WaitForPendingChangesResult": {
        "required": [
          "status"
        ],
        "type": "object",
        "properties": {
          "success": {
            "description": "Returns true if all currently scheduled transactions finished within the timeout.",
            "type": "boolean"
          }
        }
      },
      "DeletionMode": {
        "type": "string",
        "enum": [
          "delete",
          "softdelete"
        ],
        "description": "Allows to select how deletions should be applied. The mode 'delete' deletes a category (no undo). The mode 'softdelete' merely marks the category for deletion ('hidden') such that it is gone for all practical purposes (except background maintenance tasks).",
        "default": "delete"
      },
      "InsertRemoveRequest": {
        "type": "object",
        "properties": {
          "newRecords": {
            "type": "array",
            "description": "0-n records to be created.",
            "items": {
              "$ref": "#/components/schemas/RecordData"
            }
          },
          "recordsToDelete": {
            "type": "array",
            "description": "0-n record ids which are to be deleted. Specifying a record as both 'delete' and 'new' has re-insert semantics. The specified ids have to be in one of the following formats: 1. they can resemble the 'id' as returned by the records endpoint (aka url-safe string). 2. they can be a folder id (only applicable if you are connected against a folder collection). 3. If connected with a folder collection, you can optionally specify folderId::VALUE here where VALUE is a folder id.",
            "items": {
              "type": "string"
            }
          },
          "recordsToDeleteByQuery": {
            "$ref": "#/components/schemas/SearchRequest"
          },
          "deletionMode": {
            "$ref": "#/components/schemas/DeletionMode"
          }
        }
      },
      "StartTransactionRequest": {
        "type": "object",
        "properties": {
          "dataSourceId": {
            "type": "string",
            "description": "The name of the data source. It is used to track the transaction."
          },
          "indexingBufferId": {
            "type": "string",
            "description": "An optional buffer id."
          }
        }
      },
      "JobStatusResponse": {
        "type": "object",
        "properties": {
          "jobStatus": {
            "$ref": "#/components/schemas/JobStatus"
          }
        }
      },
      "JobStatus": {
        "type": "string",
        "enum": [
          "RUNNING",
          "STOPPED",
          "STOPPING",
          "PAUSED",
          "PAUSING",
          "FATAL",
          "FINISHING",
          "FINISHED",
          "SCHEDULED",
          "WAITING",
          "CREATED"
        ]
      },
      "FinishTransactionRequest": {
        "type": "object",
        "properties": {
          "numberDocuments": {
            "type": "integer",
            "format": "int64",
            "description": "The number of documents"
          },
          "analyzedSizeGb": {
            "type": "integer",
            "format": "int64",
            "description": "The size of documents which has been analyzed"
          }
        }
      },
      "FinishTransactionResponse": {
        "type": "object",
        "properties": {
          "jobId": {
            "type": "string"
          }
        }
      },
      "StartTransactionResult": {
        "type": "object",
        "properties": {
          "indexingBufferId": {
            "type": "string",
            "description": "The buffer id to use for all subsequent operations."
          }
        }
      },
      "RecordData": {
        "type": "object",
        "properties": {
          "uniqueId": {
            "type": "string",
            "description": "A unique identifier for the record. It resembles the value of 'uniqueField' when fetching records and must be provided at insertion time."
          },
          "fieldData": {
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/FieldData"
            }
          }
        }
      },
      "SearchRequest": {
        "type": "object",
        "description": "A query.",
        "properties": {
          "query": {
            "type": "string",
            "description": "The query expression. Such a query expression (also known as 'main query') contains simple word matches like 'York', implicit phrases like 'New York', explicit phrases like '\"An explicit phrase\"', fielded searches, boolean expression like 'berlin and vacation', containing near operators etc. Multiple expressions can be submitted with AND semantics (can also be seen as a 'search in' of the second expression in the first expression). Field-based expressions must respect the fields configured for the selected projectId/collectionId. Specifying this value as GET parameter requires the usual URI encoding. Note that field-based-searches for date fields support '\u003C', '=', '\u003E', '\u003C=', '\u003E=' just as for numeric values; the accepted values are either the milli seconds since 1970 followed by 'L' or a (partial) date pattern as specified in the singleMindServer configuration. Please refer to the reference manual for more details. Use 'NOT *' if you want to match no documents. Default query is '*', i.e. no query restriction. "
          },
          "language": {
            "type": "string",
            "description": "Specify a language to be used when interpreting the query, default is to use language recognition."
          },
          "joinRestriction": {
            "type": "string",
            "description": "Specify a restriction using a joined collection. The argument is expected to be of the form \u003CtargetFieldID\u003E:\u003Cquery\u003E. Example: doc_author_id:office=\"New York\" fired against a documents collection might restrict the result to all documents for which the author has an office in New York. The \u003CtargetFieldID\u003E id refers to a field which is the target of a configured join. Joins are configured in the 'Dynamic join' section of the meta-engine configuration. The \u003Cquery\u003E is the query to be run as a restriction on the source collection in order to limit the results on the target (should be URI-encoded). Some remarks: 1. In the case of using the query for a filter value request: Since dynamic join restrictions influence the result set, and hence the counts, they should be supplied independent of the filter selected. 2. A requests of sorts { joinRestriction=doc_author_id:office=A, query=york } is equivalent to query=york AND doc_author_id.find(office=A), i.e. joinRestriction=\u003CtargetFieldID\u003E:\u003Cquery\u003E and the query syntax \u003CtargetFieldID\u003Efind(\u003Cquery\u003E) are equivalent. 3. For meta engine projects, \u003CtargetFieldID\u003E is valid independent of the collectionId (i.e. you can use collectionId=people and a \u003CtargetFieldID\u003E which is in the people collection)."
          }
        }
      },
      "FieldData": {
        "type": "object",
        "properties": {
          "fieldName": {
            "type": "string",
            "description": "The field name. A common use-case is to create new folders. In this case, rm_display_name is the field for the folder's display name. The field rm_prop_parent is the field defining the new record's parent (it resembles the parent's unique id)."
          },
          "value": {
            "type": "string",
            "description": "The value to assign to this field. Only one of 'value' or 'valueList' can be populated. If this field data is of type binary, the 'value' must be the index of the binary artifact, i.e. the index within the multipart upload form."
          },
          "valueList": {
            "description": "A list of values to assign to this field (for multi-value fields only). Only one of 'value' or 'valueList' can be populated.",
            "type": "array",
            "items": {
              "type": "string"
            }
          }
        }
      },
      "InsertRemoveResult": {
        "type": "object",
        "properties": {
          "status": {
            "$ref": "#/components/schemas/StatusObject"
          }
        }
      },
      "ChangeRequest": {
        "type": "object",
        "properties": {
          "field": {
            "description": "The field to change.",
            "type": "string"
          },
          "type": {
            "description": "The type of change (case sensitive!). The semantics is as follows:\n * ADD_FOLDERS tags all documents in scope into the folders specified in folderIds.\n * SET_FOLDERS ensures that all documents in scope have exactly the folders specified in folderIds.\n * REMOVE_FOLDERS untags all documents in scope from the folders specified in folderIds.\n * REMOVE_FOLDERS_AND_BRANCH is similiar to removeFolders. However, for a hierarchical folder collection with enabled 'full hierarchy', parent associations will also be removed unless another child references the documents.\n * MAKE_ROOT requires that the identified collection is a hierarchical folder collection (=taxonomy), and that field is set to 'rm_prop_parent'. In this case, it ensures that the documents in scope (which are folder ids) will become top-level folders (example: projectId='singleMindServer.project', collectionId='my_workspace', field='rm_prop_parent', type='MAKE_ROOT', recordId='unencoded category id').\n * SET_PARENT has the same requirements as MAKE_ROOT, but it also expects exactly one argument in 'folderIds' (the new parent). It changes the hierarchy of the specified categories.\n * SET_TEXT changes the content to the value specified in 'text' in the specified 'field' for all documents in scope. This is also possible for numeric fields or date fields. Note that date fields can be changed to a string resembling the milli seconds since 1970 suffixed with 'L' or a date pattern which must be configured in the singleMindServer configuration.\n * APPEND_TEXT is similar to SET_TEXT, but it appends to the existing text (if any). This is impossible for numeric or date fields.\n * SET_TEXTS can be used to modify folder properties by assigning a list of strings. It expects the value in 'texts'.\n Note that changes for content fields necessarily require to use content-type multipart/form-data combined with type=SET_TEXT and value an integer index referencing the associated binary XML stream.",
            "type": "string",
            "enum": [
              "ADD_FOLDERS",
              "SET_FOLDERS",
              "REMOVE_FOLDERS",
              "REMOVE_FOLDERS_AND_BRANCH",
              "MAKE_ROOT",
              "SET_PARENT",
              "SET_TEXT",
              "SET_TEXTS",
              "APPEND_TEXT",
              "SET_CONTENT"
            ]
          },
          "folderIds": {
            "description": "Zero, one, or more folder ids. This value is evaluated for the following choices of 'type': 'ADD_FOLDERS', 'SET_FOLDERS', 'REMOVE_FOLDERS', 'REMOVE_FOLDERS_AND_BRANCH', 'SET_PARENT'. A folder id can be one of two possible formats: either it is the (unencoded) folder id as returned by other software components or it is the 'id' returned by the records endpoint when searching in the folder collection.",
            "type": "array",
            "items": {
              "type": "string"
            }
          },
          "text": {
            "description": "A string text. This value is evaluated for the following choices of 'type': 'SET_TEXT', 'APPEND_TEXT'. The type 'SET_TEXT' without a 'text' clears the field. Note that changes in a numeric field also use this parameter. If 'pages' is a configured numeric type, you can use field=pages, text='42', type='SET_TEXT'",
            "type": "string"
          },
          "texts": {
            "description": "A sequence of strings. This value is evaluated for the following choices of 'type': 'SET_TEXTS'.",
            "type": "array",
            "items": {
              "type": "string"
            }
          }
        },
        "required": [
          "field",
          "type"
        ]
      },
      "DimensionRequest": {
        "type": "object",
        "properties": {
          "field": {
            "type": "string",
            "description": "The field defining the dimension."
          },
          "restrictFoldersByQuery": {
            "type": "string",
            "description": "Only for folder fields: optional query to constrain the folders returned. Examples are rm_folder_id=uniquevalue or rm_display_name=displayname"
          },
          "restrictValuesByList": {
            "type": "array",
            "description": "A list of values. For folder fields, the values are folder ids. For fields of numeric or date type, each value can have one of two forms: the inclusive variant '[start,end]' or the exclusive variant '[start,end[' . For DATE fields, both start and end are either milli seconds since 1970 in UTC time followed by the suffix 'L' or a date format in one of the formats 'yyyy-MM-dd'T'HH:mm:ss.SSSXXX', 'yyyy-MM-dd'T'HH:mm:ss.SSS', 'yyyy-MM-dd'T'HH:mm:ss', 'yyyy-MM-dd'T'HH:mm', 'yyyy-MM-dd'T'HH', 'yyyy-MM-dd', 'yyyy-MM', 'yyyy' . Example for a numeric field: [4,5] which counts documents having either value 4 and the value 5. Another example for a numeric field: [4,5[ counts documents having value 4 but not 5. Example for a date field: [2000, 2000-10-01] counts documents with value between 2000-01-01 and 200-10-01 23:59:59.999 (inclusive).    ",
            "items": {
              "type": "string"
            }
          },
          "returnEmptyMembers": {
            "type": "boolean",
            "description": "Defines if members with count=0 should be returned. Default is false."
          },
          "folderOrder": {
            "type": "string",
            "description": "Only for folder fields: provide an order criterion for the results. The syntax is the same as 'order' for the records endpoint, i.e. \u003CfieldName\u003E or \u003CfiledName\u003E:asc or \u003CfieldName\u003E:desc. In this context, \u003CfieldName\u003E is any folder property. In addition, \u003CfieldName\u003E can be one of the special names 'count', 'relevance', or 'name'. Here, 'name' is the category display name."
          },
          "returnedFields": {
            "type": "string",
            "description": "A comma-separated list of field names. It is only used for folder fields and allows to fetch properties of the returned folders (like rm_display_name)."
          },
          "page": {
            "type": "integer",
            "description": "Define which page to retrieve (counting starts at 1). Default is to return all matching members. You cannot specify both 'page' and 'offset'.",
            "format": "int32"
          },
          "pageSize": {
            "type": "integer",
            "description": "Defines the page size. Only used if 'page' or 'offset' is set.",
            "format": "int32"
          },
          "offset": {
            "type": "integer",
            "description": "Define which is the first element to return (first is 0). The system will return the next 'pageSize' following items. Default is to return all matching members. You cannot specify both 'page' and 'offset'.",
            "format": "int32"
          }
        }
      }
    },
    "responses": {
      "200": {
        "description": "Successful operation",
        "content": {

        }
      },
      "400": {
        "description": "Invalid input",
        "content": {

        }
      },
      "500": {
        "description": "Server error",
        "content": {

        }
      }
    },
    "parameters": {
      "ProjectParameter": {
        "name": "projectId",
        "in": "path",
        "description": "Identifies the MindServer to connect to. It can be a fully qualified MindServer identifier like singleMindServer.ABCD or mergingMeta.DEF . The engine type prefix (like 'singleMindServer.') can omitted in which case the API will automatically test engine prefixes until it finds a matching MindServer. It tests the following engine type prefixes (in this order): 'metaMindServer.', 'mergingMeta.', 'singleMindServer.' . If in doubt, the correct project to connect to should be the ID of the topmost engine, the one directly below the corresponding application.",
        "required": true,
        "schema": {
          "type": "string"
        }
      },
      "CollectionParameter": {
        "name": "collectionId",
        "in": "path",
        "description": "Defines the collection within the specified project. For a Decisiv Search project this corresponds to the tabs in the UI, eg. 'documents', 'people', 'matter'. In this case, valid choices are configured in the column 'Identifier' in table 'Meta engine configuration: Meta engine sources: Server definitions'. In order to connect to single engines and merging meta engines, you have to specify the string 'default'. If you want to access (properties of) a folder collection (=taxonomy), set collectionId to the folder collection's name. If you need to combine both a meta engine server identifier and a folder collection name, separate them by means of '::'.  Note that single engines provide access to any configured binary storages (SingleMindServer configuraration: Native files: storage). To this end, use collectionId=rm_storage:FILETYPE where FILETYPE is the value configured as 'Storage file type' (like 'Native files', 'Image files').",
        "required": true,
        "schema": {
          "type": "string"
        }
      },
      "FolderFieldParameter": {
        "name": "fieldId",
        "in": "path",
        "description": "Folder field (taxonomy, smart filter) for which resources should be retrieved.",
        "required": true,
        "schema": {
          "type": "string"
        }
      },
      "BlockUntilCompleteParameter": {
        "name": "blockUntilComplete",
        "in": "query",
        "description": "Specifies whether the call should block until all indexes are updated. The default value 'false' is to return immediately which means that the display of values is always correct, but search indices may be outdated until all indices have been updated eventually (that means: queries will not respect the change until the index is updated). The change is persisted as soon as the API call returns, irrespective of 'blockUntilComplete'.",
        "schema": {
          "type": "boolean"
        }
      },
      "RecordIdParameter": {
        "name": "recordId",
        "in": "path",
        "description": "An id uniquely qualifying a single record. It resembles the 'id' returned by the records endpoint. Note that this id is url-safe and needs no additional encoding.",
        "required": true,
        "schema": {
          "type": "string"
        }
      },
      "QueryParameter": {
        "name": "query",
        "in": "query",
        "description": "The query expression. Such a query expression (also known as 'main query') contains simple word matches like 'York', implicit phrases like 'New York', explicit phrases like '\"An explicit phrase\"', fielded searches, boolean expression like 'berlin and vacation', containing near operators etc. Multiple expressions can be submitted with AND semantics (can also be seen as a 'search in' of the second expression in the first expression). Field-based expressions must respect the fields configured for the selected projectId/collectionId. Specifying this value as GET parameter requires the usual URI encoding. Note that field-based-searches for date fields support '\u003C', '=', '\u003E', '\u003C=', '\u003E=' just as for numeric values; the accepted values are either the milli seconds since 1970 followed by 'L' or a (partial) date pattern as specified in the singleMindServer configuration. Please refer to the reference manual for more details. Use 'NOT *' if you want to match no documents. Default query is '*', i.e. no query restriction. ",
        "schema": {
          "type": "string"
        }
      },
      "MeasureTypeParameter": {
        "name": "measureType",
        "in": "query",
        "description": "Defines the aggregate type. The default is to compute counts (i.e. a MeasureTypeParameter with typeName 'count'). All other arguments need a JSON representation of MeasureTypeParameter (url encoded json).",
        "content": {
          "application/json": {
            "schema": {
              "$ref": "#/components/schemas/MeasureTypeParameter"
            }
          }
        }
      },
      "FolderPropertiesRequestParameter": {
        "name": "folderFieldsWithProperties",
        "in": "query",
        "description": "An (exclusive) alternative to 'folderFields': it also allows to define folder fields to fetch, but it has the ability to specify properties to fetch as well. Note that the display name is always returned. All other arguments need a JSON representation of FolderPropertiesRequestParameter (url encoded json). If you leave this empty (or do not specify the parameter), the system will fetch folder ids and folder display names for all elements specified in in 'folderFields'. If you specify elements here, it is required that 'folderFields' is empty.",
        "content": {
          "application/json": {
            "schema": {
              "$ref": "#/components/schemas/FolderPropertiesRequestParameter"
            }
          }
        }
      },
      "SearchCacheControl": {
        "in": "header",
        "name": "SWA-searchCacheControl",
        "description": "Status of the search result cache.",
        "schema": {
          "type": "string",
          "enum": [
            "enabled",
            "disabled"
          ],
          "default": "enabled"
        }
      },
      "FieldListParameter": {
        "name": "fields",
        "in": "query",
        "description": "An URI-encoded, comma-separated list of fields to be retrieved for each document. Valid field names are typically those which can be inspected using the fields endpoint (i.e. those configured in the data model). In addition, search results support to retrieve dynamic properties as field name: 'rm_is_best_bet' returns true if and only if 'Best bets boosting' is active and a match is a best bet.",
        "schema": {
          "type": "string"
        }
      },
      "FolderFieldListParameter": {
        "name": "folderFields",
        "in": "query",
        "description": "An URI-encoded, comma-separated list of folder fields (taxonomies) to be retrieved for each document. Note that folder fields can be returned by 'fields' as well in which case an array of folder ids is returned. Use 'folderFields' if you need the folder and its display name.",
        "schema": {
          "type": "string"
        }
      },
      "LanguageParameter": {
        "name": "language",
        "in": "query",
        "description": "Specify a language to be used when interpreting the query, default is to use language recognition.",
        "schema": {
          "type": "string"
        }
      },
      "JoinRestriction": {
        "name": "joinRestriction",
        "in": "query",
        "description": "Specify a restriction using a joined collection. The argument is expected to be of the form \u003CtargetFieldID\u003E:\u003Cquery\u003E. Example: doc_author_id:office=\"New York\" fired against a documents collection might restrict the result to all documents for which the author has an office in New York. The \u003CtargetFieldID\u003E id refers to a field which is the target of a configured join. Joins are configured in the 'Dynamic join' section of the meta-engine configuration. The \u003Cquery\u003E is the query to be run as a restriction on the source collection in order to limit the results on the target (should be URI-encoded). Some remarks: 1. In the case of using the query for a filter value request: Since dynamic join restrictions influence the result set, and hence the counts, they should be supplied independent of the filter selected. 2. A requests of sorts { joinRestriction=doc_author_id:office=A, query=york } is equivalent to query=york AND doc_author_id.find(office=A), i.e. joinRestriction=\u003CtargetFieldID\u003E:\u003Cquery\u003E and the query syntax \u003CtargetFieldID\u003Efind(\u003Cquery\u003E) are equivalent. 3. For meta engine projects, \u003CtargetFieldID\u003E is valid independent of the collectionId (i.e. you can use collectionId=people and a \u003CtargetFieldID\u003E which is in the people collection).",
        "schema": {
          "type": "string"
        }
      }
    }
  }
}
```