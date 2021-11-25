import * as sst from "@serverless-stack/resources";

export default class StorageStack extends sst.Stack {
    users_table;

    constructor(scope, id, props) {
        super(scope, id, props);

        this.users_table = new sst.Table(this, "Users", {
            fields: {
                uid: sst.TableFieldType.STRING,
                username: sst.TableFieldType.STRING,
            },
            primaryIndex: {
                partitionKey: "uid",
                sortKey: "username",
            },
            localIndexes: {
                usernameIndex: {
                    sortKey: "username",
                }
            }
        });
    }
}
