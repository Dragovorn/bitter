import V1Stack from "./V1Stack";
import StorageStack from "./StorageStack";

export default function main(app) {
    // Set default runtime for all functions
    app.setDefaultFunctionProps({
        runtime: "go1.x"
    });

    const storage = new StorageStack(app, "persist");
    new V1Stack(app, "v1", {
        users_table: storage.users_table
    });
}
