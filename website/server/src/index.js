import express from "express";
import dotenv from "dotenv";
import http from "http";
import cors from "cors";
import { Server } from "socket.io";
import * as mongo from "./database/mongo";
import { getAll, getTopWorker, getTopGames } from "./controller/mongo-game";
import {
  getLastGames,
  getBestPlayers,
  getPlayer,
} from "./controller/redis-game";

dotenv.config();
const app = express();
const server = http.createServer(app);
mongo.connect();

const io = new Server(server, {
  cors: {
    origin: "*",
    methods: ["GET", "POST"],
  },
});

io.on("connection", (socket) => {
  console.log(`User connected ${socket.id}`);

  setInterval(() => {
    getAll()
      .then((data) => {
        socket.emit("getDataMongo", data);
      })
      .catch((e) => console.log(e));
  }, 1500);

  setInterval(() => {
    getTopWorker()
      .then((data) => {
        socket.emit("getTopWorker", data);
      })
      .catch((e) => console.log(e));
  }, 1500);

  setInterval(() => {
    getTopGames()
      .then((data) => {
        socket.emit("getTopGames", data);
      })
      .catch((e) => console.log(e));
  }, 1500);

  setInterval(() => {
    getLastGames()
      .then((data) => {
        socket.emit("getLastGames", data);
      })
      .catch((e) => console.log(e));
  }, 1500);

  setInterval(() => {
    getBestPlayers()
      .then((data) => {
        socket.emit("getBestPlayers", data);
      })
      .catch((e) => console.log(e));
  }, 1500);

  socket.on("statplayer", (value) => {
    getPlayer(value)
      .then((data) => {
        socket.emit("getPlayer", data);
      })
      .catch((e) => console.log(e));
  });

  socket.on("disconnect", () => {
    console.log("User disconnected ", socket.id);
  });
});

app.set("port", process.env.PORT);
app.use(cors());
app.use(express.json());

app.get("/", (req, res) => {
  res.send("Hello World");
});

server.listen(app.get("port"), () => {
  console.log(`Server on port ${app.get("port")}`);
});
