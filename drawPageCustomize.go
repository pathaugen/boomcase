
package main

import (
    //"fmt"
    //"net/http"
    //"net/url"
    
    "golang.org/x/net/context"
    
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/user"
	"google.golang.org/appengine/image"
	//"appengine/blobstore"
    
    //"os"
    //"io/ioutil"
    "strings"
	"strconv"
)


// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
func drawPageCustomize(ctx context.Context, output string) (string) { // context.Context vs appengine.Context
    
    
	//output = strings.Replace(output, "<CASE>", drawCase(r), -1)
	
	
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
	// Query the datastore and get all the cases
	
	// Array to hold the results
	var caseArray []Case
	//caseArray := make([]Case, 100)
	
	// Datastore query
	q := datastore.NewQuery("Case") //.Ancestor(caseKey(c)).Order("-Date").Limit(10)
	keys, err := q.GetAll(ctx, &caseArray)
	if err != nil { /*log.Errorf(ctx, "fetching case: %v", err);return*/ /*http.Error(w, err.Error(), http.StatusInternalServerError);return*/  }
	
	
	// ========== ========== ========== ========== ==========
	// Array to hold the results as data in a wrapper (res.Data.Name) with Key (res.Key) and Id (res.Id)
	//res := make([]Return, len(keys))
	//var res []Return
	//res := make([]Return, 100)
	
	/*
	// Migrate datastore retults to Results struct
	for i, c := range caseArray {
		// Do something with Case c (c.Name) and Key k
		k := keys[i]
		y := Return { Key: k, Id: uint64(k.IntID()), Data: c, }
		res = append(res, y)
	}
	// Loop through results and generate HTML
	outputCases := ""
	for i, c := range res {
		outputCases += `
			<a href="/case/1" class="case-block">
				<img src="/caseuploads/blankcase1.jpg" />
				<span class="name-container">`+strconv.Itoa(i)+` && `+strconv.Itoa(int(c.Id))+` && `+c.Data+`</span>
				<!-- <span class="watt-container"><span class="watts">WATTS</span> Watt BoomCase</span> -->
				<span class="price-container">$<span class="price">PRICE</span></span>
			</a>
		`
		//outputCases += string(res[i].Data.Name)
	}
	*/
	// ========== ========== ========== ========== ==========
	
	
	// ========== ========== ========== ========== ==========
	outputCases := ""
	for i, c := range caseArray {
		key := keys[i]
		id := int64(key.IntID())
		
		// Pull the image from blobstore
		/*
		var caseArray []Case
		q := datastore.NewQuery("Case")
		keys, err := q.GetAll(ctx, &caseArray)
		if err != nil {  }
		*/
		//blobstore.Delete(ctx, appengine.BlobKey(c.BlobKey))
		
		// ========== ========== ========== ========== ==========
		// Get a thumbnail of a blobstore image
		// https://github.com/golang/appengine/blob/master/image/image.go
		thumbOpts := image.ServingURLOptions { Size: 200, }
		thumbKey := appengine.BlobKey(c.BlobKey)
		thumbnail, _ := image.ServingURL(ctx, thumbKey, &thumbOpts) // (*url.URL, error)
		// ========== ========== ========== ========== ==========
		
		outputCases += `
			<a href="/case/`+strconv.Itoa(int(id))+`" class="case-block">
				<!-- <img src="/serve/?blobKey=`+c.BlobKey+`" /> -->
				<img src="`+thumbnail.String()+`" />
				<span class="name-container">`+c.Name+`</span>
				<!-- <span class="watt-container"><span class="watts">`+strconv.Itoa(int(c.Watts))+`</span> Watt BoomCase</span> -->
				<span class="">Customize this Case</span>
				<span class="price-container">$<span class="price">`+strconv.Itoa(int(c.Price))+`</span></span>
			</a>
		`
	}
	// ========== ========== ========== ========== ==========
	
	
	//outputCases += `<div>len(keys):[`+string(len(keys))+`]</div>`
	output = strings.Replace(output, "<CASE>", outputCases, -1)
	// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
	
	
	// ========== ========== ========== ========== ==========
	// [START if_user]
	u := user.Current(ctx)
	if u != nil && adminEmails[u.Email] {
		formCaseButton := `
			<div style="clear:both;">
				<a href="" class="btn btn-primary" id="admin-add-case" aria-label="ADMIN: Add New Custom Case">
					<i class="fa fa-plus-circle" aria-hidden="true"></i> ADMIN: Add New Custom Case
				</a>
			</div>`
		//output = strings.Replace(output, "<FORMCASE>", formCaseButton+"<FORMCASE>", -1)
		output = strings.Replace(output, "<FORMIMAGE>", formCaseButton+"<FORMIMAGE>", -1)
	} else { output = strings.Replace(output, "<FORMIMAGE>", "", -1) }
	// [END if_user]
	// ========== ========== ========== ========== ==========
    
    
    return output
}
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========

