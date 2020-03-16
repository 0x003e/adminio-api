package policy

import (
	iris "github.com/kataras/iris/v12"
	clients "github.com/rzrbld/adminio-api/clients"
	resph "github.com/rzrbld/adminio-api/handlers-response"
	log "log"
	strconv "strconv"
)

type Policy struct {
	policyName   string `json:"policyName"`
	policyString string `json:"policyString"`
}

type policySet struct {
	policyName string `json:"policyName"`
	entityName string `json:"entityName"`
	isGroup    string `json:"isGroup"`
}

// clients
var madmClnt = clients.MadmClnt
var minioClnt = clients.MinioClnt
var err error

var List = func(ctx iris.Context) {
	lp, err := madmClnt.ListCannedPolicies()
	var res = resph.BodyResHandler(ctx, err, lp)
	ctx.JSON(res)
}

var Add = func(ctx iris.Context) {
	p := Policy{}
	p.policyName = ctx.FormValue("policyName")
	p.policyString = ctx.FormValue("policyString")

	err = madmClnt.AddCannedPolicy(p.policyName, p.policyString)
	var res = resph.DefaultResHandler(ctx, err)
	ctx.JSON(res)
}

var Delete = func(ctx iris.Context) {
	p := policySet{}
	p.policyName = ctx.FormValue("policyName")

	err = madmClnt.RemoveCannedPolicy(p.policyName)
	var res = resph.DefaultResHandler(ctx, err)
	ctx.JSON(res)
}

var Set = func(ctx iris.Context) {
	p := policySet{}
	p.policyName = ctx.FormValue("policyName")
	p.entityName = ctx.FormValue("entityName")
	p.isGroup = ctx.FormValue("isGroup")

	isGroupBool, err := strconv.ParseBool(p.isGroup)

	if err != nil {
		log.Print(err)
		ctx.JSON(iris.Map{"error": err.Error()})
	}

	err = madmClnt.SetPolicy(p.policyName, p.entityName, isGroupBool)
	var res = resph.DefaultResHandler(ctx, err)
	ctx.JSON(res)
}
