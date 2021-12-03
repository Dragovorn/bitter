import * as sst from "@serverless-stack/resources";

export default class FrontendStack extends sst.Stack {
    frontend;

    constructor(scope, id, props) {
        super(scope, id, props)

        const {
            hosted_zone,
            domain_name,
            redirect_www,
        } = props;

        this.fronend = new sst.StaticSite(this, "Frontend", {
            path: "frontend/",
            buildOutput: "dist",
            buildCommand: "npm run prod",
            customDomain: redirect_www ? {
                hostedZone: hosted_zone,
                domainName: domain_name,
                domainAlias: "www." + domain_name,
            } : {
                hostedZone: hosted_zone,
                domainName: domain_name,
            }
        })
    }
}
