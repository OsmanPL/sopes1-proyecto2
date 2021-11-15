import { game } from "../models/mongo-game";

const getAll = async () => {
  let games = await game
    .find()
    .sort({ $natural: -1 })
    .then((value) => {
      return value.map((e) => {
        return {
          id: e.id,
          gameName: e.gamename,
          winner: e.winner,
          players: e.players,
          worker: e.worker,
        };
      });
    })
    .catch((e) => {
      return e;
    });
  return games;
};

const getTopWorker = async () => {
  let games = await game
    .aggregate([
      { $group: { _id: "$worker", count: { $sum: 1 } } },
      { $sort: { count: -1 } },
    ])
    .then((value) => {
      return value;
    })
    .catch((e) => {
      return e;
    });
  return games;
};

const getTopGames = async () => {
  let games = await game
    .aggregate([
      { $group: { _id: "$gamename", count: { $sum: 1 } } },
      { $sort: { count: -1 } },
      { $limit: 3 },
    ])
    .then((value) => {
      return value;
    })
    .catch((e) => {
      return e;
    });
  return games;
};

export { getAll, getTopWorker, getTopGames };
