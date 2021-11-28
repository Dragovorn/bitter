import * as sst from "@serverless-stack/resources";

export default class StorageStack extends sst.Stack {
    users_table;
    validation_table;
    username_index = "usernameIndex";
    user_id_index = "userIdIndex";

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

        this.validation_table = new sst.Table(this, "Validation", {
            fields: {
                code: sst.TableFieldType.NUMBER,
                user_id: sst.TableFieldType.STRING,
            },
            primaryIndex: {
                partitionKey: "code",
                sortKey: "user_id"
            },
            localIndexes: {
                userIdIndex: {
                    sortKey: "user_id"
                }
            }
        })
    }
}
