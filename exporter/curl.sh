#! bash
arg1=$1
curl -X POST -H "Content-Type: application/json" -d "{\"value\": $arg1}" localhost
