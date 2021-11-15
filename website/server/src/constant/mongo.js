import dotenv from "dotenv";

dotenv.config();

const user = process.env.USER_MONGO;
const password = process.env.PASSWORD_MONGO;
const host = process.env.HOST_MONGO;
const database = process.env.DATABASE_MONGO;

export { user, password, host, database };
