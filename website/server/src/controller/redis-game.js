import { get } from "../database/redis";

const getLastGames = async () => {
  try {
    let games = await get("squid_game");

    if (games === null) return [];

    games = JSON.parse(games);
    games.length > 10 ? (games = games.slice(games.length - 10)) : null;

    return games;
  } catch (error) {
    return [];
  }
};

const getBestPlayers = async () => {
  try {
    let games = await get("squid_game");
    if (games === null) return [];
    games = JSON.parse(games);

    let top = {};

    games.forEach((game) => {
      top[game.winner] = top[game.winner] ? top[game.winner] + 1 : 1;
    });

    top = Object.keys(top).map((winner) => {
      return { winner: winner, total: top[winner] };
    });

    top = top.sort((a, b) => b.total - a.total);

    top.length > 10 ? (top = top.slice(top.length - 10)) : null;

    return top;
  } catch (error) {
    return error;
  }
};

const getPlayer = async (player) => {
  try {
    let games = await get("squid_game");

    if (games === null) return [];

    games = JSON.parse(games);
    let temp = games[games.length - 1];

    return {
      id: temp.id,
      gameName: temp.gameName,
      state: player === temp.winner,
    };
  } catch (error) {
    return [];
  }
};

export { getLastGames, getBestPlayers, getPlayer };
