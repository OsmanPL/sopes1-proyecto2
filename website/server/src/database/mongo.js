import mongoose from "mongoose";
import { host, password, user, database } from "../constant/mongo";

const url = `mongodb://${user}:${password}@${host}:27017/${database}?authSource=admin`;

const connect = async () => {
  await mongoose
    .connect(url, { useNewUrlParser: true, useUnifiedTopology: true })
    .then(() =>
      console.log("MongoDB is connected to:", mongoose.connection.name)
    )
    .catch((e) => console.log(e));
};

export { connect };
