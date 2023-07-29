import axios from 'axios';

export default {
  data () {
    return {
      info: '',
    }
  },
  methods: {
    async getPutDiscPosition(endpoint, userColor, positions) {
      try {
        const response = await axios.post(endpoint, {
          userColor: userColor,
          positions: positions
        });
        return response.data;
      } catch (error) {
        console.error(error);
      }
    }
  },
}
