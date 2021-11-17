import io from "socket.io-client";

const socket = io.connect('//34.134.225.148:8080');

export { socket };
