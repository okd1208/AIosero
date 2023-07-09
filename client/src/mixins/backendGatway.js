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
    async getPutDiscPosition(userColor, positions) {
      apiClient.get('/api/v1/othello/next-move', {
        params: {
          userColor: userColor,
          positions: positions
        }
      })
      .then(function (response) {
        console.log(response.data);
      });
    },
  },
}
