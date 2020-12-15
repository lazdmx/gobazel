package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lazdmx/gobazel/server/httpModel/computeRequest"
	"github.com/lazdmx/gobazel/server/httpModel/computeResponse"
	"github.com/xeipuuv/gojsonschema"
)

func main() {
	router := httprouter.New()
	router.POST("/compute", Compute)

	log.Print("run and listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Compute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	validator, err := gojsonschema.NewSchema(gojsonschema.NewBytesLoader(computeRequest.JSONSchema))
	if err != nil {
		panic(err)
	}

	var reqBody computeRequest.T
	var resBody computeResponse.T

	if r.Body != nil {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		res, err := validator.Validate(gojsonschema.NewBytesLoader(data))
		if err != nil {
			panic(err)
		}

		if res.Valid() {
			if err := json.Unmarshal(data, &reqBody); err != nil {
				panic(err)
			}

			switch reqBody.Op {
			case "add":
				resBody.Data = reqBody.Lh + reqBody.Rh
		
			case "sub":
				resBody.Data = reqBody.Lh - reqBody.Rh
		
			default:
				panic("illegal state")
			}
		
		} else {
			for _, desc := range res.Errors() {
				log.Printf("- %s\n", desc)
			}

			resBody.Error = "bad_request"
		}
	}


	data, err := json.Marshal(&resBody)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(data))
		return
	}
	panic(err)
}
