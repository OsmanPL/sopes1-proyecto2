import express from "express";
import dotenv from "dotenv";
import http from "http";
import cors from "cors";
import { Server } from "socket.io";

const app = express();
const server = http.createServer(app);

const io = new Server(server, {
  cors: {
    origin: "*",
    methods: ["GET", "POST"],
  },
});

io.on("connection", (socket) => {
  console.log(`User connected ${socket.id}`);
  
  socket.on("disconnect", () => {
    console.log("User disconnected ", socket.id);
  });
});

dotenv.config();
app.set("port", process.env.PORT);
app.use(cors());
app.use(express.json());

app.get("/", (req, res) => {
  res.send("Hello World");
});

server.listen(app.get("port"), () => {
  console.log(`Server on port ${app.get("port")}`);
});
