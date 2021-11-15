import mongoose from "mongoose";

const Schema = mongoose.Schema;

const gameSchema = new Schema({
  id: Number,
  gamename: String,
  winner: String,
  players: Number,
  worker: String,
});

export const game = mongoose.model("games", gameSchema);
