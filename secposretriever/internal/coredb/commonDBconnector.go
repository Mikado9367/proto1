package coredb

import "secposretriever/config"

// credit : https://stackoverflow.com/questions/28751402/how-can-i-share-database-connection-between-packages-in-go
var DBconnector *config.DBstruct
