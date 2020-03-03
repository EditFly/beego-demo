package routers

import (
    "github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["beegoweb/controllers:CMSController"] = append(beego.GlobalControllerRouter["beegoweb/controllers:CMSController"],
        beego.ControllerComments{
            Method: "AllBlock",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoweb/controllers:CMSController"] = append(beego.GlobalControllerRouter["beegoweb/controllers:CMSController"],
        beego.ControllerComments{
            Method: "StaticBlock",
            Router: `/staticblock/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoweb/controllers:MainController"] = append(beego.GlobalControllerRouter["beegoweb/controllers:MainController"],
        beego.ControllerComments{
            Method: "DealConsumeRecord",
            Router: `/block`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
