import Redis from "iorejson";
import { promisify } from "util";

const client = new Redis({
  port: 6379,
  host: "34.121.89.39",
});

const get = promisify(client.get).bind(client);

export { get };
