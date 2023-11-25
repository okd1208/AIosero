import axios from 'axios';

const apiClient = axios.create({
  baseURL: 'http://localhost:8888',
});

export default {
  data () {
    return {
      info: '',
    }
  },
  methods: {
    async getPutDiscPosition(endpoint, userColor, positions, gameId, turn, lastPut) {
      try {
        const response = await axios.post(endpoint, {
          userColor: userColor,
          positions: positions,
          gameId: gameId,
          turn: turn,
          lastPut: lastPut
        });
        return response.data;
      } catch (error) {
        console.error(error);
      }
    },
    async createGameID() {
      try {
        const response = await apiClient.get("/api/v1/othello/newGameId");
        return response.data;
      } catch (error) {
        console.error(error);
      }
    }
  },
}
